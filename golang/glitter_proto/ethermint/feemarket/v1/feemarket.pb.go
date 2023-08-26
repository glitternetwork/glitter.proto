// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: ethermint/feemarket/v1/feemarket.proto

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

// Params defines the EVM module parameters
type Params struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// no base fee forces the EIP-1559 base fee to 0 (needed for 0 price calls)
	NoBaseFee bool `protobuf:"varint,1,opt,name=no_base_fee,json=noBaseFee,proto3" json:"no_base_fee,omitempty"`
	// base fee change denominator bounds the amount the base fee can change
	// between blocks.
	BaseFeeChangeDenominator uint32 `protobuf:"varint,2,opt,name=base_fee_change_denominator,json=baseFeeChangeDenominator,proto3" json:"base_fee_change_denominator,omitempty"`
	// elasticity multiplier bounds the maximum gas limit an EIP-1559 block may
	// have.
	ElasticityMultiplier uint32 `protobuf:"varint,3,opt,name=elasticity_multiplier,json=elasticityMultiplier,proto3" json:"elasticity_multiplier,omitempty"`
	// height at which the base fee calculation is enabled.
	EnableHeight int64 `protobuf:"varint,5,opt,name=enable_height,json=enableHeight,proto3" json:"enable_height,omitempty"`
	// base fee for EIP-1559 blocks.
	BaseFee string `protobuf:"bytes,6,opt,name=base_fee,json=baseFee,proto3" json:"base_fee,omitempty"`
	// min_gas_price defines the minimum gas price value for cosmos and eth transactions
	MinGasPrice string `protobuf:"bytes,7,opt,name=min_gas_price,json=minGasPrice,proto3" json:"min_gas_price,omitempty"`
	// min gas denominator bounds the minimum gasUsed to be charged
	// to senders based on GasLimit
	MinGasMultiplier string `protobuf:"bytes,8,opt,name=min_gas_multiplier,json=minGasMultiplier,proto3" json:"min_gas_multiplier,omitempty"`
}

func (x *Params) Reset() {
	*x = Params{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethermint_feemarket_v1_feemarket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Params) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Params) ProtoMessage() {}

func (x *Params) ProtoReflect() protoreflect.Message {
	mi := &file_ethermint_feemarket_v1_feemarket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Params.ProtoReflect.Descriptor instead.
func (*Params) Descriptor() ([]byte, []int) {
	return file_ethermint_feemarket_v1_feemarket_proto_rawDescGZIP(), []int{0}
}

func (x *Params) GetNoBaseFee() bool {
	if x != nil {
		return x.NoBaseFee
	}
	return false
}

func (x *Params) GetBaseFeeChangeDenominator() uint32 {
	if x != nil {
		return x.BaseFeeChangeDenominator
	}
	return 0
}

func (x *Params) GetElasticityMultiplier() uint32 {
	if x != nil {
		return x.ElasticityMultiplier
	}
	return 0
}

func (x *Params) GetEnableHeight() int64 {
	if x != nil {
		return x.EnableHeight
	}
	return 0
}

func (x *Params) GetBaseFee() string {
	if x != nil {
		return x.BaseFee
	}
	return ""
}

func (x *Params) GetMinGasPrice() string {
	if x != nil {
		return x.MinGasPrice
	}
	return ""
}

func (x *Params) GetMinGasMultiplier() string {
	if x != nil {
		return x.MinGasMultiplier
	}
	return ""
}

var File_ethermint_feemarket_v1_feemarket_proto protoreflect.FileDescriptor

var file_ethermint_feemarket_v1_feemarket_proto_rawDesc = []byte{
	0x0a, 0x26, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x66, 0x65, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x74, 0x2e, 0x66, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x03, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x1e, 0x0a, 0x0b, 0x6e, 0x6f, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6e, 0x6f, 0x42, 0x61, 0x73, 0x65, 0x46, 0x65,
	0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x5f, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x62, 0x61, 0x73, 0x65, 0x46, 0x65, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x33, 0x0a, 0x15, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x6d,
	0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x14, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x69, 0x74, 0x79, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x69, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x49, 0x0a, 0x08, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xc8, 0xde,
	0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x07, 0x62, 0x61,
	0x73, 0x65, 0x46, 0x65, 0x65, 0x12, 0x52, 0x0a, 0x0d, 0x6d, 0x69, 0x6e, 0x5f, 0x67, 0x61, 0x73,
	0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xc8, 0xde,
	0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73,
	0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x0b, 0x6d, 0x69,
	0x6e, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x12, 0x6d, 0x69, 0x6e,
	0x5f, 0x67, 0x61, 0x73, 0x5f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x26, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x44, 0x65, 0x63, 0x52, 0x10, 0x6d, 0x69, 0x6e, 0x47, 0x61, 0x73, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x52, 0x10, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x76,
	0x6d, 0x6f, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x78, 0x2f,
	0x66, 0x65, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ethermint_feemarket_v1_feemarket_proto_rawDescOnce sync.Once
	file_ethermint_feemarket_v1_feemarket_proto_rawDescData = file_ethermint_feemarket_v1_feemarket_proto_rawDesc
)

func file_ethermint_feemarket_v1_feemarket_proto_rawDescGZIP() []byte {
	file_ethermint_feemarket_v1_feemarket_proto_rawDescOnce.Do(func() {
		file_ethermint_feemarket_v1_feemarket_proto_rawDescData = protoimpl.X.CompressGZIP(file_ethermint_feemarket_v1_feemarket_proto_rawDescData)
	})
	return file_ethermint_feemarket_v1_feemarket_proto_rawDescData
}

var file_ethermint_feemarket_v1_feemarket_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ethermint_feemarket_v1_feemarket_proto_goTypes = []interface{}{
	(*Params)(nil), // 0: ethermint.feemarket.v1.Params
}
var file_ethermint_feemarket_v1_feemarket_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ethermint_feemarket_v1_feemarket_proto_init() }
func file_ethermint_feemarket_v1_feemarket_proto_init() {
	if File_ethermint_feemarket_v1_feemarket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ethermint_feemarket_v1_feemarket_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Params); i {
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
			RawDescriptor: file_ethermint_feemarket_v1_feemarket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ethermint_feemarket_v1_feemarket_proto_goTypes,
		DependencyIndexes: file_ethermint_feemarket_v1_feemarket_proto_depIdxs,
		MessageInfos:      file_ethermint_feemarket_v1_feemarket_proto_msgTypes,
	}.Build()
	File_ethermint_feemarket_v1_feemarket_proto = out.File
	file_ethermint_feemarket_v1_feemarket_proto_rawDesc = nil
	file_ethermint_feemarket_v1_feemarket_proto_goTypes = nil
	file_ethermint_feemarket_v1_feemarket_proto_depIdxs = nil
}
