package types

import (
	"encoding/json"
	"reflect"

	log "github.com/turingchain2020/turingchain/common/log/log15"
	"github.com/turingchain2020/turingchain/types"
)

/*
 * 交易相关类型定义
 * 交易action通常有对应的log结构，用于交易回执日志记录
 * 每一种action和log需要用id数值和name名称加以区分
 */

var (
	//X2ethereumX 执行器名称定义
	X2ethereumX = "x2ethereum"
	//定义actionMap
	actionMap = map[string]int32{
		NameEth2TuringchainAction:           TyEth2TuringchainAction,
		NameWithdrawEthAction:           TyWithdrawEthAction,
		NameWithdrawTuringchainAction:       TyWithdrawTuringchainAction,
		NameTuringchainToEthAction:          TyTuringchainToEthAction,
		NameAddValidatorAction:          TyAddValidatorAction,
		NameRemoveValidatorAction:       TyRemoveValidatorAction,
		NameModifyPowerAction:           TyModifyPowerAction,
		NameSetConsensusThresholdAction: TySetConsensusThresholdAction,
		NameTransferAction:              TyTransferAction,
		NameTransferToExecAction:        TyTransferToExecAction,
		NameWithdrawFromExecAction:      TyWithdrawFromExecAction,
	}
	//定义log的id和具体log类型及名称，填入具体自定义log类型
	logMap = map[int64]*types.LogInfo{
		TyEth2TuringchainLog:           {Ty: reflect.TypeOf(ReceiptEth2Turingchain{}), Name: "LogEth2Turingchain"},
		TyWithdrawEthLog:           {Ty: reflect.TypeOf(ReceiptEth2Turingchain{}), Name: "LogWithdrawEth"},
		TyWithdrawTuringchainLog:       {Ty: reflect.TypeOf(ReceiptTuringchainToEth{}), Name: "LogWithdrawTuringchain"},
		TyTuringchainToEthLog:          {Ty: reflect.TypeOf(ReceiptTuringchainToEth{}), Name: "LogTuringchainToEth"},
		TyAddValidatorLog:          {Ty: reflect.TypeOf(ReceiptValidator{}), Name: "LogAddValidator"},
		TyRemoveValidatorLog:       {Ty: reflect.TypeOf(ReceiptValidator{}), Name: "LogRemoveValidator"},
		TyModifyPowerLog:           {Ty: reflect.TypeOf(ReceiptValidator{}), Name: "LogModifyPower"},
		TySetConsensusThresholdLog: {Ty: reflect.TypeOf(ReceiptSetConsensusThreshold{}), Name: "LogSetConsensusThreshold"},
		TyProphecyLog:              {Ty: reflect.TypeOf(ReceiptEthProphecy{}), Name: "LogEthProphecy"},
		TyTransferLog:              {Ty: reflect.TypeOf(types.ReceiptAccountTransfer{}), Name: "LogTransfer"},
		TyTransferToExecLog:        {Ty: reflect.TypeOf(types.ReceiptExecAccountTransfer{}), Name: "LogTokenExecTransfer"},
		TyWithdrawFromExecLog:      {Ty: reflect.TypeOf(types.ReceiptExecAccountTransfer{}), Name: "LogTokenExecWithdraw"},
	}
	tlog = log.New("module", "x2ethereum.types")
)

// init defines a register function
func init() {
	types.AllowUserExec = append(types.AllowUserExec, []byte(X2ethereumX))
	//注册合约启用高度
	types.RegFork(X2ethereumX, InitFork)
	types.RegExec(X2ethereumX, InitExecutor)
}

// InitFork defines register fork
func InitFork(cfg *types.TuringchainConfig) {
	cfg.RegisterDappFork(X2ethereumX, "Enable", 0)
}

// InitExecutor defines register executor
func InitExecutor(cfg *types.TuringchainConfig) {
	types.RegistorExecutor(X2ethereumX, NewType(cfg))
}

//X2ethereumType ...
type X2ethereumType struct {
	types.ExecTypeBase
}

//NewType ...
func NewType(cfg *types.TuringchainConfig) *X2ethereumType {
	c := &X2ethereumType{}
	c.SetChild(c)
	c.SetConfig(cfg)
	return c
}

//GetName ...
func (x *X2ethereumType) GetName() string {
	return X2ethereumX
}

// GetPayload 获取合约action结构
func (x *X2ethereumType) GetPayload() types.Message {
	return &X2EthereumAction{}
}

// GetTypeMap 获取合约action的id和name信息
func (x *X2ethereumType) GetTypeMap() map[string]int32 {
	return actionMap
}

// GetLogMap 获取合约log相关信息
func (x *X2ethereumType) GetLogMap() map[int64]*types.LogInfo {
	return logMap
}

// ActionName get PrivacyType action name
func (x X2ethereumType) ActionName(tx *types.Transaction) string {
	var action X2EthereumAction
	err := types.Decode(tx.Payload, &action)
	if err != nil {
		return "unknown-x2ethereum-err"
	}
	tlog.Info("ActionName", "ActionName", action.GetActionName())
	return action.GetActionName()
}

// GetActionName get action name
func (action *X2EthereumAction) GetActionName() string {
	if action.Ty == TyEth2TuringchainAction && action.GetEth2TuringchainLock() != nil {
		return "Eth2TuringchainLock"
	} else if action.Ty == TyWithdrawEthAction && action.GetEth2TuringchainBurn() != nil {
		return "Eth2TuringchainBurn"
	} else if action.Ty == TyWithdrawTuringchainAction && action.GetTuringchainToEthBurn() != nil {
		return "TuringchainToEthBurn"
	} else if action.Ty == TyTuringchainToEthAction && action.GetTuringchainToEthLock() != nil {
		return "TuringchainToEthLock"
	} else if action.Ty == TyAddValidatorAction && action.GetAddValidator() != nil {
		return "AddValidator"
	} else if action.Ty == TyRemoveValidatorAction && action.GetRemoveValidator() != nil {
		return "RemoveValidator"
	} else if action.Ty == TyModifyPowerAction && action.GetModifyPower() != nil {
		return "ModifyPower"
	} else if action.Ty == TySetConsensusThresholdAction && action.GetSetConsensusThreshold() != nil {
		return "SetConsensusThreshold"
	} else if action.Ty == TyTransferAction && action.GetTransfer() != nil {
		return "Transfer"
	} else if action.Ty == TyTransferToExecAction && action.GetTransferToExec() != nil {
		return "TransferToExec"
	} else if action.Ty == TyWithdrawFromExecAction && action.GetWithdrawFromExec() != nil {
		return "WithdrawFromExec"
	}
	return "unknown-x2ethereum"
}

// CreateTx token 创建合约
func (x *X2ethereumType) CreateTx(action string, msg json.RawMessage) (*types.Transaction, error) {
	tx, err := x.ExecTypeBase.CreateTx(action, msg)
	if err != nil {
		tlog.Error("token CreateTx failed", "err", err, "action", action, "msg", string(msg))
		return nil, err
	}
	cfg := x.GetConfig()
	if !cfg.IsPara() {
		var transfer X2EthereumAction
		err = types.Decode(tx.Payload, &transfer)
		if err != nil {
			tlog.Error("token CreateTx failed", "decode payload err", err, "action", action, "msg", string(msg))
			return nil, err
		}
		if action == "Transfer" {
			tx.To = transfer.GetTransfer().To
		} else if action == "Withdraw" {
			tx.To = transfer.GetWithdrawFromExec().To
		} else if action == "TransferToExec" {
			tx.To = transfer.GetTransferToExec().To
		}
	}
	return tx, nil
}
