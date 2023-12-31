// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: cosmos/vesting/v1beta1/tx.proto

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

// MsgCreateVestingAccount defines a message that enables creating a vesting
// account.
type MsgCreateVestingAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress string        `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress   string        `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount      []*types.Coin `protobuf:"bytes,3,rep,name=amount,proto3" json:"amount,omitempty"`
	EndTime     int64         `protobuf:"varint,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Delayed     bool          `protobuf:"varint,5,opt,name=delayed,proto3" json:"delayed,omitempty"`
}

func (x *MsgCreateVestingAccount) Reset() {
	*x = MsgCreateVestingAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateVestingAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateVestingAccount) ProtoMessage() {}

func (x *MsgCreateVestingAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreateVestingAccount.ProtoReflect.Descriptor instead.
func (*MsgCreateVestingAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgCreateVestingAccount) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *MsgCreateVestingAccount) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *MsgCreateVestingAccount) GetAmount() []*types.Coin {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *MsgCreateVestingAccount) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *MsgCreateVestingAccount) GetDelayed() bool {
	if x != nil {
		return x.Delayed
	}
	return false
}

// MsgCreateVestingAccountResponse defines the Msg/CreateVestingAccount response type.
type MsgCreateVestingAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgCreateVestingAccountResponse) Reset() {
	*x = MsgCreateVestingAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateVestingAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateVestingAccountResponse) ProtoMessage() {}

func (x *MsgCreateVestingAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreateVestingAccountResponse.ProtoReflect.Descriptor instead.
func (*MsgCreateVestingAccountResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP(), []int{1}
}

// MsgCreatePeriodicVestingAccount defines a message that enables creating a vesting
// account.
type MsgCreatePeriodicVestingAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress    string    `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	ToAddress      string    `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	StartTime      int64     `protobuf:"varint,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	VestingPeriods []*Period `protobuf:"bytes,4,rep,name=vesting_periods,json=vestingPeriods,proto3" json:"vesting_periods,omitempty"`
}

func (x *MsgCreatePeriodicVestingAccount) Reset() {
	*x = MsgCreatePeriodicVestingAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreatePeriodicVestingAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreatePeriodicVestingAccount) ProtoMessage() {}

func (x *MsgCreatePeriodicVestingAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreatePeriodicVestingAccount.ProtoReflect.Descriptor instead.
func (*MsgCreatePeriodicVestingAccount) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgCreatePeriodicVestingAccount) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

func (x *MsgCreatePeriodicVestingAccount) GetToAddress() string {
	if x != nil {
		return x.ToAddress
	}
	return ""
}

func (x *MsgCreatePeriodicVestingAccount) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *MsgCreatePeriodicVestingAccount) GetVestingPeriods() []*Period {
	if x != nil {
		return x.VestingPeriods
	}
	return nil
}

// MsgCreatePeriodicVestingAccountResponse defines the Msg/CreatePeriodicVestingAccount
// response type.
type MsgCreatePeriodicVestingAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgCreatePeriodicVestingAccountResponse) Reset() {
	*x = MsgCreatePeriodicVestingAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreatePeriodicVestingAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreatePeriodicVestingAccountResponse) ProtoMessage() {}

func (x *MsgCreatePeriodicVestingAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreatePeriodicVestingAccountResponse.ProtoReflect.Descriptor instead.
func (*MsgCreatePeriodicVestingAccountResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP(), []int{3}
}

// MsgDonateAllVestingTokens defines a message that enables donating all vesting
// token to community pool.
type MsgDonateAllVestingTokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (x *MsgDonateAllVestingTokens) Reset() {
	*x = MsgDonateAllVestingTokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDonateAllVestingTokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDonateAllVestingTokens) ProtoMessage() {}

func (x *MsgDonateAllVestingTokens) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDonateAllVestingTokens.ProtoReflect.Descriptor instead.
func (*MsgDonateAllVestingTokens) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP(), []int{4}
}

func (x *MsgDonateAllVestingTokens) GetFromAddress() string {
	if x != nil {
		return x.FromAddress
	}
	return ""
}

// MsgDonateAllVestingTokensResponse defines the Msg/MsgDonateAllVestingTokens
// response type.
type MsgDonateAllVestingTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgDonateAllVestingTokensResponse) Reset() {
	*x = MsgDonateAllVestingTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDonateAllVestingTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDonateAllVestingTokensResponse) ProtoMessage() {}

func (x *MsgDonateAllVestingTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cosmos_vesting_v1beta1_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDonateAllVestingTokensResponse.ProtoReflect.Descriptor instead.
func (*MsgDonateAllVestingTokensResponse) Descriptor() ([]byte, []int) {
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP(), []int{5}
}

var File_cosmos_vesting_v1beta1_tx_proto protoreflect.FileDescriptor

var file_cosmos_vesting_v1beta1_tx_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x17, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x3a, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xf2, 0xde, 0x1f, 0x13, 0x79, 0x61, 0x6d,
	0x6c, 0x3a, 0x22, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22,
	0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a,
	0x0a, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x15, 0xf2, 0xde, 0x1f, 0x11, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x6f, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x63, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x30,
	0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x13, 0xf2, 0xde, 0x1f, 0x0f,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x52,
	0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x65, 0x64, 0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x01, 0x22, 0x21, 0x0a, 0x1f, 0x4d, 0x73, 0x67, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x9e, 0x02, 0x0a, 0x1f,
	0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69,
	0x63, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x3a, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xf2, 0xde, 0x1f, 0x13, 0x79, 0x61, 0x6d, 0x6c, 0x3a,
	0x22, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x52, 0x0b,
	0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x34, 0x0a, 0x0a, 0x74,
	0x6f, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x15, 0xf2, 0xde, 0x1f, 0x11, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x6f, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x52, 0x09, 0x74, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x34, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x15, 0xf2, 0xde, 0x1f, 0x11, 0x79, 0x61, 0x6d, 0x6c, 0x3a,
	0x22, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x0f, 0x76, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x50,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x73, 0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x29, 0x0a, 0x27,
	0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69,
	0x63, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x0a, 0x19, 0x4d, 0x73, 0x67, 0x44, 0x6f,
	0x6e, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xf2, 0xde, 0x1f, 0x13,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x52, 0x0b, 0x66, 0x72, 0x6f, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x3a, 0x04, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x23, 0x0a, 0x21, 0x4d, 0x73, 0x67, 0x44, 0x6f, 0x6e,
	0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xac, 0x03, 0x0a, 0x03,
	0x4d, 0x73, 0x67, 0x12, 0x80, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2f, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x37, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x98, 0x01, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x69, 0x63, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x1a, 0x3f, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x69, 0x63, 0x56, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x86, 0x01, 0x0a, 0x16, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x56,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x31, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x6f, 0x6e, 0x61, 0x74, 0x65, 0x41,
	0x6c, 0x6c, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a,
	0x39, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x6f, 0x6e, 0x61,
	0x74, 0x65, 0x41, 0x6c, 0x6c, 0x56, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x78, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x76, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_vesting_v1beta1_tx_proto_rawDescOnce sync.Once
	file_cosmos_vesting_v1beta1_tx_proto_rawDescData = file_cosmos_vesting_v1beta1_tx_proto_rawDesc
)

func file_cosmos_vesting_v1beta1_tx_proto_rawDescGZIP() []byte {
	file_cosmos_vesting_v1beta1_tx_proto_rawDescOnce.Do(func() {
		file_cosmos_vesting_v1beta1_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_vesting_v1beta1_tx_proto_rawDescData)
	})
	return file_cosmos_vesting_v1beta1_tx_proto_rawDescData
}

var file_cosmos_vesting_v1beta1_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cosmos_vesting_v1beta1_tx_proto_goTypes = []interface{}{
	(*MsgCreateVestingAccount)(nil),                 // 0: cosmos.vesting.v1beta1.MsgCreateVestingAccount
	(*MsgCreateVestingAccountResponse)(nil),         // 1: cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse
	(*MsgCreatePeriodicVestingAccount)(nil),         // 2: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount
	(*MsgCreatePeriodicVestingAccountResponse)(nil), // 3: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse
	(*MsgDonateAllVestingTokens)(nil),               // 4: cosmos.vesting.v1beta1.MsgDonateAllVestingTokens
	(*MsgDonateAllVestingTokensResponse)(nil),       // 5: cosmos.vesting.v1beta1.MsgDonateAllVestingTokensResponse
	(*types.Coin)(nil),                              // 6: cosmos.base.v1beta1.Coin
	(*Period)(nil),                                  // 7: cosmos.vesting.v1beta1.Period
}
var file_cosmos_vesting_v1beta1_tx_proto_depIdxs = []int32{
	6, // 0: cosmos.vesting.v1beta1.MsgCreateVestingAccount.amount:type_name -> cosmos.base.v1beta1.Coin
	7, // 1: cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount.vesting_periods:type_name -> cosmos.vesting.v1beta1.Period
	0, // 2: cosmos.vesting.v1beta1.Msg.CreateVestingAccount:input_type -> cosmos.vesting.v1beta1.MsgCreateVestingAccount
	2, // 3: cosmos.vesting.v1beta1.Msg.CreatePeriodicVestingAccount:input_type -> cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccount
	4, // 4: cosmos.vesting.v1beta1.Msg.DonateAllVestingTokens:input_type -> cosmos.vesting.v1beta1.MsgDonateAllVestingTokens
	1, // 5: cosmos.vesting.v1beta1.Msg.CreateVestingAccount:output_type -> cosmos.vesting.v1beta1.MsgCreateVestingAccountResponse
	3, // 6: cosmos.vesting.v1beta1.Msg.CreatePeriodicVestingAccount:output_type -> cosmos.vesting.v1beta1.MsgCreatePeriodicVestingAccountResponse
	5, // 7: cosmos.vesting.v1beta1.Msg.DonateAllVestingTokens:output_type -> cosmos.vesting.v1beta1.MsgDonateAllVestingTokensResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cosmos_vesting_v1beta1_tx_proto_init() }
func file_cosmos_vesting_v1beta1_tx_proto_init() {
	if File_cosmos_vesting_v1beta1_tx_proto != nil {
		return
	}
	file_cosmos_vesting_v1beta1_vesting_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cosmos_vesting_v1beta1_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateVestingAccount); i {
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
		file_cosmos_vesting_v1beta1_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateVestingAccountResponse); i {
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
		file_cosmos_vesting_v1beta1_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreatePeriodicVestingAccount); i {
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
		file_cosmos_vesting_v1beta1_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreatePeriodicVestingAccountResponse); i {
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
		file_cosmos_vesting_v1beta1_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDonateAllVestingTokens); i {
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
		file_cosmos_vesting_v1beta1_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDonateAllVestingTokensResponse); i {
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
			RawDescriptor: file_cosmos_vesting_v1beta1_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cosmos_vesting_v1beta1_tx_proto_goTypes,
		DependencyIndexes: file_cosmos_vesting_v1beta1_tx_proto_depIdxs,
		MessageInfos:      file_cosmos_vesting_v1beta1_tx_proto_msgTypes,
	}.Build()
	File_cosmos_vesting_v1beta1_tx_proto = out.File
	file_cosmos_vesting_v1beta1_tx_proto_rawDesc = nil
	file_cosmos_vesting_v1beta1_tx_proto_goTypes = nil
	file_cosmos_vesting_v1beta1_tx_proto_depIdxs = nil
}
