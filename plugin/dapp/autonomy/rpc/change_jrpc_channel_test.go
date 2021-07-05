// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc_test

import (
	"testing"

	"encoding/json"

	"github.com/turingchain2020/turingchain/rpc/jsonclient"
	rpctypes "github.com/turingchain2020/turingchain/rpc/types"
	_ "github.com/turingchain2020/turingchain/system"
	"github.com/turingchain2020/turingchain/types"
	_ "github.com/turingchain2020/plugin/plugin"
	auty "github.com/turingchain2020/plugin/plugin/dapp/autonomy/types"
)

func testPropChangeTxCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := &auty.ProposalChange{}
	payLoad, err := json.Marshal(params)
	if err != nil {
		return err
	}
	pm := &rpctypes.CreateTxIn{
		Execer:     chainTestCfg.ExecName(auty.AutonomyX),
		ActionName: "PropChange",
		Payload:    payLoad,
	}
	var res string
	return jrpc.Call("Turingchain.CreateTransaction", pm, &res)
}

func testRevokeProposalChangeTxCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := &auty.RevokeProposalChange{}
	payLoad, err := json.Marshal(params)
	if err != nil {
		return err
	}
	pm := &rpctypes.CreateTxIn{
		Execer:     chainTestCfg.ExecName(auty.AutonomyX),
		ActionName: "RvkPropChange",
		Payload:    payLoad,
	}
	var res string
	return jrpc.Call("Turingchain.CreateTransaction", pm, &res)
}

func testVoteProposalChangeTxCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := &auty.VoteProposalChange{}
	payLoad, err := json.Marshal(params)
	if err != nil {
		return err
	}
	pm := &rpctypes.CreateTxIn{
		Execer:     chainTestCfg.ExecName(auty.AutonomyX),
		ActionName: "VotePropChange",
		Payload:    payLoad,
	}
	var res string
	return jrpc.Call("Turingchain.CreateTransaction", pm, &res)
}

func testTerminateProposalChangeTxCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	params := &auty.TerminateProposalChange{}
	payLoad, err := json.Marshal(params)
	if err != nil {
		return err
	}
	pm := &rpctypes.CreateTxIn{
		Execer:     chainTestCfg.ExecName(auty.AutonomyX),
		ActionName: "TmintPropChange",
		Payload:    payLoad,
	}
	var res string
	return jrpc.Call("Turingchain.CreateTransaction", pm, &res)
}

func testGetProposalChangeCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var rep interface{}
	var params rpctypes.Query4Jrpc
	req := &types.ReqString{}
	params.FuncName = auty.GetProposalChange
	params.Payload = types.MustPBToJSON(req)
	rep = &auty.ReplyQueryProposalChange{}
	return jrpc.Call("Turingchain.Query", params, rep)
}

func testListProposalChangeCmd(t *testing.T, jrpc *jsonclient.JSONClient) error {
	var rep interface{}
	var params rpctypes.Query4Jrpc
	req := &auty.ReqQueryProposalChange{}
	params.FuncName = auty.ListProposalChange
	params.Payload = types.MustPBToJSON(req)
	rep = &auty.ReplyQueryProposalChange{}
	return jrpc.Call("Turingchain.Query", params, rep)
}
