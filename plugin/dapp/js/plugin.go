package js

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/js/executor"
	ptypes "github.com/turingchain2020/plugin/plugin/dapp/js/types"

	// init auto test
	_ "github.com/turingchain2020/plugin/plugin/dapp/js/autotest"
	"github.com/turingchain2020/plugin/plugin/dapp/js/command"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     ptypes.JsX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      command.JavaScriptCmd,
		RPC:      nil,
	})
}
