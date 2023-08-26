// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: cosmwasm/wasm/v1/tx.proto

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
	Msg_StoreCode_FullMethodName           = "/cosmwasm.wasm.v1.Msg/StoreCode"
	Msg_InstantiateContract_FullMethodName = "/cosmwasm.wasm.v1.Msg/InstantiateContract"
	Msg_ExecuteContract_FullMethodName     = "/cosmwasm.wasm.v1.Msg/ExecuteContract"
	Msg_MigrateContract_FullMethodName     = "/cosmwasm.wasm.v1.Msg/MigrateContract"
	Msg_UpdateAdmin_FullMethodName         = "/cosmwasm.wasm.v1.Msg/UpdateAdmin"
	Msg_ClearAdmin_FullMethodName          = "/cosmwasm.wasm.v1.Msg/ClearAdmin"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// StoreCode to submit Wasm code to the system
	StoreCode(ctx context.Context, in *MsgStoreCode, opts ...grpc.CallOption) (*MsgStoreCodeResponse, error)
	// Instantiate creates a new smart contract instance for the given code id.
	InstantiateContract(ctx context.Context, in *MsgInstantiateContract, opts ...grpc.CallOption) (*MsgInstantiateContractResponse, error)
	// Execute submits the given message data to a smart contract
	ExecuteContract(ctx context.Context, in *MsgExecuteContract, opts ...grpc.CallOption) (*MsgExecuteContractResponse, error)
	// Migrate runs a code upgrade/ downgrade for a smart contract
	MigrateContract(ctx context.Context, in *MsgMigrateContract, opts ...grpc.CallOption) (*MsgMigrateContractResponse, error)
	// UpdateAdmin sets a new   admin for a smart contract
	UpdateAdmin(ctx context.Context, in *MsgUpdateAdmin, opts ...grpc.CallOption) (*MsgUpdateAdminResponse, error)
	// ClearAdmin removes any admin stored for a smart contract
	ClearAdmin(ctx context.Context, in *MsgClearAdmin, opts ...grpc.CallOption) (*MsgClearAdminResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) StoreCode(ctx context.Context, in *MsgStoreCode, opts ...grpc.CallOption) (*MsgStoreCodeResponse, error) {
	out := new(MsgStoreCodeResponse)
	err := c.cc.Invoke(ctx, Msg_StoreCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) InstantiateContract(ctx context.Context, in *MsgInstantiateContract, opts ...grpc.CallOption) (*MsgInstantiateContractResponse, error) {
	out := new(MsgInstantiateContractResponse)
	err := c.cc.Invoke(ctx, Msg_InstantiateContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExecuteContract(ctx context.Context, in *MsgExecuteContract, opts ...grpc.CallOption) (*MsgExecuteContractResponse, error) {
	out := new(MsgExecuteContractResponse)
	err := c.cc.Invoke(ctx, Msg_ExecuteContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MigrateContract(ctx context.Context, in *MsgMigrateContract, opts ...grpc.CallOption) (*MsgMigrateContractResponse, error) {
	out := new(MsgMigrateContractResponse)
	err := c.cc.Invoke(ctx, Msg_MigrateContract_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateAdmin(ctx context.Context, in *MsgUpdateAdmin, opts ...grpc.CallOption) (*MsgUpdateAdminResponse, error) {
	out := new(MsgUpdateAdminResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ClearAdmin(ctx context.Context, in *MsgClearAdmin, opts ...grpc.CallOption) (*MsgClearAdminResponse, error) {
	out := new(MsgClearAdminResponse)
	err := c.cc.Invoke(ctx, Msg_ClearAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// StoreCode to submit Wasm code to the system
	StoreCode(context.Context, *MsgStoreCode) (*MsgStoreCodeResponse, error)
	// Instantiate creates a new smart contract instance for the given code id.
	InstantiateContract(context.Context, *MsgInstantiateContract) (*MsgInstantiateContractResponse, error)
	// Execute submits the given message data to a smart contract
	ExecuteContract(context.Context, *MsgExecuteContract) (*MsgExecuteContractResponse, error)
	// Migrate runs a code upgrade/ downgrade for a smart contract
	MigrateContract(context.Context, *MsgMigrateContract) (*MsgMigrateContractResponse, error)
	// UpdateAdmin sets a new   admin for a smart contract
	UpdateAdmin(context.Context, *MsgUpdateAdmin) (*MsgUpdateAdminResponse, error)
	// ClearAdmin removes any admin stored for a smart contract
	ClearAdmin(context.Context, *MsgClearAdmin) (*MsgClearAdminResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) StoreCode(context.Context, *MsgStoreCode) (*MsgStoreCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreCode not implemented")
}
func (UnimplementedMsgServer) InstantiateContract(context.Context, *MsgInstantiateContract) (*MsgInstantiateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstantiateContract not implemented")
}
func (UnimplementedMsgServer) ExecuteContract(context.Context, *MsgExecuteContract) (*MsgExecuteContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteContract not implemented")
}
func (UnimplementedMsgServer) MigrateContract(context.Context, *MsgMigrateContract) (*MsgMigrateContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MigrateContract not implemented")
}
func (UnimplementedMsgServer) UpdateAdmin(context.Context, *MsgUpdateAdmin) (*MsgUpdateAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAdmin not implemented")
}
func (UnimplementedMsgServer) ClearAdmin(context.Context, *MsgClearAdmin) (*MsgClearAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearAdmin not implemented")
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

func _Msg_StoreCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStoreCode)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StoreCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_StoreCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StoreCode(ctx, req.(*MsgStoreCode))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_InstantiateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgInstantiateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).InstantiateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_InstantiateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).InstantiateContract(ctx, req.(*MsgInstantiateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExecuteContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecuteContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecuteContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ExecuteContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecuteContract(ctx, req.(*MsgExecuteContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MigrateContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMigrateContract)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MigrateContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_MigrateContract_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MigrateContract(ctx, req.(*MsgMigrateContract))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateAdmin(ctx, req.(*MsgUpdateAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ClearAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgClearAdmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ClearAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ClearAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ClearAdmin(ctx, req.(*MsgClearAdmin))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmwasm.wasm.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreCode",
			Handler:    _Msg_StoreCode_Handler,
		},
		{
			MethodName: "InstantiateContract",
			Handler:    _Msg_InstantiateContract_Handler,
		},
		{
			MethodName: "ExecuteContract",
			Handler:    _Msg_ExecuteContract_Handler,
		},
		{
			MethodName: "MigrateContract",
			Handler:    _Msg_MigrateContract_Handler,
		},
		{
			MethodName: "UpdateAdmin",
			Handler:    _Msg_UpdateAdmin_Handler,
		},
		{
			MethodName: "ClearAdmin",
			Handler:    _Msg_ClearAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmwasm/wasm/v1/tx.proto",
}