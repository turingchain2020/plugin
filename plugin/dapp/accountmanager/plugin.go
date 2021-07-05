package types

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/accountmanager/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/accountmanager/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/accountmanager/rpc"
	accountmanagertypes "github.com/turingchain2020/plugin/plugin/dapp/accountmanager/types"
)

/*
 * 初始化dapp相关的组件
 */

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     accountmanagertypes.AccountmanagerX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.Cmd,
		RPC:      rpc.Init,
	})
}
