// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"github.com/turingchain2020/turingchain/rpc/types"
	vt "github.com/turingchain2020/plugin/plugin/dapp/valnode/types"
)

// Jrpc valnode jrpc interface
type Jrpc struct {
	cli *channelClient
}

// Grpc valnode Grpc interface
type Grpc struct {
	*channelClient
}

type channelClient struct {
	types.ChannelClient
}

// Init valnode rpc register
func Init(name string, s types.RPCServer) {
	cli := &channelClient{}
	grpc := &Grpc{channelClient: cli}
	cli.Init(name, s, &Jrpc{cli: cli}, grpc)

	vt.RegisterValnodeServer(s.GRPC(), grpc)
}
