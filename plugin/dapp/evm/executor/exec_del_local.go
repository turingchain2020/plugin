// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"github.com/turingchain2020/turingchain/types"
	evmtypes "github.com/turingchain2020/plugin/plugin/dapp/evm/types"
)

// ExecDelLocal 处理区块回滚
func (evm *EVMExecutor) ExecDelLocal(tx *types.Transaction, receipt *types.ReceiptData, index int) (*types.LocalDBSet, error) {
	set, err := evm.DriverBase.ExecDelLocal(tx, receipt, index)
	if err != nil {
		return nil, err
	}
	if receipt.GetTy() != types.ExecOk {
		return set, nil
	}
	cfg := evm.GetAPI().GetConfig()
	if cfg.IsDappFork(evm.GetHeight(), "evm", evmtypes.ForkEVMState) {
		kvs, err := evm.DelRollbackKV(tx, []byte(evmtypes.ExecutorName))
		if err != nil {
			return nil, err
		}
		set.KV = kvs
	}
	return set, err
}
