// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: cosmos/bank/v1beta1/tx.proto

package types

import (
	types "github.com/cosmos/cosmos-sdk/types"
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

// MsgSend represents a message to send coins from one account to another.
type MsgSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress string        `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress   string        `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount      []*types.Coin `protobuf:"bytes,3,rep,name=amount,proto3" json:"amount,omitempty"`
}

func (x *MsgSend) Reset() {
	*x = MsgSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSend) ProtoMessage() {}

func (x *MsgSend) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSend.ProtoReflect.Descriptor instead.
func (*MsgSend) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgSend) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *MsgSend) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *MsgSend) GetAmount() []*types.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

// MsgSendResponse defines the Msg/Send response type.
type MsgSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgSendResponse) Reset() {
	*x = MsgSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSendResponse) ProtoMessage() {}

func (x *MsgSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSendResponse.ProtoReflect.Descriptor instead.
func (*MsgSendResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

// MsgMultiSend represents an arbitrary multi-in, multi-out send message.
type MsgMultiSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Inputs  []*Input  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs []*Output `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs,omitempty"`
}

func (x *MsgMultiSend) Reset() {
	*x = MsgMultiSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgMultiSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgMultiSend) ProtoMessage() {}

func (x *MsgMultiSend) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgMultiSend.ProtoReflect.Descriptor instead.
func (*MsgMultiSend) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v1beta1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgMultiSend) GetInputs() []*Input {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *MsgMultiSend) GetOutputs() []*Output {
	if x != nil {
		return x.Outputs
	}
	return nil
}

// MsgMultiSendResponse defines the Msg/MultiSend response type.
type MsgMultiSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgMultiSendResponse) Reset() {
	*x = MsgMultiSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgMultiSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgMultiSendResponse) ProtoMessage() {}

func (x *MsgMultiSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v1beta1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgMultiSendResponse.ProtoReflect.Descriptor instead.
func (*MsgMultiSendResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v1beta1_tx_proto_rawDescGZIP(), []int{3}
}

var File_cosmos_bank_v1beta1_tx_proto protoreflect.FileDescriptor

var file_cosmos_bank_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63,
	0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xea, 0x01, 0x0a, 0x07, 0x4d, 0x73,
	0x67, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x3a, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xf2, 0xde, 0x1f,
	0x13, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x34, 0x0a, 0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xf2, 0xde, 0x1f, 0x11, 0x79, 0x61, 0x6d, 0x6c, 0x3a,
	0x22, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x52, 0x09, 0x74, 0x6f,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x63, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x42, 0x30, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x73, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x08, 0x88, 0xa0,
	0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x11, 0x0a, 0x0f, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0c, 0x4d, 0x73,
	0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x38, 0x0a, 0x06, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x06, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x73, 0x12, 0x3b, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62,
	0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x73, 0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x16, 0x0a, 0x14, 0x4d, 0x73, 0x67, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xac, 0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x4a, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12,
	0x1c, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x1a, 0x24, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x09, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x65, 0x6e, 0x64,
	0x12, 0x21, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53,
	0x65, 0x6e, 0x64, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e,
	0x6b, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b,
	0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x78,
	0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cosmos_bank_v1beta1_tx_proto_rawDescOnce sync.Once
	file_cosmos_bank_v1beta1_tx_proto_rawDescData = file_cosmos_bank_v1beta1_tx_proto_rawDesc
)

func file_cosmos_bank_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_cosmos_bank_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_bank_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_bank_v1beta1_tx_proto_rawDescData)
	})
	return file_cosmos_bank_v1beta1_tx_proto_rawDescData
}

var file_cosmos_bank_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cosmos_bank_v1beta1_tx_proto_goTypes = []interface{}{
	(*MsgSend)(nil),              // 0: cosmos.bank.v1beta1.MsgSend
	(*MsgSendResponse)(nil),      // 1: cosmos.bank.v1beta1.MsgSendResponse
	(*MsgMultiSend)(nil),         // 2: cosmos.bank.v1beta1.MsgMultiSend
	(*MsgMultiSendResponse)(nil), // 3: cosmos.bank.v1beta1.MsgMultiSendResponse
	(*types.Coin)(nil),           // 4: cosmos.base.v1beta1.Coin
	(*Input)(nil),                // 5: cosmos.bank.v1beta1.Input
	(*Output)(nil),               // 6: cosmos.bank.v1beta1.Output
}
var file_cosmos_bank_v1beta1_tx_proto_depIdxs = []int32{
	4, // 0: cosmos.bank.v1beta1.MsgSend.amount:type_name -> cosmos.base.v1beta1.Coin
	5, // 1: cosmos.bank.v1beta1.MsgMultiSend.inputs:type_name -> cosmos.bank.v1beta1.Input
	6, // 2: cosmos.bank.v1beta1.MsgMultiSend.outputs:type_name -> cosmos.bank.v1beta1.Output
	0, // 3: cosmos.bank.v1beta1.Msg.Send:input_type -> cosmos.bank.v1beta1.MsgSend
	2, // 4: cosmos.bank.v1beta1.Msg.MultiSend:input_type -> cosmos.bank.v1beta1.MsgMultiSend
	1, // 5: cosmos.bank.v1beta1.Msg.Send:output_type -> cosmos.bank.v1beta1.MsgSendResponse
	3, // 6: cosmos.bank.v1beta1.Msg.MultiSend:output_type -> cosmos.bank.v1beta1.MsgMultiSendResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cosmos_bank_v1beta1_tx_proto_init() }
func file_cosmos_bank_v1beta1_tx_proto_init() {
	if File_cosmos_bank_v1beta1_tx_proto != nil {
		return
	}
	file_cosmos_bank_v1beta1_bank_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_bank_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSend); i {
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
		file_cosmos_bank_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSendResponse); i {
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
		file_cosmos_bank_v1beta1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgMultiSend); i {
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
		file_cosmos_bank_v1beta1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgMultiSendResponse); i {
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
			RawDescriptor: file_cosmos_bank_v1beta1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_bank_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_bank_v1beta1_tx_proto_depIdxs,
		MessageInfos:      file_cosmos_bank_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_cosmos_bank_v1beta1_tx_proto = out.File
	file_cosmos_bank_v1beta1_tx_proto_rawDesc = nil
	file_cosmos_bank_v1beta1_tx_proto_goTypes = nil
	file_cosmos_bank_v1beta1_tx_proto_depIdxs = nil
}
