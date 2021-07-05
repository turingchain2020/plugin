// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package collateralize

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/collateralize/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/collateralize/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/collateralize/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.CollateralizeX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.CollateralizeCmd,
	})
}
