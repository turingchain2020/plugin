// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package retrieve

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/retrieve/cmd"
	"github.com/turingchain2020/plugin/plugin/dapp/retrieve/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/retrieve/rpc"
	"github.com/turingchain2020/plugin/plugin/dapp/retrieve/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.RetrieveX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      cmd.RetrieveCmd,
		RPC:      rpc.Init,
	})
}
