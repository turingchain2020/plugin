// Code generated by protoc-gen-go.
// source: relayer.proto
// DO NOT EDIT!

package types

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

// 以太坊账户信息
// 	 privkey : 账户地址对应的私钥
// 	 addr :账户地址
type Account4Relayer struct {
	Privkey []byte `protobuf:"bytes,1,opt,name=privkey,proto3" json:"privkey,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
}

func (m *Account4Relayer) Reset()         { *m = Account4Relayer{} }
func (m *Account4Relayer) String() string { return proto.CompactTextString(m) }
func (*Account4Relayer) ProtoMessage()    {}

type ValidatorAddr4EthRelayer struct {
	TuringchainValidator string `protobuf:"bytes,1,opt,name=turingchainValidator" json:"turingchainValidator,omitempty"`
}

func (m *ValidatorAddr4EthRelayer) Reset()         { *m = ValidatorAddr4EthRelayer{} }
func (m *ValidatorAddr4EthRelayer) String() string { return proto.CompactTextString(m) }
func (*ValidatorAddr4EthRelayer) ProtoMessage()    {}

type Txhashes struct {
	Txhash []string `protobuf:"bytes,1,rep,name=txhash" json:"txhash,omitempty"`
}

func (m *Txhashes) Reset()         { *m = Txhashes{} }
func (m *Txhashes) String() string { return proto.CompactTextString(m) }
func (*Txhashes) ProtoMessage()    {}

type ReqChangePasswd struct {
	OldPassphase string `protobuf:"bytes,1,opt,name=oldPassphase" json:"oldPassphase,omitempty"`
	NewPassphase string `protobuf:"bytes,2,opt,name=newPassphase" json:"newPassphase,omitempty"`
}

func (m *ReqChangePasswd) Reset()         { *m = ReqChangePasswd{} }
func (m *ReqChangePasswd) String() string { return proto.CompactTextString(m) }
func (*ReqChangePasswd) ProtoMessage()    {}

type ReqSetPasswd struct {
	Passphase string `protobuf:"bytes,1,opt" json:"Passphase,omitempty"`
}

func (m *ReqSetPasswd) Reset()         { *m = ReqSetPasswd{} }
func (m *ReqSetPasswd) String() string { return proto.CompactTextString(m) }
func (*ReqSetPasswd) ProtoMessage()    {}

type Account4Show struct {
	Privkey string `protobuf:"bytes,1,opt,name=privkey" json:"privkey,omitempty"`
	Addr    string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
}

func (m *Account4Show) Reset()         { *m = Account4Show{} }
func (m *Account4Show) String() string { return proto.CompactTextString(m) }
func (*Account4Show) ProtoMessage()    {}

type AssetType struct {
	Chain         string `protobuf:"bytes,1,opt,name=chain" json:"chain,omitempty"`
	IssueContract string `protobuf:"bytes,2,opt,name=issueContract" json:"issueContract,omitempty"`
	Symbol        string `protobuf:"bytes,3,opt,name=symbol" json:"symbol,omitempty"`
}

func (m *AssetType) Reset()         { *m = AssetType{} }
func (m *AssetType) String() string { return proto.CompactTextString(m) }
func (*AssetType) ProtoMessage()    {}

type EthBridgeClaim struct {
	EthereumChainID int64  `protobuf:"varint,1,opt,name=ethereumChainID" json:"ethereumChainID,omitempty"`
	BridgeBrankAddr string `protobuf:"bytes,2,opt,name=bridgeBrankAddr" json:"bridgeBrankAddr,omitempty"`
	Nonce           int64  `protobuf:"varint,3,opt,name=nonce" json:"nonce,omitempty"`
	TokenAddr       string `protobuf:"bytes,4,opt,name=tokenAddr" json:"tokenAddr,omitempty"`
	Symbol          string `protobuf:"bytes,5,opt,name=symbol" json:"symbol,omitempty"`
	EthereumSender  string `protobuf:"bytes,6,opt,name=ethereumSender" json:"ethereumSender,omitempty"`
	TuringchainReceiver string `protobuf:"bytes,7,opt,name=turingchainReceiver" json:"turingchainReceiver,omitempty"`
	Amount          string `protobuf:"bytes,9,opt,name=amount" json:"amount,omitempty"`
	ClaimType       int32  `protobuf:"varint,10,opt,name=claimType" json:"claimType,omitempty"`
	ChainName       string `protobuf:"bytes,11,opt,name=chainName" json:"chainName,omitempty"`
	Decimal         int64  `protobuf:"varint,12,opt,name=decimal" json:"decimal,omitempty"`
}

func (m *EthBridgeClaim) Reset()         { *m = EthBridgeClaim{} }
func (m *EthBridgeClaim) String() string { return proto.CompactTextString(m) }
func (*EthBridgeClaim) ProtoMessage()    {}

type ImportKeyReq struct {
	PrivateKey string `protobuf:"bytes,1,opt,name=privateKey" json:"privateKey,omitempty"`
}

func (m *ImportKeyReq) Reset()         { *m = ImportKeyReq{} }
func (m *ImportKeyReq) String() string { return proto.CompactTextString(m) }
func (*ImportKeyReq) ProtoMessage()    {}

type RelayerRunStatus struct {
	Status  int32  `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details" json:"details,omitempty"`
}

func (m *RelayerRunStatus) Reset()         { *m = RelayerRunStatus{} }
func (m *RelayerRunStatus) String() string { return proto.CompactTextString(m) }
func (*RelayerRunStatus) ProtoMessage()    {}

type NewProphecyClaim struct {
	ClaimType     uint32 `protobuf:"varint,1,opt,name=claimType" json:"claimType,omitempty"`
	TuringchainSender string `protobuf:"bytes,2,opt,name=turingchainSender" json:"turingchainSender,omitempty"`
	TokenAddr     string `protobuf:"bytes,3,opt,name=tokenAddr" json:"tokenAddr,omitempty"`
	Symbol        string `protobuf:"bytes,4,opt,name=symbol" json:"symbol,omitempty"`
	EthReceiver   string `protobuf:"bytes,5,opt,name=ethReceiver" json:"ethReceiver,omitempty"`
	Amount        string `protobuf:"bytes,6,opt,name=amount" json:"amount,omitempty"`
	TxHash        string `protobuf:"bytes,7,opt,name=txHash" json:"txHash,omitempty"`
}

func (m *NewProphecyClaim) Reset()         { *m = NewProphecyClaim{} }
func (m *NewProphecyClaim) String() string { return proto.CompactTextString(m) }
func (*NewProphecyClaim) ProtoMessage()    {}

type BalanceAddr struct {
	Owner     string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	TokenAddr string `protobuf:"bytes,2,opt,name=tokenAddr" json:"tokenAddr,omitempty"`
}

func (m *BalanceAddr) Reset()         { *m = BalanceAddr{} }
func (m *BalanceAddr) String() string { return proto.CompactTextString(m) }
func (*BalanceAddr) ProtoMessage()    {}

type MintToken struct {
	Owner     string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	TokenAddr string `protobuf:"bytes,2,opt,name=tokenAddr" json:"tokenAddr,omitempty"`
	Amount    string `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *MintToken) Reset()         { *m = MintToken{} }
func (m *MintToken) String() string { return proto.CompactTextString(m) }
func (*MintToken) ProtoMessage()    {}

type ApproveAllowance struct {
	OwnerKey  string `protobuf:"bytes,1,opt,name=ownerKey" json:"ownerKey,omitempty"`
	TokenAddr string `protobuf:"bytes,2,opt,name=tokenAddr" json:"tokenAddr,omitempty"`
	Amount    string `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *ApproveAllowance) Reset()         { *m = ApproveAllowance{} }
func (m *ApproveAllowance) String() string { return proto.CompactTextString(m) }
func (*ApproveAllowance) ProtoMessage()    {}

type LockEthErc20 struct {
	OwnerKey  string `protobuf:"bytes,1,opt,name=ownerKey" json:"ownerKey,omitempty"`
	TokenAddr string `protobuf:"bytes,2,opt,name=tokenAddr" json:"tokenAddr,omitempty"`
	Amount    string `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
	// 将lock住的资产跨链转移到turingchain的该账户名下
	TuringchainReceiver string `protobuf:"bytes,4,opt,name=turingchainReceiver" json:"turingchainReceiver,omitempty"`
}

func (m *LockEthErc20) Reset()         { *m = LockEthErc20{} }
func (m *LockEthErc20) String() string { return proto.CompactTextString(m) }
func (*LockEthErc20) ProtoMessage()    {}

type ReplyAddr struct {
	IsOK bool   `protobuf:"varint,1,opt,name=isOK" json:"isOK,omitempty"`
	Addr string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
}

func (m *ReplyAddr) Reset()         { *m = ReplyAddr{} }
func (m *ReplyAddr) String() string { return proto.CompactTextString(m) }
func (*ReplyAddr) ProtoMessage()    {}

type ReplyBalance struct {
	IsOK    bool   `protobuf:"varint,1,opt,name=isOK" json:"isOK,omitempty"`
	Balance string `protobuf:"bytes,2,opt,name=balance" json:"balance,omitempty"`
}

func (m *ReplyBalance) Reset()         { *m = ReplyBalance{} }
func (m *ReplyBalance) String() string { return proto.CompactTextString(m) }
func (*ReplyBalance) ProtoMessage()    {}

type Burn struct {
	OwnerKey        string `protobuf:"bytes,1,opt,name=ownerKey" json:"ownerKey,omitempty"`
	TokenAddr       string `protobuf:"bytes,2,opt,name=tokenAddr" json:"tokenAddr,omitempty"`
	Amount          string `protobuf:"bytes,3,opt,name=amount" json:"amount,omitempty"`
	TuringchainReceiver string `protobuf:"bytes,4,opt,name=turingchainReceiver" json:"turingchainReceiver,omitempty"`
}

func (m *Burn) Reset()         { *m = Burn{} }
func (m *Burn) String() string { return proto.CompactTextString(m) }
func (*Burn) ProtoMessage()    {}

type StaticsRequest struct {
	Owner     string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	TokenAddr string `protobuf:"bytes,2,opt,name=tokenAddr" json:"tokenAddr,omitempty"`
}

func (m *StaticsRequest) Reset()         { *m = StaticsRequest{} }
func (m *StaticsRequest) String() string { return proto.CompactTextString(m) }
func (*StaticsRequest) ProtoMessage()    {}

type StaticsAll struct {
}

func (m *StaticsAll) Reset()         { *m = StaticsAll{} }
func (m *StaticsAll) String() string { return proto.CompactTextString(m) }
func (*StaticsAll) ProtoMessage()    {}

type StaticsSingle struct {
}

func (m *StaticsSingle) Reset()         { *m = StaticsSingle{} }
func (m *StaticsSingle) String() string { return proto.CompactTextString(m) }
func (*StaticsSingle) ProtoMessage()    {}

type StaticsLockResponse struct {
	All    *StaticsLock       `protobuf:"bytes,1,opt,name=all" json:"all,omitempty"`
	Single *StaticsLockSingle `protobuf:"bytes,2,opt,name=single" json:"single,omitempty"`
}

func (m *StaticsLockResponse) Reset()         { *m = StaticsLockResponse{} }
func (m *StaticsLockResponse) String() string { return proto.CompactTextString(m) }
func (*StaticsLockResponse) ProtoMessage()    {}

func (m *StaticsLockResponse) GetAll() *StaticsLock {
	if m != nil {
		return m.All
	}
	return nil
}

func (m *StaticsLockResponse) GetSingle() *StaticsLockSingle {
	if m != nil {
		return m.Single
	}
	return nil
}

type StaticsResponse struct {
}

func (m *StaticsResponse) Reset()         { *m = StaticsResponse{} }
func (m *StaticsResponse) String() string { return proto.CompactTextString(m) }
func (*StaticsResponse) ProtoMessage()    {}

type StaticsLock struct {
	Balance string `protobuf:"bytes,1,opt,name=balance" json:"balance,omitempty"`
}

func (m *StaticsLock) Reset()         { *m = StaticsLock{} }
func (m *StaticsLock) String() string { return proto.CompactTextString(m) }
func (*StaticsLock) ProtoMessage()    {}

type StaticsDeposit struct {
	Supply string `protobuf:"bytes,1,opt,name=supply" json:"supply,omitempty"`
}

func (m *StaticsDeposit) Reset()         { *m = StaticsDeposit{} }
func (m *StaticsDeposit) String() string { return proto.CompactTextString(m) }
func (*StaticsDeposit) ProtoMessage()    {}

type StaticsLockSingle struct {
	TotalLockedAccumated int64   `protobuf:"varint,1,opt,name=totalLockedAccumated" json:"totalLockedAccumated,omitempty"`
	Locked               []int64 `protobuf:"varint,2,rep,name=locked" json:"locked,omitempty"`
}

func (m *StaticsLockSingle) Reset()         { *m = StaticsLockSingle{} }
func (m *StaticsLockSingle) String() string { return proto.CompactTextString(m) }
func (*StaticsLockSingle) ProtoMessage()    {}

type TransferToken struct {
	TokenAddr string `protobuf:"bytes,1,opt,name=tokenAddr" json:"tokenAddr,omitempty"`
	FromKey   string `protobuf:"bytes,2,opt,name=fromKey" json:"fromKey,omitempty"`
	ToAddr    string `protobuf:"bytes,3,opt,name=toAddr" json:"toAddr,omitempty"`
	Amount    string `protobuf:"bytes,4,opt,name=amount" json:"amount,omitempty"`
}

func (m *TransferToken) Reset()         { *m = TransferToken{} }
func (m *TransferToken) String() string { return proto.CompactTextString(m) }
func (*TransferToken) ProtoMessage()    {}

type Uint64 struct {
	Data uint64 `protobuf:"varint,1,opt,name=data" json:"data,omitempty"`
}

func (m *Uint64) Reset()         { *m = Uint64{} }
func (m *Uint64) String() string { return proto.CompactTextString(m) }
func (*Uint64) ProtoMessage()    {}

type TokenStatics struct {
	TokenAddr string `protobuf:"bytes,1,opt,name=tokenAddr" json:"tokenAddr,omitempty"`
}

func (m *TokenStatics) Reset()         { *m = TokenStatics{} }
func (m *TokenStatics) String() string { return proto.CompactTextString(m) }
func (*TokenStatics) ProtoMessage()    {}

type EventLogIndex struct {
	Height uint64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	Index  uint32 `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
}

func (m *EventLogIndex) Reset()         { *m = EventLogIndex{} }
func (m *EventLogIndex) String() string { return proto.CompactTextString(m) }
func (*EventLogIndex) ProtoMessage()    {}

type EthTxStatus struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Txhash string `protobuf:"bytes,2,opt,name=txhash" json:"txhash,omitempty"`
}

func (m *EthTxStatus) Reset()         { *m = EthTxStatus{} }
func (m *EthTxStatus) String() string { return proto.CompactTextString(m) }
func (*EthTxStatus) ProtoMessage()    {}

func init() {
}
