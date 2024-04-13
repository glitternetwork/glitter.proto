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
	// Parameters queries the parameters of the module.
	ListSchema(ctx context.Context, in *ListSchemaRequest, opts ...grpc.CallOption) (*ListSchemaResponse, error)
	// Parameters queries the parameters of the module.
	GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error)
	// Parameters queries the parameters of the module.
	GetDoc(ctx context.Context, in *GetDocRequest, opts ...grpc.CallOption) (*GetDocResponse, error)
	// Parameters queries the parameters of the module.
	SearchDoc(ctx context.Context, in *SearchDocRequest, opts ...grpc.CallOption) (*SearchDocResponse, error)
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// SQLQuery do queries
	SQLQuery(ctx context.Context, in *SQLQueryRequest, opts ...grpc.CallOption) (*SQLQueryResponse, error)
	// SimpleSQLQuery do queries
	SimpleSQLQuery(ctx context.Context, in *SQLQueryRequest, opts ...grpc.CallOption) (*SimpleSQLQueryResponse, error)
	// Parameters queries the parameters of the module.
	SQLListTables(ctx context.Context, in *SQLListTablesRequest, opts ...grpc.CallOption) (*SQLListTablesResponse, error)
	SQLListDatabases(ctx context.Context, in *SQLListDatabasesRequest, opts ...grpc.CallOption) (*SQLListDatabasesResponse, error)
	ShowCreateTable(ctx context.Context, in *ShowCreateTableRequest, opts ...grpc.CallOption) (*ShowCreateTableResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ListSchema(ctx context.Context, in *ListSchemaRequest, opts ...grpc.CallOption) (*ListSchemaResponse, error) {
	out := new(ListSchemaResponse)
	err := c.cc.Invoke(ctx, "/blockved.glitterchain.index.Query/ListSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error) {
	out := new(GetSchemaResponse)
	err := c.cc.Invoke(ctx, "/blockved.glitterchain.index.Query/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetDoc(ctx context.Context, in *GetDocRequest, opts ...grpc.CallOption) (*GetDocResponse, error) {
	out := new(GetDocResponse)
	err := c.cc.Invoke(ctx, "/blockved.glitterchain.index.Query/GetDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SearchDoc(ctx context.Context, in *SearchDocRequest, opts ...grpc.CallOption) (*SearchDocResponse, error) {
	out := new(SearchDocResponse)
	err := c.cc.Invoke(ctx, "/blockved.glitterchain.index.Query/SearchDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/blockved.glitterchain.index.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SQLQuery(ctx context.Context, in *SQLQueryRequest, opts ...grpc.CallOption) (*SQLQueryResponse, error) {
	out := new(SQLQueryResponse)
	err := c.cc.Invoke(ctx, "/blockved.glitterchain.index.Query/SQLQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SimpleSQLQuery(ctx context.Context, in *SQLQueryRequest, opts ...grpc.CallOption) (*SimpleSQLQueryResponse, error) {
	out := new(SimpleSQLQueryResponse)
	err := c.cc.Invoke(ctx, "/blockved.glitterchain.index.Query/SimpleSQLQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SQLListTables(ctx context.Context, in *SQLListTablesRequest, opts ...grpc.CallOption) (*SQLListTablesResponse, error) {
	out := new(SQLListTablesResponse)
	err := c.cc.Invoke(ctx, "/blockved.glitterchain.index.Query/SQLListTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SQLListDatabases(ctx context.Context, in *SQLListDatabasesRequest, opts ...grpc.CallOption) (*SQLListDatabasesResponse, error) {
	out := new(SQLListDatabasesResponse)
	err := c.cc.Invoke(ctx, "/blockved.glitterchain.index.Query/SQLListDatabases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ShowCreateTable(ctx context.Context, in *ShowCreateTableRequest, opts ...grpc.CallOption) (*ShowCreateTableResponse, error) {
	out := new(ShowCreateTableResponse)
	err := c.cc.Invoke(ctx, "/blockved.glitterchain.index.Query/ShowCreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	ListSchema(context.Context, *ListSchemaRequest) (*ListSchemaResponse, error)
	// Parameters queries the parameters of the module.
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)
	// Parameters queries the parameters of the module.
	GetDoc(context.Context, *GetDocRequest) (*GetDocResponse, error)
	// Parameters queries the parameters of the module.
	SearchDoc(context.Context, *SearchDocRequest) (*SearchDocResponse, error)
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// SQLQuery do queries
	SQLQuery(context.Context, *SQLQueryRequest) (*SQLQueryResponse, error)
	// SimpleSQLQuery do queries
	SimpleSQLQuery(context.Context, *SQLQueryRequest) (*SimpleSQLQueryResponse, error)
	// Parameters queries the parameters of the module.
	SQLListTables(context.Context, *SQLListTablesRequest) (*SQLListTablesResponse, error)
	SQLListDatabases(context.Context, *SQLListDatabasesRequest) (*SQLListDatabasesResponse, error)
	ShowCreateTable(context.Context, *ShowCreateTableRequest) (*ShowCreateTableResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) ListSchema(context.Context, *ListSchemaRequest) (*ListSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSchema not implemented")
}
func (UnimplementedQueryServer) GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (UnimplementedQueryServer) GetDoc(context.Context, *GetDocRequest) (*GetDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDoc not implemented")
}
func (UnimplementedQueryServer) SearchDoc(context.Context, *SearchDocRequest) (*SearchDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchDoc not implemented")
}
func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) SQLQuery(context.Context, *SQLQueryRequest) (*SQLQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SQLQuery not implemented")
}
func (UnimplementedQueryServer) SimpleSQLQuery(context.Context, *SQLQueryRequest) (*SimpleSQLQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimpleSQLQuery not implemented")
}
func (UnimplementedQueryServer) SQLListTables(context.Context, *SQLListTablesRequest) (*SQLListTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SQLListTables not implemented")
}
func (UnimplementedQueryServer) SQLListDatabases(context.Context, *SQLListDatabasesRequest) (*SQLListDatabasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SQLListDatabases not implemented")
}
func (UnimplementedQueryServer) ShowCreateTable(context.Context, *ShowCreateTableRequest) (*ShowCreateTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShowCreateTable not implemented")
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

func _Query_ListSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ListSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockved.glitterchain.index.Query/ListSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ListSchema(ctx, req.(*ListSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockved.glitterchain.index.Query/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetSchema(ctx, req.(*GetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockved.glitterchain.index.Query/GetDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetDoc(ctx, req.(*GetDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SearchDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchDocRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SearchDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockved.glitterchain.index.Query/SearchDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SearchDoc(ctx, req.(*SearchDocRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockved.glitterchain.index.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SQLQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SQLQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SQLQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockved.glitterchain.index.Query/SQLQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SQLQuery(ctx, req.(*SQLQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SimpleSQLQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SQLQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SimpleSQLQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockved.glitterchain.index.Query/SimpleSQLQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SimpleSQLQuery(ctx, req.(*SQLQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SQLListTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SQLListTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SQLListTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockved.glitterchain.index.Query/SQLListTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SQLListTables(ctx, req.(*SQLListTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SQLListDatabases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SQLListDatabasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SQLListDatabases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockved.glitterchain.index.Query/SQLListDatabases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SQLListDatabases(ctx, req.(*SQLListDatabasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ShowCreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShowCreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ShowCreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockved.glitterchain.index.Query/ShowCreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ShowCreateTable(ctx, req.(*ShowCreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blockved.glitterchain.index.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSchema",
			Handler:    _Query_ListSchema_Handler,
		},
		{
			MethodName: "GetSchema",
			Handler:    _Query_GetSchema_Handler,
		},
		{
			MethodName: "GetDoc",
			Handler:    _Query_GetDoc_Handler,
		},
		{
			MethodName: "SearchDoc",
			Handler:    _Query_SearchDoc_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SQLQuery",
			Handler:    _Query_SQLQuery_Handler,
		},
		{
			MethodName: "SimpleSQLQuery",
			Handler:    _Query_SimpleSQLQuery_Handler,
		},
		{
			MethodName: "SQLListTables",
			Handler:    _Query_SQLListTables_Handler,
		},
		{
			MethodName: "SQLListDatabases",
			Handler:    _Query_SQLListDatabases_Handler,
		},
		{
			MethodName: "ShowCreateTable",
			Handler:    _Query_ShowCreateTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "index/query.proto",
}
