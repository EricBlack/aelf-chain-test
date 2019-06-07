// Code generated by protoc-gen-go. DO NOT EDIT.
// source: aelf/core.proto

package aelf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
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

type TransactionResultStatus int32

const (
	TransactionResultStatus_NotExisted   TransactionResultStatus = 0
	TransactionResultStatus_Pending      TransactionResultStatus = 1
	TransactionResultStatus_Failed       TransactionResultStatus = 2
	TransactionResultStatus_Mined        TransactionResultStatus = 3
	TransactionResultStatus_Unexecutable TransactionResultStatus = 4
)

var TransactionResultStatus_name = map[int32]string{
	0: "NotExisted",
	1: "Pending",
	2: "Failed",
	3: "Mined",
	4: "Unexecutable",
}

var TransactionResultStatus_value = map[string]int32{
	"NotExisted":   0,
	"Pending":      1,
	"Failed":       2,
	"Mined":        3,
	"Unexecutable": 4,
}

func (x TransactionResultStatus) String() string {
	return proto.EnumName(TransactionResultStatus_name, int32(x))
}

func (TransactionResultStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{0}
}

type Transaction struct {
	From                 *Address `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	To                   *Address `protobuf:"bytes,2,opt,name=To,proto3" json:"To,omitempty"`
	RefBlockNumber       int64    `protobuf:"varint,3,opt,name=RefBlockNumber,proto3" json:"RefBlockNumber,omitempty"`
	RefBlockPrefix       []byte   `protobuf:"bytes,4,opt,name=RefBlockPrefix,proto3" json:"RefBlockPrefix,omitempty"`
	MethodName           string   `protobuf:"bytes,5,opt,name=MethodName,proto3" json:"MethodName,omitempty"`
	Params               []byte   `protobuf:"bytes,6,opt,name=Params,proto3" json:"Params,omitempty"`
	Signature            []byte   `protobuf:"bytes,10000,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{0}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetFrom() *Address {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Transaction) GetTo() *Address {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *Transaction) GetRefBlockNumber() int64 {
	if m != nil {
		return m.RefBlockNumber
	}
	return 0
}

func (m *Transaction) GetRefBlockPrefix() []byte {
	if m != nil {
		return m.RefBlockPrefix
	}
	return nil
}

func (m *Transaction) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *Transaction) GetParams() []byte {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *Transaction) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type StatePath struct {
	Parts                []string `protobuf:"bytes,1,rep,name=parts,proto3" json:"parts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatePath) Reset()         { *m = StatePath{} }
func (m *StatePath) String() string { return proto.CompactTextString(m) }
func (*StatePath) ProtoMessage()    {}
func (*StatePath) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{1}
}

func (m *StatePath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatePath.Unmarshal(m, b)
}
func (m *StatePath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatePath.Marshal(b, m, deterministic)
}
func (m *StatePath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatePath.Merge(m, src)
}
func (m *StatePath) XXX_Size() int {
	return xxx_messageInfo_StatePath.Size(m)
}
func (m *StatePath) XXX_DiscardUnknown() {
	xxx_messageInfo_StatePath.DiscardUnknown(m)
}

var xxx_messageInfo_StatePath proto.InternalMessageInfo

func (m *StatePath) GetParts() []string {
	if m != nil {
		return m.Parts
	}
	return nil
}

type ScopedStatePath struct {
	Address              *Address   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Path                 *StatePath `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ScopedStatePath) Reset()         { *m = ScopedStatePath{} }
func (m *ScopedStatePath) String() string { return proto.CompactTextString(m) }
func (*ScopedStatePath) ProtoMessage()    {}
func (*ScopedStatePath) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{2}
}

func (m *ScopedStatePath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScopedStatePath.Unmarshal(m, b)
}
func (m *ScopedStatePath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScopedStatePath.Marshal(b, m, deterministic)
}
func (m *ScopedStatePath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScopedStatePath.Merge(m, src)
}
func (m *ScopedStatePath) XXX_Size() int {
	return xxx_messageInfo_ScopedStatePath.Size(m)
}
func (m *ScopedStatePath) XXX_DiscardUnknown() {
	xxx_messageInfo_ScopedStatePath.DiscardUnknown(m)
}

var xxx_messageInfo_ScopedStatePath proto.InternalMessageInfo

func (m *ScopedStatePath) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ScopedStatePath) GetPath() *StatePath {
	if m != nil {
		return m.Path
	}
	return nil
}

type TransactionResult struct {
	TransactionId *Hash                   `protobuf:"bytes,1,opt,name=TransactionId,proto3" json:"TransactionId,omitempty"`
	Status        TransactionResultStatus `protobuf:"varint,2,opt,name=Status,proto3,enum=aelf.TransactionResultStatus" json:"Status,omitempty"`
	Logs          []*LogEvent             `protobuf:"bytes,3,rep,name=Logs,proto3" json:"Logs,omitempty"`
	Bloom         []byte                  `protobuf:"bytes,4,opt,name=Bloom,proto3" json:"Bloom,omitempty"`
	ReturnValue   []byte                  `protobuf:"bytes,5,opt,name=ReturnValue,proto3" json:"ReturnValue,omitempty"`
	BlockNumber   int64                   `protobuf:"varint,6,opt,name=BlockNumber,proto3" json:"BlockNumber,omitempty"`
	BlockHash     *Hash                   `protobuf:"bytes,7,opt,name=BlockHash,proto3" json:"BlockHash,omitempty"`
	Index         int32                   `protobuf:"varint,8,opt,name=Index,proto3" json:"Index,omitempty"`
	StateHash     *Hash                   `protobuf:"bytes,9,opt,name=StateHash,proto3" json:"StateHash,omitempty"`
	// Merkle proof path for this transaction
	// TODO: Remove DeferredTxnId
	DeferredTransactions []*Transaction `protobuf:"bytes,10,rep,name=DeferredTransactions,proto3" json:"DeferredTransactions,omitempty"`
	DeferredTxnId        *Hash          `protobuf:"bytes,11,opt,name=DeferredTxnId,proto3" json:"DeferredTxnId,omitempty"`
	Error                string         `protobuf:"bytes,12,opt,name=Error,proto3" json:"Error,omitempty"`
	ReadableReturnValue  string         `protobuf:"bytes,13,opt,name=ReadableReturnValue,proto3" json:"ReadableReturnValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TransactionResult) Reset()         { *m = TransactionResult{} }
func (m *TransactionResult) String() string { return proto.CompactTextString(m) }
func (*TransactionResult) ProtoMessage()    {}
func (*TransactionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{3}
}

func (m *TransactionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionResult.Unmarshal(m, b)
}
func (m *TransactionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionResult.Marshal(b, m, deterministic)
}
func (m *TransactionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionResult.Merge(m, src)
}
func (m *TransactionResult) XXX_Size() int {
	return xxx_messageInfo_TransactionResult.Size(m)
}
func (m *TransactionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionResult.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionResult proto.InternalMessageInfo

func (m *TransactionResult) GetTransactionId() *Hash {
	if m != nil {
		return m.TransactionId
	}
	return nil
}

func (m *TransactionResult) GetStatus() TransactionResultStatus {
	if m != nil {
		return m.Status
	}
	return TransactionResultStatus_NotExisted
}

func (m *TransactionResult) GetLogs() []*LogEvent {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *TransactionResult) GetBloom() []byte {
	if m != nil {
		return m.Bloom
	}
	return nil
}

func (m *TransactionResult) GetReturnValue() []byte {
	if m != nil {
		return m.ReturnValue
	}
	return nil
}

func (m *TransactionResult) GetBlockNumber() int64 {
	if m != nil {
		return m.BlockNumber
	}
	return 0
}

func (m *TransactionResult) GetBlockHash() *Hash {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *TransactionResult) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TransactionResult) GetStateHash() *Hash {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *TransactionResult) GetDeferredTransactions() []*Transaction {
	if m != nil {
		return m.DeferredTransactions
	}
	return nil
}

func (m *TransactionResult) GetDeferredTxnId() *Hash {
	if m != nil {
		return m.DeferredTxnId
	}
	return nil
}

func (m *TransactionResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *TransactionResult) GetReadableReturnValue() string {
	if m != nil {
		return m.ReadableReturnValue
	}
	return ""
}

type LogEvent struct {
	Address              *Address `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Indexed              [][]byte `protobuf:"bytes,3,rep,name=Indexed,proto3" json:"Indexed,omitempty"`
	NonIndexed           []byte   `protobuf:"bytes,4,opt,name=NonIndexed,proto3" json:"NonIndexed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogEvent) Reset()         { *m = LogEvent{} }
func (m *LogEvent) String() string { return proto.CompactTextString(m) }
func (*LogEvent) ProtoMessage()    {}
func (*LogEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{4}
}

func (m *LogEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEvent.Unmarshal(m, b)
}
func (m *LogEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEvent.Marshal(b, m, deterministic)
}
func (m *LogEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEvent.Merge(m, src)
}
func (m *LogEvent) XXX_Size() int {
	return xxx_messageInfo_LogEvent.Size(m)
}
func (m *LogEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEvent.DiscardUnknown(m)
}

var xxx_messageInfo_LogEvent proto.InternalMessageInfo

func (m *LogEvent) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *LogEvent) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogEvent) GetIndexed() [][]byte {
	if m != nil {
		return m.Indexed
	}
	return nil
}

func (m *LogEvent) GetNonIndexed() []byte {
	if m != nil {
		return m.NonIndexed
	}
	return nil
}

type SmartContractRegistration struct {
	Category             int32    `protobuf:"varint,1,opt,name=Category,proto3" json:"Category,omitempty"`
	Code                 []byte   `protobuf:"bytes,2,opt,name=Code,proto3" json:"Code,omitempty"`
	CodeHash             *Hash    `protobuf:"bytes,3,opt,name=CodeHash,proto3" json:"CodeHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmartContractRegistration) Reset()         { *m = SmartContractRegistration{} }
func (m *SmartContractRegistration) String() string { return proto.CompactTextString(m) }
func (*SmartContractRegistration) ProtoMessage()    {}
func (*SmartContractRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{5}
}

func (m *SmartContractRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmartContractRegistration.Unmarshal(m, b)
}
func (m *SmartContractRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmartContractRegistration.Marshal(b, m, deterministic)
}
func (m *SmartContractRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmartContractRegistration.Merge(m, src)
}
func (m *SmartContractRegistration) XXX_Size() int {
	return xxx_messageInfo_SmartContractRegistration.Size(m)
}
func (m *SmartContractRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_SmartContractRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_SmartContractRegistration proto.InternalMessageInfo

func (m *SmartContractRegistration) GetCategory() int32 {
	if m != nil {
		return m.Category
	}
	return 0
}

func (m *SmartContractRegistration) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *SmartContractRegistration) GetCodeHash() *Hash {
	if m != nil {
		return m.CodeHash
	}
	return nil
}

type HashList struct {
	Values               []*Hash  `protobuf:"bytes,1,rep,name=Values,proto3" json:"Values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashList) Reset()         { *m = HashList{} }
func (m *HashList) String() string { return proto.CompactTextString(m) }
func (*HashList) ProtoMessage()    {}
func (*HashList) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{6}
}

func (m *HashList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashList.Unmarshal(m, b)
}
func (m *HashList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashList.Marshal(b, m, deterministic)
}
func (m *HashList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashList.Merge(m, src)
}
func (m *HashList) XXX_Size() int {
	return xxx_messageInfo_HashList.Size(m)
}
func (m *HashList) XXX_DiscardUnknown() {
	xxx_messageInfo_HashList.DiscardUnknown(m)
}

var xxx_messageInfo_HashList proto.InternalMessageInfo

func (m *HashList) GetValues() []*Hash {
	if m != nil {
		return m.Values
	}
	return nil
}

type TransactionExecutingStateSet struct {
	Version              int64             `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Writes               map[string][]byte `protobuf:"bytes,2,rep,name=Writes,proto3" json:"Writes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Reads                map[string]bool   `protobuf:"bytes,3,rep,name=Reads,proto3" json:"Reads,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TransactionExecutingStateSet) Reset()         { *m = TransactionExecutingStateSet{} }
func (m *TransactionExecutingStateSet) String() string { return proto.CompactTextString(m) }
func (*TransactionExecutingStateSet) ProtoMessage()    {}
func (*TransactionExecutingStateSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{7}
}

func (m *TransactionExecutingStateSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionExecutingStateSet.Unmarshal(m, b)
}
func (m *TransactionExecutingStateSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionExecutingStateSet.Marshal(b, m, deterministic)
}
func (m *TransactionExecutingStateSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionExecutingStateSet.Merge(m, src)
}
func (m *TransactionExecutingStateSet) XXX_Size() int {
	return xxx_messageInfo_TransactionExecutingStateSet.Size(m)
}
func (m *TransactionExecutingStateSet) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionExecutingStateSet.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionExecutingStateSet proto.InternalMessageInfo

func (m *TransactionExecutingStateSet) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *TransactionExecutingStateSet) GetWrites() map[string][]byte {
	if m != nil {
		return m.Writes
	}
	return nil
}

func (m *TransactionExecutingStateSet) GetReads() map[string]bool {
	if m != nil {
		return m.Reads
	}
	return nil
}

type ActionResult struct {
	Success              bool     `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=ErrorMessage,proto3" json:"ErrorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActionResult) Reset()         { *m = ActionResult{} }
func (m *ActionResult) String() string { return proto.CompactTextString(m) }
func (*ActionResult) ProtoMessage()    {}
func (*ActionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{8}
}

func (m *ActionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionResult.Unmarshal(m, b)
}
func (m *ActionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionResult.Marshal(b, m, deterministic)
}
func (m *ActionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionResult.Merge(m, src)
}
func (m *ActionResult) XXX_Size() int {
	return xxx_messageInfo_ActionResult.Size(m)
}
func (m *ActionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionResult.DiscardUnknown(m)
}

var xxx_messageInfo_ActionResult proto.InternalMessageInfo

func (m *ActionResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ActionResult) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type Address struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{9}
}

func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (m *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(m, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Hash struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hash) Reset()         { *m = Hash{} }
func (m *Hash) String() string { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()    {}
func (*Hash) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{10}
}

func (m *Hash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hash.Unmarshal(m, b)
}
func (m *Hash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hash.Marshal(b, m, deterministic)
}
func (m *Hash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hash.Merge(m, src)
}
func (m *Hash) XXX_Size() int {
	return xxx_messageInfo_Hash.Size(m)
}
func (m *Hash) XXX_DiscardUnknown() {
	xxx_messageInfo_Hash.DiscardUnknown(m)
}

var xxx_messageInfo_Hash proto.InternalMessageInfo

func (m *Hash) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type SInt32Value struct {
	Value                int32    `protobuf:"zigzag32,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SInt32Value) Reset()         { *m = SInt32Value{} }
func (m *SInt32Value) String() string { return proto.CompactTextString(m) }
func (*SInt32Value) ProtoMessage()    {}
func (*SInt32Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{11}
}

func (m *SInt32Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SInt32Value.Unmarshal(m, b)
}
func (m *SInt32Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SInt32Value.Marshal(b, m, deterministic)
}
func (m *SInt32Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SInt32Value.Merge(m, src)
}
func (m *SInt32Value) XXX_Size() int {
	return xxx_messageInfo_SInt32Value.Size(m)
}
func (m *SInt32Value) XXX_DiscardUnknown() {
	xxx_messageInfo_SInt32Value.DiscardUnknown(m)
}

var xxx_messageInfo_SInt32Value proto.InternalMessageInfo

func (m *SInt32Value) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type SInt64Value struct {
	Value                int64    `protobuf:"zigzag64,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SInt64Value) Reset()         { *m = SInt64Value{} }
func (m *SInt64Value) String() string { return proto.CompactTextString(m) }
func (*SInt64Value) ProtoMessage()    {}
func (*SInt64Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_d466aeb18442da3b, []int{12}
}

func (m *SInt64Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SInt64Value.Unmarshal(m, b)
}
func (m *SInt64Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SInt64Value.Marshal(b, m, deterministic)
}
func (m *SInt64Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SInt64Value.Merge(m, src)
}
func (m *SInt64Value) XXX_Size() int {
	return xxx_messageInfo_SInt64Value.Size(m)
}
func (m *SInt64Value) XXX_DiscardUnknown() {
	xxx_messageInfo_SInt64Value.DiscardUnknown(m)
}

var xxx_messageInfo_SInt64Value proto.InternalMessageInfo

func (m *SInt64Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("aelf.TransactionResultStatus", TransactionResultStatus_name, TransactionResultStatus_value)
	proto.RegisterType((*Transaction)(nil), "aelf.Transaction")
	proto.RegisterType((*StatePath)(nil), "aelf.StatePath")
	proto.RegisterType((*ScopedStatePath)(nil), "aelf.ScopedStatePath")
	proto.RegisterType((*TransactionResult)(nil), "aelf.TransactionResult")
	proto.RegisterType((*LogEvent)(nil), "aelf.LogEvent")
	proto.RegisterType((*SmartContractRegistration)(nil), "aelf.SmartContractRegistration")
	proto.RegisterType((*HashList)(nil), "aelf.HashList")
	proto.RegisterType((*TransactionExecutingStateSet)(nil), "aelf.TransactionExecutingStateSet")
	proto.RegisterMapType((map[string]bool)(nil), "aelf.TransactionExecutingStateSet.ReadsEntry")
	proto.RegisterMapType((map[string][]byte)(nil), "aelf.TransactionExecutingStateSet.WritesEntry")
	proto.RegisterType((*ActionResult)(nil), "aelf.ActionResult")
	proto.RegisterType((*Address)(nil), "aelf.Address")
	proto.RegisterType((*Hash)(nil), "aelf.Hash")
	proto.RegisterType((*SInt32Value)(nil), "aelf.SInt32Value")
	proto.RegisterType((*SInt64Value)(nil), "aelf.SInt64Value")
}

func init() { proto.RegisterFile("aelf/core.proto", fileDescriptor_d466aeb18442da3b) }

var fileDescriptor_d466aeb18442da3b = []byte{
	// 881 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xe7, 0x7c, 0xe7, 0x7f, 0x63, 0x27, 0x71, 0x96, 0x0a, 0x8e, 0xa8, 0xa1, 0xee, 0x55, 0x2a,
	0x16, 0x12, 0x4e, 0x95, 0x02, 0x2a, 0xbc, 0x25, 0xc1, 0x11, 0x91, 0x92, 0x28, 0x5a, 0x87, 0xf2,
	0x58, 0x6d, 0x7c, 0xe3, 0xcb, 0xa9, 0xf6, 0xae, 0xb5, 0xbb, 0x57, 0x39, 0x8f, 0x7c, 0x03, 0xbe,
	0x04, 0x2f, 0x7c, 0x04, 0x3e, 0x18, 0xcf, 0x68, 0x77, 0xef, 0xec, 0x73, 0xce, 0x40, 0x9f, 0xbc,
	0x33, 0xf3, 0x9b, 0xb9, 0xd9, 0xf9, 0xfd, 0x66, 0x0d, 0x7b, 0x0c, 0x67, 0xd3, 0xa3, 0x89, 0x90,
	0x38, 0x5c, 0x48, 0xa1, 0x05, 0x09, 0x8c, 0xe3, 0xe0, 0x59, 0x22, 0x44, 0x32, 0xc3, 0x23, 0xeb,
	0xbb, 0xcb, 0xa6, 0x47, 0x3a, 0x9d, 0xa3, 0xd2, 0x6c, 0xbe, 0x70, 0xb0, 0xe8, 0x6f, 0x0f, 0x3a,
	0xb7, 0x92, 0x71, 0xc5, 0x26, 0x3a, 0x15, 0x9c, 0x3c, 0x87, 0xe0, 0x5c, 0x8a, 0x79, 0xe8, 0xf5,
	0xbd, 0x41, 0xe7, 0x78, 0x67, 0x68, 0xaa, 0x0c, 0x4f, 0xe2, 0x58, 0xa2, 0x52, 0xd4, 0x86, 0xc8,
	0x21, 0xd4, 0x6e, 0x45, 0x58, 0xdb, 0x06, 0xa8, 0xdd, 0x0a, 0xf2, 0x12, 0x76, 0x29, 0x4e, 0x4f,
	0x67, 0x62, 0xf2, 0xfe, 0x3a, 0x9b, 0xdf, 0xa1, 0x0c, 0xfd, 0xbe, 0x37, 0xf0, 0xe9, 0x23, 0x6f,
	0x19, 0x77, 0x23, 0x71, 0x9a, 0x2e, 0xc3, 0xa0, 0xef, 0x0d, 0xba, 0xf4, 0x91, 0x97, 0x7c, 0x09,
	0x70, 0x85, 0xfa, 0x5e, 0xc4, 0xd7, 0x6c, 0x8e, 0x61, 0xbd, 0xef, 0x0d, 0xda, 0xb4, 0xe4, 0x21,
	0x9f, 0x41, 0xe3, 0x86, 0x49, 0x36, 0x57, 0x61, 0xc3, 0xe6, 0xe7, 0x16, 0x39, 0x84, 0xf6, 0x38,
	0x4d, 0x38, 0xd3, 0x99, 0xc4, 0xf0, 0xf7, 0x6b, 0x1b, 0x5b, 0x7b, 0xa2, 0xe7, 0xd0, 0x1e, 0x6b,
	0xa6, 0xf1, 0x86, 0xe9, 0x7b, 0xf2, 0x04, 0xea, 0x0b, 0x26, 0xb5, 0x0a, 0xbd, 0xbe, 0x3f, 0x68,
	0x53, 0x67, 0x44, 0xef, 0x60, 0x6f, 0x3c, 0x11, 0x0b, 0x8c, 0xd7, 0xc0, 0xaf, 0xa0, 0xc9, 0xdc,
	0x5d, 0xb7, 0x4f, 0xa8, 0x88, 0x92, 0x17, 0x10, 0x2c, 0x98, 0xbe, 0xcf, 0xc7, 0xb4, 0xe7, 0x50,
	0xab, 0x3a, 0xd4, 0x06, 0xa3, 0x3f, 0x02, 0xd8, 0x2f, 0x0d, 0x9f, 0xa2, 0xca, 0x66, 0x9a, 0xbc,
	0x82, 0x9d, 0x92, 0xf3, 0x22, 0xce, 0xbf, 0x04, 0xae, 0xc6, 0xcf, 0x4c, 0xdd, 0xd3, 0x4d, 0x00,
	0xf9, 0x0e, 0x1a, 0xa6, 0x74, 0xa6, 0xec, 0xe7, 0x76, 0x8f, 0x0f, 0x1d, 0xb4, 0x52, 0xda, 0x81,
	0x68, 0x0e, 0x26, 0x11, 0x04, 0x97, 0x22, 0x51, 0xa1, 0xdf, 0xf7, 0x07, 0x9d, 0xe3, 0x5d, 0x97,
	0x74, 0x29, 0x92, 0xd1, 0x07, 0xe4, 0x9a, 0xda, 0x98, 0x99, 0xcc, 0xe9, 0x4c, 0x88, 0x79, 0x4e,
	0x8e, 0x33, 0x48, 0x1f, 0x3a, 0x14, 0x75, 0x26, 0xf9, 0x5b, 0x36, 0xcb, 0x1c, 0x29, 0x5d, 0x5a,
	0x76, 0x19, 0x44, 0x59, 0x02, 0x0d, 0x2b, 0x81, 0xb2, 0x8b, 0x0c, 0xa0, 0x6d, 0x4d, 0x73, 0xa1,
	0xb0, 0x59, 0xb9, 0xe2, 0x3a, 0x68, 0x7a, 0xb8, 0xe0, 0x31, 0x2e, 0xc3, 0x56, 0xdf, 0x1b, 0xd4,
	0xa9, 0x33, 0x4c, 0xbe, 0x9d, 0xa7, 0xcd, 0x6f, 0x57, 0xf3, 0x57, 0x41, 0x32, 0x82, 0x27, 0x3f,
	0xe1, 0x14, 0xa5, 0xc4, 0xb8, 0x34, 0x12, 0x15, 0x82, 0xbd, 0xf7, 0x7e, 0x75, 0x58, 0x5b, 0xe1,
	0x86, 0x97, 0x95, 0x7f, 0x69, 0x78, 0xe9, 0x54, 0x79, 0xd9, 0x00, 0x98, 0xc6, 0x47, 0x52, 0x0a,
	0x19, 0x76, 0xad, 0x6a, 0x9d, 0x41, 0x5e, 0xc1, 0xa7, 0x14, 0x59, 0xcc, 0xee, 0x66, 0x58, 0x1e,
	0xe2, 0x8e, 0xc5, 0x6c, 0x0b, 0x45, 0xbf, 0x79, 0xd0, 0x2a, 0x78, 0x31, 0x12, 0x3c, 0xf9, 0x4f,
	0x09, 0xe6, 0x07, 0x42, 0x20, 0xb0, 0x2b, 0x53, 0xb3, 0x85, 0xed, 0x99, 0x84, 0xd0, 0xb4, 0xd3,
	0xc3, 0xd8, 0xb2, 0xde, 0xa5, 0x85, 0x69, 0xd6, 0xec, 0x5a, 0xf0, 0x22, 0xe8, 0xd8, 0x2e, 0x79,
	0x22, 0x05, 0x5f, 0x8c, 0xe7, 0x4c, 0xea, 0x33, 0xc1, 0xb5, 0x64, 0x13, 0x4d, 0x31, 0x49, 0x95,
	0x96, 0xcc, 0xbe, 0x1a, 0x07, 0xd0, 0x3a, 0x63, 0x1a, 0x13, 0x21, 0x1f, 0x6c, 0x53, 0x75, 0xba,
	0xb2, 0x4d, 0x1b, 0x67, 0x22, 0x76, 0x6d, 0x74, 0xa9, 0x3d, 0x93, 0x97, 0xd0, 0x32, 0xbf, 0x96,
	0x3a, 0xbf, 0x32, 0xc5, 0x55, 0x2c, 0x1a, 0x42, 0xcb, 0xfc, 0x5e, 0xa6, 0x4a, 0x93, 0x08, 0x1a,
	0x76, 0x1a, 0x6e, 0x49, 0x37, 0x33, 0xf2, 0x48, 0xf4, 0x57, 0x0d, 0x9e, 0x96, 0x38, 0x1b, 0x2d,
	0x71, 0x92, 0xe9, 0x94, 0x27, 0x56, 0x0b, 0x63, 0xd4, 0xe6, 0xfe, 0x6f, 0x51, 0xaa, 0x54, 0x70,
	0xdb, 0xa7, 0x4f, 0x0b, 0x93, 0x9c, 0x43, 0xe3, 0x57, 0x99, 0x6a, 0x34, 0x3b, 0x64, 0xca, 0x0f,
	0x2b, 0xb2, 0xa8, 0x54, 0x1b, 0xba, 0x84, 0x11, 0xd7, 0xf2, 0x81, 0xe6, 0xd9, 0xe4, 0x0c, 0xea,
	0x86, 0xc2, 0x62, 0xab, 0xbe, 0xf9, 0x88, 0x32, 0x16, 0xef, 0xaa, 0xb8, 0xdc, 0x83, 0x1f, 0xa0,
	0x53, 0xaa, 0x4d, 0x7a, 0xe0, 0xbf, 0x47, 0x37, 0xd9, 0x36, 0x35, 0x47, 0xa3, 0xac, 0x0f, 0x56,
	0x35, 0x6e, 0xaa, 0xce, 0xf8, 0xb1, 0xf6, 0xc6, 0x3b, 0x78, 0x03, 0xb0, 0xae, 0xf7, 0x7f, 0x99,
	0xad, 0x52, 0x66, 0x74, 0x09, 0xdd, 0x93, 0xf2, 0x3b, 0x14, 0x42, 0x73, 0x9c, 0x4d, 0x26, 0x85,
	0xd0, 0x5a, 0xb4, 0x30, 0x49, 0x04, 0x5d, 0x2b, 0xe5, 0x2b, 0x54, 0x8a, 0x25, 0x85, 0xc2, 0x36,
	0x7c, 0xd1, 0xb3, 0x95, 0x4c, 0xcd, 0x27, 0x9d, 0xc4, 0x3d, 0xd7, 0xac, 0x13, 0xf5, 0x53, 0x08,
	0x8a, 0xed, 0xde, 0x12, 0x7d, 0x01, 0x9d, 0xf1, 0x05, 0xd7, 0xaf, 0x8f, 0xdd, 0x73, 0xb2, 0x01,
	0xda, 0x7f, 0x04, 0xfa, 0xfe, 0xdb, 0x2d, 0x20, 0x92, 0x83, 0xbe, 0x7e, 0x07, 0x9f, 0xff, 0xcb,
	0x43, 0x48, 0x76, 0x8d, 0xe6, 0xf5, 0x68, 0x99, 0x2a, 0x8d, 0x71, 0xef, 0x13, 0xd2, 0x81, 0xe6,
	0x0d, 0xf2, 0x38, 0xe5, 0x49, 0xcf, 0x23, 0x00, 0x8d, 0x73, 0x96, 0xce, 0x30, 0xee, 0xd5, 0x48,
	0x1b, 0xea, 0x57, 0x29, 0xc7, 0xb8, 0xe7, 0x93, 0x1e, 0x74, 0x7f, 0xe1, 0x68, 0x39, 0x34, 0x6b,
	0xda, 0x0b, 0x4e, 0x77, 0xfe, 0xac, 0xc1, 0xc9, 0xc8, 0x90, 0xfc, 0xb0, 0x40, 0x75, 0xd7, 0xb0,
	0x7f, 0xac, 0xaf, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x7c, 0xab, 0x00, 0x92, 0x07, 0x00,
	0x00,
}