package types

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/vote/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/vote/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/vote/rpc"
	votetypes "github.com/turingchain2020/plugin/plugin/dapp/vote/types"
)

/*
 * 初始化dapp相关的组件
 */

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     votetypes.VoteX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.Cmd,
		RPC:      rpc.Init,
	})
}
