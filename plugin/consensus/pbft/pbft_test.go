// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pbft

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/turingchain2020/turingchain/blockchain"
	"github.com/turingchain2020/turingchain/common"
	"github.com/turingchain2020/turingchain/common/crypto"
	"github.com/turingchain2020/turingchain/common/limits"
	"github.com/turingchain2020/turingchain/common/log"
	"github.com/turingchain2020/turingchain/executor"
	"github.com/turingchain2020/turingchain/p2p"
	"github.com/turingchain2020/turingchain/queue"
	"github.com/turingchain2020/turingchain/store"
	cty "github.com/turingchain2020/turingchain/system/dapp/coins/types"
	"github.com/turingchain2020/turingchain/types"
	"github.com/turingchain2020/turingchain/wallet"

	_ "github.com/turingchain2020/turingchain/system"
	_ "github.com/turingchain2020/plugin/plugin/dapp/init"
	_ "github.com/turingchain2020/plugin/plugin/store/init"
)

var (
	random       *rand.Rand
	transactions []*types.Transaction
	txSize       = 1000
)

func init() {
	err := limits.SetLimits()
	if err != nil {
		panic(err)
	}
	random = rand.New(rand.NewSource(types.Now().UnixNano()))
	log.SetLogLevel("info")
}
func TestPbft(t *testing.T) {
	q, chain, p2pnet, s, exec, cs, wallet := initEnvPbft()
	defer chain.Close()
	defer p2pnet.Close()
	defer exec.Close()
	defer s.Close()
	defer cs.Close()
	defer q.Close()
	defer wallet.Close()
	time.Sleep(5 * time.Second)

	sendReplyList(q)
	clearTestData()
}

func initEnvPbft() (queue.Queue, *blockchain.BlockChain, *p2p.Manager, queue.Module, *executor.Executor, queue.Module, queue.Module) {
	flag.Parse()
	turingchainCfg := types.NewTuringchainConfig(types.ReadFile("turingchain.test.toml"))
	var q = queue.New("channel")
	q.SetConfig(turingchainCfg)
	cfg := turingchainCfg.GetModuleConfig()
	sub := turingchainCfg.GetSubConfig()

	chain := blockchain.New(turingchainCfg)
	chain.SetQueueClient(q.Client())
	exec := executor.New(turingchainCfg)
	exec.SetQueueClient(q.Client())
	turingchainCfg.SetMinFee(0)
	s := store.New(turingchainCfg)
	s.SetQueueClient(q.Client())
	cs := NewPbft(cfg.Consensus, sub.Consensus["pbft"])
	cs.SetQueueClient(q.Client())
	p2pnet := p2p.NewP2PMgr(turingchainCfg)
	p2pnet.SetQueueClient(q.Client())
	walletm := wallet.New(turingchainCfg)
	walletm.SetQueueClient(q.Client())

	return q, chain, p2pnet, s, exec, cs, walletm

}

func sendReplyList(q queue.Queue) {
	client := q.Client()
	client.Sub("mempool")
	var count int
	for msg := range client.Recv() {
		if msg.Ty == types.EventTxList {
			count++
			createReplyList(client.GetConfig(), "test"+strconv.Itoa(count))
			msg.Reply(client.NewMessage("consensus", types.EventReplyTxList,
				&types.ReplyTxList{Txs: transactions}))
			if count == 5 {
				time.Sleep(5 * time.Second)
				break
			}
		} else if msg.Ty == types.EventGetMempoolSize {
			msg.Reply(client.NewMessage("", 0, &types.MempoolSize{}))
		}
	}
}

func getprivkey(key string) crypto.PrivKey {
	cr, err := crypto.New(types.GetSignName("", types.SECP256K1))
	if err != nil {
		panic(err)
	}
	bkey, err := common.FromHex(key)
	if err != nil {
		panic(err)
	}
	priv, err := cr.PrivKeyFromBytes(bkey)
	if err != nil {
		panic(err)
	}
	return priv
}

func createReplyList(cfg *types.TuringchainConfig, account string) {

	var result []*types.Transaction
	for j := 0; j < txSize; j++ {
		//tx := &types.Transaction{}
		val := &cty.CoinsAction_Transfer{Transfer: &types.AssetsTransfer{Amount: 10}}
		action := &cty.CoinsAction{Value: val, Ty: cty.CoinsActionTransfer}
		tx := &types.Transaction{Execer: []byte("coins"), Payload: types.Encode(action), Fee: 0}
		tx.To = "14qViLJfdGaP4EeHnDyJbEGQysnCpwn1gZ"

		tx.Nonce = random.Int63()
		tx.ChainID = cfg.GetChainID()

		tx.Sign(types.SECP256K1, getprivkey("CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944"))
		result = append(result, tx)
	}
	//result = append(result, tx)
	transactions = result
}

func clearTestData() {
	err := os.RemoveAll("datadir")
	if err != nil {
		fmt.Println("delete datadir have a err:", err.Error())
	}
	err = os.RemoveAll("wallet")
	if err != nil {
		fmt.Println("delete wallet have a err:", err.Error())
	}
	fmt.Println("test data clear successfully!")
}
