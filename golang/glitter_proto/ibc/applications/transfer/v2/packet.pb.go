// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: ibc/applications/transfer/v2/packet.proto

package types

import (
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

// FungibleTokenPacketData defines a struct for the packet payload
// See FungibleTokenPacketData spec:
// https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer#data-structures
type FungibleTokenPacketData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the token denomination to be transferred
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// the token amount to be transferred
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// the sender address
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (x *FungibleTokenPacketData) Reset() {
	*x = FungibleTokenPacketData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ibc_applications_transfer_v2_packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FungibleTokenPacketData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FungibleTokenPacketData) ProtoMessage() {}

func (x *FungibleTokenPacketData) ProtoReflect() protoreflect.Message {
	mi := &file_ibc_applications_transfer_v2_packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FungibleTokenPacketData.ProtoReflect.Descriptor instead.
func (*FungibleTokenPacketData) Descriptor() ([]byte, []int) {
	return file_ibc_applications_transfer_v2_packet_proto_rawDescGZIP(), []int{0}
}

func (x *FungibleTokenPacketData) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *FungibleTokenPacketData) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *FungibleTokenPacketData) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *FungibleTokenPacketData) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

var File_ibc_applications_transfer_v2_packet_proto protoreflect.FileDescriptor

var file_ibc_applications_transfer_v2_packet_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x62, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x70,
	0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x62, 0x63,
	0x2e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x22, 0x7b, 0x0a, 0x17, 0x46, 0x75, 0x6e,
	0x67, 0x69, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x69, 0x62, 0x63, 0x2d,
	0x67, 0x6f, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ibc_applications_transfer_v2_packet_proto_rawDescOnce sync.Once
	file_ibc_applications_transfer_v2_packet_proto_rawDescData = file_ibc_applications_transfer_v2_packet_proto_rawDesc
)

func file_ibc_applications_transfer_v2_packet_proto_rawDescGZIP() []byte {
	file_ibc_applications_transfer_v2_packet_proto_rawDescOnce.Do(func() {
		file_ibc_applications_transfer_v2_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_ibc_applications_transfer_v2_packet_proto_rawDescData)
	})
	return file_ibc_applications_transfer_v2_packet_proto_rawDescData
}

var file_ibc_applications_transfer_v2_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ibc_applications_transfer_v2_packet_proto_goTypes = []interface{}{
	(*FungibleTokenPacketData)(nil), // 0: ibc.applications.transfer.v2.FungibleTokenPacketData
}
var file_ibc_applications_transfer_v2_packet_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ibc_applications_transfer_v2_packet_proto_init() }
func file_ibc_applications_transfer_v2_packet_proto_init() {
	if File_ibc_applications_transfer_v2_packet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ibc_applications_transfer_v2_packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FungibleTokenPacketData); i {
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
			RawDescriptor: file_ibc_applications_transfer_v2_packet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ibc_applications_transfer_v2_packet_proto_goTypes,
		DependencyIndexes: file_ibc_applications_transfer_v2_packet_proto_depIdxs,
		MessageInfos:      file_ibc_applications_transfer_v2_packet_proto_msgTypes,
	}.Build()
	File_ibc_applications_transfer_v2_packet_proto = out.File
	file_ibc_applications_transfer_v2_packet_proto_rawDesc = nil
	file_ibc_applications_transfer_v2_packet_proto_goTypes = nil
	file_ibc_applications_transfer_v2_packet_proto_depIdxs = nil
}
