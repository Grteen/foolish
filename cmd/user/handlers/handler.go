package handlers

import (
	"be/cmd/user/pack"
	"be/cmd/user/service"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"
)

// implements the service interface defined in IDL
type UserServiceImpl struct {
	userdemo.UnimplementedUserServiceServer
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (*userdemo.CreateUserResponse, error) {
	resp := new(userdemo.CreateUserResponse)

	// 名称 密码 Email 有空
	if len(req.UserName) == 0 || len(req.PassWord) == 0 || len(req.Email) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewUserService(ctx).CreateUser(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (*userdemo.CheckUserResponse, error) {
	resp := new(userdemo.CheckUserResponse)

	// 账户 密码 有空
	if len(req.UserNameOrEmail) == 0 || len(req.PassWord) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewUserService(ctx).CheckUser(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}
