package ethtxs

// --------------------------------------------------------
//      Parser
//
//      Parses structs containing event information into
//      unsigned transactions for validators to sign, then
//      relays the data packets as transactions on the
//      turingchain Bridge.
// --------------------------------------------------------

import (
	"math/big"
	"strings"

	turingchainTypes "github.com/turingchain2020/turingchain/types"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/events"
	ebrelayerTypes "github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/types"
	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/types"
	"github.com/ethereum/go-ethereum/common"
)

// LogLockToEthBridgeClaim : parses and packages a LockEvent struct with a validator address in an EthBridgeClaim msg
func LogLockToEthBridgeClaim(event *events.LockEvent, ethereumChainID int64, bridgeBrankAddr string, decimal int64) (*ebrelayerTypes.EthBridgeClaim, error) {
	recipient := event.To
	if 0 == len(recipient) {
		return nil, ebrelayerTypes.ErrEmptyAddress
	}
	// Symbol formatted to lowercase
	symbol := strings.ToLower(event.Symbol)
	if symbol == "eth" && event.Token != common.HexToAddress("0x0000000000000000000000000000000000000000") {
		return nil, ebrelayerTypes.ErrAddress4Eth
	}

	witnessClaim := &ebrelayerTypes.EthBridgeClaim{}
	witnessClaim.EthereumChainID = ethereumChainID
	witnessClaim.BridgeBrankAddr = bridgeBrankAddr
	witnessClaim.Nonce = event.Nonce.Int64()
	witnessClaim.TokenAddr = event.Token.String()
	witnessClaim.Symbol = event.Symbol
	witnessClaim.EthereumSender = event.From.String()
	witnessClaim.TuringchainReceiver = string(recipient)

	if decimal > 8 {
		event.Value = event.Value.Quo(event.Value, big.NewInt(int64(types.MultiplySpecifyTimes(1, decimal-8))))
	} else {
		event.Value = event.Value.Mul(event.Value, big.NewInt(int64(types.MultiplySpecifyTimes(1, 8-decimal))))
	}
	witnessClaim.Amount = event.Value.String()

	witnessClaim.ClaimType = types.LockClaimType
	witnessClaim.ChainName = types.LockClaim
	witnessClaim.Decimal = decimal

	return witnessClaim, nil
}

//LogBurnToEthBridgeClaim ...
func LogBurnToEthBridgeClaim(event *events.BurnEvent, ethereumChainID int64, bridgeBrankAddr string, decimal int64) (*ebrelayerTypes.EthBridgeClaim, error) {
	recipient := event.TuringchainReceiver
	if 0 == len(recipient) {
		return nil, ebrelayerTypes.ErrEmptyAddress
	}

	witnessClaim := &ebrelayerTypes.EthBridgeClaim{}
	witnessClaim.EthereumChainID = ethereumChainID
	witnessClaim.BridgeBrankAddr = bridgeBrankAddr
	witnessClaim.Nonce = event.Nonce.Int64()
	witnessClaim.TokenAddr = event.Token.String()
	witnessClaim.Symbol = event.Symbol
	witnessClaim.EthereumSender = event.OwnerFrom.String()
	witnessClaim.TuringchainReceiver = string(recipient)
	witnessClaim.Amount = event.Amount.String()
	witnessClaim.ClaimType = types.BurnClaimType
	witnessClaim.ChainName = types.BurnClaim
	witnessClaim.Decimal = decimal

	return witnessClaim, nil
}

// ParseBurnLockTxReceipt : parses data from a Burn/Lock event witnessed on turingchain into a TuringchainMsg struct
func ParseBurnLockTxReceipt(claimType events.Event, receipt *turingchainTypes.ReceiptData) *events.TuringchainMsg {
	// Set up variables
	var turingchainSender []byte
	var ethereumReceiver, tokenContractAddress common.Address
	var symbol string
	var amount *big.Int

	// Iterate over attributes
	for _, log := range receipt.Logs {
		if log.Ty == types.TyTuringchainToEthLog || log.Ty == types.TyWithdrawTuringchainLog {
			txslog.Debug("ParseBurnLockTxReceipt", "value", string(log.Log))
			var turingchainToEth types.ReceiptTuringchainToEth
			err := turingchainTypes.Decode(log.Log, &turingchainToEth)
			if err != nil {
				return nil
			}
			turingchainSender = []byte(turingchainToEth.TuringchainSender)
			ethereumReceiver = common.HexToAddress(turingchainToEth.EthereumReceiver)
			tokenContractAddress = common.HexToAddress(turingchainToEth.TokenContract)
			symbol = turingchainToEth.IssuerDotSymbol
			turingchainToEth.Amount = types.TrimZeroAndDot(turingchainToEth.Amount)
			amount = big.NewInt(1)
			amount, _ = amount.SetString(turingchainToEth.Amount, 10)
			if turingchainToEth.Decimals > 8 {
				amount = amount.Mul(amount, big.NewInt(int64(types.MultiplySpecifyTimes(1, turingchainToEth.Decimals-8))))
			} else {
				amount = amount.Quo(amount, big.NewInt(int64(types.MultiplySpecifyTimes(1, 8-turingchainToEth.Decimals))))
			}

			txslog.Info("ParseBurnLockTxReceipt", "turingchainSender", turingchainSender, "ethereumReceiver", ethereumReceiver.String(), "tokenContractAddress", tokenContractAddress.String(), "symbol", symbol, "amount", amount.String())
			// Package the event data into a TuringchainMsg
			turingchainMsg := events.NewTuringchainMsg(claimType, turingchainSender, ethereumReceiver, symbol, amount, tokenContractAddress)
			return &turingchainMsg
		}
	}
	return nil
}

// TuringchainMsgToProphecyClaim : parses event data from a TuringchainMsg, packaging it as a ProphecyClaim
func TuringchainMsgToProphecyClaim(event events.TuringchainMsg) ProphecyClaim {
	claimType := event.ClaimType
	turingchainSender := event.TuringchainSender
	ethereumReceiver := event.EthereumReceiver
	tokenContractAddress := event.TokenContractAddress
	symbol := strings.ToLower(event.Symbol)
	amount := event.Amount

	prophecyClaim := ProphecyClaim{
		ClaimType:            claimType,
		TuringchainSender:        turingchainSender,
		EthereumReceiver:     ethereumReceiver,
		TokenContractAddress: tokenContractAddress,
		Symbol:               symbol,
		Amount:               amount,
	}

	return prophecyClaim
}
