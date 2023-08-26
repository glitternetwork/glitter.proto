// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: ethermint/types/v1/dynamic_fee.proto

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

// ExtensionOptionDynamicFeeTx is an extension option that specify the maxPrioPrice for cosmos tx
type ExtensionOptionDynamicFeeTx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the same as `max_priority_fee_per_gas` in eip-1559 spec
	MaxPriorityPrice string `protobuf:"bytes,1,opt,name=max_priority_price,json=maxPriorityPrice,proto3" json:"max_priority_price,omitempty"`
}

func (x *ExtensionOptionDynamicFeeTx) Reset() {
	*x = ExtensionOptionDynamicFeeTx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethermint_types_v1_dynamic_fee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtensionOptionDynamicFeeTx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionOptionDynamicFeeTx) ProtoMessage() {}

func (x *ExtensionOptionDynamicFeeTx) ProtoReflect() protoreflect.Message {
	mi := &file_ethermint_types_v1_dynamic_fee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionOptionDynamicFeeTx.ProtoReflect.Descriptor instead.
func (*ExtensionOptionDynamicFeeTx) Descriptor() ([]byte, []int) {
	return file_ethermint_types_v1_dynamic_fee_proto_rawDescGZIP(), []int{0}
}

func (x *ExtensionOptionDynamicFeeTx) GetMaxPriorityPrice() string {
	if x != nil {
		return x.MaxPriorityPrice
	}
	return ""
}

var File_ethermint_types_v1_dynamic_fee_proto protoreflect.FileDescriptor

var file_ethermint_types_v1_dynamic_fee_proto_rawDesc = []byte{
	0x0a, 0x24, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x74, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x7b, 0x0a, 0x1b, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x46, 0x65, 0x65, 0x54, 0x78, 0x12,
	0x5c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xc8, 0xde, 0x1f,
	0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64,
	0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x10, 0x6d, 0x61, 0x78,
	0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x22, 0x5a,
	0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x76, 0x6d, 0x6f,
	0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ethermint_types_v1_dynamic_fee_proto_rawDescOnce sync.Once
	file_ethermint_types_v1_dynamic_fee_proto_rawDescData = file_ethermint_types_v1_dynamic_fee_proto_rawDesc
)

func file_ethermint_types_v1_dynamic_fee_proto_rawDescGZIP() []byte {
	file_ethermint_types_v1_dynamic_fee_proto_rawDescOnce.Do(func() {
		file_ethermint_types_v1_dynamic_fee_proto_rawDescData = protoimpl.X.CompressGZIP(file_ethermint_types_v1_dynamic_fee_proto_rawDescData)
	})
	return file_ethermint_types_v1_dynamic_fee_proto_rawDescData
}

var file_ethermint_types_v1_dynamic_fee_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ethermint_types_v1_dynamic_fee_proto_goTypes = []interface{}{
	(*ExtensionOptionDynamicFeeTx)(nil), // 0: ethermint.types.v1.ExtensionOptionDynamicFeeTx
}
var file_ethermint_types_v1_dynamic_fee_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ethermint_types_v1_dynamic_fee_proto_init() }
func file_ethermint_types_v1_dynamic_fee_proto_init() {
	if File_ethermint_types_v1_dynamic_fee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ethermint_types_v1_dynamic_fee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtensionOptionDynamicFeeTx); i {
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
			RawDescriptor: file_ethermint_types_v1_dynamic_fee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ethermint_types_v1_dynamic_fee_proto_goTypes,
		DependencyIndexes: file_ethermint_types_v1_dynamic_fee_proto_depIdxs,
		MessageInfos:      file_ethermint_types_v1_dynamic_fee_proto_msgTypes,
	}.Build()
	File_ethermint_types_v1_dynamic_fee_proto = out.File
	file_ethermint_types_v1_dynamic_fee_proto_rawDesc = nil
	file_ethermint_types_v1_dynamic_fee_proto_goTypes = nil
	file_ethermint_types_v1_dynamic_fee_proto_depIdxs = nil
}
