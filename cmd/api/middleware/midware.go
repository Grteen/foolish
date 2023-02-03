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

// 查看是否具有 登陆过后的 Auth Cookie 如果没有则停止访问并返回 AuthenticationCookieExpirationErr
func AuthMidWare(ctx *gin.Context) {
	cookie, err := ctx.Cookie(constants.AuthCookieName)
	if err == http.ErrNoCookie {
		pack.SendResponse(ctx, errno.AuthenticationCookieExpirationErr)
		ctx.Abort()
		return
	}

	_, err = rpc.QueryAuthCookie(context.Background(), &userdemo.QueryAuthCookieRequest{
		Key: cookie,
	})

	if err != nil {
		pack.SendResponse(ctx, err)
		ctx.Abort()
		return
	}

	ctx.Next()
	return
}
