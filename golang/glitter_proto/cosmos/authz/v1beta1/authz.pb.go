// Since: cosmos-sdk 0.43

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: cosmos/authz/v1beta1/authz.proto

package authz

import (
	_ "github.com/gogo/protobuf/gogoproto"
	any1 "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/regen-network/cosmos-proto"
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

// GenericAuthorization gives the grantee unrestricted permissions to execute
// the provided method on behalf of the granter's account.
type GenericAuthorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Msg, identified by it's type URL, to grant unrestricted permissions to execute
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GenericAuthorization) Reset() {
	*x = GenericAuthorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_authz_v1beta1_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericAuthorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericAuthorization) ProtoMessage() {}

func (x *GenericAuthorization) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_authz_v1beta1_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericAuthorization.ProtoReflect.Descriptor instead.
func (*GenericAuthorization) Descriptor() ([]byte, []int) {
	return file_cosmos_authz_v1beta1_authz_proto_rawDescGZIP(), []int{0}
}

func (x *GenericAuthorization) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// Grant gives permissions to execute
// the provide method with expiration time.
type Grant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authorization *any1.Any            `protobuf:"bytes,1,opt,name=authorization,proto3" json:"authorization,omitempty"`
	Expiration    *timestamp.Timestamp `protobuf:"bytes,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *Grant) Reset() {
	*x = Grant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_authz_v1beta1_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Grant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Grant) ProtoMessage() {}

func (x *Grant) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_authz_v1beta1_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Grant.ProtoReflect.Descriptor instead.
func (*Grant) Descriptor() ([]byte, []int) {
	return file_cosmos_authz_v1beta1_authz_proto_rawDescGZIP(), []int{1}
}

func (x *Grant) GetAuthorization() *any1.Any {
	if x != nil {
		return x.Authorization
	}
	return nil
}

func (x *Grant) GetExpiration() *timestamp.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// GrantAuthorization extends a grant with both the addresses of the grantee and granter.
// It is used in genesis.proto and query.proto
//
// Since: cosmos-sdk 0.45.2
type GrantAuthorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Granter       string               `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee       string               `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	Authorization *any1.Any            `protobuf:"bytes,3,opt,name=authorization,proto3" json:"authorization,omitempty"`
	Expiration    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
}

func (x *GrantAuthorization) Reset() {
	*x = GrantAuthorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_authz_v1beta1_authz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrantAuthorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrantAuthorization) ProtoMessage() {}

func (x *GrantAuthorization) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_authz_v1beta1_authz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrantAuthorization.ProtoReflect.Descriptor instead.
func (*GrantAuthorization) Descriptor() ([]byte, []int) {
	return file_cosmos_authz_v1beta1_authz_proto_rawDescGZIP(), []int{2}
}

func (x *GrantAuthorization) GetGranter() string {
	if x != nil {
		return x.Granter
	}
	return ""
}

func (x *GrantAuthorization) GetGrantee() string {
	if x != nil {
		return x.Grantee
	}
	return ""
}

func (x *GrantAuthorization) GetAuthorization() *any1.Any {
	if x != nil {
		return x.Authorization
	}
	return nil
}

func (x *GrantAuthorization) GetExpiration() *timestamp.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

var File_cosmos_authz_v1beta1_authz_proto protoreflect.FileDescriptor

var file_cosmos_authz_v1beta1_authz_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x14, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x3a,
	0x11, 0xd2, 0xb4, 0x2d, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x9c, 0x01, 0x0a, 0x05, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x4d, 0x0a, 0x0d,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x11, 0xca, 0xb4, 0x2d, 0x0d, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f,
	0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xdd, 0x01, 0x0a, 0x12, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e,
	0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x12, 0x4d, 0x0a, 0x0d,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x11, 0xca, 0xb4, 0x2d, 0x0d, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0a, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f,
	0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x2a, 0xc8, 0xe1, 0x1e, 0x00, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x78, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_authz_v1beta1_authz_proto_rawDescOnce sync.Once
	file_cosmos_authz_v1beta1_authz_proto_rawDescData = file_cosmos_authz_v1beta1_authz_proto_rawDesc
)

func file_cosmos_authz_v1beta1_authz_proto_rawDescGZIP() []byte {
	file_cosmos_authz_v1beta1_authz_proto_rawDescOnce.Do(func() {
		file_cosmos_authz_v1beta1_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_authz_v1beta1_authz_proto_rawDescData)
	})
	return file_cosmos_authz_v1beta1_authz_proto_rawDescData
}

var file_cosmos_authz_v1beta1_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cosmos_authz_v1beta1_authz_proto_goTypes = []interface{}{
	(*GenericAuthorization)(nil), // 0: cosmos.authz.v1beta1.GenericAuthorization
	(*Grant)(nil),                // 1: cosmos.authz.v1beta1.Grant
	(*GrantAuthorization)(nil),   // 2: cosmos.authz.v1beta1.GrantAuthorization
	(*any1.Any)(nil),             // 3: google.protobuf.Any
	(*timestamp.Timestamp)(nil),  // 4: google.protobuf.Timestamp
}
var file_cosmos_authz_v1beta1_authz_proto_depIdxs = []int32{
	3, // 0: cosmos.authz.v1beta1.Grant.authorization:type_name -> google.protobuf.Any
	4, // 1: cosmos.authz.v1beta1.Grant.expiration:type_name -> google.protobuf.Timestamp
	3, // 2: cosmos.authz.v1beta1.GrantAuthorization.authorization:type_name -> google.protobuf.Any
	4, // 3: cosmos.authz.v1beta1.GrantAuthorization.expiration:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cosmos_authz_v1beta1_authz_proto_init() }
func file_cosmos_authz_v1beta1_authz_proto_init() {
	if File_cosmos_authz_v1beta1_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_authz_v1beta1_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericAuthorization); i {
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
		file_cosmos_authz_v1beta1_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Grant); i {
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
		file_cosmos_authz_v1beta1_authz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrantAuthorization); i {
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
			RawDescriptor: file_cosmos_authz_v1beta1_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_authz_v1beta1_authz_proto_goTypes,
		DependencyIndexes: file_cosmos_authz_v1beta1_authz_proto_depIdxs,
		MessageInfos:      file_cosmos_authz_v1beta1_authz_proto_msgTypes,
	}.Build()
	File_cosmos_authz_v1beta1_authz_proto = out.File
	file_cosmos_authz_v1beta1_authz_proto_rawDesc = nil
	file_cosmos_authz_v1beta1_authz_proto_goTypes = nil
	file_cosmos_authz_v1beta1_authz_proto_depIdxs = nil
}
