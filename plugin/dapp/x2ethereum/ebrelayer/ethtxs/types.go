package ethtxs

import (
	"math/big"

	"github.com/turingchain2020/plugin/plugin/dapp/x2ethereum/ebrelayer/events"
	"github.com/ethereum/go-ethereum/common"
)

//const ...
const (
	X2Eth      = "x2ethereum"
	BurnAction = "TuringchainToEthBurn"
	LockAction = "TuringchainToEthLock"
)

// OracleClaim : contains data required to make an OracleClaim
type OracleClaim struct {
	ProphecyID *big.Int
	Message    [32]byte
	Signature  []byte
}

// ProphecyClaim : contains data required to make an ProphecyClaim
type ProphecyClaim struct {
	ClaimType            events.Event
	TuringchainSender        []byte
	EthereumReceiver     common.Address
	TokenContractAddress common.Address
	Symbol               string
	Amount               *big.Int
}
