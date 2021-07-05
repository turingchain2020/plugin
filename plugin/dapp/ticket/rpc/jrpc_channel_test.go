// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc_test

import (
	"strings"
	"testing"

	"github.com/turingchain2020/turingchain/rpc/jsonclient"
	rpctypes "github.com/turingchain2020/turingchain/rpc/types"
	_ "github.com/turingchain2020/turingchain/system"
	"github.com/turingchain2020/turingchain/types"
	"github.com/turingchain2020/turingchain/util/testnode"
	_ "github.com/turingchain2020/plugin/plugin"
	"github.com/stretchr/testify/assert"
)

func TestJRPCChannel(t *testing.T) {
	cfg := types.NewTuringchainConfig(types.GetDefaultCfgstring())
	cfg.GetModuleConfig().Consensus.Name = "ticket"
	mocker := testnode.NewWithConfig(cfg, nil)
	mocker.Listen()
	defer mocker.Close()

	jrpcClient := mocker.GetJSONC()

	testCases := []struct {
		fn func(*testing.T, *jsonclient.JSONClient) error
	}{
		{fn: testCountTicketCmd},
		{fn: testCloseTicketCmd},
		{fn: testGetColdAddrByMinerCmd},
	}
	for index, testCase := range testCases {
		err := testCase.fn(t, jrpcClient)
		if err == nil {
			continue
		}
		assert.NotEqualf(t, err, types.ErrActionNotSupport, "test index %d", index)
		if strings.Contains(err.Error(), "rpc: can't find") {
			assert.FailNowf(t, err.Error(), "test index %d", index)
		}
		t.Log(err.Error())
	}
}

func testCountTicketCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var res int64
	return jrpc.Call("ticket.GetTicketCount", nil, &res)
}

func testCloseTicketCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var res types.ReplyHashes
	return jrpc.Call("ticket.CloseTickets", nil, &res)
}

func testGetColdAddrByMinerCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var rep interface{}
	var params rpctypes.Query4Jrpc
	req := &types.ReqString{}
	params.Execer = "ticket"
	params.FuncName = "MinerSourceList"
	params.Payload = types.MustPBToJSON(req)
	rep = &types.ReplyStrings{}
	return jrpc.Call("Turingchain.Query", params, rep)
}
