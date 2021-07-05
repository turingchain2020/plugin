package executor

import (
	"strconv"

	"github.com/turingchain2020/turingchain/types"
	x2eTy "github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/types"
)

/*
 * 实现交易相关数据本地执行，数据不上链
 * 非关键数据，本地存储(localDB), 用于辅助查询，效率高
 */

func (x *x2ethereum) ExecLocal_Eth2TuringchainLock(payload *x2eTy.Eth2Turingchain, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := x.execLocal(receiptData)
	if err != nil {
		return set, err
	}
	return x.addAutoRollBack(tx, set.KV), nil
}

func (x *x2ethereum) ExecLocal_Eth2TuringchainBurn(payload *x2eTy.Eth2Turingchain, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := x.execLocal(receiptData)
	if err != nil {
		return set, err
	}
	return x.addAutoRollBack(tx, set.KV), nil
}

func (x *x2ethereum) ExecLocal_TuringchainToEthBurn(payload *x2eTy.TuringchainToEth, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := x.execLocal(receiptData)
	if err != nil {
		return set, err
	}
	return x.addAutoRollBack(tx, set.KV), nil
}

func (x *x2ethereum) ExecLocal_TuringchainToEthLock(payload *x2eTy.TuringchainToEth, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := x.execLocal(receiptData)
	if err != nil {
		return set, err
	}
	return x.addAutoRollBack(tx, set.KV), nil
}

func (x *x2ethereum) ExecLocal_AddValidator(payload *x2eTy.MsgValidator, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	dbSet := &types.LocalDBSet{}
	//implement code
	return x.addAutoRollBack(tx, dbSet.KV), nil
}

func (x *x2ethereum) ExecLocal_RemoveValidator(payload *x2eTy.MsgValidator, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	dbSet := &types.LocalDBSet{}
	//implement code
	return x.addAutoRollBack(tx, dbSet.KV), nil
}

func (x *x2ethereum) ExecLocal_ModifyPower(payload *x2eTy.MsgValidator, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	dbSet := &types.LocalDBSet{}
	//implement code
	return x.addAutoRollBack(tx, dbSet.KV), nil
}

func (x *x2ethereum) ExecLocal_SetConsensusThreshold(payload *x2eTy.MsgConsensusThreshold, tx *types.Transaction, receiptData *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	dbSet := &types.LocalDBSet{}
	//implement code
	return x.addAutoRollBack(tx, dbSet.KV), nil
}

//设置自动回滚
func (x *x2ethereum) addAutoRollBack(tx *types.Transaction, kv []*types.KeyValue) *types.LocalDBSet {
	dbSet := &types.LocalDBSet{}
	dbSet.KV = x.AddRollbackKV(tx, tx.Execer, kv)
	return dbSet
}

func (x *x2ethereum) execLocal(receiptData *types.ReceiptData) (*types.LocalDBSet, error) {
	dbSet := &types.LocalDBSet{}
	for _, log := range receiptData.Logs {
		switch log.Ty {
		case x2eTy.TyEth2TuringchainLog:
			var receiptEth2Turingchain x2eTy.ReceiptEth2Turingchain
			err := types.Decode(log.Log, &receiptEth2Turingchain)
			if err != nil {
				return nil, err
			}

			nb, err := x.GetLocalDB().Get(x2eTy.CalTokenSymbolTotalLockOrBurnAmount(receiptEth2Turingchain.IssuerDotSymbol, receiptEth2Turingchain.TokenAddress, x2eTy.DirEth2Turingchain, "lock"))
			if err != nil && err != types.ErrNotFound {
				return nil, err
			}
			var now x2eTy.ReceiptQuerySymbolAssetsByTxType
			err = types.Decode(nb, &now)
			if err != nil {
				return nil, err
			}
			preAmount, _ := strconv.ParseFloat(x2eTy.TrimZeroAndDot(now.TotalAmount), 64)
			nowAmount, _ := strconv.ParseFloat(x2eTy.TrimZeroAndDot(receiptEth2Turingchain.Amount), 64)
			TokenAssetsByTxTypeBytes := types.Encode(&x2eTy.ReceiptQuerySymbolAssetsByTxType{
				TokenSymbol: receiptEth2Turingchain.IssuerDotSymbol,
				TxType:      "lock",
				TotalAmount: strconv.FormatFloat(preAmount+nowAmount, 'f', 4, 64),
				Direction:   1,
			})
			dbSet.KV = append(dbSet.KV, &types.KeyValue{
				Key:   x2eTy.CalTokenSymbolTotalLockOrBurnAmount(receiptEth2Turingchain.IssuerDotSymbol, receiptEth2Turingchain.TokenAddress, x2eTy.DirEth2Turingchain, "lock"),
				Value: TokenAssetsByTxTypeBytes,
			})

			nb, err = x.GetLocalDB().Get(x2eTy.CalTokenSymbolToTokenAddress(receiptEth2Turingchain.IssuerDotSymbol))
			if err != nil && err != types.ErrNotFound {
				return nil, err
			}
			var t x2eTy.ReceiptTokenToTokenAddress
			err = types.Decode(nb, &t)
			if err != nil {
				return nil, err
			}
			var exist bool
			for _, addr := range t.TokenAddress {
				if addr == receiptEth2Turingchain.TokenAddress {
					exist = true
				}
			}
			if !exist {
				t.TokenAddress = append(t.TokenAddress, receiptEth2Turingchain.TokenAddress)
			}
			TokenToTokenAddressBytes := types.Encode(&x2eTy.ReceiptTokenToTokenAddress{
				TokenAddress: t.TokenAddress,
			})
			dbSet.KV = append(dbSet.KV, &types.KeyValue{
				Key:   x2eTy.CalTokenSymbolToTokenAddress(receiptEth2Turingchain.IssuerDotSymbol),
				Value: TokenToTokenAddressBytes,
			})
		case x2eTy.TyWithdrawEthLog:
			var receiptEth2Turingchain x2eTy.ReceiptEth2Turingchain
			err := types.Decode(log.Log, &receiptEth2Turingchain)
			if err != nil {
				return nil, err
			}

			nb, err := x.GetLocalDB().Get(x2eTy.CalTokenSymbolTotalLockOrBurnAmount(receiptEth2Turingchain.IssuerDotSymbol, receiptEth2Turingchain.TokenAddress, x2eTy.DirEth2Turingchain, "withdraw"))
			if err != nil && err != types.ErrNotFound {
				return nil, err
			}
			var now x2eTy.ReceiptQuerySymbolAssetsByTxType
			err = types.Decode(nb, &now)
			if err != nil {
				return nil, err
			}

			preAmount, _ := strconv.ParseFloat(x2eTy.TrimZeroAndDot(now.TotalAmount), 64)
			nowAmount, _ := strconv.ParseFloat(x2eTy.TrimZeroAndDot(receiptEth2Turingchain.Amount), 64)
			TokenAssetsByTxTypeBytes := types.Encode(&x2eTy.ReceiptQuerySymbolAssetsByTxType{
				TokenSymbol: receiptEth2Turingchain.IssuerDotSymbol,
				TxType:      "withdraw",
				TotalAmount: strconv.FormatFloat(preAmount+nowAmount, 'f', 4, 64),
				Direction:   2,
			})
			dbSet.KV = append(dbSet.KV, &types.KeyValue{
				Key:   x2eTy.CalTokenSymbolTotalLockOrBurnAmount(receiptEth2Turingchain.IssuerDotSymbol, receiptEth2Turingchain.TokenAddress, x2eTy.DirEth2Turingchain, "withdraw"),
				Value: TokenAssetsByTxTypeBytes,
			})
		case x2eTy.TyTuringchainToEthLog:
			var receiptTuringchainToEth x2eTy.ReceiptTuringchainToEth
			err := types.Decode(log.Log, &receiptTuringchainToEth)
			if err != nil {
				return nil, err
			}

			nb, err := x.GetLocalDB().Get(x2eTy.CalTokenSymbolTotalLockOrBurnAmount(receiptTuringchainToEth.IssuerDotSymbol, receiptTuringchainToEth.TokenContract, x2eTy.DirTuringchainToEth, "lock"))
			if err != nil && err != types.ErrNotFound {
				return nil, err
			}
			var now x2eTy.ReceiptQuerySymbolAssetsByTxType
			err = types.Decode(nb, &now)
			if err != nil {
				return nil, err
			}

			preAmount, _ := strconv.ParseFloat(x2eTy.TrimZeroAndDot(now.TotalAmount), 64)
			nowAmount, _ := strconv.ParseFloat(x2eTy.TrimZeroAndDot(receiptTuringchainToEth.Amount), 64)
			TokenAssetsByTxTypeBytes := types.Encode(&x2eTy.ReceiptQuerySymbolAssetsByTxType{
				TokenSymbol: receiptTuringchainToEth.IssuerDotSymbol,
				TxType:      "lock",
				TotalAmount: strconv.FormatFloat(preAmount+nowAmount, 'f', 4, 64),
				Direction:   1,
			})
			dbSet.KV = append(dbSet.KV, &types.KeyValue{
				Key:   x2eTy.CalTokenSymbolTotalLockOrBurnAmount(receiptTuringchainToEth.IssuerDotSymbol, receiptTuringchainToEth.TokenContract, x2eTy.DirTuringchainToEth, "lock"),
				Value: TokenAssetsByTxTypeBytes,
			})
		case x2eTy.TyWithdrawTuringchainLog:
			var receiptTuringchainToEth x2eTy.ReceiptTuringchainToEth
			err := types.Decode(log.Log, &receiptTuringchainToEth)
			if err != nil {
				return nil, err
			}

			nb, err := x.GetLocalDB().Get(x2eTy.CalTokenSymbolTotalLockOrBurnAmount(receiptTuringchainToEth.IssuerDotSymbol, receiptTuringchainToEth.TokenContract, x2eTy.DirTuringchainToEth, ""))
			if err != nil && err != types.ErrNotFound {
				return nil, err
			}
			var now x2eTy.ReceiptQuerySymbolAssetsByTxType
			err = types.Decode(nb, &now)
			if err != nil {
				return nil, err
			}

			preAmount, _ := strconv.ParseFloat(x2eTy.TrimZeroAndDot(now.TotalAmount), 64)
			nowAmount, _ := strconv.ParseFloat(x2eTy.TrimZeroAndDot(receiptTuringchainToEth.Amount), 64)
			TokenAssetsByTxTypeBytes := types.Encode(&x2eTy.ReceiptQuerySymbolAssetsByTxType{
				TokenSymbol: receiptTuringchainToEth.IssuerDotSymbol,
				TxType:      "withdraw",
				TotalAmount: strconv.FormatFloat(preAmount+nowAmount, 'f', 4, 64),
				Direction:   2,
			})
			dbSet.KV = append(dbSet.KV, &types.KeyValue{
				Key:   x2eTy.CalTokenSymbolTotalLockOrBurnAmount(receiptTuringchainToEth.IssuerDotSymbol, receiptTuringchainToEth.TokenContract, x2eTy.DirTuringchainToEth, "withdraw"),
				Value: TokenAssetsByTxTypeBytes,
			})
		default:
			continue
		}
	}
	return dbSet, nil
}
