// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: ibc/core/connection/v1/genesis.proto

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

// GenesisState defines the ibc connection submodule's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connections           []*IdentifiedConnection `protobuf:"bytes,1,rep,name=connections,proto3" json:"connections,omitempty"`
	ClientConnectionPaths []*ConnectionPaths      `protobuf:"bytes,2,rep,name=client_connection_paths,json=clientConnectionPaths,proto3" json:"client_connection_paths,omitempty"`
	// the sequence for the next generated connection identifier
	NextConnectionSequence uint64  `protobuf:"varint,3,opt,name=next_connection_sequence,json=nextConnectionSequence,proto3" json:"next_connection_sequence,omitempty"`
	Params                 *Params `protobuf:"bytes,4,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_core_connection_v1_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_core_connection_v1_genesis_proto_msgTypes[0]
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
	return file_ibc_core_connection_v1_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetConnections() []*IdentifiedConnection {
	if x != nil {
		return x.Connections
	}
	return nil
}

func (x *GenesisState) GetClientConnectionPaths() []*ConnectionPaths {
	if x != nil {
		return x.ClientConnectionPaths
	}
	return nil
}

func (x *GenesisState) GetNextConnectionSequence() uint64 {
	if x != nil {
		return x.NextConnectionSequence
	}
	return 0
}

func (x *GenesisState) GetParams() *Params {
	if x != nil {
		return x.Params
	}
	return nil
}

var File_ibc_core_connection_v1_genesis_proto protoreflect.FileDescriptor

var file_ibc_core_connection_v1_genesis_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x62, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x69, 0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x14,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x69, 0x62, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x03,
	0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x54,
	0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x87, 0x01, 0x0a, 0x17, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x73, 0x42,
	0x26, 0xc8, 0xde, 0x1f, 0x00, 0xf2, 0xde, 0x1f, 0x1e, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x22, 0x52, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x73, 0x12, 0x5d,
	0x0a, 0x18, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x23, 0xf2, 0xde, 0x1f, 0x1f, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x22, 0x52, 0x16, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e,
	0x69, 0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x04, 0xc8,
	0xde, 0x1f, 0x00, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x3e, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x69, 0x62, 0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x73, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x30, 0x33, 0x2d, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_ibc_core_connection_v1_genesis_proto_rawDescOnce sync.Once
	file_ibc_core_connection_v1_genesis_proto_rawDescData = file_ibc_core_connection_v1_genesis_proto_rawDesc
)

func file_ibc_core_connection_v1_genesis_proto_rawDescGZIP() []byte {
	file_ibc_core_connection_v1_genesis_proto_rawDescOnce.Do(func() {
		file_ibc_core_connection_v1_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_ibc_core_connection_v1_genesis_proto_rawDescData)
	})
	return file_ibc_core_connection_v1_genesis_proto_rawDescData
}

var file_ibc_core_connection_v1_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ibc_core_connection_v1_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil),         // 0: ibc.core.connection.v1.GenesisState
	(*IdentifiedConnection)(nil), // 1: ibc.core.connection.v1.IdentifiedConnection
	(*ConnectionPaths)(nil),      // 2: ibc.core.connection.v1.ConnectionPaths
	(*Params)(nil),               // 3: ibc.core.connection.v1.Params
}
var file_ibc_core_connection_v1_genesis_proto_depIdxs = []int32{
	1, // 0: ibc.core.connection.v1.GenesisState.connections:type_name -> ibc.core.connection.v1.IdentifiedConnection
	2, // 1: ibc.core.connection.v1.GenesisState.client_connection_paths:type_name -> ibc.core.connection.v1.ConnectionPaths
	3, // 2: ibc.core.connection.v1.GenesisState.params:type_name -> ibc.core.connection.v1.Params
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ibc_core_connection_v1_genesis_proto_init() }
func file_ibc_core_connection_v1_genesis_proto_init() {
	if File_ibc_core_connection_v1_genesis_proto != nil {
		return
	}
	file_ibc_core_connection_v1_connection_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ibc_core_connection_v1_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ibc_core_connection_v1_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ibc_core_connection_v1_genesis_proto_goTypes,
		DependencyIndexes: file_ibc_core_connection_v1_genesis_proto_depIdxs,
		MessageInfos:      file_ibc_core_connection_v1_genesis_proto_msgTypes,
	}.Build()
	File_ibc_core_connection_v1_genesis_proto = out.File
	file_ibc_core_connection_v1_genesis_proto_rawDesc = nil
	file_ibc_core_connection_v1_genesis_proto_goTypes = nil
	file_ibc_core_connection_v1_genesis_proto_depIdxs = nil
}
