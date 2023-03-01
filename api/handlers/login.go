package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/userdemo"
	"be/pkg/constants"
	"be/pkg/errno"
	"be/pkg/uuid"
	"context"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var u LoginParma
	if err := ctx.ShouldBind(&u); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 账户或密码为空
	if len(u.NameOrEmail) == 0 || len(u.PassWord) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	username, err := rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{
		UserNameOrEmail: u.NameOrEmail,
		PassWord:        u.PassWord,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	uuid := uuid.GetUUid()
	// 检查是否有 cookie 了
	oldval, err := rpc.QueryAuthCookie(context.Background(), &userdemo.QueryAuthCookieRequest{
		Key: username,
	})

	// 错误
	if err != nil && err != errno.AuthenticationCookieExpirationErr {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	if err == errno.AuthenticationCookieExpirationErr {
		// 无 cookie
		setAuthCookie(ctx, uuid, username, constants.LoginCookieTime)
	} else {
		// 有 cookie
		// 设置新 cookie 并删除原来的 cookie
		setAuthCookie(ctx, uuid, username, constants.LoginCookieTime)
		rpc.DeleteAuthCookie(ctx, &userdemo.DeleteAuthCookieRequest{
			Key: oldval,
		})
	}

	pack.SendResponse(ctx, errno.Success)
}

func DeLogin(ctx *gin.Context) {
	var p DeLoginParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 目标账户必须与 username 相同
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 查询cookie
	cok, err := rpc.QueryAuthCookie(context.Background(), &userdemo.QueryAuthCookieRequest{
		Key: p.UserName,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 删除cookie
	err = rpc.DeleteAuthCookie(context.Background(), &userdemo.DeleteAuthCookieRequest{
		Key: cok,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func setAuthCookie(ctx *gin.Context, key, value string, maxAge int) {
	// 设置 session
	rpc.SetAuthCookie(context.Background(), &userdemo.SetAuthCookieRequest{
		Key:    key,
		Value:  value,
		MaxAge: int64(maxAge) * constants.ChangeToRedis,
	})

	rpc.SetAuthCookie(context.Background(), &userdemo.SetAuthCookieRequest{
		Key:    value,
		Value:  key,
		MaxAge: int64(maxAge) * constants.ChangeToRedis,
	})

	// 设置 cookie
	pack.SetCookie(ctx, constants.AuthCookieName, key, maxAge)
}
