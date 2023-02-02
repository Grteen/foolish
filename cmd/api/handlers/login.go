package handlers

import (
	"be/cmd/api/rpc"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var u LoginParma
	if err := ctx.ShouldBind(&u); err != nil {
		SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 账户或密码为空
	if len(u.NameOrEmail) == 0 || len(u.PassWord) == 0 {
		SendResponse(ctx, errno.ParamErr)
		return
	}

	err := rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{
		UserNameOrEmail: u.NameOrEmail,
		PassWord:        u.PassWord,
	})

	if err != nil {
		SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	SendResponse(ctx, errno.Success)
}
