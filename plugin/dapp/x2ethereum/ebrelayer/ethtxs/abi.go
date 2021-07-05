package ethtxs

import (
	"strings"

	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/ethcontract/generated"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

//const
const (
	BridgeBankABI    = "BridgeBankABI"
	TuringchainBankABI   = "TuringchainBankABI"
	TuringchainBridgeABI = "TuringchainBridgeABI"
	EthereumBankABI  = "EthereumBankABI"
)

//LoadABI ...
func LoadABI(contractName string) abi.ABI {
	var abiJSON string
	switch contractName {
	case BridgeBankABI:
		abiJSON = generated.BridgeBankABI
	case TuringchainBankABI:
		abiJSON = generated.TuringchainBankABI
	case TuringchainBridgeABI:
		abiJSON = generated.TuringchainBridgeABI
	case EthereumBankABI:
		abiJSON = generated.EthereumBankABI
	default:
		panic("No abi matched")
	}

	// Convert the raw abi into a usable format
	contractABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic(err)
	}

	return contractABI
}
