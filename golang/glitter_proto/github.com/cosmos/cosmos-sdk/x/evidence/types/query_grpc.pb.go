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
	// Evidence queries evidence based on evidence hash.
	Evidence(ctx context.Context, in *QueryEvidenceRequest, opts ...grpc.CallOption) (*QueryEvidenceResponse, error)
	// AllEvidence queries all evidence.
	AllEvidence(ctx context.Context, in *QueryAllEvidenceRequest, opts ...grpc.CallOption) (*QueryAllEvidenceResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Evidence(ctx context.Context, in *QueryEvidenceRequest, opts ...grpc.CallOption) (*QueryEvidenceResponse, error) {
	out := new(QueryEvidenceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.evidence.v1beta1.Query/Evidence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllEvidence(ctx context.Context, in *QueryAllEvidenceRequest, opts ...grpc.CallOption) (*QueryAllEvidenceResponse, error) {
	out := new(QueryAllEvidenceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.evidence.v1beta1.Query/AllEvidence", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Evidence queries evidence based on evidence hash.
	Evidence(context.Context, *QueryEvidenceRequest) (*QueryEvidenceResponse, error)
	// AllEvidence queries all evidence.
	AllEvidence(context.Context, *QueryAllEvidenceRequest) (*QueryAllEvidenceResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Evidence(context.Context, *QueryEvidenceRequest) (*QueryEvidenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Evidence not implemented")
}
func (UnimplementedQueryServer) AllEvidence(context.Context, *QueryAllEvidenceRequest) (*QueryAllEvidenceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllEvidence not implemented")
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

func _Query_Evidence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEvidenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Evidence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.evidence.v1beta1.Query/Evidence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Evidence(ctx, req.(*QueryEvidenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllEvidence_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllEvidenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllEvidence(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.evidence.v1beta1.Query/AllEvidence",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllEvidence(ctx, req.(*QueryAllEvidenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.evidence.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Evidence",
			Handler:    _Query_Evidence_Handler,
		},
		{
			MethodName: "AllEvidence",
			Handler:    _Query_AllEvidence_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/evidence/v1beta1/query.proto",
}
