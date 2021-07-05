package executor

import "github.com/turingchain2020/turingchain/types"

// CheckTx 本执行器不做任何校验
func (h *Echo) CheckTx(tx *types.Transaction, index int) error {
	return nil
}
