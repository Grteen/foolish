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

func QueryAvator(ctx context.Context, req *userdemo.QueryAvatorRequest) ([]string, error) {
	resp, err := client.QueryAvator(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Avator, nil
}

func RdbIncreaseItfUser(ctx context.Context, req *userdemo.RdbIncreaseItfRequest) error {
	resp, err := client.RdbIncreaseItf(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}
