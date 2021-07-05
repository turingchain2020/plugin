// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package guess

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/guess/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/guess/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/guess/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.GuessX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.GuessCmd,
	})
}
