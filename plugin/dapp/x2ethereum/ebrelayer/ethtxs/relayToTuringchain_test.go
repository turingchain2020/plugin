package ethtxs

import (
	"fmt"
	"testing"

	"github.com/turingchain2020/turingchain/client/mocks"
	turingchainCommon "github.com/turingchain2020/turingchain/common"
	_ "github.com/turingchain2020/turingchain/system"
	"github.com/turingchain2020/turingchain/system/crypto/secp256k1"
	turingchainTypes "github.com/turingchain2020/turingchain/types"
	"github.com/turingchain2020/turingchain/util/testnode"
	ebrelayerTypes "github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	chainTestCfg = turingchainTypes.NewTuringchainConfig(turingchainTypes.GetDefaultCfgstring())
)

func Test_RelayToTuringchain(t *testing.T) {
	var tx turingchainTypes.Transaction
	var ret turingchainTypes.Reply
	ret.IsOk = true

	mockapi := &mocks.QueueProtocolAPI{}
	// 这里对需要mock的方法打桩,Close是必须的，其它方法根据需要
	mockapi.On("Close").Return()
	mockapi.On("AddPushSubscribe", mock.Anything).Return(&ret, nil)
	mockapi.On("CreateTransaction", mock.Anything).Return(&tx, nil)
	mockapi.On("SendTx", mock.Anything).Return(&ret, nil)
	mockapi.On("SendTransaction", mock.Anything).Return(&ret, nil)
	mockapi.On("GetConfig", mock.Anything).Return(chainTestCfg, nil)

	mock33 := testnode.New("", mockapi)
	defer mock33.Close()
	rpcCfg := mock33.GetCfg().RPC
	// 这里必须设置监听端口，默认的是无效值
	rpcCfg.JrpcBindAddr = "127.0.0.1:9671"
	mock33.GetRPC().Listen()

	turingchainPrivateKeyStr := "0xd627968e445f2a41c92173225791bae1ba42126ae96c32f28f97ff8f226e5c68"
	var driver secp256k1.Driver
	privateKeySli, err := turingchainCommon.FromHex(turingchainPrivateKeyStr)
	require.Nil(t, err)

	priKey, err := driver.PrivKeyFromBytes(privateKeySli)
	require.Nil(t, err)

	claim := &ebrelayerTypes.EthBridgeClaim{}

	fmt.Println("======================= testRelayLockToTuringchain =======================")
	_, err = RelayLockToTuringchain(priKey, claim, "http://127.0.0.1:9671")
	require.Nil(t, err)

	fmt.Println("======================= testRelayBurnToTuringchain =======================")
	_, err = RelayBurnToTuringchain(priKey, claim, "http://127.0.0.1:9671")
	require.Nil(t, err)
}
