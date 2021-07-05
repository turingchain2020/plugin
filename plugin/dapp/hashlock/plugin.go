// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashlock

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/hashlock/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/hashlock/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/hashlock/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.HashlockX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.HashlockCmd,
		RPC:      nil,
	})
}
