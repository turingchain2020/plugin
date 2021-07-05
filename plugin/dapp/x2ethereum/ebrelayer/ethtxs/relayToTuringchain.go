package ethtxs

// ------------------------------------------------------------
//	Relay : Builds and encodes EthBridgeClaim Msgs with the
//  	specified variables, before presenting the unsigned
//      transaction to validators for optional signing.
//      Once signed, the data packets are sent as transactions
//      on the turingchain Bridge.
// ------------------------------------------------------------

import (
	"github.com/turingchain2020/turingchain/common"
	turingchainCrypto "github.com/turingchain2020/turingchain/common/crypto"
	"github.com/turingchain2020/turingchain/rpc/jsonclient"
	rpctypes "github.com/turingchain2020/turingchain/rpc/types"
	turingchainTypes "github.com/turingchain2020/turingchain/types"
	ebrelayerTypes "github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/types"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/types"
)

// RelayLockToTuringchain : RelayLockToTuringchain applies validator's signature to an EthBridgeClaim message
//		containing information about an event on the Ethereum blockchain before relaying to the Bridge
func RelayLockToTuringchain(privateKey turingchainCrypto.PrivKey, claim *ebrelayerTypes.EthBridgeClaim, rpcURL string) (string, error) {
	var res string

	params := &types.Eth2Turingchain{
		EthereumChainID:       claim.EthereumChainID,
		BridgeContractAddress: claim.BridgeBrankAddr,
		Nonce:                 claim.Nonce,
		IssuerDotSymbol:       claim.Symbol,
		TokenContractAddress:  claim.TokenAddr,
		EthereumSender:        claim.EthereumSender,
		TuringchainReceiver:       claim.TuringchainReceiver,
		Amount:                claim.Amount,
		ClaimType:             int64(claim.ClaimType),
		Decimals:              claim.Decimal,
	}

	pm := rpctypes.CreateTxIn{
		Execer:     X2Eth,
		ActionName: types.NameEth2TuringchainAction,
		Payload:    turingchainTypes.MustPBToJSON(params),
	}
	ctx := jsonclient.NewRPCCtx(rpcURL, "Turingchain.CreateTransaction", pm, &res)
	_, _ = ctx.RunResult()

	data, err := common.FromHex(res)
	if err != nil {
		return "", err
	}
	var tx turingchainTypes.Transaction
	err = turingchainTypes.Decode(data, &tx)
	if err != nil {
		return "", err
	}

	if tx.Fee == 0 {
		tx.Fee, err = tx.GetRealFee(1e5)
		if err != nil {
			return "", err
		}
	}
	//构建交易，验证人validator用来向turingchain合约证明自己验证了该笔从以太坊向turingchain跨链转账的交易
	tx.Sign(turingchainTypes.SECP256K1, privateKey)

	txData := turingchainTypes.Encode(&tx)
	dataStr := common.ToHex(txData)
	pms := rpctypes.RawParm{
		Token: "TRC",
		Data:  dataStr,
	}
	var txhash string

	ctx = jsonclient.NewRPCCtx(rpcURL, "Turingchain.SendTransaction", pms, &txhash)
	_, err = ctx.RunResult()
	return txhash, err
}

//RelayBurnToTuringchain ...
func RelayBurnToTuringchain(privateKey turingchainCrypto.PrivKey, claim *ebrelayerTypes.EthBridgeClaim, rpcURL string) (string, error) {
	var res string

	params := &types.Eth2Turingchain{
		EthereumChainID:       claim.EthereumChainID,
		BridgeContractAddress: claim.BridgeBrankAddr,
		Nonce:                 claim.Nonce,
		IssuerDotSymbol:       claim.Symbol,
		TokenContractAddress:  claim.TokenAddr,
		EthereumSender:        claim.EthereumSender,
		TuringchainReceiver:       claim.TuringchainReceiver,
		Amount:                claim.Amount,
		ClaimType:             int64(claim.ClaimType),
		Decimals:              claim.Decimal,
	}

	pm := rpctypes.CreateTxIn{
		Execer:     X2Eth,
		ActionName: types.NameWithdrawEthAction,
		Payload:    turingchainTypes.MustPBToJSON(params),
	}
	ctx := jsonclient.NewRPCCtx(rpcURL, "Turingchain.CreateTransaction", pm, &res)
	_, _ = ctx.RunResult()

	data, err := common.FromHex(res)
	if err != nil {
		return "", err
	}
	var tx turingchainTypes.Transaction
	err = turingchainTypes.Decode(data, &tx)
	if err != nil {
		return "", err
	}

	if tx.Fee == 0 {
		tx.Fee, err = tx.GetRealFee(1e5)
		if err != nil {
			return "", err
		}
	}
	//构建交易，验证人validator用来向turingchain合约证明自己验证了该笔从以太坊向turingchain跨链转账的交易
	tx.Sign(turingchainTypes.SECP256K1, privateKey)

	txData := turingchainTypes.Encode(&tx)
	dataStr := common.ToHex(txData)
	pms := rpctypes.RawParm{
		Token: "TRC",
		Data:  dataStr,
	}
	var txhash string

	ctx = jsonclient.NewRPCCtx(rpcURL, "Turingchain.SendTransaction", pms, &txhash)
	_, err = ctx.RunResult()
	return txhash, err
}
