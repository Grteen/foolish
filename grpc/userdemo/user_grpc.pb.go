// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: user.proto

package userdemo

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	CheckUser(ctx context.Context, in *CheckUserRequest, opts ...grpc.CallOption) (*CheckUserResponse, error)
	QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error)
	QueryUserInfo(ctx context.Context, in *QueryUserInfoRequest, opts ...grpc.CallOption) (*QueryUserInfoResponse, error)
	QueryAvator(ctx context.Context, in *QueryAvatorRequest, opts ...grpc.CallOption) (*QueryAvatorResponse, error)
	SetAuthCookie(ctx context.Context, in *SetAuthCookieRequest, opts ...grpc.CallOption) (*SetAuthCookieResponse, error)
	QueryAuthCookie(ctx context.Context, in *QueryAuthCookieRequest, opts ...grpc.CallOption) (*QueryAuthCookieResponse, error)
	DeleteAuthCookie(ctx context.Context, in *DeleteAuthCookieRequest, opts ...grpc.CallOption) (*DeleteAuthCookieResponse, error)
	CreateSubscribe(ctx context.Context, in *CreateSubscribeRequest, opts ...grpc.CallOption) (*CreateSubscribeResponse, error)
	DeleteSubscribe(ctx context.Context, in *DeleteSubscribeRequest, opts ...grpc.CallOption) (*DeleteSubscribeResponse, error)
	QuerySubscribe(ctx context.Context, in *QuerySubscribeRequest, opts ...grpc.CallOption) (*QuerySubscribeResponse, error)
	QueryAllSubscribe(ctx context.Context, in *QueryAllSubscribeRequest, opts ...grpc.CallOption) (*QueryAllSubscribeResponse, error)
	QueryAllFans(ctx context.Context, in *QueryAllFansRequest, opts ...grpc.CallOption) (*QueryAllFansResponse, error)
	RdbSetUser(ctx context.Context, in *RdbSetUserRequest, opts ...grpc.CallOption) (*RdbSetUserResponse, error)
	RdbGetUser(ctx context.Context, in *RdbGetUserRequest, opts ...grpc.CallOption) (*RdbGetUserResponse, error)
	RdbIncreaseItf(ctx context.Context, in *RdbIncreaseItfRequest, opts ...grpc.CallOption) (*RdbIncreaseItfResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckUser(ctx context.Context, in *CheckUserRequest, opts ...grpc.CallOption) (*CheckUserResponse, error) {
	out := new(CheckUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CheckUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error) {
	out := new(QueryUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/QueryUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*UpdateUserInfoResponse, error) {
	out := new(UpdateUserInfoResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QueryUserInfo(ctx context.Context, in *QueryUserInfoRequest, opts ...grpc.CallOption) (*QueryUserInfoResponse, error) {
	out := new(QueryUserInfoResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/QueryUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QueryAvator(ctx context.Context, in *QueryAvatorRequest, opts ...grpc.CallOption) (*QueryAvatorResponse, error) {
	out := new(QueryAvatorResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/QueryAvator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetAuthCookie(ctx context.Context, in *SetAuthCookieRequest, opts ...grpc.CallOption) (*SetAuthCookieResponse, error) {
	out := new(SetAuthCookieResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/SetAuthCookie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QueryAuthCookie(ctx context.Context, in *QueryAuthCookieRequest, opts ...grpc.CallOption) (*QueryAuthCookieResponse, error) {
	out := new(QueryAuthCookieResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/QueryAuthCookie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteAuthCookie(ctx context.Context, in *DeleteAuthCookieRequest, opts ...grpc.CallOption) (*DeleteAuthCookieResponse, error) {
	out := new(DeleteAuthCookieResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteAuthCookie", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CreateSubscribe(ctx context.Context, in *CreateSubscribeRequest, opts ...grpc.CallOption) (*CreateSubscribeResponse, error) {
	out := new(CreateSubscribeResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteSubscribe(ctx context.Context, in *DeleteSubscribeRequest, opts ...grpc.CallOption) (*DeleteSubscribeResponse, error) {
	out := new(DeleteSubscribeResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QuerySubscribe(ctx context.Context, in *QuerySubscribeRequest, opts ...grpc.CallOption) (*QuerySubscribeResponse, error) {
	out := new(QuerySubscribeResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/QuerySubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QueryAllSubscribe(ctx context.Context, in *QueryAllSubscribeRequest, opts ...grpc.CallOption) (*QueryAllSubscribeResponse, error) {
	out := new(QueryAllSubscribeResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/QueryAllSubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) QueryAllFans(ctx context.Context, in *QueryAllFansRequest, opts ...grpc.CallOption) (*QueryAllFansResponse, error) {
	out := new(QueryAllFansResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/QueryAllFans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RdbSetUser(ctx context.Context, in *RdbSetUserRequest, opts ...grpc.CallOption) (*RdbSetUserResponse, error) {
	out := new(RdbSetUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/RdbSetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RdbGetUser(ctx context.Context, in *RdbGetUserRequest, opts ...grpc.CallOption) (*RdbGetUserResponse, error) {
	out := new(RdbGetUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/RdbGetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RdbIncreaseItf(ctx context.Context, in *RdbIncreaseItfRequest, opts ...grpc.CallOption) (*RdbIncreaseItfResponse, error) {
	out := new(RdbIncreaseItfResponse)
	err := c.cc.Invoke(ctx, "/user.UserService/RdbIncreaseItf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	CheckUser(context.Context, *CheckUserRequest) (*CheckUserResponse, error)
	QueryUser(context.Context, *QueryUserRequest) (*QueryUserResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error)
	QueryUserInfo(context.Context, *QueryUserInfoRequest) (*QueryUserInfoResponse, error)
	QueryAvator(context.Context, *QueryAvatorRequest) (*QueryAvatorResponse, error)
	SetAuthCookie(context.Context, *SetAuthCookieRequest) (*SetAuthCookieResponse, error)
	QueryAuthCookie(context.Context, *QueryAuthCookieRequest) (*QueryAuthCookieResponse, error)
	DeleteAuthCookie(context.Context, *DeleteAuthCookieRequest) (*DeleteAuthCookieResponse, error)
	CreateSubscribe(context.Context, *CreateSubscribeRequest) (*CreateSubscribeResponse, error)
	DeleteSubscribe(context.Context, *DeleteSubscribeRequest) (*DeleteSubscribeResponse, error)
	QuerySubscribe(context.Context, *QuerySubscribeRequest) (*QuerySubscribeResponse, error)
	QueryAllSubscribe(context.Context, *QueryAllSubscribeRequest) (*QueryAllSubscribeResponse, error)
	QueryAllFans(context.Context, *QueryAllFansRequest) (*QueryAllFansResponse, error)
	RdbSetUser(context.Context, *RdbSetUserRequest) (*RdbSetUserResponse, error)
	RdbGetUser(context.Context, *RdbGetUserRequest) (*RdbGetUserResponse, error)
	RdbIncreaseItf(context.Context, *RdbIncreaseItfRequest) (*RdbIncreaseItfResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) CheckUser(context.Context, *CheckUserRequest) (*CheckUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUser not implemented")
}
func (UnimplementedUserServiceServer) QueryUser(context.Context, *QueryUserRequest) (*QueryUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*UpdateUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedUserServiceServer) QueryUserInfo(context.Context, *QueryUserInfoRequest) (*QueryUserInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUserInfo not implemented")
}
func (UnimplementedUserServiceServer) QueryAvator(context.Context, *QueryAvatorRequest) (*QueryAvatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAvator not implemented")
}
func (UnimplementedUserServiceServer) SetAuthCookie(context.Context, *SetAuthCookieRequest) (*SetAuthCookieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAuthCookie not implemented")
}
func (UnimplementedUserServiceServer) QueryAuthCookie(context.Context, *QueryAuthCookieRequest) (*QueryAuthCookieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAuthCookie not implemented")
}
func (UnimplementedUserServiceServer) DeleteAuthCookie(context.Context, *DeleteAuthCookieRequest) (*DeleteAuthCookieResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthCookie not implemented")
}
func (UnimplementedUserServiceServer) CreateSubscribe(context.Context, *CreateSubscribeRequest) (*CreateSubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubscribe not implemented")
}
func (UnimplementedUserServiceServer) DeleteSubscribe(context.Context, *DeleteSubscribeRequest) (*DeleteSubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSubscribe not implemented")
}
func (UnimplementedUserServiceServer) QuerySubscribe(context.Context, *QuerySubscribeRequest) (*QuerySubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySubscribe not implemented")
}
func (UnimplementedUserServiceServer) QueryAllSubscribe(context.Context, *QueryAllSubscribeRequest) (*QueryAllSubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAllSubscribe not implemented")
}
func (UnimplementedUserServiceServer) QueryAllFans(context.Context, *QueryAllFansRequest) (*QueryAllFansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAllFans not implemented")
}
func (UnimplementedUserServiceServer) RdbSetUser(context.Context, *RdbSetUserRequest) (*RdbSetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RdbSetUser not implemented")
}
func (UnimplementedUserServiceServer) RdbGetUser(context.Context, *RdbGetUserRequest) (*RdbGetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RdbGetUser not implemented")
}
func (UnimplementedUserServiceServer) RdbIncreaseItf(context.Context, *RdbIncreaseItfRequest) (*RdbIncreaseItfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RdbIncreaseItf not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CheckUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckUser(ctx, req.(*CheckUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QueryUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QueryUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/QueryUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QueryUser(ctx, req.(*QueryUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QueryUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QueryUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/QueryUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QueryUserInfo(ctx, req.(*QueryUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QueryAvator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAvatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QueryAvator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/QueryAvator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QueryAvator(ctx, req.(*QueryAvatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetAuthCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAuthCookieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetAuthCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/SetAuthCookie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetAuthCookie(ctx, req.(*SetAuthCookieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QueryAuthCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAuthCookieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QueryAuthCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/QueryAuthCookie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QueryAuthCookie(ctx, req.(*QueryAuthCookieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteAuthCookie_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthCookieRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteAuthCookie(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteAuthCookie",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteAuthCookie(ctx, req.(*DeleteAuthCookieRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CreateSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateSubscribe(ctx, req.(*CreateSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteSubscribe(ctx, req.(*DeleteSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QuerySubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QuerySubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/QuerySubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QuerySubscribe(ctx, req.(*QuerySubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QueryAllSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllSubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QueryAllSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/QueryAllSubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QueryAllSubscribe(ctx, req.(*QueryAllSubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_QueryAllFans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllFansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).QueryAllFans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/QueryAllFans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).QueryAllFans(ctx, req.(*QueryAllFansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RdbSetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbSetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RdbSetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/RdbSetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RdbSetUser(ctx, req.(*RdbSetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RdbGetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbGetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RdbGetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/RdbGetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RdbGetUser(ctx, req.(*RdbGetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RdbIncreaseItf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbIncreaseItfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RdbIncreaseItf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/RdbIncreaseItf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RdbIncreaseItf(ctx, req.(*RdbIncreaseItfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "CheckUser",
			Handler:    _UserService_CheckUser_Handler,
		},
		{
			MethodName: "QueryUser",
			Handler:    _UserService_QueryUser_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _UserService_UpdateUserInfo_Handler,
		},
		{
			MethodName: "QueryUserInfo",
			Handler:    _UserService_QueryUserInfo_Handler,
		},
		{
			MethodName: "QueryAvator",
			Handler:    _UserService_QueryAvator_Handler,
		},
		{
			MethodName: "SetAuthCookie",
			Handler:    _UserService_SetAuthCookie_Handler,
		},
		{
			MethodName: "QueryAuthCookie",
			Handler:    _UserService_QueryAuthCookie_Handler,
		},
		{
			MethodName: "DeleteAuthCookie",
			Handler:    _UserService_DeleteAuthCookie_Handler,
		},
		{
			MethodName: "CreateSubscribe",
			Handler:    _UserService_CreateSubscribe_Handler,
		},
		{
			MethodName: "DeleteSubscribe",
			Handler:    _UserService_DeleteSubscribe_Handler,
		},
		{
			MethodName: "QuerySubscribe",
			Handler:    _UserService_QuerySubscribe_Handler,
		},
		{
			MethodName: "QueryAllSubscribe",
			Handler:    _UserService_QueryAllSubscribe_Handler,
		},
		{
			MethodName: "QueryAllFans",
			Handler:    _UserService_QueryAllFans_Handler,
		},
		{
			MethodName: "RdbSetUser",
			Handler:    _UserService_RdbSetUser_Handler,
		},
		{
			MethodName: "RdbGetUser",
			Handler:    _UserService_RdbGetUser_Handler,
		},
		{
			MethodName: "RdbIncreaseItf",
			Handler:    _UserService_RdbIncreaseItf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
