package wasm

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/wasm/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/wasm/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/wasm/rpc"
	"github.com/turingchain2020/plugin/plugin/dapp/wasm/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.WasmX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.Cmd,
		RPC:      rpc.Init,
	})
}
