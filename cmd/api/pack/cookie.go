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
	val, _ := GetContextValue(ctx, string(constants.AuthCookieUserName))

	if val != target {
		// 检查是否是管理员
		isa, _ := GetContextValue(ctx, string(constants.AuthCookieIsAdministrator))
		isai, ok := isa.(int32)
		if !ok {
			return errno.ServiceFault
		}
		if isai == 1 {
			return nil
		}
		return errno.PermissionDeniedErr
	}

	return nil
}

// 获取 ctx 中的 Key 对应的 value
func GetContextValue(ctx *gin.Context, Key string) (interface{}, error) {
	val, exist := ctx.Get(Key)
	if !exist {
		return nil, nil
	}
	return val, nil
}
