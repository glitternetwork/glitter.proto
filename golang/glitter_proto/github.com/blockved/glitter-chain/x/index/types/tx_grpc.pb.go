// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: index/tx.proto

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Msg_ManageSchema_FullMethodName = "/blockved.glitterchain.index.Msg/ManageSchema"
	Msg_ManageDoc_FullMethodName    = "/blockved.glitterchain.index.Msg/ManageDoc"
	Msg_SQLExec_FullMethodName      = "/blockved.glitterchain.index.Msg/SQLExec"
	Msg_SQLGrant_FullMethodName     = "/blockved.glitterchain.index.Msg/SQLGrant"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	ManageSchema(ctx context.Context, in *MsgSchema, opts ...grpc.CallOption) (*MsgSchemaResponse, error)
	ManageDoc(ctx context.Context, in *MsgDoc, opts ...grpc.CallOption) (*MsgDocResponse, error)
	SQLExec(ctx context.Context, in *SQLExecRequest, opts ...grpc.CallOption) (*SQLExecResponse, error)
	SQLGrant(ctx context.Context, in *SQLGrantRequest, opts ...grpc.CallOption) (*SQLGrantResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) ManageSchema(ctx context.Context, in *MsgSchema, opts ...grpc.CallOption) (*MsgSchemaResponse, error) {
	out := new(MsgSchemaResponse)
	err := c.cc.Invoke(ctx, Msg_ManageSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ManageDoc(ctx context.Context, in *MsgDoc, opts ...grpc.CallOption) (*MsgDocResponse, error) {
	out := new(MsgDocResponse)
	err := c.cc.Invoke(ctx, Msg_ManageDoc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SQLExec(ctx context.Context, in *SQLExecRequest, opts ...grpc.CallOption) (*SQLExecResponse, error) {
	out := new(SQLExecResponse)
	err := c.cc.Invoke(ctx, Msg_SQLExec_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SQLGrant(ctx context.Context, in *SQLGrantRequest, opts ...grpc.CallOption) (*SQLGrantResponse, error) {
	out := new(SQLGrantResponse)
	err := c.cc.Invoke(ctx, Msg_SQLGrant_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	ManageSchema(context.Context, *MsgSchema) (*MsgSchemaResponse, error)
	ManageDoc(context.Context, *MsgDoc) (*MsgDocResponse, error)
	SQLExec(context.Context, *SQLExecRequest) (*SQLExecResponse, error)
	SQLGrant(context.Context, *SQLGrantRequest) (*SQLGrantResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) ManageSchema(context.Context, *MsgSchema) (*MsgSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageSchema not implemented")
}
func (UnimplementedMsgServer) ManageDoc(context.Context, *MsgDoc) (*MsgDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManageDoc not implemented")
}
func (UnimplementedMsgServer) SQLExec(context.Context, *SQLExecRequest) (*SQLExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SQLExec not implemented")
}
func (UnimplementedMsgServer) SQLGrant(context.Context, *SQLGrantRequest) (*SQLGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SQLGrant not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_ManageSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSchema)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ManageSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ManageSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ManageSchema(ctx, req.(*MsgSchema))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ManageDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDoc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ManageDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ManageDoc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ManageDoc(ctx, req.(*MsgDoc))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SQLExec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SQLExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SQLExec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SQLExec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SQLExec(ctx, req.(*SQLExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SQLGrant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SQLGrantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SQLGrant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SQLGrant_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SQLGrant(ctx, req.(*SQLGrantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blockved.glitterchain.index.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ManageSchema",
			Handler:    _Msg_ManageSchema_Handler,
		},
		{
			MethodName: "ManageDoc",
			Handler:    _Msg_ManageDoc_Handler,
		},
		{
			MethodName: "SQLExec",
			Handler:    _Msg_SQLExec_Handler,
		},
		{
			MethodName: "SQLGrant",
			Handler:    _Msg_SQLGrant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index/tx.proto",
}
