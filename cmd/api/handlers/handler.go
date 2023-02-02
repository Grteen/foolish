package handlers

import (
	"be/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

func SendResponse(ctx *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	ctx.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}

type UserParma struct {
	Name     string `form:"name"`
	PassWord string `form:"password"`
	Email    string `form:"email"`
}

type LoginParma struct {
	NameOrEmail string `form:"account"`
	PassWord    string `form:"password"`
}
