// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raft

import (
	"fmt"
	"os"
	"testing"
	"time"

	//加载系统内置store, 不要依赖plugin
	_ "github.com/turingchain2020/turingchain/system/dapp/init"
	_ "github.com/turingchain2020/turingchain/system/mempool/init"
	_ "github.com/turingchain2020/turingchain/system/store/init"
	"github.com/turingchain2020/turingchain/util"
	"github.com/turingchain2020/turingchain/util/testnode"

	_ "github.com/turingchain2020/turingchain/system"
	_ "github.com/turingchain2020/plugin/plugin/dapp/init"
	_ "github.com/turingchain2020/plugin/plugin/store/init"
)

// 执行： go test -cover
func TestRaft(t *testing.T) {
	mock33 := testnode.New("turingchain.test.toml", nil)
	cfg := mock33.GetClient().GetConfig()
	defer mock33.Close()
	mock33.Listen()
	t.Log(mock33.GetGenesisAddress())
	time.Sleep(10 * time.Second)
	txs := util.GenNoneTxs(cfg, mock33.GetGenesisKey(), 10)
	for i := 0; i < len(txs); i++ {
		mock33.GetAPI().SendTx(txs[i])
	}
	mock33.WaitHeight(1)
	txs = util.GenNoneTxs(cfg, mock33.GetGenesisKey(), 10)
	for i := 0; i < len(txs); i++ {
		mock33.GetAPI().SendTx(txs[i])
	}
	mock33.WaitHeight(2)
	clearTestData()
}

func clearTestData() {
	err := os.RemoveAll("turingchain_raft-1")
	if err != nil {
		fmt.Println("delete turingchain_raft dir have a err:", err.Error())
	}
	fmt.Println("test data clear successfully!")
}
