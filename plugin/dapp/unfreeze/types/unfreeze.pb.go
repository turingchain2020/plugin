// Code generated by protoc-gen-go. DO NOT EDIT.
// source: unfreeze.proto

package types

import (
	context "context"
	fmt "fmt"
	math "math"

	types "github.com/turingchain2020/turingchain/types"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Unfreeze struct {
	//解冻交易ID（唯一识别码）
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID,proto3" json:"unfreezeID,omitempty"`
	//开始时间
	StartTime int64 `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	//币种
	AssetExec   string `protobuf:"bytes,3,opt,name=assetExec,proto3" json:"assetExec,omitempty"`
	AssetSymbol string `protobuf:"bytes,4,opt,name=assetSymbol,proto3" json:"assetSymbol,omitempty"`
	//冻结总额
	TotalCount int64 `protobuf:"varint,5,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	//发币人地址
	Initiator string `protobuf:"bytes,6,opt,name=initiator,proto3" json:"initiator,omitempty"`
	//收币人地址
	Beneficiary string `protobuf:"bytes,7,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	//解冻剩余币数
	Remaining int64 `protobuf:"varint,8,opt,name=remaining,proto3" json:"remaining,omitempty"`
	//解冻方式（百分比；固额）
	Means string `protobuf:"bytes,9,opt,name=means,proto3" json:"means,omitempty"`
	// Types that are valid to be assigned to MeansOpt:
	//	*Unfreeze_FixAmount
	//	*Unfreeze_LeftProportion
	MeansOpt             isUnfreeze_MeansOpt `protobuf_oneof:"meansOpt"`
	Terminated           bool                `protobuf:"varint,12,opt,name=terminated,proto3" json:"terminated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Unfreeze) Reset()         { *m = Unfreeze{} }
func (m *Unfreeze) String() string { return proto.CompactTextString(m) }
func (*Unfreeze) ProtoMessage()    {}
func (*Unfreeze) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{0}
}

func (m *Unfreeze) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unfreeze.Unmarshal(m, b)
}
func (m *Unfreeze) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unfreeze.Marshal(b, m, deterministic)
}
func (m *Unfreeze) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unfreeze.Merge(m, src)
}
func (m *Unfreeze) XXX_Size() int {
	return xxx_messageInfo_Unfreeze.Size(m)
}
func (m *Unfreeze) XXX_DiscardUnknown() {
	xxx_messageInfo_Unfreeze.DiscardUnknown(m)
}

var xxx_messageInfo_Unfreeze proto.InternalMessageInfo

func (m *Unfreeze) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *Unfreeze) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Unfreeze) GetAssetExec() string {
	if m != nil {
		return m.AssetExec
	}
	return ""
}

func (m *Unfreeze) GetAssetSymbol() string {
	if m != nil {
		return m.AssetSymbol
	}
	return ""
}

func (m *Unfreeze) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *Unfreeze) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *Unfreeze) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *Unfreeze) GetRemaining() int64 {
	if m != nil {
		return m.Remaining
	}
	return 0
}

func (m *Unfreeze) GetMeans() string {
	if m != nil {
		return m.Means
	}
	return ""
}

type isUnfreeze_MeansOpt interface {
	isUnfreeze_MeansOpt()
}

type Unfreeze_FixAmount struct {
	FixAmount *FixAmount `protobuf:"bytes,10,opt,name=fixAmount,proto3,oneof"`
}

type Unfreeze_LeftProportion struct {
	LeftProportion *LeftProportion `protobuf:"bytes,11,opt,name=leftProportion,proto3,oneof"`
}

func (*Unfreeze_FixAmount) isUnfreeze_MeansOpt() {}

func (*Unfreeze_LeftProportion) isUnfreeze_MeansOpt() {}

func (m *Unfreeze) GetMeansOpt() isUnfreeze_MeansOpt {
	if m != nil {
		return m.MeansOpt
	}
	return nil
}

func (m *Unfreeze) GetFixAmount() *FixAmount {
	if x, ok := m.GetMeansOpt().(*Unfreeze_FixAmount); ok {
		return x.FixAmount
	}
	return nil
}

func (m *Unfreeze) GetLeftProportion() *LeftProportion {
	if x, ok := m.GetMeansOpt().(*Unfreeze_LeftProportion); ok {
		return x.LeftProportion
	}
	return nil
}

func (m *Unfreeze) GetTerminated() bool {
	if m != nil {
		return m.Terminated
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Unfreeze) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Unfreeze_FixAmount)(nil),
		(*Unfreeze_LeftProportion)(nil),
	}
}

// 按时间固定额度解冻
type FixAmount struct {
	Period               int64    `protobuf:"varint,1,opt,name=period,proto3" json:"period,omitempty"`
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FixAmount) Reset()         { *m = FixAmount{} }
func (m *FixAmount) String() string { return proto.CompactTextString(m) }
func (*FixAmount) ProtoMessage()    {}
func (*FixAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{1}
}

func (m *FixAmount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixAmount.Unmarshal(m, b)
}
func (m *FixAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixAmount.Marshal(b, m, deterministic)
}
func (m *FixAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixAmount.Merge(m, src)
}
func (m *FixAmount) XXX_Size() int {
	return xxx_messageInfo_FixAmount.Size(m)
}
func (m *FixAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_FixAmount.DiscardUnknown(m)
}

var xxx_messageInfo_FixAmount proto.InternalMessageInfo

func (m *FixAmount) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *FixAmount) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// 固定时间间隔按余量百分比解冻
type LeftProportion struct {
	Period               int64    `protobuf:"varint,1,opt,name=period,proto3" json:"period,omitempty"`
	TenThousandth        int64    `protobuf:"varint,2,opt,name=tenThousandth,proto3" json:"tenThousandth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeftProportion) Reset()         { *m = LeftProportion{} }
func (m *LeftProportion) String() string { return proto.CompactTextString(m) }
func (*LeftProportion) ProtoMessage()    {}
func (*LeftProportion) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{2}
}

func (m *LeftProportion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeftProportion.Unmarshal(m, b)
}
func (m *LeftProportion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeftProportion.Marshal(b, m, deterministic)
}
func (m *LeftProportion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeftProportion.Merge(m, src)
}
func (m *LeftProportion) XXX_Size() int {
	return xxx_messageInfo_LeftProportion.Size(m)
}
func (m *LeftProportion) XXX_DiscardUnknown() {
	xxx_messageInfo_LeftProportion.DiscardUnknown(m)
}

var xxx_messageInfo_LeftProportion proto.InternalMessageInfo

func (m *LeftProportion) GetPeriod() int64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *LeftProportion) GetTenThousandth() int64 {
	if m != nil {
		return m.TenThousandth
	}
	return 0
}

// message for execs.unfreeze
type UnfreezeAction struct {
	// Types that are valid to be assigned to Value:
	//	*UnfreezeAction_Create
	//	*UnfreezeAction_Withdraw
	//	*UnfreezeAction_Terminate
	Value                isUnfreezeAction_Value `protobuf_oneof:"value"`
	Ty                   int32                  `protobuf:"varint,4,opt,name=ty,proto3" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UnfreezeAction) Reset()         { *m = UnfreezeAction{} }
func (m *UnfreezeAction) String() string { return proto.CompactTextString(m) }
func (*UnfreezeAction) ProtoMessage()    {}
func (*UnfreezeAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{3}
}

func (m *UnfreezeAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnfreezeAction.Unmarshal(m, b)
}
func (m *UnfreezeAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnfreezeAction.Marshal(b, m, deterministic)
}
func (m *UnfreezeAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnfreezeAction.Merge(m, src)
}
func (m *UnfreezeAction) XXX_Size() int {
	return xxx_messageInfo_UnfreezeAction.Size(m)
}
func (m *UnfreezeAction) XXX_DiscardUnknown() {
	xxx_messageInfo_UnfreezeAction.DiscardUnknown(m)
}

var xxx_messageInfo_UnfreezeAction proto.InternalMessageInfo

type isUnfreezeAction_Value interface {
	isUnfreezeAction_Value()
}

type UnfreezeAction_Create struct {
	Create *UnfreezeCreate `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type UnfreezeAction_Withdraw struct {
	Withdraw *UnfreezeWithdraw `protobuf:"bytes,2,opt,name=withdraw,proto3,oneof"`
}

type UnfreezeAction_Terminate struct {
	Terminate *UnfreezeTerminate `protobuf:"bytes,3,opt,name=terminate,proto3,oneof"`
}

func (*UnfreezeAction_Create) isUnfreezeAction_Value() {}

func (*UnfreezeAction_Withdraw) isUnfreezeAction_Value() {}

func (*UnfreezeAction_Terminate) isUnfreezeAction_Value() {}

func (m *UnfreezeAction) GetValue() isUnfreezeAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *UnfreezeAction) GetCreate() *UnfreezeCreate {
	if x, ok := m.GetValue().(*UnfreezeAction_Create); ok {
		return x.Create
	}
	return nil
}

func (m *UnfreezeAction) GetWithdraw() *UnfreezeWithdraw {
	if x, ok := m.GetValue().(*UnfreezeAction_Withdraw); ok {
		return x.Withdraw
	}
	return nil
}

func (m *UnfreezeAction) GetTerminate() *UnfreezeTerminate {
	if x, ok := m.GetValue().(*UnfreezeAction_Terminate); ok {
		return x.Terminate
	}
	return nil
}

func (m *UnfreezeAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UnfreezeAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UnfreezeAction_Create)(nil),
		(*UnfreezeAction_Withdraw)(nil),
		(*UnfreezeAction_Terminate)(nil),
	}
}

// action
type UnfreezeCreate struct {
	StartTime   int64  `protobuf:"varint,1,opt,name=startTime,proto3" json:"startTime,omitempty"`
	AssetExec   string `protobuf:"bytes,2,opt,name=assetExec,proto3" json:"assetExec,omitempty"`
	AssetSymbol string `protobuf:"bytes,3,opt,name=assetSymbol,proto3" json:"assetSymbol,omitempty"`
	TotalCount  int64  `protobuf:"varint,4,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	Beneficiary string `protobuf:"bytes,5,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	Means       string `protobuf:"bytes,6,opt,name=means,proto3" json:"means,omitempty"`
	// Types that are valid to be assigned to MeansOpt:
	//	*UnfreezeCreate_FixAmount
	//	*UnfreezeCreate_LeftProportion
	MeansOpt             isUnfreezeCreate_MeansOpt `protobuf_oneof:"meansOpt"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UnfreezeCreate) Reset()         { *m = UnfreezeCreate{} }
func (m *UnfreezeCreate) String() string { return proto.CompactTextString(m) }
func (*UnfreezeCreate) ProtoMessage()    {}
func (*UnfreezeCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{4}
}

func (m *UnfreezeCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnfreezeCreate.Unmarshal(m, b)
}
func (m *UnfreezeCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnfreezeCreate.Marshal(b, m, deterministic)
}
func (m *UnfreezeCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnfreezeCreate.Merge(m, src)
}
func (m *UnfreezeCreate) XXX_Size() int {
	return xxx_messageInfo_UnfreezeCreate.Size(m)
}
func (m *UnfreezeCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_UnfreezeCreate.DiscardUnknown(m)
}

var xxx_messageInfo_UnfreezeCreate proto.InternalMessageInfo

func (m *UnfreezeCreate) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *UnfreezeCreate) GetAssetExec() string {
	if m != nil {
		return m.AssetExec
	}
	return ""
}

func (m *UnfreezeCreate) GetAssetSymbol() string {
	if m != nil {
		return m.AssetSymbol
	}
	return ""
}

func (m *UnfreezeCreate) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *UnfreezeCreate) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *UnfreezeCreate) GetMeans() string {
	if m != nil {
		return m.Means
	}
	return ""
}

type isUnfreezeCreate_MeansOpt interface {
	isUnfreezeCreate_MeansOpt()
}

type UnfreezeCreate_FixAmount struct {
	FixAmount *FixAmount `protobuf:"bytes,7,opt,name=fixAmount,proto3,oneof"`
}

type UnfreezeCreate_LeftProportion struct {
	LeftProportion *LeftProportion `protobuf:"bytes,8,opt,name=leftProportion,proto3,oneof"`
}

func (*UnfreezeCreate_FixAmount) isUnfreezeCreate_MeansOpt() {}

func (*UnfreezeCreate_LeftProportion) isUnfreezeCreate_MeansOpt() {}

func (m *UnfreezeCreate) GetMeansOpt() isUnfreezeCreate_MeansOpt {
	if m != nil {
		return m.MeansOpt
	}
	return nil
}

func (m *UnfreezeCreate) GetFixAmount() *FixAmount {
	if x, ok := m.GetMeansOpt().(*UnfreezeCreate_FixAmount); ok {
		return x.FixAmount
	}
	return nil
}

func (m *UnfreezeCreate) GetLeftProportion() *LeftProportion {
	if x, ok := m.GetMeansOpt().(*UnfreezeCreate_LeftProportion); ok {
		return x.LeftProportion
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*UnfreezeCreate) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*UnfreezeCreate_FixAmount)(nil),
		(*UnfreezeCreate_LeftProportion)(nil),
	}
}

type UnfreezeWithdraw struct {
	UnfreezeID           string   `protobuf:"bytes,1,opt,name=unfreezeID,proto3" json:"unfreezeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnfreezeWithdraw) Reset()         { *m = UnfreezeWithdraw{} }
func (m *UnfreezeWithdraw) String() string { return proto.CompactTextString(m) }
func (*UnfreezeWithdraw) ProtoMessage()    {}
func (*UnfreezeWithdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{5}
}

func (m *UnfreezeWithdraw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnfreezeWithdraw.Unmarshal(m, b)
}
func (m *UnfreezeWithdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnfreezeWithdraw.Marshal(b, m, deterministic)
}
func (m *UnfreezeWithdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnfreezeWithdraw.Merge(m, src)
}
func (m *UnfreezeWithdraw) XXX_Size() int {
	return xxx_messageInfo_UnfreezeWithdraw.Size(m)
}
func (m *UnfreezeWithdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_UnfreezeWithdraw.DiscardUnknown(m)
}

var xxx_messageInfo_UnfreezeWithdraw proto.InternalMessageInfo

func (m *UnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

type UnfreezeTerminate struct {
	UnfreezeID           string   `protobuf:"bytes,1,opt,name=unfreezeID,proto3" json:"unfreezeID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnfreezeTerminate) Reset()         { *m = UnfreezeTerminate{} }
func (m *UnfreezeTerminate) String() string { return proto.CompactTextString(m) }
func (*UnfreezeTerminate) ProtoMessage()    {}
func (*UnfreezeTerminate) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{6}
}

func (m *UnfreezeTerminate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnfreezeTerminate.Unmarshal(m, b)
}
func (m *UnfreezeTerminate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnfreezeTerminate.Marshal(b, m, deterministic)
}
func (m *UnfreezeTerminate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnfreezeTerminate.Merge(m, src)
}
func (m *UnfreezeTerminate) XXX_Size() int {
	return xxx_messageInfo_UnfreezeTerminate.Size(m)
}
func (m *UnfreezeTerminate) XXX_DiscardUnknown() {
	xxx_messageInfo_UnfreezeTerminate.DiscardUnknown(m)
}

var xxx_messageInfo_UnfreezeTerminate proto.InternalMessageInfo

func (m *UnfreezeTerminate) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

// receipt
type ReceiptUnfreeze struct {
	Prev                 *Unfreeze `protobuf:"bytes,1,opt,name=prev,proto3" json:"prev,omitempty"`
	Current              *Unfreeze `protobuf:"bytes,2,opt,name=current,proto3" json:"current,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReceiptUnfreeze) Reset()         { *m = ReceiptUnfreeze{} }
func (m *ReceiptUnfreeze) String() string { return proto.CompactTextString(m) }
func (*ReceiptUnfreeze) ProtoMessage()    {}
func (*ReceiptUnfreeze) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{7}
}

func (m *ReceiptUnfreeze) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReceiptUnfreeze.Unmarshal(m, b)
}
func (m *ReceiptUnfreeze) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReceiptUnfreeze.Marshal(b, m, deterministic)
}
func (m *ReceiptUnfreeze) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReceiptUnfreeze.Merge(m, src)
}
func (m *ReceiptUnfreeze) XXX_Size() int {
	return xxx_messageInfo_ReceiptUnfreeze.Size(m)
}
func (m *ReceiptUnfreeze) XXX_DiscardUnknown() {
	xxx_messageInfo_ReceiptUnfreeze.DiscardUnknown(m)
}

var xxx_messageInfo_ReceiptUnfreeze proto.InternalMessageInfo

func (m *ReceiptUnfreeze) GetPrev() *Unfreeze {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptUnfreeze) GetCurrent() *Unfreeze {
	if m != nil {
		return m.Current
	}
	return nil
}

type LocalUnfreeze struct {
	Unfreeze             *Unfreeze `protobuf:"bytes,1,opt,name=unfreeze,proto3" json:"unfreeze,omitempty"`
	TxIndex              string    `protobuf:"bytes,2,opt,name=txIndex,proto3" json:"txIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LocalUnfreeze) Reset()         { *m = LocalUnfreeze{} }
func (m *LocalUnfreeze) String() string { return proto.CompactTextString(m) }
func (*LocalUnfreeze) ProtoMessage()    {}
func (*LocalUnfreeze) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{8}
}

func (m *LocalUnfreeze) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalUnfreeze.Unmarshal(m, b)
}
func (m *LocalUnfreeze) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalUnfreeze.Marshal(b, m, deterministic)
}
func (m *LocalUnfreeze) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalUnfreeze.Merge(m, src)
}
func (m *LocalUnfreeze) XXX_Size() int {
	return xxx_messageInfo_LocalUnfreeze.Size(m)
}
func (m *LocalUnfreeze) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalUnfreeze.DiscardUnknown(m)
}

var xxx_messageInfo_LocalUnfreeze proto.InternalMessageInfo

func (m *LocalUnfreeze) GetUnfreeze() *Unfreeze {
	if m != nil {
		return m.Unfreeze
	}
	return nil
}

func (m *LocalUnfreeze) GetTxIndex() string {
	if m != nil {
		return m.TxIndex
	}
	return ""
}

// query
type ReplyQueryUnfreezeWithdraw struct {
	UnfreezeID           string   `protobuf:"bytes,1,opt,name=unfreezeID,proto3" json:"unfreezeID,omitempty"`
	AvailableAmount      int64    `protobuf:"varint,2,opt,name=availableAmount,proto3" json:"availableAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReplyQueryUnfreezeWithdraw) Reset()         { *m = ReplyQueryUnfreezeWithdraw{} }
func (m *ReplyQueryUnfreezeWithdraw) String() string { return proto.CompactTextString(m) }
func (*ReplyQueryUnfreezeWithdraw) ProtoMessage()    {}
func (*ReplyQueryUnfreezeWithdraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{9}
}

func (m *ReplyQueryUnfreezeWithdraw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyQueryUnfreezeWithdraw.Unmarshal(m, b)
}
func (m *ReplyQueryUnfreezeWithdraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyQueryUnfreezeWithdraw.Marshal(b, m, deterministic)
}
func (m *ReplyQueryUnfreezeWithdraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyQueryUnfreezeWithdraw.Merge(m, src)
}
func (m *ReplyQueryUnfreezeWithdraw) XXX_Size() int {
	return xxx_messageInfo_ReplyQueryUnfreezeWithdraw.Size(m)
}
func (m *ReplyQueryUnfreezeWithdraw) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyQueryUnfreezeWithdraw.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyQueryUnfreezeWithdraw proto.InternalMessageInfo

func (m *ReplyQueryUnfreezeWithdraw) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *ReplyQueryUnfreezeWithdraw) GetAvailableAmount() int64 {
	if m != nil {
		return m.AvailableAmount
	}
	return 0
}

type ReqUnfreezes struct {
	Direction            int32    `protobuf:"varint,1,opt,name=direction,proto3" json:"direction,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	FromKey              string   `protobuf:"bytes,3,opt,name=fromKey,proto3" json:"fromKey,omitempty"`
	Initiator            string   `protobuf:"bytes,4,opt,name=initiator,proto3" json:"initiator,omitempty"`
	Beneficiary          string   `protobuf:"bytes,5,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUnfreezes) Reset()         { *m = ReqUnfreezes{} }
func (m *ReqUnfreezes) String() string { return proto.CompactTextString(m) }
func (*ReqUnfreezes) ProtoMessage()    {}
func (*ReqUnfreezes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{10}
}

func (m *ReqUnfreezes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUnfreezes.Unmarshal(m, b)
}
func (m *ReqUnfreezes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUnfreezes.Marshal(b, m, deterministic)
}
func (m *ReqUnfreezes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUnfreezes.Merge(m, src)
}
func (m *ReqUnfreezes) XXX_Size() int {
	return xxx_messageInfo_ReqUnfreezes.Size(m)
}
func (m *ReqUnfreezes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUnfreezes.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUnfreezes proto.InternalMessageInfo

func (m *ReqUnfreezes) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *ReqUnfreezes) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqUnfreezes) GetFromKey() string {
	if m != nil {
		return m.FromKey
	}
	return ""
}

func (m *ReqUnfreezes) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *ReqUnfreezes) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

type ReplyUnfreeze struct {
	//解冻交易ID（唯一识别码）
	UnfreezeID string `protobuf:"bytes,1,opt,name=unfreezeID,proto3" json:"unfreezeID,omitempty"`
	//开始时间
	StartTime int64 `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	//币种
	AssetExec   string `protobuf:"bytes,3,opt,name=assetExec,proto3" json:"assetExec,omitempty"`
	AssetSymbol string `protobuf:"bytes,4,opt,name=assetSymbol,proto3" json:"assetSymbol,omitempty"`
	//冻结总额
	TotalCount int64 `protobuf:"varint,5,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	//发币人地址
	Initiator string `protobuf:"bytes,6,opt,name=initiator,proto3" json:"initiator,omitempty"`
	//收币人地址
	Beneficiary string `protobuf:"bytes,7,opt,name=beneficiary,proto3" json:"beneficiary,omitempty"`
	//解冻剩余币数
	Remaining int64 `protobuf:"varint,8,opt,name=remaining,proto3" json:"remaining,omitempty"`
	//解冻方式（百分比；固额）
	Means string `protobuf:"bytes,9,opt,name=means,proto3" json:"means,omitempty"`
	// Types that are valid to be assigned to MeansOpt:
	//	*ReplyUnfreeze_FixAmount
	//	*ReplyUnfreeze_LeftProportion
	MeansOpt             isReplyUnfreeze_MeansOpt `protobuf_oneof:"meansOpt"`
	Terminated           bool                     `protobuf:"varint,12,opt,name=terminated,proto3" json:"terminated,omitempty"`
	Key                  string                   `protobuf:"bytes,13,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ReplyUnfreeze) Reset()         { *m = ReplyUnfreeze{} }
func (m *ReplyUnfreeze) String() string { return proto.CompactTextString(m) }
func (*ReplyUnfreeze) ProtoMessage()    {}
func (*ReplyUnfreeze) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{11}
}

func (m *ReplyUnfreeze) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyUnfreeze.Unmarshal(m, b)
}
func (m *ReplyUnfreeze) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyUnfreeze.Marshal(b, m, deterministic)
}
func (m *ReplyUnfreeze) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyUnfreeze.Merge(m, src)
}
func (m *ReplyUnfreeze) XXX_Size() int {
	return xxx_messageInfo_ReplyUnfreeze.Size(m)
}
func (m *ReplyUnfreeze) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyUnfreeze.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyUnfreeze proto.InternalMessageInfo

func (m *ReplyUnfreeze) GetUnfreezeID() string {
	if m != nil {
		return m.UnfreezeID
	}
	return ""
}

func (m *ReplyUnfreeze) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *ReplyUnfreeze) GetAssetExec() string {
	if m != nil {
		return m.AssetExec
	}
	return ""
}

func (m *ReplyUnfreeze) GetAssetSymbol() string {
	if m != nil {
		return m.AssetSymbol
	}
	return ""
}

func (m *ReplyUnfreeze) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ReplyUnfreeze) GetInitiator() string {
	if m != nil {
		return m.Initiator
	}
	return ""
}

func (m *ReplyUnfreeze) GetBeneficiary() string {
	if m != nil {
		return m.Beneficiary
	}
	return ""
}

func (m *ReplyUnfreeze) GetRemaining() int64 {
	if m != nil {
		return m.Remaining
	}
	return 0
}

func (m *ReplyUnfreeze) GetMeans() string {
	if m != nil {
		return m.Means
	}
	return ""
}

type isReplyUnfreeze_MeansOpt interface {
	isReplyUnfreeze_MeansOpt()
}

type ReplyUnfreeze_FixAmount struct {
	FixAmount *FixAmount `protobuf:"bytes,10,opt,name=fixAmount,proto3,oneof"`
}

type ReplyUnfreeze_LeftProportion struct {
	LeftProportion *LeftProportion `protobuf:"bytes,11,opt,name=leftProportion,proto3,oneof"`
}

func (*ReplyUnfreeze_FixAmount) isReplyUnfreeze_MeansOpt() {}

func (*ReplyUnfreeze_LeftProportion) isReplyUnfreeze_MeansOpt() {}

func (m *ReplyUnfreeze) GetMeansOpt() isReplyUnfreeze_MeansOpt {
	if m != nil {
		return m.MeansOpt
	}
	return nil
}

func (m *ReplyUnfreeze) GetFixAmount() *FixAmount {
	if x, ok := m.GetMeansOpt().(*ReplyUnfreeze_FixAmount); ok {
		return x.FixAmount
	}
	return nil
}

func (m *ReplyUnfreeze) GetLeftProportion() *LeftProportion {
	if x, ok := m.GetMeansOpt().(*ReplyUnfreeze_LeftProportion); ok {
		return x.LeftProportion
	}
	return nil
}

func (m *ReplyUnfreeze) GetTerminated() bool {
	if m != nil {
		return m.Terminated
	}
	return false
}

func (m *ReplyUnfreeze) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ReplyUnfreeze) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ReplyUnfreeze_FixAmount)(nil),
		(*ReplyUnfreeze_LeftProportion)(nil),
	}
}

type ReplyUnfreezes struct {
	Unfreeze             []*ReplyUnfreeze `protobuf:"bytes,1,rep,name=unfreeze,proto3" json:"unfreeze,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ReplyUnfreezes) Reset()         { *m = ReplyUnfreezes{} }
func (m *ReplyUnfreezes) String() string { return proto.CompactTextString(m) }
func (*ReplyUnfreezes) ProtoMessage()    {}
func (*ReplyUnfreezes) Descriptor() ([]byte, []int) {
	return fileDescriptor_6caa0554cb0b9167, []int{12}
}

func (m *ReplyUnfreezes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyUnfreezes.Unmarshal(m, b)
}
func (m *ReplyUnfreezes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyUnfreezes.Marshal(b, m, deterministic)
}
func (m *ReplyUnfreezes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyUnfreezes.Merge(m, src)
}
func (m *ReplyUnfreezes) XXX_Size() int {
	return xxx_messageInfo_ReplyUnfreezes.Size(m)
}
func (m *ReplyUnfreezes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyUnfreezes.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyUnfreezes proto.InternalMessageInfo

func (m *ReplyUnfreezes) GetUnfreeze() []*ReplyUnfreeze {
	if m != nil {
		return m.Unfreeze
	}
	return nil
}

func init() {
	proto.RegisterType((*Unfreeze)(nil), "types.Unfreeze")
	proto.RegisterType((*FixAmount)(nil), "types.FixAmount")
	proto.RegisterType((*LeftProportion)(nil), "types.LeftProportion")
	proto.RegisterType((*UnfreezeAction)(nil), "types.UnfreezeAction")
	proto.RegisterType((*UnfreezeCreate)(nil), "types.UnfreezeCreate")
	proto.RegisterType((*UnfreezeWithdraw)(nil), "types.UnfreezeWithdraw")
	proto.RegisterType((*UnfreezeTerminate)(nil), "types.UnfreezeTerminate")
	proto.RegisterType((*ReceiptUnfreeze)(nil), "types.ReceiptUnfreeze")
	proto.RegisterType((*LocalUnfreeze)(nil), "types.LocalUnfreeze")
	proto.RegisterType((*ReplyQueryUnfreezeWithdraw)(nil), "types.ReplyQueryUnfreezeWithdraw")
	proto.RegisterType((*ReqUnfreezes)(nil), "types.ReqUnfreezes")
	proto.RegisterType((*ReplyUnfreeze)(nil), "types.ReplyUnfreeze")
	proto.RegisterType((*ReplyUnfreezes)(nil), "types.ReplyUnfreezes")
}

func init() {
	proto.RegisterFile("unfreeze.proto", fileDescriptor_6caa0554cb0b9167)
}

var fileDescriptor_6caa0554cb0b9167 = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x56, 0xcb, 0x6e, 0x13, 0x4b,
	0x10, 0xf5, 0x78, 0x3c, 0x7e, 0x94, 0x63, 0xc7, 0xb7, 0x6f, 0xee, 0x65, 0x14, 0x21, 0x64, 0x06,
	0x16, 0x46, 0x48, 0x01, 0x39, 0x20, 0x21, 0xb1, 0x40, 0x49, 0x78, 0x38, 0x22, 0xe2, 0x31, 0x09,
	0xb0, 0x6e, 0x8f, 0xcb, 0x49, 0x8b, 0x79, 0xa5, 0xdd, 0x4e, 0x3c, 0x7c, 0x04, 0x5f, 0xc0, 0x9f,
	0xb0, 0x66, 0xc5, 0x8e, 0x2f, 0x42, 0xd3, 0xf3, 0x1e, 0x27, 0x58, 0x61, 0x87, 0xc4, 0xce, 0x75,
	0xaa, 0xea, 0x74, 0x75, 0xd5, 0xa9, 0x1e, 0x43, 0x77, 0xee, 0x4e, 0x39, 0xe2, 0x27, 0xdc, 0xf2,
	0xb9, 0x27, 0x3c, 0xa2, 0x89, 0xc0, 0xc7, 0xd9, 0xe6, 0x9a, 0xe5, 0x39, 0x8e, 0xe7, 0x46, 0xa0,
	0xf1, 0x55, 0x85, 0xe6, 0xbb, 0x38, 0x8e, 0xdc, 0x00, 0x48, 0x72, 0xf6, 0x9f, 0xea, 0x4a, 0x5f,
	0x19, 0xb4, 0xcc, 0x1c, 0x42, 0xae, 0x43, 0x6b, 0x26, 0x28, 0x17, 0x47, 0xcc, 0x41, 0xbd, 0xda,
	0x57, 0x06, 0xaa, 0x99, 0x01, 0xa1, 0x97, 0xce, 0x66, 0x28, 0x9e, 0x2d, 0xd0, 0xd2, 0x55, 0x99,
	0x9c, 0x01, 0xa4, 0x0f, 0x6d, 0x69, 0x1c, 0x06, 0xce, 0xd8, 0xb3, 0xf5, 0x9a, 0xf4, 0xe7, 0xa1,
	0xf0, 0x74, 0xe1, 0x09, 0x6a, 0xef, 0x79, 0x73, 0x57, 0xe8, 0x9a, 0xa4, 0xcf, 0x21, 0x21, 0x3f,
	0x73, 0x99, 0x60, 0x54, 0x78, 0x5c, 0xaf, 0x47, 0xfc, 0x29, 0x10, 0xf2, 0x8f, 0xd1, 0xc5, 0x29,
	0xb3, 0x18, 0xe5, 0x81, 0xde, 0x88, 0xf8, 0x73, 0x50, 0x98, 0xcf, 0xd1, 0xa1, 0xcc, 0x65, 0xee,
	0xb1, 0xde, 0x8c, 0xaa, 0x4f, 0x01, 0xb2, 0x01, 0x9a, 0x83, 0xd4, 0x9d, 0xe9, 0x2d, 0x99, 0x19,
	0x19, 0xe4, 0x3e, 0xb4, 0xa6, 0x6c, 0xb1, 0xe3, 0xc8, 0x92, 0xa0, 0xaf, 0x0c, 0xda, 0xc3, 0xde,
	0x96, 0xec, 0xe3, 0xd6, 0xf3, 0x04, 0x1f, 0x55, 0xcc, 0x2c, 0x88, 0x3c, 0x81, 0xae, 0x8d, 0x53,
	0xf1, 0x86, 0x7b, 0xbe, 0xc7, 0x05, 0xf3, 0x5c, 0xbd, 0x2d, 0xd3, 0xfe, 0x8b, 0xd3, 0x0e, 0x0a,
	0xce, 0x51, 0xc5, 0x2c, 0x85, 0xcb, 0x36, 0x20, 0x77, 0x98, 0x4b, 0x05, 0x4e, 0xf4, 0xb5, 0xbe,
	0x32, 0x68, 0x9a, 0x39, 0x64, 0x17, 0xa0, 0x29, 0x6b, 0x7b, 0xed, 0x0b, 0xe3, 0x31, 0xb4, 0xd2,
	0x32, 0xc8, 0xff, 0x50, 0xf7, 0x91, 0x33, 0x6f, 0x22, 0x27, 0xa7, 0x9a, 0xb1, 0x15, 0xe2, 0x34,
	0xba, 0x40, 0x34, 0xb2, 0xd8, 0x32, 0x5e, 0x41, 0xb7, 0x58, 0xcc, 0xa5, 0x0c, 0xb7, 0xa1, 0x23,
	0xd0, 0x3d, 0x3a, 0xf1, 0xe6, 0x33, 0xea, 0x4e, 0xc4, 0x49, 0x4c, 0x54, 0x04, 0x8d, 0xef, 0x0a,
	0x74, 0x13, 0x29, 0xed, 0x58, 0x92, 0xf0, 0x1e, 0xd4, 0x2d, 0x8e, 0x54, 0xa0, 0x24, 0xcc, 0x9a,
	0x90, 0x84, 0xed, 0x49, 0xe7, 0xa8, 0x62, 0xc6, 0x61, 0xe4, 0x21, 0x34, 0xcf, 0x99, 0x38, 0x99,
	0x70, 0x7a, 0x2e, 0x0f, 0x69, 0x0f, 0xaf, 0x95, 0x52, 0x3e, 0xc4, 0xee, 0x51, 0xc5, 0x4c, 0x43,
	0xc9, 0x23, 0x68, 0xa5, 0x1d, 0x92, 0xd2, 0x6b, 0x0f, 0xf5, 0x52, 0xde, 0x51, 0xe2, 0x0f, 0xc7,
	0x95, 0x06, 0x93, 0x2e, 0x54, 0x45, 0x20, 0xd5, 0xa8, 0x99, 0x55, 0x11, 0xec, 0x36, 0x40, 0x3b,
	0xa3, 0xf6, 0x1c, 0x8d, 0x6f, 0xd5, 0xec, 0x36, 0x51, 0x99, 0x45, 0xf9, 0x2b, 0xbf, 0x94, 0x7f,
	0x75, 0x85, 0xfc, 0xd5, 0x55, 0xf2, 0xaf, 0x2d, 0xc9, 0xbf, 0x24, 0x70, 0x6d, 0x59, 0xe0, 0xa9,
	0x84, 0xeb, 0x97, 0x4a, 0xb8, 0xf1, 0x7b, 0x12, 0x6e, 0x5e, 0x49, 0xc2, 0x05, 0x89, 0x0e, 0xa1,
	0x57, 0x1e, 0xdd, 0xaa, 0x77, 0xc6, 0xd8, 0x86, 0x7f, 0x96, 0xc6, 0xb6, 0x32, 0x89, 0xc2, 0xba,
	0x89, 0x16, 0x32, 0x5f, 0xa4, 0xef, 0xd9, 0x2d, 0xa8, 0xf9, 0x1c, 0xcf, 0x62, 0xf1, 0xad, 0x97,
	0x14, 0x61, 0x4a, 0x27, 0xb9, 0x03, 0x0d, 0x6b, 0xce, 0x39, 0xc6, 0xfb, 0x71, 0x41, 0x5c, 0xe2,
	0x37, 0xde, 0x43, 0xe7, 0xc0, 0xb3, 0xa8, 0x9d, 0x1e, 0x70, 0x17, 0x9a, 0x49, 0x05, 0x97, 0x1d,
	0x92, 0x06, 0x10, 0x1d, 0x1a, 0x62, 0xb1, 0xef, 0x4e, 0x70, 0x11, 0xcb, 0x23, 0x31, 0x8d, 0x29,
	0x6c, 0x9a, 0xe8, 0xdb, 0xc1, 0xdb, 0x39, 0xf2, 0xe0, 0xaa, 0xdd, 0x22, 0x03, 0x58, 0xa7, 0x67,
	0x94, 0xd9, 0x74, 0x6c, 0xe3, 0x4e, 0x7e, 0xd1, 0xcb, 0xb0, 0xf1, 0x45, 0x81, 0x35, 0x13, 0x4f,
	0x93, 0x13, 0x66, 0xa1, 0x66, 0x27, 0x8c, 0xa3, 0x5c, 0x56, 0xc9, 0xac, 0x99, 0x19, 0x10, 0xea,
	0xc9, 0x4a, 0xe9, 0x34, 0x33, 0x32, 0xc2, 0x6b, 0x4c, 0xb9, 0xe7, 0xbc, 0xc4, 0x20, 0x56, 0x71,
	0x62, 0x16, 0x1f, 0xe8, 0xda, 0x8a, 0x07, 0x7a, 0x59, 0xbf, 0xc6, 0x0f, 0x15, 0x3a, 0xb2, 0x0f,
	0x7f, 0x3f, 0x48, 0x7f, 0xd0, 0x07, 0x89, 0xf4, 0x40, 0xfd, 0x88, 0x81, 0xde, 0x91, 0x65, 0x86,
	0x3f, 0x0b, 0xfb, 0xbf, 0x0b, 0xdd, 0xc2, 0x4c, 0xc3, 0x2b, 0xe4, 0x97, 0x46, 0x1d, 0xb4, 0x87,
	0x1b, 0x71, 0x29, 0x85, 0xc0, 0x6c, 0x73, 0x86, 0x9f, 0x95, 0x2c, 0x85, 0x1c, 0xc0, 0xbf, 0x2f,
	0x50, 0x2c, 0x6d, 0x49, 0x2f, 0xe5, 0x38, 0x3d, 0x14, 0x9c, 0xb9, 0xc7, 0x9b, 0x37, 0xf3, 0xac,
	0x17, 0xae, 0x96, 0x51, 0x21, 0x0f, 0xa0, 0x53, 0x70, 0x5d, 0xc0, 0x53, 0x5e, 0x69, 0xa3, 0x32,
	0xae, 0xcb, 0x3f, 0x4f, 0xdb, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd2, 0x37, 0x34, 0x48, 0x63,
	0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UnfreezeClient is the client API for Unfreeze service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UnfreezeClient interface {
	GetUnfreezeWithdraw(ctx context.Context, in *types.ReqString, opts ...grpc.CallOption) (*ReplyQueryUnfreezeWithdraw, error)
	QueryUnfreeze(ctx context.Context, in *types.ReqString, opts ...grpc.CallOption) (*Unfreeze, error)
}

type unfreezeClient struct {
	cc grpc.ClientConnInterface
}

func NewUnfreezeClient(cc grpc.ClientConnInterface) UnfreezeClient {
	return &unfreezeClient{cc}
}

func (c *unfreezeClient) GetUnfreezeWithdraw(ctx context.Context, in *types.ReqString, opts ...grpc.CallOption) (*ReplyQueryUnfreezeWithdraw, error) {
	out := new(ReplyQueryUnfreezeWithdraw)
	err := c.cc.Invoke(ctx, "/types.unfreeze/GetUnfreezeWithdraw", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *unfreezeClient) QueryUnfreeze(ctx context.Context, in *types.ReqString, opts ...grpc.CallOption) (*Unfreeze, error) {
	out := new(Unfreeze)
	err := c.cc.Invoke(ctx, "/types.unfreeze/QueryUnfreeze", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UnfreezeServer is the server API for Unfreeze service.
type UnfreezeServer interface {
	GetUnfreezeWithdraw(context.Context, *types.ReqString) (*ReplyQueryUnfreezeWithdraw, error)
	QueryUnfreeze(context.Context, *types.ReqString) (*Unfreeze, error)
}

// UnimplementedUnfreezeServer can be embedded to have forward compatible implementations.
type UnimplementedUnfreezeServer struct {
}

func (*UnimplementedUnfreezeServer) GetUnfreezeWithdraw(ctx context.Context, req *types.ReqString) (*ReplyQueryUnfreezeWithdraw, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnfreezeWithdraw not implemented")
}
func (*UnimplementedUnfreezeServer) QueryUnfreeze(ctx context.Context, req *types.ReqString) (*Unfreeze, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUnfreeze not implemented")
}

func RegisterUnfreezeServer(s *grpc.Server, srv UnfreezeServer) {
	s.RegisterService(&_Unfreeze_serviceDesc, srv)
}

func _Unfreeze_GetUnfreezeWithdraw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ReqString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnfreezeServer).GetUnfreezeWithdraw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.unfreeze/GetUnfreezeWithdraw",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnfreezeServer).GetUnfreezeWithdraw(ctx, req.(*types.ReqString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Unfreeze_QueryUnfreeze_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.ReqString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UnfreezeServer).QueryUnfreeze(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.unfreeze/QueryUnfreeze",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UnfreezeServer).QueryUnfreeze(ctx, req.(*types.ReqString))
	}
	return interceptor(ctx, in, info, handler)
}

var _Unfreeze_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.unfreeze",
	HandlerType: (*UnfreezeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUnfreezeWithdraw",
			Handler:    _Unfreeze_GetUnfreezeWithdraw_Handler,
		},
		{
			MethodName: "QueryUnfreeze",
			Handler:    _Unfreeze_QueryUnfreeze_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "unfreeze.proto",
}
