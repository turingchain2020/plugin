// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"encoding/json"
	"fmt"

	"github.com/turingchain2020/turingchain/common/address"
	"github.com/turingchain2020/turingchain/common/log/log15"
	"github.com/turingchain2020/turingchain/types"
)

var tlog = log15.New("module", ParaX)

// paracross 执行器的日志类型
const (

	// TyLogParacrossCommit commit log key
	TyLogParacrossCommit = 650
	// TyLogParacrossCommitDone commit down key
	TyLogParacrossCommitDone = 651
	// record 和 commit 不一样， 对应高度完成共识后收到commit 交易
	// 这个交易就不参与共识, 只做记录
	// TyLogParacrossCommitRecord commit record key
	TyLogParacrossCommitRecord = 652
	// TyLogParaAssetTransfer asset transfer log key
	TyLogParaAssetTransfer = 653
	// TyLogParaAssetWithdraw asset withdraw log key
	TyLogParaAssetWithdraw = 654
	//在平行链上保存节点参与共识的数据
	// TyLogParacrossMiner miner log key
	TyLogParacrossMiner = 655
	// TyLogParaAssetDeposit asset deposit log key
	TyLogParaAssetDeposit = 656
	// TyLogParaNodeConfig config super node log key
	TyLogParaNodeConfig            = 657
	TyLogParaNodeVoteDone          = 658
	TyLogParaNodeGroupAddrsUpdate  = 659
	TyLogParaNodeGroupConfig       = 660
	TyLogParaNodeStatusUpdate      = 661
	TyLogParaNodeGroupStatusUpdate = 664
	TyLogParaSelfConsStageConfig   = 665
	TyLogParaStageVoteDone         = 666
	TyLogParaStageGroupUpdate      = 667
	//TyLogParaCrossAssetTransfer 统一的跨链资产转移
	TyLogParaCrossAssetTransfer = 670
	TyLogParaBindMinerAddr      = 671
	TyLogParaBindMinerNode      = 672
)

// action type
const (
	// ParacrossActionCommit paracross consensus commit action
	ParacrossActionCommit = iota
	// ParacrossActionMiner paracross consensus miner action
	ParacrossActionMiner
	// ParacrossActionTransfer paracross asset transfer action
	ParacrossActionTransfer
	// ParacrossActionWithdraw paracross asset withdraw action
	ParacrossActionWithdraw
	// ParacrossActionTransferToExec asset transfer to exec
	ParacrossActionTransferToExec
	// ParacrossActionParaBindMiner  para chain bind super node miner
	ParacrossActionParaBindMiner
)

const (
	paraCrossTransferActionTypeStart = 10000
	//paraCrossTransferActionTypeEnd   = 10100
)

const (
	// ParacrossActionAssetTransfer mainchain paracross asset transfer key
	ParacrossActionAssetTransfer = iota + paraCrossTransferActionTypeStart
	// ParacrossActionAssetWithdraw mainchain paracross asset withdraw key
	ParacrossActionAssetWithdraw
	//ParacrossActionNodeConfig para super node config
	ParacrossActionNodeConfig
	//ParacrossActionNodeGroupApply apply for node group initially
	ParacrossActionNodeGroupApply
	//ParacrossActionSelfStageConfig apply for self consensus stage config
	ParacrossActionSelfStageConfig
	// ParacrossActionCrossAssetTransfer crossChain asset transfer key
	ParacrossActionCrossAssetTransfer
)

//paracross asset porcess
const (
	ParacrossNoneTransfer = iota
	ParacrossMainAssetTransfer
	ParacrossMainAssetWithdraw
	ParacrossParaAssetTransfer
	ParacrossParaAssetWithdraw
)

// status
const (
	// ParacrossStatusCommiting commit status
	ParacrossStatusCommiting = iota
	// ParacrossStatusCommitDone commit done status
	ParacrossStatusCommitDone
)

// config op
const (
	ParaOpNewApply = iota + 1
	ParaOpVote
	ParaOpQuit
	ParaOpCancel
	ParaOpModify
)

// node vote op
const (
	ParaVoteInvalid = iota
	ParaVoteYes
	ParaVoteNo
	ParaVoteEnd
)

//config yes or no
const (
	ParaConfigInvalid = iota
	ParaConfigYes
	ParaConfigNo
)

// ParaNodeVoteStr ...
var ParaNodeVoteStr = []string{"invalid", "yes", "no"}

//针对addr申请的id的生命周期
const (
	// ParaApplyJoining apply for join group
	ParaApplyJoining = iota + 1
	// ParaApplyQuiting apply for quiting group
	ParaApplyQuiting
	// ParaApplyClosed id voting closed
	ParaApplyClosed
	// ParaApplyCanceled to cancel apply of joining or quiting
	ParaApplyCanceled
	// ParaApplyVoting record voting status
	ParaApplyVoting
)

//针对addr本身的生命周期，addr维护了申请id和quit id，方便查询如coinfrozen等额外信息
const (
	// ParaApplyJoined pass to add by votes
	ParaApplyJoined = iota + 10
	// ParaApplyQuited pass to quite by votes
	ParaApplyQuited
)

const (
	//ParacrossNodeGroupApply apply for para chain node group initially
	ParacrossNodeGroupApply = iota + 1
	//ParacrossNodeGroupApprove super manager approve the apply
	ParacrossNodeGroupApprove
	//ParacrossNodeGroupQuit applyer quit the apply when not be approved
	ParacrossNodeGroupQuit
	//ParacrossNodeGroupModify applyer modify some parameters
	ParacrossNodeGroupModify
)

var (
	// ParacrossActionCommitStr Commit string
	ParacrossActionCommitStr = string("Commit")
	paracrossTransferPerfix  = "crossPara."
	// ParacrossActionAssetTransferStr asset transfer key
	ParacrossActionAssetTransferStr = paracrossTransferPerfix + string("AssetTransfer")
	// ParacrossActionAssetWithdrawStr asset withdraw key
	ParacrossActionAssetWithdrawStr = paracrossTransferPerfix + string("AssetWithdraw")
	// ParacrossActionTransferStr trasfer key
	ParacrossActionTransferStr = paracrossTransferPerfix + string("Transfer")
	// ParacrossActionTransferToExecStr transfer to exec key
	ParacrossActionTransferToExecStr = paracrossTransferPerfix + string("TransferToExec")
	// ParacrossActionWithdrawStr withdraw key
	ParacrossActionWithdrawStr = paracrossTransferPerfix + string("Withdraw")
)

// CalcMinerHeightKey get miner key
func CalcMinerHeightKey(title string, height int64) []byte {
	paraVoteHeightKey := "LODB-paracross-titleVoteHeight-"
	return []byte(fmt.Sprintf(paraVoteHeightKey+"%s-%012d", title, height))
}

// CreateRawCommitTx4MainChain create commit tx to main chain
func CreateRawCommitTx4MainChain(cfg *types.TuringchainConfig, status *ParacrossCommitAction, name string, fee int64) (*types.Transaction, error) {
	return createRawCommitTx(cfg, status, name, fee)
}

func createRawCommitTx(cfg *types.TuringchainConfig, commit *ParacrossCommitAction, name string, feeRate int64) (*types.Transaction, error) {
	action := &ParacrossAction{
		Ty:    ParacrossActionCommit,
		Value: &ParacrossAction_Commit{commit},
	}
	tx := &types.Transaction{
		Execer:  []byte(name),
		Payload: types.Encode(action),
		To:      address.ExecAddress(name),
		Expire:  types.Now().Unix() + int64(120), //120s
	}
	tx, err := types.FormatTx(cfg, name, tx)
	if err != nil {
		return nil, err
	}
	if feeRate != 0 {
		tx.Fee, err = tx.GetRealFee(feeRate)
		if err != nil {
			return nil, err
		}
	}
	return tx, nil
}

// CreateRawAssetTransferTx create asset transfer tx
func CreateRawAssetTransferTx(cfg *types.TuringchainConfig, param *types.CreateTx) (*types.Transaction, error) {
	// 跨链交易需要在主链和平行链上执行， 所以应该可以在主链和平行链上构建
	if !types.IsParaExecName(param.GetExecName()) {
		tlog.Error("CreateRawAssetTransferTx", "exec", param.GetExecName())
		return nil, types.ErrInvalidParam
	}

	transfer := &ParacrossAction{}
	if !param.IsWithdraw {
		v := &ParacrossAction_AssetTransfer{AssetTransfer: &types.AssetsTransfer{
			Amount: param.Amount, Note: param.GetNote(), To: param.GetTo(), Cointoken: param.TokenSymbol}}
		transfer.Value = v
		transfer.Ty = ParacrossActionAssetTransfer
	} else {
		v := &ParacrossAction_AssetWithdraw{AssetWithdraw: &types.AssetsWithdraw{
			Amount: param.Amount, Note: param.GetNote(), To: param.GetTo(), Cointoken: param.TokenSymbol, ExecName: param.ExecName}}
		transfer.Value = v
		transfer.Ty = ParacrossActionAssetWithdraw
	}
	tx := &types.Transaction{
		Execer:  []byte(param.GetExecName()),
		Payload: types.Encode(transfer),
		To:      address.ExecAddress(param.GetExecName()),
		Fee:     param.Fee,
	}
	tx, err := types.FormatTx(cfg, param.GetExecName(), tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// CreateRawMinerTx create miner tx
func CreateRawMinerTx(cfg *types.TuringchainConfig, value *ParacrossMinerAction) (*types.Transaction, error) {

	action := &ParacrossAction{
		Ty:    ParacrossActionMiner,
		Value: &ParacrossAction_Miner{value},
	}
	tx := &types.Transaction{
		Execer:  []byte(cfg.ExecName(ParaX)),
		Payload: types.Encode(action),
		Nonce:   0, //for consensus purpose, block hash need same, different auth node need keep totally same vote tx
		To:      address.ExecAddress(cfg.ExecName(ParaX)),
		ChainID: cfg.GetChainID(),
	}
	err := tx.SetRealFee(cfg.GetMinTxFeeRate())
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// CreateRawTransferTx create paracross asset transfer tx with transfer and withdraw
func (p ParacrossType) CreateRawTransferTx(action string, param json.RawMessage) (*types.Transaction, error) {
	tlog.Info("ParacrossType CreateTx", "action", action, "msg", string(param))
	tx, err := p.ExecTypeBase.CreateTx(action, param)
	if err != nil {
		tlog.Error("ParacrossType CreateTx failed", "err", err, "action", action, "msg", string(param))
		return nil, err
	}
	cfg := p.GetConfig()
	if !cfg.IsPara() {
		var transfer ParacrossAction
		err = types.Decode(tx.Payload, &transfer)
		if err != nil {
			tlog.Error("ParacrossType CreateTx failed", "decode payload err", err, "action", action, "msg", string(param))
			return nil, err
		}
		if action == "Transfer" {
			tx.To = transfer.GetTransfer().To
		} else if action == "Withdraw" {
			tx.To = transfer.GetWithdraw().To
		} else if action == "TransferToExec" {
			tx.To = transfer.GetTransferToExec().To
		}
	}

	return tx, nil
}

//GetDappForkHeight get paracross dapp fork height
func GetDappForkHeight(cfg *types.TuringchainConfig, forkKey string) int64 {
	var forkHeight int64
	if cfg.IsPara() {
		key := forkKey
		switch forkKey {
		case ForkCommitTx:
			key = MainForkParacrossCommitTx
		case ForkLoopCheckCommitTxDone:
			key = MainLoopCheckCommitTxDoneForkHeight
		}

		forkHeight = types.Conf(cfg, ParaPrefixConsSubConf).GInt(key)
		if forkHeight <= 0 {
			forkHeight = types.MaxHeight
		}
	} else {
		forkHeight = cfg.GetDappFork(ParaX, forkKey)

		// CI特殊处理，主链是local，fork都是0，平行链有些配置项需要设置为非0，不然获取到的高度为MaxHeight
		if cfg.IsLocal() {
			switch forkKey {
			case ForkCommitTx:
				forkHeight = 10
			case ForkLoopCheckCommitTxDone:
				forkHeight = 60
			}
		}
	}
	return forkHeight
}

// IsParaForkHeight check height more than fork height
func IsParaForkHeight(cfg *types.TuringchainConfig, height int64, forkKey string) bool {
	return height >= GetDappForkHeight(cfg, forkKey)
}
