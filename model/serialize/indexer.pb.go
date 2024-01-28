// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: indexer.proto

package serialize

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProtoToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tick        string `protobuf:"bytes,1,opt,name=tick,proto3" json:"tick,omitempty"`
	Number      uint64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Precision   uint32 `protobuf:"varint,3,opt,name=precision,proto3" json:"precision,omitempty"`
	Max         string `protobuf:"bytes,4,opt,name=max,proto3" json:"max,omitempty"`
	Limit       string `protobuf:"bytes,5,opt,name=limit,proto3" json:"limit,omitempty"`
	Minted      string `protobuf:"bytes,6,opt,name=minted,proto3" json:"minted,omitempty"`
	Progress    uint32 `protobuf:"varint,7,opt,name=progress,proto3" json:"progress,omitempty"`
	Holders     uint32 `protobuf:"varint,8,opt,name=holders,proto3" json:"holders,omitempty"`
	Trxs        uint32 `protobuf:"varint,9,opt,name=trxs,proto3" json:"trxs,omitempty"`
	CreatedAt   uint64 `protobuf:"varint,10,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	CompletedAt uint64 `protobuf:"varint,11,opt,name=completedAt,proto3" json:"completedAt,omitempty"`
	Hash        []byte `protobuf:"bytes,12,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *ProtoToken) Reset() {
	*x = ProtoToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoToken) ProtoMessage() {}

func (x *ProtoToken) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoToken.ProtoReflect.Descriptor instead.
func (*ProtoToken) Descriptor() ([]byte, []int) {
	return file_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoToken) GetTick() string {
	if x != nil {
		return x.Tick
	}
	return ""
}

func (x *ProtoToken) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ProtoToken) GetPrecision() uint32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *ProtoToken) GetMax() string {
	if x != nil {
		return x.Max
	}
	return ""
}

func (x *ProtoToken) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *ProtoToken) GetMinted() string {
	if x != nil {
		return x.Minted
	}
	return ""
}

func (x *ProtoToken) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *ProtoToken) GetHolders() uint32 {
	if x != nil {
		return x.Holders
	}
	return 0
}

func (x *ProtoToken) GetTrxs() uint32 {
	if x != nil {
		return x.Trxs
	}
	return 0
}

func (x *ProtoToken) GetCreatedAt() uint64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ProtoToken) GetCompletedAt() uint64 {
	if x != nil {
		return x.CompletedAt
	}
	return 0
}

func (x *ProtoToken) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

type ProtoRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Number    uint64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Tick      string `protobuf:"bytes,3,opt,name=tick,proto3" json:"tick,omitempty"`
	From      []byte `protobuf:"bytes,4,opt,name=from,proto3" json:"from,omitempty"`
	To        []byte `protobuf:"bytes,5,opt,name=to,proto3" json:"to,omitempty"`
	Operation string `protobuf:"bytes,6,opt,name=operation,proto3" json:"operation,omitempty"`
	Precision uint32 `protobuf:"varint,7,opt,name=precision,proto3" json:"precision,omitempty"`
	Amount    string `protobuf:"bytes,8,opt,name=amount,proto3" json:"amount,omitempty"`
	Limit     string `protobuf:"bytes,9,opt,name=limit,proto3" json:"limit,omitempty"`
	Hash      []byte `protobuf:"bytes,10,opt,name=hash,proto3" json:"hash,omitempty"`
	Block     uint64 `protobuf:"varint,11,opt,name=block,proto3" json:"block,omitempty"`
	Timestamp uint64 `protobuf:"varint,12,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Valid     int32  `protobuf:"varint,13,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *ProtoRecord) Reset() {
	*x = ProtoRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoRecord) ProtoMessage() {}

func (x *ProtoRecord) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoRecord.ProtoReflect.Descriptor instead.
func (*ProtoRecord) Descriptor() ([]byte, []int) {
	return file_indexer_proto_rawDescGZIP(), []int{1}
}

func (x *ProtoRecord) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProtoRecord) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *ProtoRecord) GetTick() string {
	if x != nil {
		return x.Tick
	}
	return ""
}

func (x *ProtoRecord) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *ProtoRecord) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *ProtoRecord) GetOperation() string {
	if x != nil {
		return x.Operation
	}
	return ""
}

func (x *ProtoRecord) GetPrecision() uint32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *ProtoRecord) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *ProtoRecord) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *ProtoRecord) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *ProtoRecord) GetBlock() uint64 {
	if x != nil {
		return x.Block
	}
	return 0
}

func (x *ProtoRecord) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ProtoRecord) GetValid() int32 {
	if x != nil {
		return x.Valid
	}
	return 0
}

type ProtoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InsId     string `protobuf:"bytes,1,opt,name=insId,proto3" json:"insId,omitempty"`
	Owner     []byte `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Exchange  []byte `protobuf:"bytes,3,opt,name=exchange,proto3" json:"exchange,omitempty"`
	Tick      string `protobuf:"bytes,4,opt,name=tick,proto3" json:"tick,omitempty"`
	Amount    string `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
	Precision uint32 `protobuf:"varint,6,opt,name=precision,proto3" json:"precision,omitempty"`
}

func (x *ProtoList) Reset() {
	*x = ProtoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoList) ProtoMessage() {}

func (x *ProtoList) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoList.ProtoReflect.Descriptor instead.
func (*ProtoList) Descriptor() ([]byte, []int) {
	return file_indexer_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoList) GetInsId() string {
	if x != nil {
		return x.InsId
	}
	return ""
}

func (x *ProtoList) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *ProtoList) GetExchange() []byte {
	if x != nil {
		return x.Exchange
	}
	return nil
}

func (x *ProtoList) GetTick() string {
	if x != nil {
		return x.Tick
	}
	return ""
}

func (x *ProtoList) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *ProtoList) GetPrecision() uint32 {
	if x != nil {
		return x.Precision
	}
	return 0
}

type TickBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tick   string `protobuf:"bytes,1,opt,name=tick,proto3" json:"tick,omitempty"`
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TickBalance) Reset() {
	*x = TickBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TickBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickBalance) ProtoMessage() {}

func (x *TickBalance) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TickBalance.ProtoReflect.Descriptor instead.
func (*TickBalance) Descriptor() ([]byte, []int) {
	return file_indexer_proto_rawDescGZIP(), []int{3}
}

func (x *TickBalance) GetTick() string {
	if x != nil {
		return x.Tick
	}
	return ""
}

func (x *TickBalance) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type UserBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  []byte         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balances []*TickBalance `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances,omitempty"`
}

func (x *UserBalance) Reset() {
	*x = UserBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBalance) ProtoMessage() {}

func (x *UserBalance) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBalance.ProtoReflect.Descriptor instead.
func (*UserBalance) Descriptor() ([]byte, []int) {
	return file_indexer_proto_rawDescGZIP(), []int{4}
}

func (x *UserBalance) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *UserBalance) GetBalances() []*TickBalance {
	if x != nil {
		return x.Balances
	}
	return nil
}

type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block    uint64         `protobuf:"varint,1,opt,name=block,proto3" json:"block,omitempty"`
	Number   uint64         `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	RecordId uint64         `protobuf:"varint,3,opt,name=recordId,proto3" json:"recordId,omitempty"`
	Tokens   []*ProtoToken  `protobuf:"bytes,4,rep,name=tokens,proto3" json:"tokens,omitempty"`
	Lists    []*ProtoList   `protobuf:"bytes,5,rep,name=lists,proto3" json:"lists,omitempty"`
	Balances []*UserBalance `protobuf:"bytes,6,rep,name=balances,proto3" json:"balances,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_indexer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_indexer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_indexer_proto_rawDescGZIP(), []int{5}
}

func (x *Snapshot) GetBlock() uint64 {
	if x != nil {
		return x.Block
	}
	return 0
}

func (x *Snapshot) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Snapshot) GetRecordId() uint64 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

func (x *Snapshot) GetTokens() []*ProtoToken {
	if x != nil {
		return x.Tokens
	}
	return nil
}

func (x *Snapshot) GetLists() []*ProtoList {
	if x != nil {
		return x.Lists
	}
	return nil
}

func (x *Snapshot) GetBalances() []*UserBalance {
	if x != nil {
		return x.Balances
	}
	return nil
}

var File_indexer_proto protoreflect.FileDescriptor

var file_indexer_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xb4, 0x02, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x61,
	0x78, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x74, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x69, 0x6e, 0x74, 0x65, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x68,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x68, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x72, 0x78, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x72, 0x78, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0xb5, 0x02,
	0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a,
	0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1c, 0x0a,
	0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x73, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x73, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x39, 0x0a, 0x0b, 0x54, 0x69, 0x63, 0x6b, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x57, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x54, 0x69, 0x63, 0x6b, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x22, 0xd7, 0x01, 0x0a, 0x08, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64,
	0x12, 0x29, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x05, 0x6c,
	0x69, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x05, 0x6c, 0x69,
	0x73, 0x74, 0x73, 0x12, 0x2e, 0x0a, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x73, 0x42, 0x1e, 0x5a, 0x1c, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_indexer_proto_rawDescOnce sync.Once
	file_indexer_proto_rawDescData = file_indexer_proto_rawDesc
)

func file_indexer_proto_rawDescGZIP() []byte {
	file_indexer_proto_rawDescOnce.Do(func() {
		file_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_indexer_proto_rawDescData)
	})
	return file_indexer_proto_rawDescData
}

var file_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_indexer_proto_goTypes = []interface{}{
	(*ProtoToken)(nil),  // 0: model.ProtoToken
	(*ProtoRecord)(nil), // 1: model.ProtoRecord
	(*ProtoList)(nil),   // 2: model.ProtoList
	(*TickBalance)(nil), // 3: model.TickBalance
	(*UserBalance)(nil), // 4: model.UserBalance
	(*Snapshot)(nil),    // 5: model.Snapshot
}
var file_indexer_proto_depIdxs = []int32{
	3, // 0: model.UserBalance.balances:type_name -> model.TickBalance
	0, // 1: model.Snapshot.tokens:type_name -> model.ProtoToken
	2, // 2: model.Snapshot.lists:type_name -> model.ProtoList
	4, // 3: model.Snapshot.balances:type_name -> model.UserBalance
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_indexer_proto_init() }
func file_indexer_proto_init() {
	if File_indexer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_indexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoToken); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_indexer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_indexer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_indexer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TickBalance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_indexer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBalance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_indexer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_indexer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_indexer_proto_goTypes,
		DependencyIndexes: file_indexer_proto_depIdxs,
		MessageInfos:      file_indexer_proto_msgTypes,
	}.Build()
	File_indexer_proto = out.File
	file_indexer_proto_rawDesc = nil
	file_indexer_proto_goTypes = nil
	file_indexer_proto_depIdxs = nil
}
