// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package autonomy

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/autonomy/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/autonomy/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/autonomy/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.AutonomyX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.AutonomyCmd,
	})
}
