// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// RegisterCounterpartyAddress defines a rpc handler method for MsgRegisterCounterpartyAddress
	// RegisterCounterpartyAddress is called by the relayer on each channelEnd and allows them to specify their
	// counterparty address before relaying. This ensures they will be properly compensated for forward relaying since
	// destination chain must send back relayer's source address (counterparty address) in acknowledgement. This function
	// may be called more than once by a relayer, in which case, latest counterparty address is always used.
	RegisterCounterpartyAddress(ctx context.Context, in *MsgRegisterCounterpartyAddress, opts ...grpc.CallOption) (*MsgRegisterCounterpartyAddressResponse, error)
	// PayPacketFee defines a rpc handler method for MsgPayPacketFee
	// PayPacketFee is an open callback that may be called by any module/user that wishes to escrow funds in order to
	// incentivize the relaying of the packet at the next sequence
	// NOTE: This method is intended to be used within a multi msg transaction, where the subsequent msg that follows
	// initiates the lifecycle of the incentivized packet
	PayPacketFee(ctx context.Context, in *MsgPayPacketFee, opts ...grpc.CallOption) (*MsgPayPacketFeeResponse, error)
	// PayPacketFeeAsync defines a rpc handler method for MsgPayPacketFeeAsync
	// PayPacketFeeAsync is an open callback that may be called by any module/user that wishes to escrow funds in order to
	// incentivize the relaying of a known packet (i.e. at a particular sequence)
	PayPacketFeeAsync(ctx context.Context, in *MsgPayPacketFeeAsync, opts ...grpc.CallOption) (*MsgPayPacketFeeAsyncResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterCounterpartyAddress(ctx context.Context, in *MsgRegisterCounterpartyAddress, opts ...grpc.CallOption) (*MsgRegisterCounterpartyAddressResponse, error) {
	out := new(MsgRegisterCounterpartyAddressResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Msg/RegisterCounterpartyAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PayPacketFee(ctx context.Context, in *MsgPayPacketFee, opts ...grpc.CallOption) (*MsgPayPacketFeeResponse, error) {
	out := new(MsgPayPacketFeeResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Msg/PayPacketFee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) PayPacketFeeAsync(ctx context.Context, in *MsgPayPacketFeeAsync, opts ...grpc.CallOption) (*MsgPayPacketFeeAsyncResponse, error) {
	out := new(MsgPayPacketFeeAsyncResponse)
	err := c.cc.Invoke(ctx, "/ibc.applications.fee.v1.Msg/PayPacketFeeAsync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// RegisterCounterpartyAddress defines a rpc handler method for MsgRegisterCounterpartyAddress
	// RegisterCounterpartyAddress is called by the relayer on each channelEnd and allows them to specify their
	// counterparty address before relaying. This ensures they will be properly compensated for forward relaying since
	// destination chain must send back relayer's source address (counterparty address) in acknowledgement. This function
	// may be called more than once by a relayer, in which case, latest counterparty address is always used.
	RegisterCounterpartyAddress(context.Context, *MsgRegisterCounterpartyAddress) (*MsgRegisterCounterpartyAddressResponse, error)
	// PayPacketFee defines a rpc handler method for MsgPayPacketFee
	// PayPacketFee is an open callback that may be called by any module/user that wishes to escrow funds in order to
	// incentivize the relaying of the packet at the next sequence
	// NOTE: This method is intended to be used within a multi msg transaction, where the subsequent msg that follows
	// initiates the lifecycle of the incentivized packet
	PayPacketFee(context.Context, *MsgPayPacketFee) (*MsgPayPacketFeeResponse, error)
	// PayPacketFeeAsync defines a rpc handler method for MsgPayPacketFeeAsync
	// PayPacketFeeAsync is an open callback that may be called by any module/user that wishes to escrow funds in order to
	// incentivize the relaying of a known packet (i.e. at a particular sequence)
	PayPacketFeeAsync(context.Context, *MsgPayPacketFeeAsync) (*MsgPayPacketFeeAsyncResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) RegisterCounterpartyAddress(context.Context, *MsgRegisterCounterpartyAddress) (*MsgRegisterCounterpartyAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCounterpartyAddress not implemented")
}
func (UnimplementedMsgServer) PayPacketFee(context.Context, *MsgPayPacketFee) (*MsgPayPacketFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayPacketFee not implemented")
}
func (UnimplementedMsgServer) PayPacketFeeAsync(context.Context, *MsgPayPacketFeeAsync) (*MsgPayPacketFeeAsyncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayPacketFeeAsync not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterCounterpartyAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterCounterpartyAddress)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterCounterpartyAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Msg/RegisterCounterpartyAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterCounterpartyAddress(ctx, req.(*MsgRegisterCounterpartyAddress))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PayPacketFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPayPacketFee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PayPacketFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Msg/PayPacketFee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PayPacketFee(ctx, req.(*MsgPayPacketFee))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_PayPacketFeeAsync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgPayPacketFeeAsync)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).PayPacketFeeAsync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.applications.fee.v1.Msg/PayPacketFeeAsync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).PayPacketFeeAsync(ctx, req.(*MsgPayPacketFeeAsync))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.applications.fee.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterCounterpartyAddress",
			Handler:    _Msg_RegisterCounterpartyAddress_Handler,
		},
		{
			MethodName: "PayPacketFee",
			Handler:    _Msg_PayPacketFee_Handler,
		},
		{
			MethodName: "PayPacketFeeAsync",
			Handler:    _Msg_PayPacketFeeAsync_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/applications/fee/v1/tx.proto",
}
