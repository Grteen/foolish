// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: smtp.proto

package msmtpdemo

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
	SmtpService_SendSmtp_FullMethodName    = "/msmtp.SmtpService/SendSmtp"
	SmtpService_QueryVerify_FullMethodName = "/msmtp.SmtpService/QueryVerify"
)

// SmtpServiceClient is the client API for SmtpService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmtpServiceClient interface {
	SendSmtp(ctx context.Context, in *SendSmtpRequest, opts ...grpc.CallOption) (*SendSmtpResponse, error)
	QueryVerify(ctx context.Context, in *QueryVerifyRequest, opts ...grpc.CallOption) (*QueryVerifyResponse, error)
}

type smtpServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmtpServiceClient(cc grpc.ClientConnInterface) SmtpServiceClient {
	return &smtpServiceClient{cc}
}

func (c *smtpServiceClient) SendSmtp(ctx context.Context, in *SendSmtpRequest, opts ...grpc.CallOption) (*SendSmtpResponse, error) {
	out := new(SendSmtpResponse)
	err := c.cc.Invoke(ctx, SmtpService_SendSmtp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smtpServiceClient) QueryVerify(ctx context.Context, in *QueryVerifyRequest, opts ...grpc.CallOption) (*QueryVerifyResponse, error) {
	out := new(QueryVerifyResponse)
	err := c.cc.Invoke(ctx, SmtpService_QueryVerify_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmtpServiceServer is the server API for SmtpService service.
// All implementations must embed UnimplementedSmtpServiceServer
// for forward compatibility
type SmtpServiceServer interface {
	SendSmtp(context.Context, *SendSmtpRequest) (*SendSmtpResponse, error)
	QueryVerify(context.Context, *QueryVerifyRequest) (*QueryVerifyResponse, error)
	mustEmbedUnimplementedSmtpServiceServer()
}

// UnimplementedSmtpServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSmtpServiceServer struct {
}

func (UnimplementedSmtpServiceServer) SendSmtp(context.Context, *SendSmtpRequest) (*SendSmtpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmtp not implemented")
}
func (UnimplementedSmtpServiceServer) QueryVerify(context.Context, *QueryVerifyRequest) (*QueryVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryVerify not implemented")
}
func (UnimplementedSmtpServiceServer) mustEmbedUnimplementedSmtpServiceServer() {}

// UnsafeSmtpServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmtpServiceServer will
// result in compilation errors.
type UnsafeSmtpServiceServer interface {
	mustEmbedUnimplementedSmtpServiceServer()
}

func RegisterSmtpServiceServer(s grpc.ServiceRegistrar, srv SmtpServiceServer) {
	s.RegisterService(&SmtpService_ServiceDesc, srv)
}

func _SmtpService_SendSmtp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmtpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmtpServiceServer).SendSmtp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmtpService_SendSmtp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmtpServiceServer).SendSmtp(ctx, req.(*SendSmtpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmtpService_QueryVerify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmtpServiceServer).QueryVerify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmtpService_QueryVerify_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmtpServiceServer).QueryVerify(ctx, req.(*QueryVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmtpService_ServiceDesc is the grpc.ServiceDesc for SmtpService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmtpService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msmtp.SmtpService",
	HandlerType: (*SmtpServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSmtp",
			Handler:    _SmtpService_SendSmtp_Handler,
		},
		{
			MethodName: "QueryVerify",
			Handler:    _SmtpService_QueryVerify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "smtp.proto",
}
