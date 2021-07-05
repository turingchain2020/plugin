// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pokerbull

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/pokerbull/cmd"
	"github.com/turingchain2020/plugin/plugin/dapp/pokerbull/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/pokerbull/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.PokerBullX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      cmd.PokerBullCmd,
	})
}
