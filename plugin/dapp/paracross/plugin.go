// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paracross

import (
	"github.com/turingchain2020/turingchain/pluginmgr"
	_ "github.com/turingchain2020/plugin/plugin/crypto/bls"              // register bls package for ut usage
	_ "github.com/turingchain2020/plugin/plugin/dapp/paracross/autotest" // register autotest package
	"github.com/turingchain2020/plugin/plugin/dapp/paracross/commands"
	"github.com/turingchain2020/plugin/plugin/dapp/paracross/executor"
	"github.com/turingchain2020/plugin/plugin/dapp/paracross/rpc"
	"github.com/turingchain2020/plugin/plugin/dapp/paracross/types"
	_ "github.com/turingchain2020/plugin/plugin/dapp/paracross/wallet" // register wallet package
)

func init() {
	pluginmgr.Register(&pluginmgr.PluginBase{
		Name:     types.ParaX,
		ExecName: executor.GetName(),
		Exec:     executor.Init,
		Cmd:      commands.ParcCmd,
		RPC:      rpc.Init,
	})
}
