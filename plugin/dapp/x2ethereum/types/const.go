package types

//key
var (
	ProphecyKey                         = []byte("prefix_for_Prophecy")
	Eth2TuringchainKey                      = []byte("prefix_for_Eth2Turingchain")
	WithdrawEthKey                      = []byte("prefix_for_WithdrawEth")
	TuringchainToEthKey                     = []byte("prefix_for_TuringchainToEth")
	WithdrawTuringchainKey                  = []byte("prefix_for_WithdrawTuringchain")
	LastTotalPowerKey                   = []byte("prefix_for_LastTotalPower")
	ValidatorMapsKey                    = []byte("prefix_for_ValidatorMaps")
	ConsensusThresholdKey               = []byte("prefix_for_ConsensusThreshold")
	TokenSymbolTotalLockOrBurnAmountKey = []byte("prefix_for_TokenSymbolTotalLockOrBurnAmount-")
	TokenSymbolToTokenAddressKey        = []byte("prefix_for_TokenSymbolToTokenAddress-")
)

// log for x2ethereum
// log类型id值
const (
	TyUnknownLog = iota + 100
	TyEth2TuringchainLog
	TyWithdrawEthLog
	TyWithdrawTuringchainLog
	TyTuringchainToEthLog
	TyAddValidatorLog
	TyRemoveValidatorLog
	TyModifyPowerLog
	TySetConsensusThresholdLog
	TyProphecyLog
	TyTransferLog
	TyTransferToExecLog
	TyWithdrawFromExecLog
)

// action类型id和name，这些常量可以自定义修改
const (
	TyUnknowAction = iota + 100
	TyEth2TuringchainAction
	TyWithdrawEthAction
	TyWithdrawTuringchainAction
	TyTuringchainToEthAction
	TyAddValidatorAction
	TyRemoveValidatorAction
	TyModifyPowerAction
	TySetConsensusThresholdAction
	TyTransferAction
	TyTransferToExecAction
	TyWithdrawFromExecAction

	NameEth2TuringchainAction           = "Eth2TuringchainLock"
	NameWithdrawEthAction           = "Eth2TuringchainBurn"
	NameWithdrawTuringchainAction       = "TuringchainToEthBurn"
	NameTuringchainToEthAction          = "TuringchainToEthLock"
	NameAddValidatorAction          = "AddValidator"
	NameRemoveValidatorAction       = "RemoveValidator"
	NameModifyPowerAction           = "ModifyPower"
	NameSetConsensusThresholdAction = "SetConsensusThreshold"
	NameTransferAction              = "Transfer"
	NameTransferToExecAction        = "TransferToExec"
	NameWithdrawFromExecAction      = "WithdrawFromExec"
)

//DefaultConsensusNeeded ...
const DefaultConsensusNeeded = int64(70)

//direct ...
const (
	DirEth2Turingchain  = "eth2turingchain"
	DirTuringchainToEth = "turingchaintoeth"
	LockClaim       = "lock"
	BurnClaim       = "burn"
)

//DirectionType type
var DirectionType = [3]string{"", DirEth2Turingchain, DirTuringchainToEth}

// query function name
const (
	FuncQueryEthProphecy               = "GetEthProphecy"
	FuncQueryValidators                = "GetValidators"
	FuncQueryTotalPower                = "GetTotalPower"
	FuncQueryConsensusThreshold        = "GetConsensusThreshold"
	FuncQuerySymbolTotalAmountByTxType = "GetSymbolTotalAmountByTxType"
	FuncQueryRelayerBalance            = "GetRelayerBalance"
)

//lock type
const (
	LockClaimType = int32(1)
	BurnClaimType = int32(2)
)
