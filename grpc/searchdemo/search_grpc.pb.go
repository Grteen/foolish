// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: search.proto

package searchdemo

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
	SearchService_SearchArtical_FullMethodName  = "/searchdemo.SearchService/SearchArtical"
	SearchService_SearchUserZoom_FullMethodName = "/searchdemo.SearchService/SearchUserZoom"
)

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	SearchArtical(ctx context.Context, in *SearchArticalRequest, opts ...grpc.CallOption) (*SearchArticalResponse, error)
	SearchUserZoom(ctx context.Context, in *SearchUserZoomRequest, opts ...grpc.CallOption) (*SearchUserZoomResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) SearchArtical(ctx context.Context, in *SearchArticalRequest, opts ...grpc.CallOption) (*SearchArticalResponse, error) {
	out := new(SearchArticalResponse)
	err := c.cc.Invoke(ctx, SearchService_SearchArtical_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) SearchUserZoom(ctx context.Context, in *SearchUserZoomRequest, opts ...grpc.CallOption) (*SearchUserZoomResponse, error) {
	out := new(SearchUserZoomResponse)
	err := c.cc.Invoke(ctx, SearchService_SearchUserZoom_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations must embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	SearchArtical(context.Context, *SearchArticalRequest) (*SearchArticalResponse, error)
	SearchUserZoom(context.Context, *SearchUserZoomRequest) (*SearchUserZoomResponse, error)
	mustEmbedUnimplementedSearchServiceServer()
}

// UnimplementedSearchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) SearchArtical(context.Context, *SearchArticalRequest) (*SearchArticalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchArtical not implemented")
}
func (UnimplementedSearchServiceServer) SearchUserZoom(context.Context, *SearchUserZoomRequest) (*SearchUserZoomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserZoom not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_SearchArtical_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchArticalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchArtical(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_SearchArtical_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchArtical(ctx, req.(*SearchArticalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_SearchUserZoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserZoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).SearchUserZoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_SearchUserZoom_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).SearchUserZoom(ctx, req.(*SearchUserZoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "searchdemo.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchArtical",
			Handler:    _SearchService_SearchArtical_Handler,
		},
		{
			MethodName: "SearchUserZoom",
			Handler:    _SearchService_SearchUserZoom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}
