package events

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// TuringchainMsg : contains data from MsgBurn and MsgLock events
type TuringchainMsg struct {
	ClaimType            Event
	TuringchainSender        []byte
	EthereumReceiver     common.Address
	TokenContractAddress common.Address
	Symbol               string
	Amount               *big.Int
}

// NewTuringchainMsg : creates a new TuringchainMsg
func NewTuringchainMsg(
	claimType Event,
	turingchainSender []byte,
	ethereumReceiver common.Address,
	symbol string,
	amount *big.Int,
	tokenContractAddress common.Address,
) TuringchainMsg {
	// Package data into a TuringchainMsg
	turingchainMsg := TuringchainMsg{
		ClaimType:            claimType,
		TuringchainSender:        turingchainSender,
		EthereumReceiver:     ethereumReceiver,
		Symbol:               symbol,
		Amount:               amount,
		TokenContractAddress: tokenContractAddress,
	}

	return turingchainMsg
}
