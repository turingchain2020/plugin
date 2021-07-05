// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package echo

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	"github.com/turingchain2020/plugin/plugin/dapp/echo/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/echo/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/echo/rpc"
	echotypes "github.com/turingchain2020/plugin/plugin/dapp/echo/types/echo"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     echotypes.EchoX,
		ExecName: echotypes.EchoX,
		Exec:     executor.Init,
		Cmd:      commands.EchoCmd,
		RPC:      rpc.Init,
	})
}
