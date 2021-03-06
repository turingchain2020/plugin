// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gossip

import (
	"time"
)

// time limit for timeout
var (
	UpdateState                 = 2 * time.Second
	PingTimeout                 = 14 * time.Second
	DefaultSendTimeout          = 10 * time.Second
	DialTimeout                 = 5 * time.Second
	mapUpdateInterval           = 45 * time.Hour
	StreamPingTimeout           = 20 * time.Second
	MonitorPeerInfoInterval     = 10 * time.Second
	MonitorPeerNumInterval      = 30 * time.Second
	MonitorReBalanceInterval    = 15 * time.Minute
	GetAddrFromAddrBookInterval = 5 * time.Second
	GetAddrFromOnlineInterval   = 5 * time.Second
	GetAddrFromGitHubInterval   = 5 * time.Minute
	CheckActivePeersInterVal    = 5 * time.Second
	CheckBlackListInterVal      = 30 * time.Second
	CheckCfgSeedsInterVal       = 1 * time.Minute
)

const (
	msgTx           = 1
	msgBlock        = 2
	tryMapPortTimes = 20
	maxSamIPNum     = 20
)

const (
	//defalutNatPort  = 23802
	maxOutBoundNum  = 25
	stableBoundNum  = 15
	maxAttemps      = 5
	protocol        = "tcp"
	externalPortTag = "externalport"
)

const (
	nodeNetwork = 1
	nodeGetUTXO = 2
	nodeBloom   = 4
)

const (
	// Service service number
	Service int32 = nodeBloom + nodeNetwork + nodeGetUTXO
)

// leveldb 中p2p privkey,addrkey
const (
	addrkeyTag = "addrs"
	privKeyTag = "privkey"
)

//TTL
const (
	DefaultLtTxBroadCastTTL  = 3
	DefaultMaxTxBroadCastTTL = 25
	// 100KB
	defaultMinLtBlockSize = 100
)

// P2pCacheTxSize p2pcache size of transaction
const (
	PeerAddrCacheNum = 1000
	//接收的交易哈希过滤缓存设为mempool最大接收交易量
	TxRecvFilterCacheNum = 10240
	BlockFilterCacheNum  = 50
	//发送过滤主要用于发送时冗余检测, 发送完即可以被删除, 维护较小缓存数
	TxSendFilterCacheNum  = 500
	BlockCacheNum         = 10
	MaxBlockCacheByteSize = 100 * 1024 * 1024
)

// TestNetSeeds test seeds of net
var TestNetSeeds = []string{
	"47.97.223.101:9675",
}

// MainNetSeeds built-in list of seed
var MainNetSeeds = []string{
	"39.105.112.116:9675",
	"101.200.37.68:9675",
}
