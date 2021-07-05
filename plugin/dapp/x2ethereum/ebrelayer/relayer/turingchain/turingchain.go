package turingchain

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	dbm "github.com/turingchain2020/turingchain/common/db"
	log "github.com/turingchain2020/turingchain/common/log/log15"
	"github.com/turingchain2020/turingchain/rpc/jsonclient"
	rpctypes "github.com/turingchain2020/turingchain/rpc/types"
	turingchainTypes "github.com/turingchain2020/turingchain/types"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/ethcontract/generated"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/ethinterface"
	relayerTx "github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/ethtxs"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/events"
	syncTx "github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/relayer/turingchain/transceiver/sync"
	ebTypes "github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/types"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/utils"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/types"
	ethCommon "github.com/ethereum/go-ethereum/common"
)

var relayerLog = log.New("module", "turingchain_relayer")

//Relayer4Turingchain ...
type Relayer4Turingchain struct {
	syncTxReceipts      *syncTx.TxReceipts
	ethClient           ethinterface.EthClientSpec
	rpcLaddr            string //用户向指定的blockchain节点进行rpc调用
	fetchHeightPeriodMs int64
	db                  dbm.DB
	lastHeight4Tx       int64 //等待被处理的具有相应的交易回执的高度
	matDegree           int32 //成熟度         heightSync2App    matDegress   height
	//passphase            string
	privateKey4Ethereum  *ecdsa.PrivateKey
	ethSender            ethCommon.Address
	bridgeRegistryAddr   ethCommon.Address
	oracleInstance       *generated.Oracle
	totalTx4TuringchainToEth int64
	statusCheckedIndex   int64
	ctx                  context.Context
	rwLock               sync.RWMutex
	unlock               chan int
}

// StartTuringchainRelayer : initializes a relayer which witnesses events on the turingchain network and relays them to Ethereum
func StartTuringchainRelayer(ctx context.Context, syncTxConfig *ebTypes.SyncTxConfig, registryAddr, provider string, db dbm.DB) *Relayer4Turingchain {
	chian33Relayer := &Relayer4Turingchain{
		rpcLaddr:            syncTxConfig.TuringchainHost,
		fetchHeightPeriodMs: syncTxConfig.FetchHeightPeriodMs,
		unlock:              make(chan int),
		db:                  db,
		ctx:                 ctx,
		bridgeRegistryAddr:  ethCommon.HexToAddress(registryAddr),
	}

	syncCfg := &ebTypes.SyncTxReceiptConfig{
		TuringchainHost:       syncTxConfig.TuringchainHost,
		PushHost:          syncTxConfig.PushHost,
		PushName:          syncTxConfig.PushName,
		PushBind:          syncTxConfig.PushBind,
		StartSyncHeight:   syncTxConfig.StartSyncHeight,
		StartSyncSequence: syncTxConfig.StartSyncSequence,
		StartSyncHash:     syncTxConfig.StartSyncHash,
	}

	client, err := relayerTx.SetupWebsocketEthClient(provider)
	if err != nil {
		panic(err)
	}
	chian33Relayer.ethClient = client
	chian33Relayer.totalTx4TuringchainToEth = chian33Relayer.getTotalTxAmount2Eth()
	chian33Relayer.statusCheckedIndex = chian33Relayer.getStatusCheckedIndex()

	go chian33Relayer.syncProc(syncCfg)
	return chian33Relayer
}

//QueryTxhashRelay2Eth ...
func (turingchainRelayer *Relayer4Turingchain) QueryTxhashRelay2Eth() ebTypes.Txhashes {
	txhashs := utils.QueryTxhashes([]byte(turingchainToEthBurnLockTxHashPrefix), turingchainRelayer.db)
	return ebTypes.Txhashes{Txhash: txhashs}
}

func (turingchainRelayer *Relayer4Turingchain) syncProc(syncCfg *ebTypes.SyncTxReceiptConfig) {
	_, _ = fmt.Fprintln(os.Stdout, "Pls unlock or import private key for Turingchain relayer")
	<-turingchainRelayer.unlock
	_, _ = fmt.Fprintln(os.Stdout, "Turingchain relayer starts to run...")

	turingchainRelayer.syncTxReceipts = syncTx.StartSyncTxReceipt(syncCfg, turingchainRelayer.db)
	turingchainRelayer.lastHeight4Tx = turingchainRelayer.loadLastSyncHeight()

	oracleInstance, err := relayerTx.RecoverOracleInstance(turingchainRelayer.ethClient, turingchainRelayer.bridgeRegistryAddr, turingchainRelayer.bridgeRegistryAddr)
	if err != nil {
		panic(err.Error())
	}
	turingchainRelayer.oracleInstance = oracleInstance

	timer := time.NewTicker(time.Duration(turingchainRelayer.fetchHeightPeriodMs) * time.Millisecond)
	for {
		select {
		case <-timer.C:
			height := turingchainRelayer.getCurrentHeight()
			relayerLog.Debug("syncProc", "getCurrentHeight", height)
			turingchainRelayer.onNewHeightProc(height)

		case <-turingchainRelayer.ctx.Done():
			timer.Stop()
			return
		}
	}
}

func (turingchainRelayer *Relayer4Turingchain) getCurrentHeight() int64 {
	var res rpctypes.Header
	ctx := jsonclient.NewRPCCtx(turingchainRelayer.rpcLaddr, "Turingchain.GetLastHeader", nil, &res)
	_, err := ctx.RunResult()
	if nil != err {
		relayerLog.Error("getCurrentHeight", "Failede due to:", err.Error())
	}
	return res.Height
}

func (turingchainRelayer *Relayer4Turingchain) onNewHeightProc(currentHeight int64) {
	//检查已经提交的交易结果
	turingchainRelayer.rwLock.Lock()
	for turingchainRelayer.statusCheckedIndex < turingchainRelayer.totalTx4TuringchainToEth {
		index := turingchainRelayer.statusCheckedIndex + 1
		txhash, err := turingchainRelayer.getEthTxhash(index)
		if nil != err {
			relayerLog.Error("onNewHeightProc", "getEthTxhash for index ", index, "error", err.Error())
			break
		}
		status := relayerTx.GetEthTxStatus(turingchainRelayer.ethClient, txhash)
		//按照提交交易的先后顺序检查交易，只要出现当前交易还在pending状态，就不再检查后续交易，等到下个区块再从该交易进行检查
		//TODO:可能会由于网络和打包挖矿的原因，使得交易执行顺序和提交顺序有差别，后续完善该检查逻辑
		if status == relayerTx.EthTxPending.String() {
			break
		}
		_ = turingchainRelayer.setLastestRelay2EthTxhash(status, txhash.Hex(), index)
		atomic.AddInt64(&turingchainRelayer.statusCheckedIndex, 1)
		_ = turingchainRelayer.setStatusCheckedIndex(turingchainRelayer.statusCheckedIndex)
	}
	turingchainRelayer.rwLock.Unlock()
	//未达到足够的成熟度，不进行处理
	//  +++++++++||++++++++++++||++++++++++||
	//           ^             ^           ^
	// lastHeight4Tx    matDegress   currentHeight
	for turingchainRelayer.lastHeight4Tx+int64(turingchainRelayer.matDegree)+1 <= currentHeight {
		relayerLog.Info("onNewHeightProc", "currHeight", currentHeight, "lastHeight4Tx", turingchainRelayer.lastHeight4Tx)

		lastHeight4Tx := turingchainRelayer.lastHeight4Tx
		TxReceipts, err := turingchainRelayer.syncTxReceipts.GetNextValidTxReceipts(lastHeight4Tx)
		if nil == TxReceipts || nil != err {
			if err != nil {
				relayerLog.Error("onNewHeightProc", "Failed to GetNextValidTxReceipts due to:", err.Error())
			}
			break
		}
		relayerLog.Debug("onNewHeightProc", "currHeight", currentHeight, "valid tx receipt with height:", TxReceipts.Height)

		txs := TxReceipts.Tx
		for i, tx := range txs {
			//检查是否为lns的交易(包括平行链：user.p.xxx.lns)，将闪电网络交易进行收集
			if 0 != bytes.Compare(tx.Execer, []byte(relayerTx.X2Eth)) &&
				(len(tx.Execer) > 4 && string(tx.Execer[(len(tx.Execer)-4):]) != "."+relayerTx.X2Eth) {
				relayerLog.Debug("onNewHeightProc, the tx is not x2ethereum", "Execer", string(tx.Execer), "height:", TxReceipts.Height)
				continue
			}
			var ss types.X2EthereumAction
			_ = turingchainTypes.Decode(tx.Payload, &ss)
			actionName := ss.GetActionName()
			if relayerTx.BurnAction == actionName || relayerTx.LockAction == actionName {
				relayerLog.Debug("^_^ ^_^ Processing turingchain tx receipt", "ActionName", actionName, "fromAddr", tx.From(), "exec", string(tx.Execer))
				actionEvent := getOracleClaimType(actionName)
				if err := turingchainRelayer.handleBurnLockMsg(actionEvent, TxReceipts.ReceiptData[i], tx.Hash()); nil != err {
					errInfo := fmt.Sprintf("Failed to handleBurnLockMsg due to:%s", err.Error())
					panic(errInfo)
				}
			}
		}
		turingchainRelayer.lastHeight4Tx = TxReceipts.Height
		turingchainRelayer.setLastSyncHeight(turingchainRelayer.lastHeight4Tx)
	}
}

// getOracleClaimType : sets the OracleClaim's claim type based upon the witnessed event type
func getOracleClaimType(eventType string) events.Event {
	var claimType events.Event

	switch eventType {
	case events.MsgBurn.String():
		claimType = events.Event(events.ClaimTypeBurn)
	case events.MsgLock.String():
		claimType = events.Event(events.ClaimTypeLock)
	default:
		panic(errors.New("eventType invalid"))
	}

	return claimType
}

// handleBurnLockMsg : parse event data as a TuringchainMsg, package it into a ProphecyClaim, then relay tx to the Ethereum Network
func (turingchainRelayer *Relayer4Turingchain) handleBurnLockMsg(claimEvent events.Event, receipt *turingchainTypes.ReceiptData, turingchainTxHash []byte) error {
	relayerLog.Info("handleBurnLockMsg", "Received tx with hash", ethCommon.Bytes2Hex(turingchainTxHash))

	// Parse the witnessed event's data into a new TuringchainMsg
	turingchainMsg := relayerTx.ParseBurnLockTxReceipt(claimEvent, receipt)
	if nil == turingchainMsg {
		//收到执行失败的交易，直接跳过
		relayerLog.Error("handleBurnLockMsg", "Received failed tx with hash", ethCommon.Bytes2Hex(turingchainTxHash))
		return nil
	}

	// Parse the TuringchainMsg into a ProphecyClaim for relay to Ethereum
	prophecyClaim := relayerTx.TuringchainMsgToProphecyClaim(*turingchainMsg)

	// Relay the TuringchainMsg to the Ethereum network
	txhash, err := relayerTx.RelayOracleClaimToEthereum(turingchainRelayer.oracleInstance, turingchainRelayer.ethClient, turingchainRelayer.ethSender, claimEvent, prophecyClaim, turingchainRelayer.privateKey4Ethereum, turingchainTxHash)
	if nil != err {
		return err
	}

	//保存交易hash，方便查询
	atomic.AddInt64(&turingchainRelayer.totalTx4TuringchainToEth, 1)
	txIndex := atomic.LoadInt64(&turingchainRelayer.totalTx4TuringchainToEth)
	if err = turingchainRelayer.updateTotalTxAmount2Eth(txIndex); nil != err {
		relayerLog.Error("handleLogNewProphecyClaimEvent", "Failed to RelayLockToTuringchain due to:", err.Error())
		return err
	}
	if err = turingchainRelayer.setLastestRelay2EthTxhash(relayerTx.EthTxPending.String(), txhash, txIndex); nil != err {
		relayerLog.Error("handleLogNewProphecyClaimEvent", "Failed to RelayLockToTuringchain due to:", err.Error())
		return err
	}
	return nil
}
