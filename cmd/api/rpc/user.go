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

func QueryUser(ctx context.Context, req *userdemo.QueryUserRequest) ([]*userdemo.User, error) {
	resp, err := client.QueryUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.User, err
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

func CreateSubscribe(ctx context.Context, req *userdemo.CreateSubscribeRequest) error {
	resp, err := client.CreateSubscribe(ctx, req)
	if err != nil {
		return errno.ConvertErr(err)
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func DeleteSubscribe(ctx context.Context, req *userdemo.DeleteSubscribeRequest) error {
	resp, err := client.DeleteSubscribe(ctx, req)
	if err != nil {
		return errno.ConvertErr(err)
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func QuerySubscribe(ctx context.Context, req *userdemo.QuerySubscribeRequest) (*userdemo.UserSub, error) {
	resp, err := client.QuerySubscribe(ctx, req)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Usersub, nil
}

func QueryAllSubscribe(ctx context.Context, req *userdemo.QueryAllSubscribeRequest) ([]string, error) {
	resp, err := client.QueryAllSubscribe(ctx, req)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Subs, nil
}

func QueryALLFans(ctx context.Context, req *userdemo.QueryAllFansRequest) ([]string, error) {
	resp, err := client.QueryAllFans(ctx, req)

	if err != nil {
		return nil, errno.ConvertErr(err)
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Fans, nil
}

func SearchArtAct(ctx context.Context, req *userdemo.SearchArtActRequest) ([]*userdemo.ArtAct, error) {
	resp, err := client.SearchArtAct(ctx, req)
	if err != nil {
		return nil, errno.ConvertErr(err)
	}

	if resp.Resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.Artact, nil

}

func RdbSetUser(ctx context.Context, req *userdemo.RdbSetUserRequest) error {
	resp, err := client.RdbSetUser(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}

func RdbGetUser(ctx context.Context, req *userdemo.RdbGetUserRequest) ([]*userdemo.RdbUser, []string, error) {
	resp, err := client.RdbGetUser(ctx, req)
	if err != nil {
		return nil, nil, err
	}

	if resp.Resp.StatusCode != 0 {
		return nil, nil, errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return resp.RdbUsers, resp.Ungot, nil
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

func UpdateUserPublic(ctx context.Context, req *userdemo.UpdateUserPublicRequest) error {
	resp, err := client.UpdateUserPublic(ctx, req)
	if err != nil {
		return err
	}

	if resp.Resp.StatusCode != 0 {
		return errno.NewErrNo(resp.Resp.StatusCode, resp.Resp.StatusMessage)
	}

	return nil
}
