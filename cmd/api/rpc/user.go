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

func CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (string, error) {
	resp, err := client.CheckUser(ctx, req)
	if err != nil {
		return resp.UserName, err
	}

	if resp.Resp.StatusCode != 0 {
		return resp.UserName, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.UserName, err
}

func UpdateUserInfo(ctx context.Context, req *userdemo.UpdateUserInfoRequest) error {
	resp, err := client.UpdateUserInfo(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func QueryUserInfo(ctx context.Context, req *userdemo.QueryUserInfoRequest) (*userdemo.UserInfo, error) {
	resp, err := client.QueryUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.UserInfo, nil
}

func SetAuthCookie(ctx context.Context, req *userdemo.SetAuthCookieRequest) error {
	resp, err := client.SetAuthCookie(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func QueryAuthCookie(ctx context.Context, req *userdemo.QueryAuthCookieRequest) (string, error) {
	resp, err := client.QueryAuthCookie(ctx, req)
	if err != nil {
		return "", err
	}

	if resp.Resp.StatusCode != 0 {
		return "", errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Value, nil
}

func DeleteAuthCookie(ctx context.Context, req *userdemo.DeleteAuthCookieRequest) error {
	resp, err := client.DeleteAuthCookie(ctx, req)
	if err != nil {
		return errno.ConvertErr(err)
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}
