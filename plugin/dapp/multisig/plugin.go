package multisig

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	_ "github.com/turingchain2020/plugin/plugin/dapp/multisig/autotest" //register auto test
	"github.com/turingchain2020/plugin/plugin/dapp/multisig/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/multisig/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/multisig/rpc"
	mty "github.com/turingchain2020/plugin/plugin/dapp/multisig/types"
	_ "github.com/turingchain2020/plugin/plugin/dapp/multisig/wallet" // register wallet package
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     mty.MultiSigX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.MultiSigCmd,
		RPC:      rpc.Init,
	})
}
