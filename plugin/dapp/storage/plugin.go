package types

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/storage/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/storage/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/storage/rpc"
	storagetypes "github.com/turingchain2020/plugin/plugin/dapp/storage/types"
)

/*
 * 初始化dapp相关的组件
 */

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     storagetypes.StorageX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.Cmd,
		RPC:      rpc.Init,
	})
}
