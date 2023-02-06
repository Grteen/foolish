package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/userdemo"
	"be/pkg/constants"
	"be/pkg/errno"
	"be/pkg/uuid"
	"context"
	"fmt"

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
	fmt.Println(username)

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	uuid := uuid.GetUUid()
	// 设置 Cookie 和 Seesion
	setAuthCookie(ctx, uuid, username, constants.LoginCookieTime)
	pack.SendResponse(ctx, errno.Success)
}

func setAuthCookie(ctx *gin.Context, key, value string, maxAge int) {
	// 设置 session
	rpc.SetAuthCookie(context.Background(), &userdemo.SetAuthCookieRequest{
		Key:    key,
		Value:  value,
		MaxAge: int64(maxAge) * constants.ChangeToRedis,
	})

	// 设置 cookie
	pack.SetCookie(ctx, constants.AuthCookieName, key, maxAge)
}
