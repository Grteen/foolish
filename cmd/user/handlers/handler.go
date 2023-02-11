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

// 设置 登录 cookie
func (s *UserServiceImpl) SetAuthCookie(ctx context.Context, req *userdemo.SetAuthCookieRequest) (*userdemo.SetAuthCookieResponse, error) {
	resp := new(userdemo.SetAuthCookieResponse)

	// cookie 无值
	if len(req.Key) == 0 || len(req.Value) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewUserService(ctx).SetAuthCookie(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil

}

// 查询 登录 cookie 并返回 cookie 对应的值  如果不存在 则返回 cookie 过期错误
func (s *UserServiceImpl) QueryAuthCookie(ctx context.Context, req *userdemo.QueryAuthCookieRequest) (*userdemo.QueryAuthCookieResponse, error) {
	resp := new(userdemo.QueryAuthCookieResponse)

	res, err := service.NewUserService(ctx).QueryAuthCookie(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Value = res
	return resp, nil
}

// 删除 登录 cookie
func (s *UserServiceImpl) DeleteAuthCookie(ctx context.Context, req *userdemo.DeleteAuthCookieRequest) (*userdemo.DeleteAuthCookieResponse, error) {
	resp := new(userdemo.DeleteAuthCookieResponse)

	err := service.NewUserService(ctx).DeleteAuthCookie(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 创建一个用户 如果失败返回对应错误
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

// 检测用户和密码是否匹配 如果不匹配返回相应错误
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (*userdemo.CheckUserResponse, error) {
	resp := new(userdemo.CheckUserResponse)

	// 账户 密码 有空
	if len(req.UserNameOrEmail) == 0 || len(req.PassWord) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	username, err := service.NewUserService(ctx).CheckUser(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.UserName = username
	return resp, nil
}

// 更新用户信息 如果更新失败 返回对应错误
func (s *UserServiceImpl) UpdateUserInfo(ctx context.Context, req *userdemo.UpdateUserInfoRequest) (*userdemo.UpdateUserInfoResponse, error) {
	resp := new(userdemo.UpdateUserInfoResponse)

	// 账户 为空
	if len(req.UserName) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewUserService(ctx).UpdateUserInfo(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询用户信息 如果用户不存在 返回用户还未注册错误
func (s *UserServiceImpl) QueryUserInfo(ctx context.Context, req *userdemo.QueryUserInfoRequest) (*userdemo.QueryUserInfoResponse, error) {
	resp := new(userdemo.QueryUserInfoResponse)

	// 账户 为空
	if len(req.UserName) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	ufs, err := service.NewUserService(ctx).QueryUserInfo(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// 用户存在且没查到用户信息
	if len(ufs) == 0 {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.UserInfo = &userdemo.UserInfo{
		UserName:    ufs[0].UserName,
		NickName:    ufs[0].NickName,
		UserAvator:  ufs[0].UserAvator,
		Description: ufs[0].Description,
	}

	return resp, nil
}

// 查询用户头像
func (s *UserServiceImpl) QueryAvator(ctx context.Context, req *userdemo.QueryAvatorRequest) (*userdemo.QueryAvatorResponse, error) {
	resp := new(userdemo.QueryAvatorResponse)

	// 用户名空
	if len(req.UserName) == 0 {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}

	res, err := service.NewUserService(ctx).QueryAvator(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Avator = res

	return resp, nil
}
