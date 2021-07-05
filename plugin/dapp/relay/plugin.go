// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package relay

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/relay/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/relay/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/relay/rpc"
	"github.com/turingchain2020/plugin/plugin/dapp/relay/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.RelayX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.RelayCmd,
		RPC:      rpc.Init,
	})
}
