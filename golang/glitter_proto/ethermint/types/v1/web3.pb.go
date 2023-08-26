// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: ethermint/types/v1/web3.proto

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

type ExtensionOptionsWeb3Tx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// typed data chain id used only in EIP712 Domain and should match
	// Ethereum network ID in a Web3 provider (e.g. Metamask).
	TypedDataChainId uint64 `protobuf:"varint,1,opt,name=typed_data_chain_id,json=typedDataChainId,proto3" json:"typed_data_chain_id,omitempty"`
	// fee payer is an account address for the fee payer. It will be validated
	// during EIP712 signature checking.
	FeePayer string `protobuf:"bytes,2,opt,name=fee_payer,json=feePayer,proto3" json:"fee_payer,omitempty"`
	// fee payer sig is a signature data from the fee paying account,
	// allows to perform fee delegation when using EIP712 Domain.
	FeePayerSig []byte `protobuf:"bytes,3,opt,name=fee_payer_sig,json=feePayerSig,proto3" json:"fee_payer_sig,omitempty"`
}

func (x *ExtensionOptionsWeb3Tx) Reset() {
	*x = ExtensionOptionsWeb3Tx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethermint_types_v1_web3_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionOptionsWeb3Tx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionOptionsWeb3Tx) ProtoMessage() {}

func (x *ExtensionOptionsWeb3Tx) ProtoReflect() protoreflect.Message {
	mi := &file_ethermint_types_v1_web3_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionOptionsWeb3Tx.ProtoReflect.Descriptor instead.
func (*ExtensionOptionsWeb3Tx) Descriptor() ([]byte, []int) {
	return file_ethermint_types_v1_web3_proto_rawDescGZIP(), []int{0}
}

func (x *ExtensionOptionsWeb3Tx) GetTypedDataChainId() uint64 {
	if x != nil {
		return x.TypedDataChainId
	}
	return 0
}

func (x *ExtensionOptionsWeb3Tx) GetFeePayer() string {
	if x != nil {
		return x.FeePayer
	}
	return ""
}

func (x *ExtensionOptionsWeb3Tx) GetFeePayerSig() []byte {
	if x != nil {
		return x.FeePayerSig
	}
	return nil
}

var File_ethermint_types_v1_web3_proto protoreflect.FileDescriptor

var file_ethermint_types_v1_web3_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x16, 0x45, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x57, 0x65,
	0x62, 0x33, 0x54, 0x78, 0x12, 0x61, 0x0a, 0x13, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x32, 0xe2, 0xde, 0x1f, 0x10, 0x54, 0x79, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0xea, 0xde, 0x1f, 0x1a, 0x74, 0x79, 0x70, 0x65, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x44, 0x2c, 0x6f, 0x6d, 0x69, 0x74,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x10, 0x74, 0x79, 0x70, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x65, 0x65, 0x5f, 0x70,
	0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xea, 0xde, 0x1f, 0x12,
	0x66, 0x65, 0x65, 0x50, 0x61, 0x79, 0x65, 0x72, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x08, 0x66, 0x65, 0x65, 0x50, 0x61, 0x79, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x0d,
	0x66, 0x65, 0x65, 0x5f, 0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x69, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x19, 0xea, 0xde, 0x1f, 0x15, 0x66, 0x65, 0x65, 0x50, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x69, 0x67, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x0b,
	0x66, 0x65, 0x65, 0x50, 0x61, 0x79, 0x65, 0x72, 0x53, 0x69, 0x67, 0x3a, 0x04, 0x88, 0xa0, 0x1f,
	0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x76, 0x6d, 0x6f, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ethermint_types_v1_web3_proto_rawDescOnce sync.Once
	file_ethermint_types_v1_web3_proto_rawDescData = file_ethermint_types_v1_web3_proto_rawDesc
)

func file_ethermint_types_v1_web3_proto_rawDescGZIP() []byte {
	file_ethermint_types_v1_web3_proto_rawDescOnce.Do(func() {
		file_ethermint_types_v1_web3_proto_rawDescData = protoimpl.X.CompressGZIP(file_ethermint_types_v1_web3_proto_rawDescData)
	})
	return file_ethermint_types_v1_web3_proto_rawDescData
}

var file_ethermint_types_v1_web3_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ethermint_types_v1_web3_proto_goTypes = []interface{}{
	(*ExtensionOptionsWeb3Tx)(nil), // 0: ethermint.types.v1.ExtensionOptionsWeb3Tx
}
var file_ethermint_types_v1_web3_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ethermint_types_v1_web3_proto_init() }
func file_ethermint_types_v1_web3_proto_init() {
	if File_ethermint_types_v1_web3_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ethermint_types_v1_web3_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionOptionsWeb3Tx); i {
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
			RawDescriptor: file_ethermint_types_v1_web3_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ethermint_types_v1_web3_proto_goTypes,
		DependencyIndexes: file_ethermint_types_v1_web3_proto_depIdxs,
		MessageInfos:      file_ethermint_types_v1_web3_proto_msgTypes,
	}.Build()
	File_ethermint_types_v1_web3_proto = out.File
	file_ethermint_types_v1_web3_proto_rawDesc = nil
	file_ethermint_types_v1_web3_proto_goTypes = nil
	file_ethermint_types_v1_web3_proto_depIdxs = nil
}
