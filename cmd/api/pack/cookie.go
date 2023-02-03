package pack

import (
	"github.com/gin-gonic/gin"
)

// 设置 cookie
func SetCookie(ctx *gin.Context, name, value string, maxAge int) {
	ctx.SetCookie(name, value, maxAge, "/", "", false, true)
}
