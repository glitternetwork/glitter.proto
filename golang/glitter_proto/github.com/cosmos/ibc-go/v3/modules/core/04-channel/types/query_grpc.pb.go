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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Channel queries an IBC Channel.
	Channel(ctx context.Context, in *QueryChannelRequest, opts ...grpc.CallOption) (*QueryChannelResponse, error)
	// Channels queries all the IBC channels of a chain.
	Channels(ctx context.Context, in *QueryChannelsRequest, opts ...grpc.CallOption) (*QueryChannelsResponse, error)
	// ConnectionChannels queries all the channels associated with a connection
	// end.
	ConnectionChannels(ctx context.Context, in *QueryConnectionChannelsRequest, opts ...grpc.CallOption) (*QueryConnectionChannelsResponse, error)
	// ChannelClientState queries for the client state for the channel associated
	// with the provided channel identifiers.
	ChannelClientState(ctx context.Context, in *QueryChannelClientStateRequest, opts ...grpc.CallOption) (*QueryChannelClientStateResponse, error)
	// ChannelConsensusState queries for the consensus state for the channel
	// associated with the provided channel identifiers.
	ChannelConsensusState(ctx context.Context, in *QueryChannelConsensusStateRequest, opts ...grpc.CallOption) (*QueryChannelConsensusStateResponse, error)
	// PacketCommitment queries a stored packet commitment hash.
	PacketCommitment(ctx context.Context, in *QueryPacketCommitmentRequest, opts ...grpc.CallOption) (*QueryPacketCommitmentResponse, error)
	// PacketCommitments returns all the packet commitments hashes associated
	// with a channel.
	PacketCommitments(ctx context.Context, in *QueryPacketCommitmentsRequest, opts ...grpc.CallOption) (*QueryPacketCommitmentsResponse, error)
	// PacketReceipt queries if a given packet sequence has been received on the
	// queried chain
	PacketReceipt(ctx context.Context, in *QueryPacketReceiptRequest, opts ...grpc.CallOption) (*QueryPacketReceiptResponse, error)
	// PacketAcknowledgement queries a stored packet acknowledgement hash.
	PacketAcknowledgement(ctx context.Context, in *QueryPacketAcknowledgementRequest, opts ...grpc.CallOption) (*QueryPacketAcknowledgementResponse, error)
	// PacketAcknowledgements returns all the packet acknowledgements associated
	// with a channel.
	PacketAcknowledgements(ctx context.Context, in *QueryPacketAcknowledgementsRequest, opts ...grpc.CallOption) (*QueryPacketAcknowledgementsResponse, error)
	// UnreceivedPackets returns all the unreceived IBC packets associated with a
	// channel and sequences.
	UnreceivedPackets(ctx context.Context, in *QueryUnreceivedPacketsRequest, opts ...grpc.CallOption) (*QueryUnreceivedPacketsResponse, error)
	// UnreceivedAcks returns all the unreceived IBC acknowledgements associated
	// with a channel and sequences.
	UnreceivedAcks(ctx context.Context, in *QueryUnreceivedAcksRequest, opts ...grpc.CallOption) (*QueryUnreceivedAcksResponse, error)
	// NextSequenceReceive returns the next receive sequence for a given channel.
	NextSequenceReceive(ctx context.Context, in *QueryNextSequenceReceiveRequest, opts ...grpc.CallOption) (*QueryNextSequenceReceiveResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Channel(ctx context.Context, in *QueryChannelRequest, opts ...grpc.CallOption) (*QueryChannelResponse, error) {
	out := new(QueryChannelResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/Channel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Channels(ctx context.Context, in *QueryChannelsRequest, opts ...grpc.CallOption) (*QueryChannelsResponse, error) {
	out := new(QueryChannelsResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/Channels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ConnectionChannels(ctx context.Context, in *QueryConnectionChannelsRequest, opts ...grpc.CallOption) (*QueryConnectionChannelsResponse, error) {
	out := new(QueryConnectionChannelsResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/ConnectionChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChannelClientState(ctx context.Context, in *QueryChannelClientStateRequest, opts ...grpc.CallOption) (*QueryChannelClientStateResponse, error) {
	out := new(QueryChannelClientStateResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/ChannelClientState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChannelConsensusState(ctx context.Context, in *QueryChannelConsensusStateRequest, opts ...grpc.CallOption) (*QueryChannelConsensusStateResponse, error) {
	out := new(QueryChannelConsensusStateResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/ChannelConsensusState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PacketCommitment(ctx context.Context, in *QueryPacketCommitmentRequest, opts ...grpc.CallOption) (*QueryPacketCommitmentResponse, error) {
	out := new(QueryPacketCommitmentResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/PacketCommitment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PacketCommitments(ctx context.Context, in *QueryPacketCommitmentsRequest, opts ...grpc.CallOption) (*QueryPacketCommitmentsResponse, error) {
	out := new(QueryPacketCommitmentsResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/PacketCommitments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PacketReceipt(ctx context.Context, in *QueryPacketReceiptRequest, opts ...grpc.CallOption) (*QueryPacketReceiptResponse, error) {
	out := new(QueryPacketReceiptResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/PacketReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PacketAcknowledgement(ctx context.Context, in *QueryPacketAcknowledgementRequest, opts ...grpc.CallOption) (*QueryPacketAcknowledgementResponse, error) {
	out := new(QueryPacketAcknowledgementResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/PacketAcknowledgement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PacketAcknowledgements(ctx context.Context, in *QueryPacketAcknowledgementsRequest, opts ...grpc.CallOption) (*QueryPacketAcknowledgementsResponse, error) {
	out := new(QueryPacketAcknowledgementsResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/PacketAcknowledgements", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UnreceivedPackets(ctx context.Context, in *QueryUnreceivedPacketsRequest, opts ...grpc.CallOption) (*QueryUnreceivedPacketsResponse, error) {
	out := new(QueryUnreceivedPacketsResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/UnreceivedPackets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) UnreceivedAcks(ctx context.Context, in *QueryUnreceivedAcksRequest, opts ...grpc.CallOption) (*QueryUnreceivedAcksResponse, error) {
	out := new(QueryUnreceivedAcksResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/UnreceivedAcks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) NextSequenceReceive(ctx context.Context, in *QueryNextSequenceReceiveRequest, opts ...grpc.CallOption) (*QueryNextSequenceReceiveResponse, error) {
	out := new(QueryNextSequenceReceiveResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.channel.v1.Query/NextSequenceReceive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Channel queries an IBC Channel.
	Channel(context.Context, *QueryChannelRequest) (*QueryChannelResponse, error)
	// Channels queries all the IBC channels of a chain.
	Channels(context.Context, *QueryChannelsRequest) (*QueryChannelsResponse, error)
	// ConnectionChannels queries all the channels associated with a connection
	// end.
	ConnectionChannels(context.Context, *QueryConnectionChannelsRequest) (*QueryConnectionChannelsResponse, error)
	// ChannelClientState queries for the client state for the channel associated
	// with the provided channel identifiers.
	ChannelClientState(context.Context, *QueryChannelClientStateRequest) (*QueryChannelClientStateResponse, error)
	// ChannelConsensusState queries for the consensus state for the channel
	// associated with the provided channel identifiers.
	ChannelConsensusState(context.Context, *QueryChannelConsensusStateRequest) (*QueryChannelConsensusStateResponse, error)
	// PacketCommitment queries a stored packet commitment hash.
	PacketCommitment(context.Context, *QueryPacketCommitmentRequest) (*QueryPacketCommitmentResponse, error)
	// PacketCommitments returns all the packet commitments hashes associated
	// with a channel.
	PacketCommitments(context.Context, *QueryPacketCommitmentsRequest) (*QueryPacketCommitmentsResponse, error)
	// PacketReceipt queries if a given packet sequence has been received on the
	// queried chain
	PacketReceipt(context.Context, *QueryPacketReceiptRequest) (*QueryPacketReceiptResponse, error)
	// PacketAcknowledgement queries a stored packet acknowledgement hash.
	PacketAcknowledgement(context.Context, *QueryPacketAcknowledgementRequest) (*QueryPacketAcknowledgementResponse, error)
	// PacketAcknowledgements returns all the packet acknowledgements associated
	// with a channel.
	PacketAcknowledgements(context.Context, *QueryPacketAcknowledgementsRequest) (*QueryPacketAcknowledgementsResponse, error)
	// UnreceivedPackets returns all the unreceived IBC packets associated with a
	// channel and sequences.
	UnreceivedPackets(context.Context, *QueryUnreceivedPacketsRequest) (*QueryUnreceivedPacketsResponse, error)
	// UnreceivedAcks returns all the unreceived IBC acknowledgements associated
	// with a channel and sequences.
	UnreceivedAcks(context.Context, *QueryUnreceivedAcksRequest) (*QueryUnreceivedAcksResponse, error)
	// NextSequenceReceive returns the next receive sequence for a given channel.
	NextSequenceReceive(context.Context, *QueryNextSequenceReceiveRequest) (*QueryNextSequenceReceiveResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Channel(context.Context, *QueryChannelRequest) (*QueryChannelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Channel not implemented")
}
func (UnimplementedQueryServer) Channels(context.Context, *QueryChannelsRequest) (*QueryChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Channels not implemented")
}
func (UnimplementedQueryServer) ConnectionChannels(context.Context, *QueryConnectionChannelsRequest) (*QueryConnectionChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectionChannels not implemented")
}
func (UnimplementedQueryServer) ChannelClientState(context.Context, *QueryChannelClientStateRequest) (*QueryChannelClientStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelClientState not implemented")
}
func (UnimplementedQueryServer) ChannelConsensusState(context.Context, *QueryChannelConsensusStateRequest) (*QueryChannelConsensusStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChannelConsensusState not implemented")
}
func (UnimplementedQueryServer) PacketCommitment(context.Context, *QueryPacketCommitmentRequest) (*QueryPacketCommitmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PacketCommitment not implemented")
}
func (UnimplementedQueryServer) PacketCommitments(context.Context, *QueryPacketCommitmentsRequest) (*QueryPacketCommitmentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PacketCommitments not implemented")
}
func (UnimplementedQueryServer) PacketReceipt(context.Context, *QueryPacketReceiptRequest) (*QueryPacketReceiptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PacketReceipt not implemented")
}
func (UnimplementedQueryServer) PacketAcknowledgement(context.Context, *QueryPacketAcknowledgementRequest) (*QueryPacketAcknowledgementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PacketAcknowledgement not implemented")
}
func (UnimplementedQueryServer) PacketAcknowledgements(context.Context, *QueryPacketAcknowledgementsRequest) (*QueryPacketAcknowledgementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PacketAcknowledgements not implemented")
}
func (UnimplementedQueryServer) UnreceivedPackets(context.Context, *QueryUnreceivedPacketsRequest) (*QueryUnreceivedPacketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnreceivedPackets not implemented")
}
func (UnimplementedQueryServer) UnreceivedAcks(context.Context, *QueryUnreceivedAcksRequest) (*QueryUnreceivedAcksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnreceivedAcks not implemented")
}
func (UnimplementedQueryServer) NextSequenceReceive(context.Context, *QueryNextSequenceReceiveRequest) (*QueryNextSequenceReceiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextSequenceReceive not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Channel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChannelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Channel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/Channel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Channel(ctx, req.(*QueryChannelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Channels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Channels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/Channels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Channels(ctx, req.(*QueryChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ConnectionChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryConnectionChannelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ConnectionChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/ConnectionChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ConnectionChannels(ctx, req.(*QueryConnectionChannelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ChannelClientState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChannelClientStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChannelClientState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/ChannelClientState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChannelClientState(ctx, req.(*QueryChannelClientStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ChannelConsensusState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChannelConsensusStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChannelConsensusState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/ChannelConsensusState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChannelConsensusState(ctx, req.(*QueryChannelConsensusStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PacketCommitment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPacketCommitmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PacketCommitment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/PacketCommitment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PacketCommitment(ctx, req.(*QueryPacketCommitmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PacketCommitments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPacketCommitmentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PacketCommitments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/PacketCommitments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PacketCommitments(ctx, req.(*QueryPacketCommitmentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PacketReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPacketReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PacketReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/PacketReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PacketReceipt(ctx, req.(*QueryPacketReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PacketAcknowledgement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPacketAcknowledgementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PacketAcknowledgement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/PacketAcknowledgement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PacketAcknowledgement(ctx, req.(*QueryPacketAcknowledgementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PacketAcknowledgements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPacketAcknowledgementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PacketAcknowledgements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/PacketAcknowledgements",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PacketAcknowledgements(ctx, req.(*QueryPacketAcknowledgementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UnreceivedPackets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUnreceivedPacketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UnreceivedPackets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/UnreceivedPackets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UnreceivedPackets(ctx, req.(*QueryUnreceivedPacketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_UnreceivedAcks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUnreceivedAcksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).UnreceivedAcks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/UnreceivedAcks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).UnreceivedAcks(ctx, req.(*QueryUnreceivedAcksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_NextSequenceReceive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNextSequenceReceiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NextSequenceReceive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.channel.v1.Query/NextSequenceReceive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NextSequenceReceive(ctx, req.(*QueryNextSequenceReceiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.core.channel.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Channel",
			Handler:    _Query_Channel_Handler,
		},
		{
			MethodName: "Channels",
			Handler:    _Query_Channels_Handler,
		},
		{
			MethodName: "ConnectionChannels",
			Handler:    _Query_ConnectionChannels_Handler,
		},
		{
			MethodName: "ChannelClientState",
			Handler:    _Query_ChannelClientState_Handler,
		},
		{
			MethodName: "ChannelConsensusState",
			Handler:    _Query_ChannelConsensusState_Handler,
		},
		{
			MethodName: "PacketCommitment",
			Handler:    _Query_PacketCommitment_Handler,
		},
		{
			MethodName: "PacketCommitments",
			Handler:    _Query_PacketCommitments_Handler,
		},
		{
			MethodName: "PacketReceipt",
			Handler:    _Query_PacketReceipt_Handler,
		},
		{
			MethodName: "PacketAcknowledgement",
			Handler:    _Query_PacketAcknowledgement_Handler,
		},
		{
			MethodName: "PacketAcknowledgements",
			Handler:    _Query_PacketAcknowledgements_Handler,
		},
		{
			MethodName: "UnreceivedPackets",
			Handler:    _Query_UnreceivedPackets_Handler,
		},
		{
			MethodName: "UnreceivedAcks",
			Handler:    _Query_UnreceivedAcks_Handler,
		},
		{
			MethodName: "NextSequenceReceive",
			Handler:    _Query_NextSequenceReceive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/core/channel/v1/query.proto",
}
