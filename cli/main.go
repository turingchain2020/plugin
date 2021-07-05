// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package main

import (
	_ "github.com/turingchain2020/turingchain/system"
	"github.com/turingchain2020/plugin/cli/buildflags"
	_ "github.com/turingchain2020/plugin/plugin"

	"github.com/turingchain2020/turingchain/util/cli"
)

func main() {
	if buildflags.RPCAddr == "" {
		buildflags.RPCAddr = "http://localhost:9671"
	}
	cli.Run(buildflags.RPCAddr, buildflags.ParaName, "")
}
