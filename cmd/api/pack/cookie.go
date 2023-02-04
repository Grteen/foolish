package pack

import (
	"be/cmd/api/rpc"
	"be/grpc/userdemo"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 设置 cookie
func SetCookie(ctx *gin.Context, name, value string, maxAge int) {
	ctx.SetCookie(name, value, maxAge, "/", "", false, true)
}

// 查询 Authcookie 的值 并返回
func GetAuthCookie(ctx *gin.Context) (string, error) {
	cookie, err := ctx.Cookie(constants.AuthCookieName)
	if err == http.ErrNoCookie {
		return "", errno.AuthenticationCookieExpirationErr
	}

	return cookie, nil
}

// 查看 Authcookie 对应的值 是否等于 target
func CheckAuthCookie(ctx *gin.Context, target string) error {
	cookie, err := ctx.Cookie(constants.AuthCookieName)
	if err == http.ErrNoCookie {
		return errno.AuthenticationCookieExpirationErr
	}

	res, err := rpc.QueryAuthCookie(context.Background(), &userdemo.QueryAuthCookieRequest{
		Key: cookie,
	})

	if err != nil {
		return errno.ServiceFault
	}

	if res != target {
		return errno.PermissionDeniedErr
	}

	return nil
}
