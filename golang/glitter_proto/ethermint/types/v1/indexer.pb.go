// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: ethermint/types/v1/indexer.proto

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

// TxResult is the value stored in eth tx indexer
type TxResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the block height
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// cosmos tx index
	TxIndex uint32 `protobuf:"varint,2,opt,name=tx_index,json=txIndex,proto3" json:"tx_index,omitempty"`
	// the msg index in a batch tx
	MsgIndex uint32 `protobuf:"varint,3,opt,name=msg_index,json=msgIndex,proto3" json:"msg_index,omitempty"`
	// eth tx index, the index in the list of valid eth tx in the block,
	// aka. the transaction list returned by eth_getBlock api.
	EthTxIndex int32 `protobuf:"varint,4,opt,name=eth_tx_index,json=ethTxIndex,proto3" json:"eth_tx_index,omitempty"`
	// if the eth tx is failed
	Failed bool `protobuf:"varint,5,opt,name=failed,proto3" json:"failed,omitempty"`
	// gas used by tx, if exceeds block gas limit,
	// it's set to gas limit which is what's actually deducted by ante handler.
	GasUsed uint64 `protobuf:"varint,6,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	// the cumulative gas used within current batch tx
	CumulativeGasUsed uint64 `protobuf:"varint,7,opt,name=cumulative_gas_used,json=cumulativeGasUsed,proto3" json:"cumulative_gas_used,omitempty"`
}

func (x *TxResult) Reset() {
	*x = TxResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ethermint_types_v1_indexer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxResult) ProtoMessage() {}

func (x *TxResult) ProtoReflect() protoreflect.Message {
	mi := &file_ethermint_types_v1_indexer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxResult.ProtoReflect.Descriptor instead.
func (*TxResult) Descriptor() ([]byte, []int) {
	return file_ethermint_types_v1_indexer_proto_rawDescGZIP(), []int{0}
}

func (x *TxResult) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *TxResult) GetTxIndex() uint32 {
	if x != nil {
		return x.TxIndex
	}
	return 0
}

func (x *TxResult) GetMsgIndex() uint32 {
	if x != nil {
		return x.MsgIndex
	}
	return 0
}

func (x *TxResult) GetEthTxIndex() int32 {
	if x != nil {
		return x.EthTxIndex
	}
	return 0
}

func (x *TxResult) GetFailed() bool {
	if x != nil {
		return x.Failed
	}
	return false
}

func (x *TxResult) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *TxResult) GetCumulativeGasUsed() uint64 {
	if x != nil {
		return x.CumulativeGasUsed
	}
	return 0
}

var File_ethermint_types_v1_indexer_proto protoreflect.FileDescriptor

var file_ethermint_types_v1_indexer_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x74, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5, 0x01, 0x0a,
	0x08, 0x54, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x73, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x6d, 0x73, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0c, 0x65, 0x74, 0x68,
	0x5f, 0x74, 0x78, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x65, 0x74, 0x68, 0x54, 0x78, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x67, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x2e,
	0x0a, 0x13, 0x63, 0x75, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x67, 0x61, 0x73,
	0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x63, 0x75, 0x6d,
	0x75, 0x6c, 0x61, 0x74, 0x69, 0x76, 0x65, 0x47, 0x61, 0x73, 0x55, 0x73, 0x65, 0x64, 0x3a, 0x04,
	0x88, 0xa0, 0x1f, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x76, 0x6d, 0x6f, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ethermint_types_v1_indexer_proto_rawDescOnce sync.Once
	file_ethermint_types_v1_indexer_proto_rawDescData = file_ethermint_types_v1_indexer_proto_rawDesc
)

func file_ethermint_types_v1_indexer_proto_rawDescGZIP() []byte {
	file_ethermint_types_v1_indexer_proto_rawDescOnce.Do(func() {
		file_ethermint_types_v1_indexer_proto_rawDescData = protoimpl.X.CompressGZIP(file_ethermint_types_v1_indexer_proto_rawDescData)
	})
	return file_ethermint_types_v1_indexer_proto_rawDescData
}

var file_ethermint_types_v1_indexer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ethermint_types_v1_indexer_proto_goTypes = []interface{}{
	(*TxResult)(nil), // 0: ethermint.types.v1.TxResult
}
var file_ethermint_types_v1_indexer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ethermint_types_v1_indexer_proto_init() }
func file_ethermint_types_v1_indexer_proto_init() {
	if File_ethermint_types_v1_indexer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ethermint_types_v1_indexer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxResult); i {
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
			RawDescriptor: file_ethermint_types_v1_indexer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ethermint_types_v1_indexer_proto_goTypes,
		DependencyIndexes: file_ethermint_types_v1_indexer_proto_depIdxs,
		MessageInfos:      file_ethermint_types_v1_indexer_proto_msgTypes,
	}.Build()
	File_ethermint_types_v1_indexer_proto = out.File
	file_ethermint_types_v1_indexer_proto_rawDesc = nil
	file_ethermint_types_v1_indexer_proto_goTypes = nil
	file_ethermint_types_v1_indexer_proto_depIdxs = nil
}
