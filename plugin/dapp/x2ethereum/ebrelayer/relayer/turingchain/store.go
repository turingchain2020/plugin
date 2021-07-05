package turingchain

import (
	"fmt"
	"sync/atomic"

	"github.com/turingchain2020/turingchain/types"
	ebTypes "github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/types"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/utils"
	"github.com/ethereum/go-ethereum/common"
)

//key ...
var (
	lastSyncHeightPrefix              = []byte("lastSyncHeight:")
	turingchainToEthBurnLockTxHashPrefix  = "turingchainToEthBurnLockTxHash"
	turingchainToEthBurnLockTxTotalAmount = []byte("turingchainToEthBurnLockTxTotalAmount")
	EthTxStatusCheckedIndex           = []byte("EthTxStatusCheckedIndex")
)

func calcRelay2EthTxhash(txindex int64) []byte {
	return []byte(fmt.Sprintf("%s-%012d", turingchainToEthBurnLockTxHashPrefix, txindex))
}

func (turingchainRelayer *Relayer4Turingchain) updateTotalTxAmount2Eth(total int64) error {
	totalTx := &types.Int64{
		Data: atomic.LoadInt64(&turingchainRelayer.totalTx4TuringchainToEth),
	}
	//更新成功见证的交易数
	return turingchainRelayer.db.Set(turingchainToEthBurnLockTxTotalAmount, types.Encode(totalTx))
}

func (turingchainRelayer *Relayer4Turingchain) getTotalTxAmount2Eth() int64 {
	totalTx, _ := utils.LoadInt64FromDB(turingchainToEthBurnLockTxTotalAmount, turingchainRelayer.db)
	return totalTx
}

func (turingchainRelayer *Relayer4Turingchain) setLastestRelay2EthTxhash(status, txhash string, txIndex int64) error {
	key := calcRelay2EthTxhash(txIndex)
	ethTxStatus := &ebTypes.EthTxStatus{
		Status: status,
		Txhash: txhash,
	}
	data := types.Encode(ethTxStatus)
	return turingchainRelayer.db.Set(key, data)
}

func (turingchainRelayer *Relayer4Turingchain) getEthTxhash(txIndex int64) (common.Hash, error) {
	key := calcRelay2EthTxhash(txIndex)
	ethTxStatus := &ebTypes.EthTxStatus{}
	data, err := turingchainRelayer.db.Get(key)
	if nil != err {
		return common.Hash{}, err
	}
	err = types.Decode(data, ethTxStatus)
	if nil != err {
		return common.Hash{}, err
	}
	return common.HexToHash(ethTxStatus.Txhash), nil
}

func (turingchainRelayer *Relayer4Turingchain) setStatusCheckedIndex(txIndex int64) error {
	index := &types.Int64{
		Data: txIndex,
	}
	data := types.Encode(index)
	return turingchainRelayer.db.Set(EthTxStatusCheckedIndex, data)
}

func (turingchainRelayer *Relayer4Turingchain) getStatusCheckedIndex() int64 {
	index, _ := utils.LoadInt64FromDB(EthTxStatusCheckedIndex, turingchainRelayer.db)
	return index
}

//获取上次同步到app的高度
func (turingchainRelayer *Relayer4Turingchain) loadLastSyncHeight() int64 {
	height, err := utils.LoadInt64FromDB(lastSyncHeightPrefix, turingchainRelayer.db)
	if nil != err && err != types.ErrHeightNotExist {
		relayerLog.Error("loadLastSyncHeight", "err:", err.Error())
		return 0
	}
	return height
}

func (turingchainRelayer *Relayer4Turingchain) setLastSyncHeight(syncHeight int64) {
	bytes := types.Encode(&types.Int64{Data: syncHeight})
	_ = turingchainRelayer.db.Set(lastSyncHeightPrefix, bytes)
}
