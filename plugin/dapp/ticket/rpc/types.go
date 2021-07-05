// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"github.com/turingchain2020/turingchain/rpc/types"
	ty "github.com/turingchain2020/plugin/plugin/dapp/ticket/types"
)

// Jrpc json rpc type
type Jrpc struct {
	cli *channelClient
}

// Grpc grpc type
type Grpc struct {
	*channelClient
}

type channelClient struct {
	types.ChannelClient
}

// Init initial
func Init(name string, s types.RPCServer) {
	cli := &channelClient{}
	grpc := &Grpc{channelClient: cli}
	cli.Init(name, s, &Jrpc{cli: cli}, grpc)
	ty.RegisterTicketServer(s.GRPC(), grpc)
}
