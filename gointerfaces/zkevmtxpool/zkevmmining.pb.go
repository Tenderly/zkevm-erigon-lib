// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: zkevmtxpool/zkevmmining.proto

package zkevmtxpool

import (
	types "github.com/tenderly/zkevm-erigon-lib/gointerfaces/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OnPendingBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnPendingBlockRequest) Reset() {
	*x = OnPendingBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnPendingBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnPendingBlockRequest) ProtoMessage() {}

func (x *OnPendingBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnPendingBlockRequest.ProtoReflect.Descriptor instead.
func (*OnPendingBlockRequest) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{0}
}

type OnPendingBlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RplBlock []byte `protobuf:"bytes,1,opt,name=rpl_block,json=rplBlock,proto3" json:"rpl_block,omitempty"`
}

func (x *OnPendingBlockReply) Reset() {
	*x = OnPendingBlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnPendingBlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnPendingBlockReply) ProtoMessage() {}

func (x *OnPendingBlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnPendingBlockReply.ProtoReflect.Descriptor instead.
func (*OnPendingBlockReply) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{1}
}

func (x *OnPendingBlockReply) GetRplBlock() []byte {
	if x != nil {
		return x.RplBlock
	}
	return nil
}

type OnMinedBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnMinedBlockRequest) Reset() {
	*x = OnMinedBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnMinedBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnMinedBlockRequest) ProtoMessage() {}

func (x *OnMinedBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnMinedBlockRequest.ProtoReflect.Descriptor instead.
func (*OnMinedBlockRequest) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{2}
}

type OnMinedBlockReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RplBlock []byte `protobuf:"bytes,1,opt,name=rpl_block,json=rplBlock,proto3" json:"rpl_block,omitempty"`
}

func (x *OnMinedBlockReply) Reset() {
	*x = OnMinedBlockReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnMinedBlockReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnMinedBlockReply) ProtoMessage() {}

func (x *OnMinedBlockReply) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnMinedBlockReply.ProtoReflect.Descriptor instead.
func (*OnMinedBlockReply) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{3}
}

func (x *OnMinedBlockReply) GetRplBlock() []byte {
	if x != nil {
		return x.RplBlock
	}
	return nil
}

type OnPendingLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OnPendingLogsRequest) Reset() {
	*x = OnPendingLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnPendingLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnPendingLogsRequest) ProtoMessage() {}

func (x *OnPendingLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnPendingLogsRequest.ProtoReflect.Descriptor instead.
func (*OnPendingLogsRequest) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{4}
}

type OnPendingLogsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RplLogs []byte `protobuf:"bytes,1,opt,name=rpl_logs,json=rplLogs,proto3" json:"rpl_logs,omitempty"`
}

func (x *OnPendingLogsReply) Reset() {
	*x = OnPendingLogsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OnPendingLogsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnPendingLogsReply) ProtoMessage() {}

func (x *OnPendingLogsReply) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnPendingLogsReply.ProtoReflect.Descriptor instead.
func (*OnPendingLogsReply) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{5}
}

func (x *OnPendingLogsReply) GetRplLogs() []byte {
	if x != nil {
		return x.RplLogs
	}
	return nil
}

type GetWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetWorkRequest) Reset() {
	*x = GetWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkRequest) ProtoMessage() {}

func (x *GetWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkRequest.ProtoReflect.Descriptor instead.
func (*GetWorkRequest) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{6}
}

type GetWorkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeaderHash  string `protobuf:"bytes,1,opt,name=header_hash,json=headerHash,proto3" json:"header_hash,omitempty"`    // 32 bytes hex encoded current block header pow-hash
	SeedHash    string `protobuf:"bytes,2,opt,name=seed_hash,json=seedHash,proto3" json:"seed_hash,omitempty"`          // 32 bytes hex encoded seed hash used for DAG
	Target      string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`                              // 32 bytes hex encoded boundary condition ("target"), 2^256/difficulty
	BlockNumber string `protobuf:"bytes,4,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"` // hex encoded block number
}

func (x *GetWorkReply) Reset() {
	*x = GetWorkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkReply) ProtoMessage() {}

func (x *GetWorkReply) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkReply.ProtoReflect.Descriptor instead.
func (*GetWorkReply) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{7}
}

func (x *GetWorkReply) GetHeaderHash() string {
	if x != nil {
		return x.HeaderHash
	}
	return ""
}

func (x *GetWorkReply) GetSeedHash() string {
	if x != nil {
		return x.SeedHash
	}
	return ""
}

func (x *GetWorkReply) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *GetWorkReply) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

type SubmitWorkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNonce []byte `protobuf:"bytes,1,opt,name=block_nonce,json=blockNonce,proto3" json:"block_nonce,omitempty"`
	PowHash    []byte `protobuf:"bytes,2,opt,name=pow_hash,json=powHash,proto3" json:"pow_hash,omitempty"`
	Digest     []byte `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *SubmitWorkRequest) Reset() {
	*x = SubmitWorkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitWorkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitWorkRequest) ProtoMessage() {}

func (x *SubmitWorkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitWorkRequest.ProtoReflect.Descriptor instead.
func (*SubmitWorkRequest) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{8}
}

func (x *SubmitWorkRequest) GetBlockNonce() []byte {
	if x != nil {
		return x.BlockNonce
	}
	return nil
}

func (x *SubmitWorkRequest) GetPowHash() []byte {
	if x != nil {
		return x.PowHash
	}
	return nil
}

func (x *SubmitWorkRequest) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

type SubmitWorkReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SubmitWorkReply) Reset() {
	*x = SubmitWorkReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitWorkReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitWorkReply) ProtoMessage() {}

func (x *SubmitWorkReply) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitWorkReply.ProtoReflect.Descriptor instead.
func (*SubmitWorkReply) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{9}
}

func (x *SubmitWorkReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type SubmitHashRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate uint64 `protobuf:"varint,1,opt,name=rate,proto3" json:"rate,omitempty"`
	Id   []byte `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubmitHashRateRequest) Reset() {
	*x = SubmitHashRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitHashRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitHashRateRequest) ProtoMessage() {}

func (x *SubmitHashRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitHashRateRequest.ProtoReflect.Descriptor instead.
func (*SubmitHashRateRequest) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{10}
}

func (x *SubmitHashRateRequest) GetRate() uint64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *SubmitHashRateRequest) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

type SubmitHashRateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SubmitHashRateReply) Reset() {
	*x = SubmitHashRateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitHashRateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitHashRateReply) ProtoMessage() {}

func (x *SubmitHashRateReply) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitHashRateReply.ProtoReflect.Descriptor instead.
func (*SubmitHashRateReply) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{11}
}

func (x *SubmitHashRateReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type HashRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HashRateRequest) Reset() {
	*x = HashRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashRateRequest) ProtoMessage() {}

func (x *HashRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashRateRequest.ProtoReflect.Descriptor instead.
func (*HashRateRequest) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{12}
}

type HashRateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashRate uint64 `protobuf:"varint,1,opt,name=hash_rate,json=hashRate,proto3" json:"hash_rate,omitempty"`
}

func (x *HashRateReply) Reset() {
	*x = HashRateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashRateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashRateReply) ProtoMessage() {}

func (x *HashRateReply) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashRateReply.ProtoReflect.Descriptor instead.
func (*HashRateReply) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{13}
}

func (x *HashRateReply) GetHashRate() uint64 {
	if x != nil {
		return x.HashRate
	}
	return 0
}

type MiningRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MiningRequest) Reset() {
	*x = MiningRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiningRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiningRequest) ProtoMessage() {}

func (x *MiningRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiningRequest.ProtoReflect.Descriptor instead.
func (*MiningRequest) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{14}
}

type MiningReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Running bool `protobuf:"varint,2,opt,name=running,proto3" json:"running,omitempty"`
}

func (x *MiningReply) Reset() {
	*x = MiningReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MiningReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MiningReply) ProtoMessage() {}

func (x *MiningReply) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmtxpool_zkevmmining_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MiningReply.ProtoReflect.Descriptor instead.
func (*MiningReply) Descriptor() ([]byte, []int) {
	return file_zkevmtxpool_zkevmmining_proto_rawDescGZIP(), []int{15}
}

func (x *MiningReply) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *MiningReply) GetRunning() bool {
	if x != nil {
		return x.Running
	}
	return false
}

var File_zkevmtxpool_zkevmmining_proto protoreflect.FileDescriptor

var file_zkevmtxpool_zkevmmining_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x7a, 0x6b,
	0x65, 0x76, 0x6d, 0x6d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x7a, 0x6b, 0x65, 0x76, 0x6d,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x17, 0x0a, 0x15, 0x4f, 0x6e, 0x50, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x32, 0x0a, 0x13, 0x4f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x70, 0x6c, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x70, 0x6c, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x22, 0x15, 0x0a, 0x13, 0x4f, 0x6e, 0x4d, 0x69, 0x6e, 0x65, 0x64, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x11, 0x4f, 0x6e,
	0x4d, 0x69, 0x6e, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x70, 0x6c, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x72, 0x70, 0x6c, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x16, 0x0a, 0x14,
	0x4f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x12, 0x4f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x70,
	0x6c, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x70,
	0x6c, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x87, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x65,
	0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x65, 0x64, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x67, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6f, 0x77, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x6f, 0x77, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x3b, 0x0a,
	0x15, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x22, 0x25, 0x0a, 0x13, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f,
	0x6b, 0x22, 0x11, 0x0a, 0x0f, 0x48, 0x61, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x2c, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x52, 0x61,
	0x74, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x41, 0x0a, 0x0b, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72,
	0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x32, 0xb7, 0x05, 0x0a, 0x06, 0x4d, 0x69, 0x6e, 0x69, 0x6e,
	0x67, 0x12, 0x3b, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x58,
	0x0a, 0x0e, 0x4f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x22, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f,
	0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f,
	0x6f, 0x6c, 0x2e, 0x4f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x12, 0x52, 0x0a, 0x0c, 0x4f, 0x6e, 0x4d, 0x69,
	0x6e, 0x65, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x20, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d,
	0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x6e, 0x4d, 0x69, 0x6e, 0x65, 0x64, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x7a, 0x6b, 0x65,
	0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x6e, 0x4d, 0x69, 0x6e, 0x65, 0x64,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x12, 0x55, 0x0a, 0x0d,
	0x4f, 0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x21, 0x2e,
	0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f, 0x6e, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4f,
	0x6e, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x12, 0x1b,
	0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x7a, 0x6b,
	0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x4a, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74,
	0x57, 0x6f, 0x72, 0x6b, 0x12, 0x1e, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f,
	0x6f, 0x6c, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f,
	0x6f, 0x6c, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x56, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x52, 0x61, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f,
	0x6f, 0x6c, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73, 0x68, 0x52, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d,
	0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x44, 0x0a, 0x08, 0x48, 0x61,
	0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f,
	0x6f, 0x6c, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x3e, 0x0a, 0x06, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x7a, 0x6b, 0x65,
	0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78,
	0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x4d, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x1b, 0x5a, 0x19, 0x2e, 0x2f, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f,
	0x6c, 0x3b, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x78, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zkevmtxpool_zkevmmining_proto_rawDescOnce sync.Once
	file_zkevmtxpool_zkevmmining_proto_rawDescData = file_zkevmtxpool_zkevmmining_proto_rawDesc
)

func file_zkevmtxpool_zkevmmining_proto_rawDescGZIP() []byte {
	file_zkevmtxpool_zkevmmining_proto_rawDescOnce.Do(func() {
		file_zkevmtxpool_zkevmmining_proto_rawDescData = protoimpl.X.CompressGZIP(file_zkevmtxpool_zkevmmining_proto_rawDescData)
	})
	return file_zkevmtxpool_zkevmmining_proto_rawDescData
}

var file_zkevmtxpool_zkevmmining_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_zkevmtxpool_zkevmmining_proto_goTypes = []interface{}{
	(*OnPendingBlockRequest)(nil), // 0: zkevmtxpool.OnPendingBlockRequest
	(*OnPendingBlockReply)(nil),   // 1: zkevmtxpool.OnPendingBlockReply
	(*OnMinedBlockRequest)(nil),   // 2: zkevmtxpool.OnMinedBlockRequest
	(*OnMinedBlockReply)(nil),     // 3: zkevmtxpool.OnMinedBlockReply
	(*OnPendingLogsRequest)(nil),  // 4: zkevmtxpool.OnPendingLogsRequest
	(*OnPendingLogsReply)(nil),    // 5: zkevmtxpool.OnPendingLogsReply
	(*GetWorkRequest)(nil),        // 6: zkevmtxpool.GetWorkRequest
	(*GetWorkReply)(nil),          // 7: zkevmtxpool.GetWorkReply
	(*SubmitWorkRequest)(nil),     // 8: zkevmtxpool.SubmitWorkRequest
	(*SubmitWorkReply)(nil),       // 9: zkevmtxpool.SubmitWorkReply
	(*SubmitHashRateRequest)(nil), // 10: zkevmtxpool.SubmitHashRateRequest
	(*SubmitHashRateReply)(nil),   // 11: zkevmtxpool.SubmitHashRateReply
	(*HashRateRequest)(nil),       // 12: zkevmtxpool.HashRateRequest
	(*HashRateReply)(nil),         // 13: zkevmtxpool.HashRateReply
	(*MiningRequest)(nil),         // 14: zkevmtxpool.MiningRequest
	(*MiningReply)(nil),           // 15: zkevmtxpool.MiningReply
	(*emptypb.Empty)(nil),         // 16: google.protobuf.Empty
	(*types.VersionReply)(nil),    // 17: zkevmtypes.VersionReply
}
var file_zkevmtxpool_zkevmmining_proto_depIdxs = []int32{
	16, // 0: zkevmtxpool.Mining.Version:input_type -> google.protobuf.Empty
	0,  // 1: zkevmtxpool.Mining.OnPendingBlock:input_type -> zkevmtxpool.OnPendingBlockRequest
	2,  // 2: zkevmtxpool.Mining.OnMinedBlock:input_type -> zkevmtxpool.OnMinedBlockRequest
	4,  // 3: zkevmtxpool.Mining.OnPendingLogs:input_type -> zkevmtxpool.OnPendingLogsRequest
	6,  // 4: zkevmtxpool.Mining.GetWork:input_type -> zkevmtxpool.GetWorkRequest
	8,  // 5: zkevmtxpool.Mining.SubmitWork:input_type -> zkevmtxpool.SubmitWorkRequest
	10, // 6: zkevmtxpool.Mining.SubmitHashRate:input_type -> zkevmtxpool.SubmitHashRateRequest
	12, // 7: zkevmtxpool.Mining.HashRate:input_type -> zkevmtxpool.HashRateRequest
	14, // 8: zkevmtxpool.Mining.Mining:input_type -> zkevmtxpool.MiningRequest
	17, // 9: zkevmtxpool.Mining.Version:output_type -> zkevmtypes.VersionReply
	1,  // 10: zkevmtxpool.Mining.OnPendingBlock:output_type -> zkevmtxpool.OnPendingBlockReply
	3,  // 11: zkevmtxpool.Mining.OnMinedBlock:output_type -> zkevmtxpool.OnMinedBlockReply
	5,  // 12: zkevmtxpool.Mining.OnPendingLogs:output_type -> zkevmtxpool.OnPendingLogsReply
	7,  // 13: zkevmtxpool.Mining.GetWork:output_type -> zkevmtxpool.GetWorkReply
	9,  // 14: zkevmtxpool.Mining.SubmitWork:output_type -> zkevmtxpool.SubmitWorkReply
	11, // 15: zkevmtxpool.Mining.SubmitHashRate:output_type -> zkevmtxpool.SubmitHashRateReply
	13, // 16: zkevmtxpool.Mining.HashRate:output_type -> zkevmtxpool.HashRateReply
	15, // 17: zkevmtxpool.Mining.Mining:output_type -> zkevmtxpool.MiningReply
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_zkevmtxpool_zkevmmining_proto_init() }
func file_zkevmtxpool_zkevmmining_proto_init() {
	if File_zkevmtxpool_zkevmmining_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zkevmtxpool_zkevmmining_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnPendingBlockRequest); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnPendingBlockReply); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnMinedBlockRequest); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnMinedBlockReply); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnPendingLogsRequest); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OnPendingLogsReply); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkRequest); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkReply); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitWorkRequest); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitWorkReply); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitHashRateRequest); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitHashRateReply); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashRateRequest); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashRateReply); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiningRequest); i {
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
		file_zkevmtxpool_zkevmmining_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MiningReply); i {
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
			RawDescriptor: file_zkevmtxpool_zkevmmining_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zkevmtxpool_zkevmmining_proto_goTypes,
		DependencyIndexes: file_zkevmtxpool_zkevmmining_proto_depIdxs,
		MessageInfos:      file_zkevmtxpool_zkevmmining_proto_msgTypes,
	}.Build()
	File_zkevmtxpool_zkevmmining_proto = out.File
	file_zkevmtxpool_zkevmmining_proto_rawDesc = nil
	file_zkevmtxpool_zkevmmining_proto_goTypes = nil
	file_zkevmtxpool_zkevmmining_proto_depIdxs = nil
}
