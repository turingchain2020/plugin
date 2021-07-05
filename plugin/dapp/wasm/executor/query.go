package executor

import (
	"github.com/turingchain2020/turingchain/types"
	types2 "github.com/turingchain2020/plugin/plugin/dapp/wasm/types"
)

func (w *Wasm) Query_Check(query *types2.QueryCheckContract) (types.Message, error) {
	if query == nil {
		return nil, types.ErrInvalidParam
	}
	return &types.Reply{IsOk: w.contractExist(query.Name)}, nil
}
