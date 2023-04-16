package pack

import (
	"be/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type DataResponse struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(ctx *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	ctx.JSON(http.StatusOK, Response{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
	})
}

func SendData(ctx *gin.Context, err error, data interface{}) {
	Err := errno.ConvertErr(err)
	ctx.JSON(http.StatusOK, DataResponse{
		Code:    Err.ErrCode,
		Message: Err.ErrMsg,
		Data:    data,
	})
}
