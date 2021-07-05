// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package token 创建token
package token

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	_ "github.com/turingchain2020/plugin/plugin/dapp/token/autotest" // register token autotest package
	"github.com/turingchain2020/plugin/plugin/dapp/token/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/token/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/token/rpc"
	"github.com/turingchain2020/plugin/plugin/dapp/token/types"
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.TokenX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.TokenCmd,
		RPC:      rpc.Init,
	})
}
