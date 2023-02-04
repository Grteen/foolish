package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"

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
	if len(u.UserName) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与账户相匹配
	err := pack.CheckAuthCookie(ctx, u.UserName)
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

	res, err := rpc.QueryUserInfo(context.Background(), &userdemo.QueryUserInfoRequest{
		UserName: p.UserName,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, res)
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
