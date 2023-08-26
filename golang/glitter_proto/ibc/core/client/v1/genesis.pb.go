// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: ibc/core/client/v1/genesis.proto

package types

import (
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the ibc client submodule's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// client states with their corresponding identifiers
	Clients []*IdentifiedClientState `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
	// consensus states from each client
	ClientsConsensus []*ClientConsensusStates `protobuf:"bytes,2,rep,name=clients_consensus,json=clientsConsensus,proto3" json:"clients_consensus,omitempty"`
	// metadata from each client
	ClientsMetadata []*IdentifiedGenesisMetadata `protobuf:"bytes,3,rep,name=clients_metadata,json=clientsMetadata,proto3" json:"clients_metadata,omitempty"`
	Params          *Params                      `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
	// create localhost on initialization
	CreateLocalhost bool `protobuf:"varint,5,opt,name=create_localhost,json=createLocalhost,proto3" json:"create_localhost,omitempty"`
	// the sequence for the next generated client identifier
	NextClientSequence uint64 `protobuf:"varint,6,opt,name=next_client_sequence,json=nextClientSequence,proto3" json:"next_client_sequence,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_core_client_v1_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_core_client_v1_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_ibc_core_client_v1_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetClients() []*IdentifiedClientState {
	if x != nil {
		return x.Clients
	}
	return nil
}

func (x *GenesisState) GetClientsConsensus() []*ClientConsensusStates {
	if x != nil {
		return x.ClientsConsensus
	}
	return nil
}

func (x *GenesisState) GetClientsMetadata() []*IdentifiedGenesisMetadata {
	if x != nil {
		return x.ClientsMetadata
	}
	return nil
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *GenesisState) GetCreateLocalhost() bool {
	if x != nil {
		return x.CreateLocalhost
	}
	return false
}

func (x *GenesisState) GetNextClientSequence() uint64 {
	if x != nil {
		return x.NextClientSequence
	}
	return 0
}

// GenesisMetadata defines the genesis type for metadata that clients may return
// with ExportMetadata
type GenesisMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// store key of metadata without clientID-prefix
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// metadata value
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GenesisMetadata) Reset() {
	*x = GenesisMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_core_client_v1_genesis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisMetadata) ProtoMessage() {}

func (x *GenesisMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_core_client_v1_genesis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisMetadata.ProtoReflect.Descriptor instead.
func (*GenesisMetadata) Descriptor() ([]byte, []int) {
	return file_ibc_core_client_v1_genesis_proto_rawDescGZIP(), []int{1}
}

func (x *GenesisMetadata) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GenesisMetadata) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

// IdentifiedGenesisMetadata has the client metadata with the corresponding
// client id.
type IdentifiedGenesisMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId       string             `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientMetadata []*GenesisMetadata `protobuf:"bytes,2,rep,name=client_metadata,json=clientMetadata,proto3" json:"client_metadata,omitempty"`
}

func (x *IdentifiedGenesisMetadata) Reset() {
	*x = IdentifiedGenesisMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_core_client_v1_genesis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdentifiedGenesisMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdentifiedGenesisMetadata) ProtoMessage() {}

func (x *IdentifiedGenesisMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_core_client_v1_genesis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdentifiedGenesisMetadata.ProtoReflect.Descriptor instead.
func (*IdentifiedGenesisMetadata) Descriptor() ([]byte, []int) {
	return file_ibc_core_client_v1_genesis_proto_rawDescGZIP(), []int{2}
}

func (x *IdentifiedGenesisMetadata) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *IdentifiedGenesisMetadata) GetClientMetadata() []*GenesisMetadata {
	if x != nil {
		return x.ClientMetadata
	}
	return nil
}

var File_ibc_core_client_v1_genesis_proto protoreflect.FileDescriptor

var file_ibc_core_client_v1_genesis_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x62, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x69, 0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x69, 0x62, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x04,
	0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x63,
	0x0a, 0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x69, 0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x1e, 0xc8, 0xde, 0x1f, 0x00,
	0xaa, 0xdf, 0x1f, 0x16, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x07, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x92, 0x01, 0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x69, 0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x42, 0x3a, 0xc8, 0xde, 0x1f, 0x00,
	0xf2, 0xde, 0x1f, 0x18, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x22, 0xaa, 0xdf, 0x1f, 0x16,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x52, 0x10, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x43,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x12, 0x79, 0x0a, 0x10, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x1f, 0xc8, 0xde, 0x1f, 0x00, 0xf2, 0xde, 0x1f, 0x17, 0x79, 0x61, 0x6d, 0x6c, 0x3a,
	0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42,
	0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x46, 0x0a,
	0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x1b, 0xf2, 0xde, 0x1f, 0x17, 0x79, 0x61, 0x6d,
	0x6c, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68,
	0x6f, 0x73, 0x74, 0x22, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x14, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x1f, 0xf2, 0xde, 0x1f, 0x1b, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x22, 0x52, 0x12, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65,
	0x73, 0x69, 0x73, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x22, 0xbc, 0x01, 0x0a, 0x19, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0xf2, 0xde, 0x1f, 0x10,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x6c, 0x0a, 0x0f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x1e, 0xc8, 0xde, 0x1f, 0x00, 0xf2, 0xde,
	0x1f, 0x16, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0x52, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x69, 0x62,
	0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x30, 0x32, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ibc_core_client_v1_genesis_proto_rawDescOnce sync.Once
	file_ibc_core_client_v1_genesis_proto_rawDescData = file_ibc_core_client_v1_genesis_proto_rawDesc
)

func file_ibc_core_client_v1_genesis_proto_rawDescGZIP() []byte {
	file_ibc_core_client_v1_genesis_proto_rawDescOnce.Do(func() {
		file_ibc_core_client_v1_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_ibc_core_client_v1_genesis_proto_rawDescData)
	})
	return file_ibc_core_client_v1_genesis_proto_rawDescData
}

var file_ibc_core_client_v1_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ibc_core_client_v1_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil),              // 0: ibc.core.client.v1.GenesisState
	(*GenesisMetadata)(nil),           // 1: ibc.core.client.v1.GenesisMetadata
	(*IdentifiedGenesisMetadata)(nil), // 2: ibc.core.client.v1.IdentifiedGenesisMetadata
	(*IdentifiedClientState)(nil),     // 3: ibc.core.client.v1.IdentifiedClientState
	(*ClientConsensusStates)(nil),     // 4: ibc.core.client.v1.ClientConsensusStates
	(*Params)(nil),                    // 5: ibc.core.client.v1.Params
}
var file_ibc_core_client_v1_genesis_proto_depIdxs = []int32{
	3, // 0: ibc.core.client.v1.GenesisState.clients:type_name -> ibc.core.client.v1.IdentifiedClientState
	4, // 1: ibc.core.client.v1.GenesisState.clients_consensus:type_name -> ibc.core.client.v1.ClientConsensusStates
	2, // 2: ibc.core.client.v1.GenesisState.clients_metadata:type_name -> ibc.core.client.v1.IdentifiedGenesisMetadata
	5, // 3: ibc.core.client.v1.GenesisState.params:type_name -> ibc.core.client.v1.Params
	1, // 4: ibc.core.client.v1.IdentifiedGenesisMetadata.client_metadata:type_name -> ibc.core.client.v1.GenesisMetadata
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ibc_core_client_v1_genesis_proto_init() }
func file_ibc_core_client_v1_genesis_proto_init() {
	if File_ibc_core_client_v1_genesis_proto != nil {
		return
	}
	file_ibc_core_client_v1_client_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ibc_core_client_v1_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
		file_ibc_core_client_v1_genesis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisMetadata); i {
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
		file_ibc_core_client_v1_genesis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdentifiedGenesisMetadata); i {
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
			RawDescriptor: file_ibc_core_client_v1_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ibc_core_client_v1_genesis_proto_goTypes,
		DependencyIndexes: file_ibc_core_client_v1_genesis_proto_depIdxs,
		MessageInfos:      file_ibc_core_client_v1_genesis_proto_msgTypes,
	}.Build()
	File_ibc_core_client_v1_genesis_proto = out.File
	file_ibc_core_client_v1_genesis_proto_rawDesc = nil
	file_ibc_core_client_v1_genesis_proto_goTypes = nil
	file_ibc_core_client_v1_genesis_proto_depIdxs = nil
}
