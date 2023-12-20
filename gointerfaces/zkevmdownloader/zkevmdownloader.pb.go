// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.2
// source: zkevmdownloader/zkevmdownloader.proto

package zkevmdownloader

import (
	zkevmtypes "github.com/tenderly/zkevm-erigon-lib/gointerfaces/zkevmtypes"
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

// DownloadItem:
// - if Erigon created new snapshot and want seed it
// - if Erigon wnat download files - it fills only "torrent_hash" field
type DownloadItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path        string           `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	TorrentHash *zkevmtypes.H160 `protobuf:"bytes,2,opt,name=torrent_hash,json=torrentHash,proto3" json:"torrent_hash,omitempty"` // will be resolved as magnet link
}

func (x *DownloadItem) Reset() {
	*x = DownloadItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmdownloader_zkevmdownloader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadItem) ProtoMessage() {}

func (x *DownloadItem) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmdownloader_zkevmdownloader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadItem.ProtoReflect.Descriptor instead.
func (*DownloadItem) Descriptor() ([]byte, []int) {
	return file_zkevmdownloader_zkevmdownloader_proto_rawDescGZIP(), []int{0}
}

func (x *DownloadItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *DownloadItem) GetTorrentHash() *zkevmtypes.H160 {
	if x != nil {
		return x.TorrentHash
	}
	return nil
}

type DownloadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*DownloadItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // single hash will be resolved as magnet link
}

func (x *DownloadRequest) Reset() {
	*x = DownloadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmdownloader_zkevmdownloader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadRequest) ProtoMessage() {}

func (x *DownloadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmdownloader_zkevmdownloader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadRequest.ProtoReflect.Descriptor instead.
func (*DownloadRequest) Descriptor() ([]byte, []int) {
	return file_zkevmdownloader_zkevmdownloader_proto_rawDescGZIP(), []int{1}
}

func (x *DownloadRequest) GetItems() []*DownloadItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmdownloader_zkevmdownloader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmdownloader_zkevmdownloader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_zkevmdownloader_zkevmdownloader_proto_rawDescGZIP(), []int{2}
}

type StatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatsRequest) Reset() {
	*x = StatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmdownloader_zkevmdownloader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsRequest) ProtoMessage() {}

func (x *StatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmdownloader_zkevmdownloader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsRequest.ProtoReflect.Descriptor instead.
func (*StatsRequest) Descriptor() ([]byte, []int) {
	return file_zkevmdownloader_zkevmdownloader_proto_rawDescGZIP(), []int{3}
}

type StatsReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// First step on startup - "resolve metadata":
	//   - understand total amount of data to download
	//   - ensure all pieces hashes available
	//   - validate files after crush
	//   - when all metadata ready - can start download/upload
	MetadataReady    int32   `protobuf:"varint,1,opt,name=metadata_ready,json=metadataReady,proto3" json:"metadata_ready,omitempty"`
	FilesTotal       int32   `protobuf:"varint,2,opt,name=files_total,json=filesTotal,proto3" json:"files_total,omitempty"`
	PeersUnique      int32   `protobuf:"varint,4,opt,name=peers_unique,json=peersUnique,proto3" json:"peers_unique,omitempty"`
	ConnectionsTotal uint64  `protobuf:"varint,5,opt,name=connections_total,json=connectionsTotal,proto3" json:"connections_total,omitempty"`
	Completed        bool    `protobuf:"varint,6,opt,name=completed,proto3" json:"completed,omitempty"`
	Progress         float32 `protobuf:"fixed32,7,opt,name=progress,proto3" json:"progress,omitempty"`
	BytesCompleted   uint64  `protobuf:"varint,8,opt,name=bytes_completed,json=bytesCompleted,proto3" json:"bytes_completed,omitempty"`
	BytesTotal       uint64  `protobuf:"varint,9,opt,name=bytes_total,json=bytesTotal,proto3" json:"bytes_total,omitempty"`
	UploadRate       uint64  `protobuf:"varint,10,opt,name=upload_rate,json=uploadRate,proto3" json:"upload_rate,omitempty"`       // bytes/sec
	DownloadRate     uint64  `protobuf:"varint,11,opt,name=download_rate,json=downloadRate,proto3" json:"download_rate,omitempty"` // bytes/sec
}

func (x *StatsReply) Reset() {
	*x = StatsReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zkevmdownloader_zkevmdownloader_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsReply) ProtoMessage() {}

func (x *StatsReply) ProtoReflect() protoreflect.Message {
	mi := &file_zkevmdownloader_zkevmdownloader_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsReply.ProtoReflect.Descriptor instead.
func (*StatsReply) Descriptor() ([]byte, []int) {
	return file_zkevmdownloader_zkevmdownloader_proto_rawDescGZIP(), []int{4}
}

func (x *StatsReply) GetMetadataReady() int32 {
	if x != nil {
		return x.MetadataReady
	}
	return 0
}

func (x *StatsReply) GetFilesTotal() int32 {
	if x != nil {
		return x.FilesTotal
	}
	return 0
}

func (x *StatsReply) GetPeersUnique() int32 {
	if x != nil {
		return x.PeersUnique
	}
	return 0
}

func (x *StatsReply) GetConnectionsTotal() uint64 {
	if x != nil {
		return x.ConnectionsTotal
	}
	return 0
}

func (x *StatsReply) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *StatsReply) GetProgress() float32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *StatsReply) GetBytesCompleted() uint64 {
	if x != nil {
		return x.BytesCompleted
	}
	return 0
}

func (x *StatsReply) GetBytesTotal() uint64 {
	if x != nil {
		return x.BytesTotal
	}
	return 0
}

func (x *StatsReply) GetUploadRate() uint64 {
	if x != nil {
		return x.UploadRate
	}
	return 0
}

func (x *StatsReply) GetDownloadRate() uint64 {
	if x != nil {
		return x.DownloadRate
	}
	return 0
}

var File_zkevmdownloader_zkevmdownloader_proto protoreflect.FileDescriptor

var file_zkevmdownloader_zkevmdownloader_proto_rawDesc = []byte{
	0x0a, 0x25, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x2f, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x57, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x33, 0x0a, 0x0c, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x7a,
	0x6b, 0x65, 0x76, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x31, 0x36, 0x30, 0x52, 0x0b,
	0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x22, 0x46, 0x0a, 0x0f, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0xee, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x70, 0x65, 0x65, 0x72, 0x73, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x61, 0x74, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x61, 0x74, 0x65, 0x32, 0xdf, 0x01, 0x0a, 0x0a, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x20, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x65, 0x72, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x06,
	0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x1e, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d, 0x64, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x7a, 0x6b, 0x65, 0x76,
	0x6d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x7a, 0x6b, 0x65, 0x76, 0x6d,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x2e, 0x2f, 0x7a, 0x6b, 0x65,
	0x76, 0x6d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x3b, 0x7a, 0x6b, 0x65,
	0x76, 0x6d, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zkevmdownloader_zkevmdownloader_proto_rawDescOnce sync.Once
	file_zkevmdownloader_zkevmdownloader_proto_rawDescData = file_zkevmdownloader_zkevmdownloader_proto_rawDesc
)

func file_zkevmdownloader_zkevmdownloader_proto_rawDescGZIP() []byte {
	file_zkevmdownloader_zkevmdownloader_proto_rawDescOnce.Do(func() {
		file_zkevmdownloader_zkevmdownloader_proto_rawDescData = protoimpl.X.CompressGZIP(file_zkevmdownloader_zkevmdownloader_proto_rawDescData)
	})
	return file_zkevmdownloader_zkevmdownloader_proto_rawDescData
}

var file_zkevmdownloader_zkevmdownloader_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_zkevmdownloader_zkevmdownloader_proto_goTypes = []interface{}{
	(*DownloadItem)(nil),    // 0: zkevmdownloader.DownloadItem
	(*DownloadRequest)(nil), // 1: zkevmdownloader.DownloadRequest
	(*VerifyRequest)(nil),   // 2: zkevmdownloader.VerifyRequest
	(*StatsRequest)(nil),    // 3: zkevmdownloader.StatsRequest
	(*StatsReply)(nil),      // 4: zkevmdownloader.StatsReply
	(*zkevmtypes.H160)(nil), // 5: zkevmtypes.H160
	(*emptypb.Empty)(nil),   // 6: google.protobuf.Empty
}
var file_zkevmdownloader_zkevmdownloader_proto_depIdxs = []int32{
	5, // 0: zkevmdownloader.DownloadItem.torrent_hash:type_name -> zkevmtypes.H160
	0, // 1: zkevmdownloader.DownloadRequest.items:type_name -> zkevmdownloader.DownloadItem
	1, // 2: zkevmdownloader.Downloader.Download:input_type -> zkevmdownloader.DownloadRequest
	2, // 3: zkevmdownloader.Downloader.Verify:input_type -> zkevmdownloader.VerifyRequest
	3, // 4: zkevmdownloader.Downloader.Stats:input_type -> zkevmdownloader.StatsRequest
	6, // 5: zkevmdownloader.Downloader.Download:output_type -> google.protobuf.Empty
	6, // 6: zkevmdownloader.Downloader.Verify:output_type -> google.protobuf.Empty
	4, // 7: zkevmdownloader.Downloader.Stats:output_type -> zkevmdownloader.StatsReply
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_zkevmdownloader_zkevmdownloader_proto_init() }
func file_zkevmdownloader_zkevmdownloader_proto_init() {
	if File_zkevmdownloader_zkevmdownloader_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zkevmdownloader_zkevmdownloader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadItem); i {
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
		file_zkevmdownloader_zkevmdownloader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadRequest); i {
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
		file_zkevmdownloader_zkevmdownloader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_zkevmdownloader_zkevmdownloader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsRequest); i {
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
		file_zkevmdownloader_zkevmdownloader_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsReply); i {
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
			RawDescriptor: file_zkevmdownloader_zkevmdownloader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zkevmdownloader_zkevmdownloader_proto_goTypes,
		DependencyIndexes: file_zkevmdownloader_zkevmdownloader_proto_depIdxs,
		MessageInfos:      file_zkevmdownloader_zkevmdownloader_proto_msgTypes,
	}.Build()
	File_zkevmdownloader_zkevmdownloader_proto = out.File
	file_zkevmdownloader_zkevmdownloader_proto_rawDesc = nil
	file_zkevmdownloader_zkevmdownloader_proto_goTypes = nil
	file_zkevmdownloader_zkevmdownloader_proto_depIdxs = nil
}
