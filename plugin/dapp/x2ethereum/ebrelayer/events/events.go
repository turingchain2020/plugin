package events

import (
	log "github.com/turingchain2020/turingchain/common/log/log15"
)

// Event : enum containing supported contract events
type Event int

var eventsLog = log.New("module", "ethereum_relayer")

const (
	// Unsupported : unsupported Turingchain or Ethereum event
	Unsupported Event = iota
	// MsgBurn : Turingchain event 'TuringchainMsg' type MsgBurn
	MsgBurn
	// MsgLock :  Turingchain event 'TuringchainMsg' type MsgLock
	MsgLock
	// LogLock : Ethereum event 'LockEvent'
	LogLock
	// LogTuringchainTokenBurn : Ethereum event 'LogTuringchainTokenBurn' in contract turingchainBank
	LogTuringchainTokenBurn
	// LogNewProphecyClaim : Ethereum event 'NewProphecyClaimEvent'
	LogNewProphecyClaim
)

//const
const (
	ClaimTypeBurn = uint8(1)
	ClaimTypeLock = uint8(2)
)

// String : returns the event type as a string
func (d Event) String() string {
	return [...]string{"unknown-x2ethereum", "TuringchainToEthBurn", "TuringchainToEthLock", "LogLock", "LogTuringchainTokenBurn", "LogNewProphecyClaim"}[d]
}

// TuringchainMsgAttributeKey : enum containing supported attribute keys
type TuringchainMsgAttributeKey int

const (
	// UnsupportedAttributeKey : unsupported attribute key
	UnsupportedAttributeKey TuringchainMsgAttributeKey = iota
	// TuringchainSender : sender's address on Turingchain network
	TuringchainSender
	// EthereumReceiver : receiver's address on Ethereum network
	EthereumReceiver
	// Coin : coin type
	Coin
	// TokenContractAddress : coin's corresponding contract address deployed on the Ethereum network
	TokenContractAddress
)

// String : returns the event type as a string
func (d TuringchainMsgAttributeKey) String() string {
	return [...]string{"unsupported", "turingchain_sender", "ethereum_receiver", "amount", "token_contract_address"}[d]
}
