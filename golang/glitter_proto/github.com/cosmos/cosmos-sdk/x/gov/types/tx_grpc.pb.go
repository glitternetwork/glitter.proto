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
	// SubmitProposal defines a method to create new proposal given a content.
	SubmitProposal(ctx context.Context, in *MsgSubmitProposal, opts ...grpc.CallOption) (*MsgSubmitProposalResponse, error)
	// Vote defines a method to add a vote on a specific proposal.
	Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error)
	// VoteWeighted defines a method to add a weighted vote on a specific proposal.
	//
	// Since: cosmos-sdk 0.43
	VoteWeighted(ctx context.Context, in *MsgVoteWeighted, opts ...grpc.CallOption) (*MsgVoteWeightedResponse, error)
	// Deposit defines a method to add deposit on a specific proposal.
	Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitProposal(ctx context.Context, in *MsgSubmitProposal, opts ...grpc.CallOption) (*MsgSubmitProposalResponse, error) {
	out := new(MsgSubmitProposalResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/SubmitProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Vote(ctx context.Context, in *MsgVote, opts ...grpc.CallOption) (*MsgVoteResponse, error) {
	out := new(MsgVoteResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/Vote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) VoteWeighted(ctx context.Context, in *MsgVoteWeighted, opts ...grpc.CallOption) (*MsgVoteWeightedResponse, error) {
	out := new(MsgVoteWeightedResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/VoteWeighted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Deposit(ctx context.Context, in *MsgDeposit, opts ...grpc.CallOption) (*MsgDepositResponse, error) {
	out := new(MsgDepositResponse)
	err := c.cc.Invoke(ctx, "/cosmos.gov.v1beta1.Msg/Deposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// SubmitProposal defines a method to create new proposal given a content.
	SubmitProposal(context.Context, *MsgSubmitProposal) (*MsgSubmitProposalResponse, error)
	// Vote defines a method to add a vote on a specific proposal.
	Vote(context.Context, *MsgVote) (*MsgVoteResponse, error)
	// VoteWeighted defines a method to add a weighted vote on a specific proposal.
	//
	// Since: cosmos-sdk 0.43
	VoteWeighted(context.Context, *MsgVoteWeighted) (*MsgVoteWeightedResponse, error)
	// Deposit defines a method to add deposit on a specific proposal.
	Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) SubmitProposal(context.Context, *MsgSubmitProposal) (*MsgSubmitProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProposal not implemented")
}
func (UnimplementedMsgServer) Vote(context.Context, *MsgVote) (*MsgVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (UnimplementedMsgServer) VoteWeighted(context.Context, *MsgVoteWeighted) (*MsgVoteWeightedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteWeighted not implemented")
}
func (UnimplementedMsgServer) Deposit(context.Context, *MsgDeposit) (*MsgDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deposit not implemented")
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

func _Msg_SubmitProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitProposal)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/SubmitProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitProposal(ctx, req.(*MsgSubmitProposal))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/Vote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Vote(ctx, req.(*MsgVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_VoteWeighted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVoteWeighted)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VoteWeighted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/VoteWeighted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VoteWeighted(ctx, req.(*MsgVoteWeighted))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Deposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeposit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.gov.v1beta1.Msg/Deposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deposit(ctx, req.(*MsgDeposit))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.gov.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitProposal",
			Handler:    _Msg_SubmitProposal_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _Msg_Vote_Handler,
		},
		{
			MethodName: "VoteWeighted",
			Handler:    _Msg_VoteWeighted_Handler,
		},
		{
			MethodName: "Deposit",
			Handler:    _Msg_Deposit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/gov/v1beta1/tx.proto",
}
