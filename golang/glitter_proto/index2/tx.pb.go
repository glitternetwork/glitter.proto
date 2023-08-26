// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0-devel
// 	protoc        v3.14.0
// source: index/tx.proto

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

type MsgSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress string `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	SchemaName   string `protobuf:"bytes,2,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	Body         string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MsgSchema) Reset() {
	*x = MsgSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSchema) ProtoMessage() {}

func (x *MsgSchema) ProtoReflect() protoreflect.Message {
	mi := &file_index_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSchema.ProtoReflect.Descriptor instead.
func (*MsgSchema) Descriptor() ([]byte, []int) {
	return file_index_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgSchema) GetOwnerAddress() string {
	if x != nil {
		return x.OwnerAddress
	}
	return ""
}

func (x *MsgSchema) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *MsgSchema) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type MsgSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgSchemaResponse) Reset() {
	*x = MsgSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSchemaResponse) ProtoMessage() {}

func (x *MsgSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_index_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgSchemaResponse.ProtoReflect.Descriptor instead.
func (*MsgSchemaResponse) Descriptor() ([]byte, []int) {
	return file_index_tx_proto_rawDescGZIP(), []int{1}
}

type MsgDoc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress string `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	SchemaName   string `protobuf:"bytes,2,opt,name=schema_name,json=schemaName,proto3" json:"schema_name,omitempty"`
	Body         string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *MsgDoc) Reset() {
	*x = MsgDoc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDoc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDoc) ProtoMessage() {}

func (x *MsgDoc) ProtoReflect() protoreflect.Message {
	mi := &file_index_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDoc.ProtoReflect.Descriptor instead.
func (*MsgDoc) Descriptor() ([]byte, []int) {
	return file_index_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgDoc) GetOwnerAddress() string {
	if x != nil {
		return x.OwnerAddress
	}
	return ""
}

func (x *MsgDoc) GetSchemaName() string {
	if x != nil {
		return x.SchemaName
	}
	return ""
}

func (x *MsgDoc) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type MsgDocResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgDocResponse) Reset() {
	*x = MsgDocResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgDocResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgDocResponse) ProtoMessage() {}

func (x *MsgDocResponse) ProtoReflect() protoreflect.Message {
	mi := &file_index_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgDocResponse.ProtoReflect.Descriptor instead.
func (*MsgDocResponse) Descriptor() ([]byte, []int) {
	return file_index_tx_proto_rawDescGZIP(), []int{3}
}

type SQLExecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid       string      `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Sql       string      `protobuf:"bytes,2,opt,name=sql,proto3" json:"sql,omitempty"`
	Arguments []*Argument `protobuf:"bytes,3,rep,name=arguments,proto3" json:"arguments,omitempty"`
}

func (x *SQLExecRequest) Reset() {
	*x = SQLExecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_tx_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLExecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLExecRequest) ProtoMessage() {}

func (x *SQLExecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_index_tx_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLExecRequest.ProtoReflect.Descriptor instead.
func (*SQLExecRequest) Descriptor() ([]byte, []int) {
	return file_index_tx_proto_rawDescGZIP(), []int{4}
}

func (x *SQLExecRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *SQLExecRequest) GetSql() string {
	if x != nil {
		return x.Sql
	}
	return ""
}

func (x *SQLExecRequest) GetArguments() []*Argument {
	if x != nil {
		return x.Arguments
	}
	return nil
}

type SQLExecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results   []*ResultSet `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	TookTimes float32      `protobuf:"fixed32,2,opt,name=took_times,json=tookTimes,proto3" json:"took_times,omitempty"`
}

func (x *SQLExecResponse) Reset() {
	*x = SQLExecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_tx_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLExecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLExecResponse) ProtoMessage() {}

func (x *SQLExecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_index_tx_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLExecResponse.ProtoReflect.Descriptor instead.
func (*SQLExecResponse) Descriptor() ([]byte, []int) {
	return file_index_tx_proto_rawDescGZIP(), []int{5}
}

func (x *SQLExecResponse) GetResults() []*ResultSet {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *SQLExecResponse) GetTookTimes() float32 {
	if x != nil {
		return x.TookTimes
	}
	return 0
}

type SQLGrantRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid        string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	OnTable    string `protobuf:"bytes,2,opt,name=onTable,proto3" json:"onTable,omitempty"`
	ToUID      string `protobuf:"bytes,3,opt,name=toUID,proto3" json:"toUID,omitempty"`
	Role       string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	OnDatabase string `protobuf:"bytes,5,opt,name=onDatabase,proto3" json:"onDatabase,omitempty"`
}

func (x *SQLGrantRequest) Reset() {
	*x = SQLGrantRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_tx_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLGrantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLGrantRequest) ProtoMessage() {}

func (x *SQLGrantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_index_tx_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLGrantRequest.ProtoReflect.Descriptor instead.
func (*SQLGrantRequest) Descriptor() ([]byte, []int) {
	return file_index_tx_proto_rawDescGZIP(), []int{6}
}

func (x *SQLGrantRequest) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *SQLGrantRequest) GetOnTable() string {
	if x != nil {
		return x.OnTable
	}
	return ""
}

func (x *SQLGrantRequest) GetToUID() string {
	if x != nil {
		return x.ToUID
	}
	return ""
}

func (x *SQLGrantRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *SQLGrantRequest) GetOnDatabase() string {
	if x != nil {
		return x.OnDatabase
	}
	return ""
}

type SQLGrantResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SQLGrantResponse) Reset() {
	*x = SQLGrantResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_index_tx_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SQLGrantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SQLGrantResponse) ProtoMessage() {}

func (x *SQLGrantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_index_tx_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SQLGrantResponse.ProtoReflect.Descriptor instead.
func (*SQLGrantResponse) Descriptor() ([]byte, []int) {
	return file_index_tx_proto_rawDescGZIP(), []int{7}
}

var File_index_tx_proto protoreflect.FileDescriptor

var file_index_tx_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x76, 0x65, 0x64, 0x2e, 0x67, 0x6c, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x14, 0x67,
	0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x73, 0x71, 0x6c, 0x5f, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x01, 0x0a, 0x09,
	0x4d, 0x73, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x3d, 0x0a, 0x0d, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0xf2, 0xde, 0x1f, 0x14, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x37, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xf2,
	0xde, 0x1f, 0x12, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0f, 0xf2, 0xde, 0x1f, 0x0b, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x62, 0x6f, 0x64, 0x79, 0x22,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x08, 0x88, 0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00,
	0x22, 0x13, 0x0a, 0x11, 0x4d, 0x73, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x06, 0x4d, 0x73, 0x67, 0x44, 0x6f, 0x63,
	0x12, 0x3d, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xf2, 0xde, 0x1f, 0x14, 0x79, 0x61, 0x6d,
	0x6c, 0x3a, 0x22, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x37, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xf2, 0xde, 0x1f, 0x12, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xf2, 0xde, 0x1f, 0x0b, 0x79, 0x61, 0x6d, 0x6c,
	0x3a, 0x22, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x3a, 0x08, 0x88,
	0xa0, 0x1f, 0x00, 0xe8, 0xa0, 0x1f, 0x00, 0x22, 0x10, 0x0a, 0x0e, 0x4d, 0x73, 0x67, 0x44, 0x6f,
	0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x79, 0x0a, 0x0e, 0x53, 0x51, 0x4c,
	0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x73, 0x71, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x71, 0x6c, 0x12,
	0x43, 0x0a, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x76, 0x65, 0x64, 0x2e, 0x67, 0x6c,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x41, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x61, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x72, 0x0a, 0x0f, 0x53, 0x51, 0x4c, 0x45, 0x78, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x76, 0x65, 0x64, 0x2e, 0x67, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74,
	0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6f,
	0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74,
	0x6f, 0x6f, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x0f, 0x53, 0x51, 0x4c,
	0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6f, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6f, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x55, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x55, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x53, 0x51, 0x4c, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9b, 0x03, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x66,
	0x0a, 0x0c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x26,
	0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x76, 0x65, 0x64, 0x2e, 0x67, 0x6c, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x4d, 0x73, 0x67,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a, 0x2e, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x76, 0x65,
	0x64, 0x2e, 0x67, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x09, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x44, 0x6f, 0x63, 0x12, 0x23, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x76, 0x65, 0x64, 0x2e, 0x67,
	0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x6f, 0x63, 0x1a, 0x2b, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x76, 0x65, 0x64, 0x2e, 0x67, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x4d, 0x73, 0x67, 0x44, 0x6f, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x07, 0x53, 0x51, 0x4c, 0x45, 0x78, 0x65, 0x63,
	0x12, 0x2b, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x76, 0x65, 0x64, 0x2e, 0x67, 0x6c, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x53,
	0x51, 0x4c, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x76, 0x65, 0x64, 0x2e, 0x67, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x53, 0x51, 0x4c, 0x45,
	0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x08, 0x53,
	0x51, 0x4c, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x2c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x76,
	0x65, 0x64, 0x2e, 0x67, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x53, 0x51, 0x4c, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x76, 0x65, 0x64,
	0x2e, 0x67, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x53, 0x51, 0x4c, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x76, 0x65, 0x64, 0x2f, 0x67, 0x6c, 0x69, 0x74,
	0x74, 0x65, 0x72, 0x2d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x78, 0x2f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_index_tx_proto_rawDescOnce sync.Once
	file_index_tx_proto_rawDescData = file_index_tx_proto_rawDesc
)

func file_index_tx_proto_rawDescGZIP() []byte {
	file_index_tx_proto_rawDescOnce.Do(func() {
		file_index_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_index_tx_proto_rawDescData)
	})
	return file_index_tx_proto_rawDescData
}

var file_index_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_index_tx_proto_goTypes = []interface{}{
	(*MsgSchema)(nil),         // 0: blockved.glitterchain.index.MsgSchema
	(*MsgSchemaResponse)(nil), // 1: blockved.glitterchain.index.MsgSchemaResponse
	(*MsgDoc)(nil),            // 2: blockved.glitterchain.index.MsgDoc
	(*MsgDocResponse)(nil),    // 3: blockved.glitterchain.index.MsgDocResponse
	(*SQLExecRequest)(nil),    // 4: blockved.glitterchain.index.SQLExecRequest
	(*SQLExecResponse)(nil),   // 5: blockved.glitterchain.index.SQLExecResponse
	(*SQLGrantRequest)(nil),   // 6: blockved.glitterchain.index.SQLGrantRequest
	(*SQLGrantResponse)(nil),  // 7: blockved.glitterchain.index.SQLGrantResponse
	(*Argument)(nil),          // 8: blockved.glitterchain.index.Argument
	(*ResultSet)(nil),         // 9: blockved.glitterchain.index.ResultSet
}
var file_index_tx_proto_depIdxs = []int32{
	8, // 0: blockved.glitterchain.index.SQLExecRequest.arguments:type_name -> blockved.glitterchain.index.Argument
	9, // 1: blockved.glitterchain.index.SQLExecResponse.results:type_name -> blockved.glitterchain.index.ResultSet
	0, // 2: blockved.glitterchain.index.Msg.ManageSchema:input_type -> blockved.glitterchain.index.MsgSchema
	2, // 3: blockved.glitterchain.index.Msg.ManageDoc:input_type -> blockved.glitterchain.index.MsgDoc
	4, // 4: blockved.glitterchain.index.Msg.SQLExec:input_type -> blockved.glitterchain.index.SQLExecRequest
	6, // 5: blockved.glitterchain.index.Msg.SQLGrant:input_type -> blockved.glitterchain.index.SQLGrantRequest
	1, // 6: blockved.glitterchain.index.Msg.ManageSchema:output_type -> blockved.glitterchain.index.MsgSchemaResponse
	3, // 7: blockved.glitterchain.index.Msg.ManageDoc:output_type -> blockved.glitterchain.index.MsgDocResponse
	5, // 8: blockved.glitterchain.index.Msg.SQLExec:output_type -> blockved.glitterchain.index.SQLExecResponse
	7, // 9: blockved.glitterchain.index.Msg.SQLGrant:output_type -> blockved.glitterchain.index.SQLGrantResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_index_tx_proto_init() }
func file_index_tx_proto_init() {
	if File_index_tx_proto != nil {
		return
	}
	file_index_sql_engine_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_index_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSchema); i {
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
		file_index_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSchemaResponse); i {
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
		file_index_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDoc); i {
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
		file_index_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgDocResponse); i {
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
		file_index_tx_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLExecRequest); i {
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
		file_index_tx_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLExecResponse); i {
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
		file_index_tx_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLGrantRequest); i {
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
		file_index_tx_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SQLGrantResponse); i {
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
			RawDescriptor: file_index_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_index_tx_proto_goTypes,
		DependencyIndexes: file_index_tx_proto_depIdxs,
		MessageInfos:      file_index_tx_proto_msgTypes,
	}.Build()
	File_index_tx_proto = out.File
	file_index_tx_proto_rawDesc = nil
	file_index_tx_proto_goTypes = nil
	file_index_tx_proto_depIdxs = nil
}