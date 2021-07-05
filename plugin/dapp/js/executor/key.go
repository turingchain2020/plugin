package executor

import (
	"github.com/turingchain2020/turingchain/types"
	ptypes "github.com/turingchain2020/plugin/plugin/dapp/js/types"
)

func calcAllPrefix(cfg *types.TuringchainConfig, name string) ([]byte, []byte) {
	execer := cfg.ExecName("user." + ptypes.JsX + "." + name)
	state := types.CalcStatePrefix([]byte(execer))
	local := types.CalcLocalPrefix([]byte(execer))
	return state, local
}

func calcCodeKey(name string) []byte {
	return append([]byte("mavl-"+ptypes.JsX+"-code-"), []byte(name)...)
}
