// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: ibc/applications/interchain_accounts/v1/packet.proto

package types

import (
	_ "github.com/gogo/protobuf/gogoproto"
	any1 "github.com/golang/protobuf/ptypes/any"
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

// Type defines a classification of message issued from a controller chain to its associated interchain accounts
// host
type Type int32

const (
	// Default zero value enumeration
	Type_TYPE_UNSPECIFIED Type = 0
	// Execute a transaction on an interchain accounts host chain
	Type_TYPE_EXECUTE_TX Type = 1
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_EXECUTE_TX",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_EXECUTE_TX":  1,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_ibc_applications_interchain_accounts_v1_packet_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_ibc_applications_interchain_accounts_v1_packet_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_ibc_applications_interchain_accounts_v1_packet_proto_rawDescGZIP(), []int{0}
}

// InterchainAccountPacketData is comprised of a raw transaction, type of transaction and optional memo field.
type InterchainAccountPacketData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Type   `protobuf:"varint,1,opt,name=type,proto3,enum=ibc.applications.interchain_accounts.v1.Type" json:"type,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Memo string `protobuf:"bytes,3,opt,name=memo,proto3" json:"memo,omitempty"`
}

func (x *InterchainAccountPacketData) Reset() {
	*x = InterchainAccountPacketData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_applications_interchain_accounts_v1_packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterchainAccountPacketData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterchainAccountPacketData) ProtoMessage() {}

func (x *InterchainAccountPacketData) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_applications_interchain_accounts_v1_packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterchainAccountPacketData.ProtoReflect.Descriptor instead.
func (*InterchainAccountPacketData) Descriptor() ([]byte, []int) {
	return file_ibc_applications_interchain_accounts_v1_packet_proto_rawDescGZIP(), []int{0}
}

func (x *InterchainAccountPacketData) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

func (x *InterchainAccountPacketData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *InterchainAccountPacketData) GetMemo() string {
	if x != nil {
		return x.Memo
	}
	return ""
}

// CosmosTx contains a list of sdk.Msg's. It should be used when sending transactions to an SDK host chain.
type CosmosTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*any1.Any `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *CosmosTx) Reset() {
	*x = CosmosTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_applications_interchain_accounts_v1_packet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CosmosTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CosmosTx) ProtoMessage() {}

func (x *CosmosTx) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_applications_interchain_accounts_v1_packet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CosmosTx.ProtoReflect.Descriptor instead.
func (*CosmosTx) Descriptor() ([]byte, []int) {
	return file_ibc_applications_interchain_accounts_v1_packet_proto_rawDescGZIP(), []int{1}
}

func (x *CosmosTx) GetMessages() []*any1.Any {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_ibc_applications_interchain_accounts_v1_packet_proto protoreflect.FileDescriptor

var file_ibc_applications_interchain_accounts_v1_packet_proto_rawDesc = []byte{
	0x0a, 0x34, 0x69, 0x62, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x69, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x88, 0x01, 0x0a, 0x1b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x41, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d,
	0x2e, 0x69, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x22, 0x3c, 0x0a, 0x08, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x54, 0x78, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2a, 0x58, 0x0a, 0x04, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x25, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x1a, 0x0f, 0x8a, 0x9d, 0x20, 0x0b, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x12, 0x23, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x5f, 0x54, 0x58, 0x10, 0x01, 0x1a, 0x0e, 0x8a,
	0x9d, 0x20, 0x0a, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45, 0x5f, 0x54, 0x58, 0x1a, 0x04, 0x88,
	0xa3, 0x1e, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x69, 0x62, 0x63, 0x2d, 0x67, 0x6f, 0x2f,
	0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f,
	0x32, 0x37, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2d, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ibc_applications_interchain_accounts_v1_packet_proto_rawDescOnce sync.Once
	file_ibc_applications_interchain_accounts_v1_packet_proto_rawDescData = file_ibc_applications_interchain_accounts_v1_packet_proto_rawDesc
)

func file_ibc_applications_interchain_accounts_v1_packet_proto_rawDescGZIP() []byte {
	file_ibc_applications_interchain_accounts_v1_packet_proto_rawDescOnce.Do(func() {
		file_ibc_applications_interchain_accounts_v1_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_ibc_applications_interchain_accounts_v1_packet_proto_rawDescData)
	})
	return file_ibc_applications_interchain_accounts_v1_packet_proto_rawDescData
}

var file_ibc_applications_interchain_accounts_v1_packet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ibc_applications_interchain_accounts_v1_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ibc_applications_interchain_accounts_v1_packet_proto_goTypes = []interface{}{
	(Type)(0),                           // 0: ibc.applications.interchain_accounts.v1.Type
	(*InterchainAccountPacketData)(nil), // 1: ibc.applications.interchain_accounts.v1.InterchainAccountPacketData
	(*CosmosTx)(nil),                    // 2: ibc.applications.interchain_accounts.v1.CosmosTx
	(*any1.Any)(nil),                    // 3: google.protobuf.Any
}
var file_ibc_applications_interchain_accounts_v1_packet_proto_depIdxs = []int32{
	0, // 0: ibc.applications.interchain_accounts.v1.InterchainAccountPacketData.type:type_name -> ibc.applications.interchain_accounts.v1.Type
	3, // 1: ibc.applications.interchain_accounts.v1.CosmosTx.messages:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ibc_applications_interchain_accounts_v1_packet_proto_init() }
func file_ibc_applications_interchain_accounts_v1_packet_proto_init() {
	if File_ibc_applications_interchain_accounts_v1_packet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ibc_applications_interchain_accounts_v1_packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterchainAccountPacketData); i {
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
		file_ibc_applications_interchain_accounts_v1_packet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CosmosTx); i {
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
			RawDescriptor: file_ibc_applications_interchain_accounts_v1_packet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ibc_applications_interchain_accounts_v1_packet_proto_goTypes,
		DependencyIndexes: file_ibc_applications_interchain_accounts_v1_packet_proto_depIdxs,
		EnumInfos:         file_ibc_applications_interchain_accounts_v1_packet_proto_enumTypes,
		MessageInfos:      file_ibc_applications_interchain_accounts_v1_packet_proto_msgTypes,
	}.Build()
	File_ibc_applications_interchain_accounts_v1_packet_proto = out.File
	file_ibc_applications_interchain_accounts_v1_packet_proto_rawDesc = nil
	file_ibc_applications_interchain_accounts_v1_packet_proto_goTypes = nil
	file_ibc_applications_interchain_accounts_v1_packet_proto_depIdxs = nil
}
