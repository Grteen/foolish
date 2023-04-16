package handlers

import (
	"be/cmd/user/pack"
	"be/cmd/user/service"
	"be/grpc/userdemo"
	"be/pkg/check"
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

// 查询用户
func (s *UserServiceImpl) QueryUser(ctx context.Context, req *userdemo.QueryUserRequest) (*userdemo.QueryUserResponse, error) {
	resp := new(userdemo.QueryUserResponse)

	// 账户为空
	if len(req.User) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	usresp, err := s.QueryUserEx(ctx, req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}
	if usresp.Resp.StatusCode != 0 {
		resp.Resp = pack.BuildRespByResp(usresp.Resp.StatusCode, usresp.Resp.StatusMessage)
		return resp, nil
	}

	ufresp, err := s.QueryUserInfo(ctx, &userdemo.QueryUserInfoRequest{
		UserName: req.User,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}
	if ufresp.Resp.StatusCode != 0 {
		resp.Resp = pack.BuildRespByResp(ufresp.Resp.StatusCode, ufresp.Resp.StatusMessage)
		return resp, nil
	}

	// 缓存用户
	err = service.NewUserService(ctx).RdbSetUser(&userdemo.RdbSetUserRequest{
		RdbUser: &userdemo.RdbUser{
			UserName:        usresp.User[0].UserName,
			NickName:        ufresp.UserInfo.NickName,
			Description:     ufresp.UserInfo.Description,
			IsAdministrator: usresp.User[0].IsAdministrator,
			UserAvator:      ufresp.UserInfo.UserAvator,
			SubNum:          usresp.User[0].SubNum,
			FanNum:          usresp.User[0].FanNum,
			ArtNum:          usresp.User[0].ArtNum,
			FanPublic:       usresp.User[0].FanPublic,
			SubPublic:       usresp.User[0].SubPublic,
		},
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.User = append(resp.User, &userdemo.User{
		UserName:        usresp.User[0].UserName,
		SubNum:          usresp.User[0].SubNum,
		FanNum:          usresp.User[0].FanNum,
		ArtNum:          usresp.User[0].ArtNum,
		FanPublic:       usresp.User[0].FanPublic,
		SubPublic:       usresp.User[0].SubPublic,
		IsAdministrator: usresp.User[0].IsAdministrator,
		UserInfo: &userdemo.UserInfo{
			UserName:    ufresp.UserInfo.UserName,
			NickName:    ufresp.UserInfo.NickName,
			Description: ufresp.UserInfo.Description,
			UserAvator:  ufresp.UserInfo.UserAvator,
		},
	})
	return resp, nil
}

// 更改用户 关注粉丝列表 权限
func (s *UserServiceImpl) UpdateUserPublic(ctx context.Context, req *userdemo.UpdateUserPublicRequest) (*userdemo.UpdateUserPublicResponse, error) {
	resp := new(userdemo.UpdateUserPublicResponse)

	// 检测参数
	if !check.CheckUserName(req.UserName) || !check.CheckPublic(req.SubPublic) || !check.CheckPublic(req.FanPublic) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewUserService(ctx).UpdateUserPublic(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// 修改缓存
	err = service.NewUserService(ctx).RdbSetUserPublic(&userdemo.RdbSetUserPublicRequest{
		Username:  req.UserName,
		FanPublic: req.FanPublic,
		SubPublic: req.SubPublic,
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询用户的粉丝数 文章数 订阅数
func (s *UserServiceImpl) QueryUserEx(ctx context.Context, req *userdemo.QueryUserRequest) (*userdemo.QueryUserResponse, error) {
	resp := new(userdemo.QueryUserResponse)

	// 账户为空
	if len(req.User) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询缓存
	rdbus, ungot, err := service.NewUserService(ctx).RdbGetUser(&userdemo.RdbGetUserRequest{
		Users: []string{req.User},
	})

	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}
	if len(ungot) != 0 {
		req.User = ungot[0]
		us, err := service.NewUserService(ctx).QueryUser(req)
		if err != nil {
			resp.Resp = pack.BuildResp(errno.ConvertErr(err))
			return resp, nil
		}
		rdbus = append(rdbus, ChangeUserToRdbUser(us)...)
	}
	if len(rdbus) == 0 {
		resp.Resp = pack.BuildResp(errno.UserNotRegisterErr)
		return resp, nil
	}
	resp.Resp = pack.BuildResp(errno.Success)
	for _, u := range rdbus {
		resp.User = append(resp.User, &userdemo.User{
			UserName:        u.UserName,
			SubNum:          u.SubNum,
			FanNum:          u.FanNum,
			ArtNum:          u.ArtNum,
			FanPublic:       u.FanPublic,
			SubPublic:       u.SubPublic,
			IsAdministrator: u.IsAdministrator,
		})
	}

	return resp, nil
}

// 查询用户信息 如果用户不存在 返回用户还未注册错误
// UserInfo is not the RdbUser 所以不要缓存
func (s *UserServiceImpl) QueryUserInfo(ctx context.Context, req *userdemo.QueryUserInfoRequest) (*userdemo.QueryUserInfoResponse, error) {
	resp := new(userdemo.QueryUserInfoResponse)

	// 账户 为空
	if len(req.UserName) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询缓存
	rdbufs, ungot, err := service.NewUserService(ctx).RdbGetUser(&userdemo.RdbGetUserRequest{
		Users: []string{req.UserName},
	})
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}
	if len(ungot) != 0 {
		req.UserName = ungot[0]
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
		rdbufs = append(rdbufs, ChangeUserInfoToRdbUserInfo(ufs)...)
	}
	if len(rdbufs) == 0 {
		resp.Resp = pack.BuildResp(errno.UserNotRegisterErr)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.UserInfo = &userdemo.UserInfo{
		UserName:    rdbufs[0].UserName,
		NickName:    rdbufs[0].NickName,
		UserAvator:  rdbufs[0].UserAvator,
		Description: rdbufs[0].Description,
	}

	return resp, nil
}

// 更新用户信息 如果更新失败 返回对应错误
// UserInfo is not the RdbUser 所以不要缓存
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

// 订阅
func (s *UserServiceImpl) CreateSubscribe(ctx context.Context, req *userdemo.CreateSubscribeRequest) (*userdemo.CreateSubscribeResponse, error) {
	resp := new(userdemo.CreateSubscribeResponse)

	// 用户名 空
	if len(req.User) == 0 || len(req.Sub) == 0 {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}

	// 查询是否存在 user 对 sub 的订阅
	uss, err := service.NewUserService(ctx).QuerySubscribe(&userdemo.QuerySubscribeRequest{
		User: req.User,
		Sub:  req.Sub,
	})

	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// 存在 user 对于 sub 的订阅
	if len(uss) != 0 {
		resp.Resp = pack.BuildResp(errno.AlreadySubscribeErr)
		return resp, nil
	}

	err = service.NewUserService(ctx).CreateSubscribe(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 取消订阅
func (s *UserServiceImpl) DeleteSubscribe(ctx context.Context, req *userdemo.DeleteSubscribeRequest) (*userdemo.DeleteSubscribeResponse, error) {
	resp := new(userdemo.DeleteSubscribeResponse)

	// 用户名 空
	if len(req.User) == 0 || len(req.Sub) == 0 {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}

	// 查询是否存在 user 对 sub 的订阅
	uss, err := service.NewUserService(ctx).QuerySubscribe(&userdemo.QuerySubscribeRequest{
		User: req.User,
		Sub:  req.Sub,
	})

	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// 不存在 user 对于 sub 的订阅
	if len(uss) == 0 {
		resp.Resp = pack.BuildResp(errno.NoSubscribeErr)
		return resp, nil
	}

	err = service.NewUserService(ctx).DeleteSubscribe(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 查询是否有user对于sub的订阅  不存在则返回 NoSubscribeErr
func (s *UserServiceImpl) QuerySubscribe(ctx context.Context, req *userdemo.QuerySubscribeRequest) (*userdemo.QuerySubscribeResponse, error) {
	resp := new(userdemo.QuerySubscribeResponse)

	// 用户名 空
	if len(req.User) == 0 || len(req.Sub) == 0 {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}
	uss, err := service.NewUserService(ctx).QuerySubscribe(&userdemo.QuerySubscribeRequest{
		User: req.User,
		Sub:  req.Sub,
	})

	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	if len(uss) == 0 {
		resp.Resp = pack.BuildResp(errno.NoSubscribeErr)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Usersub = &userdemo.UserSub{
		User: uss[0].User,
		Sub:  uss[0].Sub,
	}
	return resp, nil
}

// 查询 用户的所有订阅 返回订阅的用户名
func (s *UserServiceImpl) QueryAllSubscribe(ctx context.Context, req *userdemo.QueryAllSubscribeRequest) (*userdemo.QueryAllSubscribeResponse, error) {
	resp := new(userdemo.QueryAllSubscribeResponse)

	// 用户名 空
	if len(req.User) == 0 {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}

	uss, err := service.NewUserService(ctx).QueryAllSubscribe(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Subs = uss
	return resp, nil
}

// 查询 用户的所有粉丝 返回粉丝的用户名
func (s *UserServiceImpl) QueryAllFans(ctx context.Context, req *userdemo.QueryAllFansRequest) (*userdemo.QueryAllFansResponse, error) {
	resp := new(userdemo.QueryAllFansResponse)

	// 用户名 空
	if len(req.User) == 0 {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}

	uss, err := service.NewUserService(ctx).QueryALLFans(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Fans = uss
	return resp, nil
}

// 查询所有动态和文章的 id 并按照时间降序返回
func (s *UserServiceImpl) SearchArtAct(ctx context.Context, req *userdemo.SearchArtActRequest) (*userdemo.SearchArtActResponse, error) {
	resp := new(userdemo.SearchArtActResponse)

	// 检测参数
	if !check.CheckUserName(req.UserName) || !check.CheckZeroOrPostive(req.Limit) || !check.CheckZeroOrPostive(req.Offset) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	ats, err := service.NewUserService(ctx).SearchArtAct(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	for _, at := range ats {
		resp.Artact = append(resp.Artact, &userdemo.ArtAct{
			ID:        int32(at.ID),
			CreatedAt: at.CreatedAt.In(pack.Tz).Format(pack.TimeLayout),
			Type:      at.Type,
		})
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 将 RdbUser 存储在 redis 中
func (s *UserServiceImpl) RdbSetUser(ctx context.Context, req *userdemo.RdbSetUserRequest) (*userdemo.RdbSetUserResponse, error) {
	resp := new(userdemo.RdbSetUserResponse)

	// 非用户输入 无需验证
	err := service.NewUserService(ctx).RdbSetUser(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

// 获取 RdbUser
func (s *UserServiceImpl) RdbGetUser(ctx context.Context, req *userdemo.RdbGetUserRequest) (*userdemo.RdbGetUserResponse, error) {
	resp := new(userdemo.RdbGetUserResponse)

	// users 为空
	if len(req.Users) == 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	users, ungot, err := service.NewUserService(ctx).RdbGetUser(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Ungot = ungot
	for _, user := range users {
		resp.RdbUsers = append(resp.RdbUsers, &userdemo.RdbUser{
			UserName:    user.UserName,
			NickName:    user.NickName,
			Description: user.Description,
			UserAvator:  user.UserAvator,
			SubNum:      user.SubNum,
			FanNum:      user.FanNum,
			ArtNum:      user.ArtNum,
		})
	}

	return resp, nil
}

// 增加 粉丝数 关注数 文章数
func (s *UserServiceImpl) RdbIncreaseItf(ctx context.Context, req *userdemo.RdbIncreaseItfRequest) (*userdemo.RdbIncreaseItfResponse, error) {
	resp := new(userdemo.RdbIncreaseItfResponse)

	// 参数检测
	if len(req.UserName) == 0 || len(req.Field) == 0 || req.Val <= -2 || req.Val >= 2 {
		resp.Resp = pack.BuildResp(errno.ServiceFault)
		return resp, nil
	}

	err := service.NewUserService(ctx).RdbIncreaseItf(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}
