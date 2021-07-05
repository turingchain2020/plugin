package x2ethereum

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/rpc"
	x2ethereumtypes "github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/types"
)

/*
 * 初始化dapp相关的组件
 */

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     x2ethereumtypes.X2ethereumX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.Cmd,
		RPC:      rpc.Init,
	})
}
