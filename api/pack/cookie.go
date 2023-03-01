package pack

import (
	"be/pkg/constants"
	"be/pkg/errno"
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
	val, err := GetContextValue(ctx, string(constants.AuthCookieUserName))
	if err != nil {
		return errno.ConvertErr(err)
	}

	if val != target {
		return errno.PermissionDeniedErr
	}

	return nil
}

// 获取 ctx 中的 Key 对应的 value
func GetContextValue(ctx *gin.Context, Key string) (interface{}, error) {
	val, exist := ctx.Get(Key)
	if !exist {
		return nil, errno.ServiceFault
	}
	return val, nil
}
