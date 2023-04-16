package middleware

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/userdemo"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 跨域问题
func AccessMidWare(ctx *gin.Context) {
	method := ctx.Request.Method

	// ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
	ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
}

// 查看是否具有 登陆过后的 Auth Cookie 如果没有则停止访问并返回 AuthenticationCookieExpirationErr
func AuthMidWare(ctx *gin.Context) {
	cookie, err := ctx.Cookie(constants.AuthCookieName)
	if err == http.ErrNoCookie {
		pack.SendResponse(ctx, errno.AuthenticationCookieExpirationErr)
		ctx.Abort()
		return
	}

	userName, err := rpc.QueryAuthCookie(context.Background(), &userdemo.QueryAuthCookieRequest{
		Key: cookie,
	})
	if err != nil {
		pack.SendResponse(ctx, err)
		ctx.Abort()
		return
	}

	// 查询是否为管理员
	ia, err := rpc.QueryUser(context.Background(), &userdemo.QueryUserRequest{
		User: userName,
	})
	if err != nil {
		pack.SendResponse(ctx, err)
		ctx.Abort()
		return
	}

	// 设置该 cookie 对应的用户名
	ctx.Set(string(constants.AuthCookieUserName), userName)
	ctx.Set(string(constants.AuthCookieIsAdministrator), ia[0].IsAdministrator)
	ctx.Next()
	return
}
