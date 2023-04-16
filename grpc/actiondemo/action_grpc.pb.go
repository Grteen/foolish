// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: action.proto

package actiondemo

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

// ActionServiceClient is the client API for ActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionServiceClient interface {
	CreateAction(ctx context.Context, in *CreateActionRequest, opts ...grpc.CallOption) (*CreateActionResponse, error)
	QueryAction(ctx context.Context, in *QueryActionRequest, opts ...grpc.CallOption) (*QueryActionResponse, error)
	DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...grpc.CallOption) (*DeleteActionResponse, error)
	QueryActionByAuthor(ctx context.Context, in *QueryActionByAuthorRequest, opts ...grpc.CallOption) (*QueryActionByAuthorResponse, error)
	CreateActionLike(ctx context.Context, in *CreateActionLikeRequest, opts ...grpc.CallOption) (*CreateActionLikeResponse, error)
	DeleteActionLike(ctx context.Context, in *DeleteActionLikeRequest, opts ...grpc.CallOption) (*DeleteActionLikeResponse, error)
	QueryActionLike(ctx context.Context, in *QueryActionLikeRequest, opts ...grpc.CallOption) (*QueryActionLikeResponse, error)
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
	QueryComment(ctx context.Context, in *QueryCommentRequest, opts ...grpc.CallOption) (*QueryCommentResponse, error)
	QueryCommentByActionID(ctx context.Context, in *QueryCommentByActionIDRequest, opts ...grpc.CallOption) (*QueryCommentByActionIDResponse, error)
	DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error)
}

type actionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActionServiceClient(cc grpc.ClientConnInterface) ActionServiceClient {
	return &actionServiceClient{cc}
}

func (c *actionServiceClient) CreateAction(ctx context.Context, in *CreateActionRequest, opts ...grpc.CallOption) (*CreateActionResponse, error) {
	out := new(CreateActionResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/CreateAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) QueryAction(ctx context.Context, in *QueryActionRequest, opts ...grpc.CallOption) (*QueryActionResponse, error) {
	out := new(QueryActionResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/QueryAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) DeleteAction(ctx context.Context, in *DeleteActionRequest, opts ...grpc.CallOption) (*DeleteActionResponse, error) {
	out := new(DeleteActionResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/DeleteAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) QueryActionByAuthor(ctx context.Context, in *QueryActionByAuthorRequest, opts ...grpc.CallOption) (*QueryActionByAuthorResponse, error) {
	out := new(QueryActionByAuthorResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/QueryActionByAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) CreateActionLike(ctx context.Context, in *CreateActionLikeRequest, opts ...grpc.CallOption) (*CreateActionLikeResponse, error) {
	out := new(CreateActionLikeResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/CreateActionLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) DeleteActionLike(ctx context.Context, in *DeleteActionLikeRequest, opts ...grpc.CallOption) (*DeleteActionLikeResponse, error) {
	out := new(DeleteActionLikeResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/DeleteActionLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) QueryActionLike(ctx context.Context, in *QueryActionLikeRequest, opts ...grpc.CallOption) (*QueryActionLikeResponse, error) {
	out := new(QueryActionLikeResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/QueryActionLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/CreateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	out := new(UpdateCommentResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/UpdateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) QueryComment(ctx context.Context, in *QueryCommentRequest, opts ...grpc.CallOption) (*QueryCommentResponse, error) {
	out := new(QueryCommentResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/QueryComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) QueryCommentByActionID(ctx context.Context, in *QueryCommentByActionIDRequest, opts ...grpc.CallOption) (*QueryCommentByActionIDResponse, error) {
	out := new(QueryCommentByActionIDResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/QueryCommentByActionID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionServiceClient) DeleteComment(ctx context.Context, in *DeleteCommentRequest, opts ...grpc.CallOption) (*DeleteCommentResponse, error) {
	out := new(DeleteCommentResponse)
	err := c.cc.Invoke(ctx, "/actiondemo.ActionService/DeleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionServiceServer is the server API for ActionService service.
// All implementations must embed UnimplementedActionServiceServer
// for forward compatibility
type ActionServiceServer interface {
	CreateAction(context.Context, *CreateActionRequest) (*CreateActionResponse, error)
	QueryAction(context.Context, *QueryActionRequest) (*QueryActionResponse, error)
	DeleteAction(context.Context, *DeleteActionRequest) (*DeleteActionResponse, error)
	QueryActionByAuthor(context.Context, *QueryActionByAuthorRequest) (*QueryActionByAuthorResponse, error)
	CreateActionLike(context.Context, *CreateActionLikeRequest) (*CreateActionLikeResponse, error)
	DeleteActionLike(context.Context, *DeleteActionLikeRequest) (*DeleteActionLikeResponse, error)
	QueryActionLike(context.Context, *QueryActionLikeRequest) (*QueryActionLikeResponse, error)
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error)
	QueryComment(context.Context, *QueryCommentRequest) (*QueryCommentResponse, error)
	QueryCommentByActionID(context.Context, *QueryCommentByActionIDRequest) (*QueryCommentByActionIDResponse, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error)
	mustEmbedUnimplementedActionServiceServer()
}

// UnimplementedActionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActionServiceServer struct {
}

func (UnimplementedActionServiceServer) CreateAction(context.Context, *CreateActionRequest) (*CreateActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAction not implemented")
}
func (UnimplementedActionServiceServer) QueryAction(context.Context, *QueryActionRequest) (*QueryActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAction not implemented")
}
func (UnimplementedActionServiceServer) DeleteAction(context.Context, *DeleteActionRequest) (*DeleteActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAction not implemented")
}
func (UnimplementedActionServiceServer) QueryActionByAuthor(context.Context, *QueryActionByAuthorRequest) (*QueryActionByAuthorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryActionByAuthor not implemented")
}
func (UnimplementedActionServiceServer) CreateActionLike(context.Context, *CreateActionLikeRequest) (*CreateActionLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActionLike not implemented")
}
func (UnimplementedActionServiceServer) DeleteActionLike(context.Context, *DeleteActionLikeRequest) (*DeleteActionLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteActionLike not implemented")
}
func (UnimplementedActionServiceServer) QueryActionLike(context.Context, *QueryActionLikeRequest) (*QueryActionLikeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryActionLike not implemented")
}
func (UnimplementedActionServiceServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedActionServiceServer) UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedActionServiceServer) QueryComment(context.Context, *QueryCommentRequest) (*QueryCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryComment not implemented")
}
func (UnimplementedActionServiceServer) QueryCommentByActionID(context.Context, *QueryCommentByActionIDRequest) (*QueryCommentByActionIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryCommentByActionID not implemented")
}
func (UnimplementedActionServiceServer) DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedActionServiceServer) mustEmbedUnimplementedActionServiceServer() {}

// UnsafeActionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionServiceServer will
// result in compilation errors.
type UnsafeActionServiceServer interface {
	mustEmbedUnimplementedActionServiceServer()
}

func RegisterActionServiceServer(s grpc.ServiceRegistrar, srv ActionServiceServer) {
	s.RegisterService(&ActionService_ServiceDesc, srv)
}

func _ActionService_CreateAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).CreateAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/CreateAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).CreateAction(ctx, req.(*CreateActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_QueryAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).QueryAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/QueryAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).QueryAction(ctx, req.(*QueryActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_DeleteAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).DeleteAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/DeleteAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).DeleteAction(ctx, req.(*DeleteActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_QueryActionByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActionByAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).QueryActionByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/QueryActionByAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).QueryActionByAuthor(ctx, req.(*QueryActionByAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_CreateActionLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActionLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).CreateActionLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/CreateActionLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).CreateActionLike(ctx, req.(*CreateActionLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_DeleteActionLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteActionLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).DeleteActionLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/DeleteActionLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).DeleteActionLike(ctx, req.(*DeleteActionLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_QueryActionLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActionLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).QueryActionLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/QueryActionLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).QueryActionLike(ctx, req.(*QueryActionLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/CreateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/UpdateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_QueryComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).QueryComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/QueryComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).QueryComment(ctx, req.(*QueryCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_QueryCommentByActionID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCommentByActionIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).QueryCommentByActionID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/QueryCommentByActionID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).QueryCommentByActionID(ctx, req.(*QueryCommentByActionIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionService_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionServiceServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actiondemo.ActionService/DeleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionServiceServer).DeleteComment(ctx, req.(*DeleteCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActionService_ServiceDesc is the grpc.ServiceDesc for ActionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "actiondemo.ActionService",
	HandlerType: (*ActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAction",
			Handler:    _ActionService_CreateAction_Handler,
		},
		{
			MethodName: "QueryAction",
			Handler:    _ActionService_QueryAction_Handler,
		},
		{
			MethodName: "DeleteAction",
			Handler:    _ActionService_DeleteAction_Handler,
		},
		{
			MethodName: "QueryActionByAuthor",
			Handler:    _ActionService_QueryActionByAuthor_Handler,
		},
		{
			MethodName: "CreateActionLike",
			Handler:    _ActionService_CreateActionLike_Handler,
		},
		{
			MethodName: "DeleteActionLike",
			Handler:    _ActionService_DeleteActionLike_Handler,
		},
		{
			MethodName: "QueryActionLike",
			Handler:    _ActionService_QueryActionLike_Handler,
		},
		{
			MethodName: "CreateComment",
			Handler:    _ActionService_CreateComment_Handler,
		},
		{
			MethodName: "UpdateComment",
			Handler:    _ActionService_UpdateComment_Handler,
		},
		{
			MethodName: "QueryComment",
			Handler:    _ActionService_QueryComment_Handler,
		},
		{
			MethodName: "QueryCommentByActionID",
			Handler:    _ActionService_QueryCommentByActionID_Handler,
		},
		{
			MethodName: "DeleteComment",
			Handler:    _ActionService_DeleteComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "action.proto",
}