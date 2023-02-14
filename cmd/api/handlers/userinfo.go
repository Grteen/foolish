package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"
	"regexp"

	"github.com/gin-gonic/gin"
)

// 更新用户信息 只更新非 0 值字段
func UpdateUserInfo(ctx *gin.Context) {
	var u UpdateUserInfoParma
	if err := ctx.ShouldBind(&u); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 账户为空
	if len(u.UserName) == 0 || len(u.Description) == 0 || len(u.NickName) == 0 || len(u.Avator) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与账户相匹配
	err := pack.CheckAuthCookie(ctx, u.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// Description 小于 500  昵称 小于 20
	if len(u.Description) > 500 || len(u.NickName) > 20 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 检测参数
	if len(u.Description) != 0 {
		descReg := regexp.MustCompile("[^\u4e00-\u9fa5a-z0-9A-Z_\\-]")
		if descReg.MatchString(u.Description) {
			pack.SendResponse(ctx, errno.ParamErr)
			return
		}
	}

	if len(u.NickName) != 0 {
		nickReg := regexp.MustCompile("[^\u4e00-\u9fa5a-z0-9A-Z_\\-]")
		if nickReg.MatchString(u.NickName) {
			pack.SendResponse(ctx, errno.ParamErr)
			return
		}
	}

	if len(u.Avator) != 0 {
		avatorReg := regexp.MustCompile("[^a-z0-9A-Z:/.]")
		if avatorReg.MatchString(u.Avator) {
			pack.SendResponse(ctx, errno.ParamErr)
			return
		}
	}

	us, err := rpc.QueryUser(context.Background(), &userdemo.QueryUserRequest{
		User: u.UserName,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.UpdateUserInfo(context.Background(), &userdemo.UpdateUserInfoRequest{
		UserName:    u.UserName,
		Description: u.Description,
		NickName:    u.NickName,
		UserAvator:  u.Avator,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 更新缓存
	err = rpc.RdbSetUser(context.Background(), &userdemo.RdbSetUserRequest{
		RdbUser: &userdemo.RdbUser{
			UserName:    u.UserName,
			NickName:    u.NickName,
			Description: u.Description,
			UserAvator:  u.Avator,
			SubNum:      us[0].SubNum,
			FanNum:      us[0].FanNum,
			ArtNum:      us[0].ArtNum,
		},
	})

	pack.SendResponse(ctx, errno.Success)
}

// 根据 UserName 查询用户信息
func QueryUserInfo(ctx *gin.Context) {
	var p QueryUserInfoParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 账户为空
	if len(p.UserName) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询缓存
	rdbres, ungot, err := rpc.RdbGetUser(context.Background(), &userdemo.RdbGetUserRequest{
		Users: []string{p.UserName},
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	if len(ungot) == 0 {
		// 全部查询到了
		pack.SendData(ctx, errno.Success, rdbres[0])
		return
	}

	res, err := rpc.QueryUserInfo(context.Background(), &userdemo.QueryUserInfoRequest{
		UserName: ungot[0],
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	uf := &struct {
		UserName    string `json:"userName"`
		NickName    string `json:"nickName"`
		Description string `json:"description"`
		Avator      string `json:"userAvator"`
		SubNum      int32  `json:"subNum"`
		FanNum      int32  `json:"fanNum"`
		ArtNum      int32  `json:"artNum"`
	}{UserName: res.UserName, NickName: res.NickName, Description: res.Description, Avator: res.UserAvator}

	// 查询 订阅数 粉丝数 和 文章数
	us, err := rpc.QueryUser(context.Background(), &userdemo.QueryUserRequest{
		User: ungot[0],
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	uf.SubNum = us[0].SubNum
	uf.FanNum = us[0].FanNum
	uf.ArtNum = us[0].ArtNum

	// 全部缓存
	err = rpc.RdbSetUser(context.Background(), &userdemo.RdbSetUserRequest{
		RdbUser: &userdemo.RdbUser{
			UserName:    uf.UserName,
			NickName:    uf.NickName,
			Description: uf.Description,
			UserAvator:  uf.Avator,
			SubNum:      uf.SubNum,
			FanNum:      uf.FanNum,
			ArtNum:      uf.ArtNum,
		},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, uf)
}

// 根据当前 Cookie 查询 UserName
func QueryUserSelf(ctx *gin.Context) {
	cookie, err := pack.GetAuthCookie(ctx)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	res, err := rpc.QueryAuthCookie(context.Background(), &userdemo.QueryAuthCookieRequest{
		Key: cookie,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
	}

	pack.SendData(ctx, errno.Success, gin.H{
		"username": res,
	})
}
