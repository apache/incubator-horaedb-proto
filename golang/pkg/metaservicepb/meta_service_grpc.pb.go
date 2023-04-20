// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: meta_service.proto

package metaservicepb

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

// CeresmetaRpcServiceClient is the client API for CeresmetaRpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CeresmetaRpcServiceClient interface {
	AllocSchemaID(ctx context.Context, in *AllocSchemaIdRequest, opts ...grpc.CallOption) (*AllocSchemaIdResponse, error)
	GetTablesOfShards(ctx context.Context, in *GetTablesOfShardsRequest, opts ...grpc.CallOption) (*GetTablesOfShardsResponse, error)
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*CreateTableResponse, error)
	DropTable(ctx context.Context, in *DropTableRequest, opts ...grpc.CallOption) (*DropTableResponse, error)
	RouteTables(ctx context.Context, in *RouteTablesRequest, opts ...grpc.CallOption) (*RouteTablesResponse, error)
	GetNodes(ctx context.Context, in *GetNodesRequest, opts ...grpc.CallOption) (*GetNodesResponse, error)
	NodeHeartbeat(ctx context.Context, in *NodeHeartbeatRequest, opts ...grpc.CallOption) (*NodeHeartbeatResponse, error)
}

type ceresmetaRpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCeresmetaRpcServiceClient(cc grpc.ClientConnInterface) CeresmetaRpcServiceClient {
	return &ceresmetaRpcServiceClient{cc}
}

func (c *ceresmetaRpcServiceClient) AllocSchemaID(ctx context.Context, in *AllocSchemaIdRequest, opts ...grpc.CallOption) (*AllocSchemaIdResponse, error) {
	out := new(AllocSchemaIdResponse)
	err := c.cc.Invoke(ctx, "/meta_service.CeresmetaRpcService/AllocSchemaID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ceresmetaRpcServiceClient) GetTablesOfShards(ctx context.Context, in *GetTablesOfShardsRequest, opts ...grpc.CallOption) (*GetTablesOfShardsResponse, error) {
	out := new(GetTablesOfShardsResponse)
	err := c.cc.Invoke(ctx, "/meta_service.CeresmetaRpcService/GetTablesOfShards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ceresmetaRpcServiceClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*CreateTableResponse, error) {
	out := new(CreateTableResponse)
	err := c.cc.Invoke(ctx, "/meta_service.CeresmetaRpcService/CreateTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ceresmetaRpcServiceClient) DropTable(ctx context.Context, in *DropTableRequest, opts ...grpc.CallOption) (*DropTableResponse, error) {
	out := new(DropTableResponse)
	err := c.cc.Invoke(ctx, "/meta_service.CeresmetaRpcService/DropTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ceresmetaRpcServiceClient) RouteTables(ctx context.Context, in *RouteTablesRequest, opts ...grpc.CallOption) (*RouteTablesResponse, error) {
	out := new(RouteTablesResponse)
	err := c.cc.Invoke(ctx, "/meta_service.CeresmetaRpcService/RouteTables", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ceresmetaRpcServiceClient) GetNodes(ctx context.Context, in *GetNodesRequest, opts ...grpc.CallOption) (*GetNodesResponse, error) {
	out := new(GetNodesResponse)
	err := c.cc.Invoke(ctx, "/meta_service.CeresmetaRpcService/GetNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ceresmetaRpcServiceClient) NodeHeartbeat(ctx context.Context, in *NodeHeartbeatRequest, opts ...grpc.CallOption) (*NodeHeartbeatResponse, error) {
	out := new(NodeHeartbeatResponse)
	err := c.cc.Invoke(ctx, "/meta_service.CeresmetaRpcService/NodeHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CeresmetaRpcServiceServer is the server API for CeresmetaRpcService service.
// All implementations must embed UnimplementedCeresmetaRpcServiceServer
// for forward compatibility
type CeresmetaRpcServiceServer interface {
	AllocSchemaID(context.Context, *AllocSchemaIdRequest) (*AllocSchemaIdResponse, error)
	GetTablesOfShards(context.Context, *GetTablesOfShardsRequest) (*GetTablesOfShardsResponse, error)
	CreateTable(context.Context, *CreateTableRequest) (*CreateTableResponse, error)
	DropTable(context.Context, *DropTableRequest) (*DropTableResponse, error)
	RouteTables(context.Context, *RouteTablesRequest) (*RouteTablesResponse, error)
	GetNodes(context.Context, *GetNodesRequest) (*GetNodesResponse, error)
	NodeHeartbeat(context.Context, *NodeHeartbeatRequest) (*NodeHeartbeatResponse, error)
	mustEmbedUnimplementedCeresmetaRpcServiceServer()
}

// UnimplementedCeresmetaRpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCeresmetaRpcServiceServer struct {
}

func (UnimplementedCeresmetaRpcServiceServer) AllocSchemaID(context.Context, *AllocSchemaIdRequest) (*AllocSchemaIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocSchemaID not implemented")
}
func (UnimplementedCeresmetaRpcServiceServer) GetTablesOfShards(context.Context, *GetTablesOfShardsRequest) (*GetTablesOfShardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTablesOfShards not implemented")
}
func (UnimplementedCeresmetaRpcServiceServer) CreateTable(context.Context, *CreateTableRequest) (*CreateTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (UnimplementedCeresmetaRpcServiceServer) DropTable(context.Context, *DropTableRequest) (*DropTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropTable not implemented")
}
func (UnimplementedCeresmetaRpcServiceServer) RouteTables(context.Context, *RouteTablesRequest) (*RouteTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RouteTables not implemented")
}
func (UnimplementedCeresmetaRpcServiceServer) GetNodes(context.Context, *GetNodesRequest) (*GetNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodes not implemented")
}
func (UnimplementedCeresmetaRpcServiceServer) NodeHeartbeat(context.Context, *NodeHeartbeatRequest) (*NodeHeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NodeHeartbeat not implemented")
}
func (UnimplementedCeresmetaRpcServiceServer) mustEmbedUnimplementedCeresmetaRpcServiceServer() {}

// UnsafeCeresmetaRpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CeresmetaRpcServiceServer will
// result in compilation errors.
type UnsafeCeresmetaRpcServiceServer interface {
	mustEmbedUnimplementedCeresmetaRpcServiceServer()
}

func RegisterCeresmetaRpcServiceServer(s grpc.ServiceRegistrar, srv CeresmetaRpcServiceServer) {
	s.RegisterService(&CeresmetaRpcService_ServiceDesc, srv)
}

func _CeresmetaRpcService_AllocSchemaID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocSchemaIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CeresmetaRpcServiceServer).AllocSchemaID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.CeresmetaRpcService/AllocSchemaID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CeresmetaRpcServiceServer).AllocSchemaID(ctx, req.(*AllocSchemaIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CeresmetaRpcService_GetTablesOfShards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTablesOfShardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CeresmetaRpcServiceServer).GetTablesOfShards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.CeresmetaRpcService/GetTablesOfShards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CeresmetaRpcServiceServer).GetTablesOfShards(ctx, req.(*GetTablesOfShardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CeresmetaRpcService_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CeresmetaRpcServiceServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.CeresmetaRpcService/CreateTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CeresmetaRpcServiceServer).CreateTable(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CeresmetaRpcService_DropTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CeresmetaRpcServiceServer).DropTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.CeresmetaRpcService/DropTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CeresmetaRpcServiceServer).DropTable(ctx, req.(*DropTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CeresmetaRpcService_RouteTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RouteTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CeresmetaRpcServiceServer).RouteTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.CeresmetaRpcService/RouteTables",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CeresmetaRpcServiceServer).RouteTables(ctx, req.(*RouteTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CeresmetaRpcService_GetNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CeresmetaRpcServiceServer).GetNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.CeresmetaRpcService/GetNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CeresmetaRpcServiceServer).GetNodes(ctx, req.(*GetNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CeresmetaRpcService_NodeHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeHeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CeresmetaRpcServiceServer).NodeHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/meta_service.CeresmetaRpcService/NodeHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CeresmetaRpcServiceServer).NodeHeartbeat(ctx, req.(*NodeHeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CeresmetaRpcService_ServiceDesc is the grpc.ServiceDesc for CeresmetaRpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CeresmetaRpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meta_service.CeresmetaRpcService",
	HandlerType: (*CeresmetaRpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllocSchemaID",
			Handler:    _CeresmetaRpcService_AllocSchemaID_Handler,
		},
		{
			MethodName: "GetTablesOfShards",
			Handler:    _CeresmetaRpcService_GetTablesOfShards_Handler,
		},
		{
			MethodName: "CreateTable",
			Handler:    _CeresmetaRpcService_CreateTable_Handler,
		},
		{
			MethodName: "DropTable",
			Handler:    _CeresmetaRpcService_DropTable_Handler,
		},
		{
			MethodName: "RouteTables",
			Handler:    _CeresmetaRpcService_RouteTables_Handler,
		},
		{
			MethodName: "GetNodes",
			Handler:    _CeresmetaRpcService_GetNodes_Handler,
		},
		{
			MethodName: "NodeHeartbeat",
			Handler:    _CeresmetaRpcService_NodeHeartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "meta_service.proto",
}
