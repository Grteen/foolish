package rpc

import (
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client userdemo.UserServiceClient

func InitUserRpc() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client = userdemo.NewUserServiceClient(conn)
}

func CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) error {
	resp, err := client.CreateUser(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) error {
	resp, err := client.CheckUser(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}
