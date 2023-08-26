// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: ibc/applications/interchain_accounts/v1/account.proto

package types

import (
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// An InterchainAccount is defined as a BaseAccount & the address of the account owner on the controller chain
type InterchainAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAccount  *types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3" json:"base_account,omitempty"`
	AccountOwner string             `protobuf:"bytes,2,opt,name=account_owner,json=accountOwner,proto3" json:"account_owner,omitempty"`
}

func (x *InterchainAccount) Reset() {
	*x = InterchainAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_applications_interchain_accounts_v1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterchainAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterchainAccount) ProtoMessage() {}

func (x *InterchainAccount) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_applications_interchain_accounts_v1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterchainAccount.ProtoReflect.Descriptor instead.
func (*InterchainAccount) Descriptor() ([]byte, []int) {
	return file_ibc_applications_interchain_accounts_v1_account_proto_rawDescGZIP(), []int{0}
}

func (x *InterchainAccount) GetBaseAccount() *types.BaseAccount {
	if x != nil {
		return x.BaseAccount
	}
	return nil
}

func (x *InterchainAccount) GetAccountOwner() string {
	if x != nil {
		return x.AccountOwner
	}
	return ""
}

var File_ibc_applications_interchain_accounts_v1_account_proto protoreflect.FileDescriptor

var file_ibc_applications_interchain_accounts_v1_account_proto_rawDesc = []byte{
	0x0a, 0x35, 0x69, 0x62, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x69, 0x62, 0x63, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xd4, 0x01, 0x0a, 0x11, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x60, 0x0a, 0x0c, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x1b, 0xd0, 0xde, 0x1f, 0x01, 0xf2, 0xde, 0x1f, 0x13, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x52, 0x0b, 0x62, 0x61,
	0x73, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0d, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xf2, 0xde, 0x1f, 0x14, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x3a, 0x1e, 0x88, 0xa0, 0x1f, 0x00, 0x98, 0xa0,
	0x1f, 0x00, 0xd2, 0xb4, 0x2d, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x69, 0x62,
	0x63, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f,
	0x61, 0x70, 0x70, 0x73, 0x2f, 0x32, 0x37, 0x2d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ibc_applications_interchain_accounts_v1_account_proto_rawDescOnce sync.Once
	file_ibc_applications_interchain_accounts_v1_account_proto_rawDescData = file_ibc_applications_interchain_accounts_v1_account_proto_rawDesc
)

func file_ibc_applications_interchain_accounts_v1_account_proto_rawDescGZIP() []byte {
	file_ibc_applications_interchain_accounts_v1_account_proto_rawDescOnce.Do(func() {
		file_ibc_applications_interchain_accounts_v1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_ibc_applications_interchain_accounts_v1_account_proto_rawDescData)
	})
	return file_ibc_applications_interchain_accounts_v1_account_proto_rawDescData
}

var file_ibc_applications_interchain_accounts_v1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ibc_applications_interchain_accounts_v1_account_proto_goTypes = []interface{}{
	(*InterchainAccount)(nil), // 0: ibc.applications.interchain_accounts.v1.InterchainAccount
	(*types.BaseAccount)(nil), // 1: cosmos.auth.v1beta1.BaseAccount
}
var file_ibc_applications_interchain_accounts_v1_account_proto_depIdxs = []int32{
	1, // 0: ibc.applications.interchain_accounts.v1.InterchainAccount.base_account:type_name -> cosmos.auth.v1beta1.BaseAccount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ibc_applications_interchain_accounts_v1_account_proto_init() }
func file_ibc_applications_interchain_accounts_v1_account_proto_init() {
	if File_ibc_applications_interchain_accounts_v1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ibc_applications_interchain_accounts_v1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterchainAccount); i {
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
			RawDescriptor: file_ibc_applications_interchain_accounts_v1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ibc_applications_interchain_accounts_v1_account_proto_goTypes,
		DependencyIndexes: file_ibc_applications_interchain_accounts_v1_account_proto_depIdxs,
		MessageInfos:      file_ibc_applications_interchain_accounts_v1_account_proto_msgTypes,
	}.Build()
	File_ibc_applications_interchain_accounts_v1_account_proto = out.File
	file_ibc_applications_interchain_accounts_v1_account_proto_rawDesc = nil
	file_ibc_applications_interchain_accounts_v1_account_proto_goTypes = nil
	file_ibc_applications_interchain_accounts_v1_account_proto_depIdxs = nil
}
