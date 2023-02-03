package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"
	"regexp"

	"github.com/gin-gonic/gin"
)

// 注册用户
func Register(ctx *gin.Context) {
	var u UserParma
	if err := ctx.ShouldBind(&u); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 如果三者中有一个为空
	if len(u.Name) == 0 || len(u.PassWord) == 0 || len(u.Email) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 检查参数
	userPwReg := regexp.MustCompile("[0-9A-Za-z_\\-]{3,18}")
	if !userPwReg.MatchString(u.Name) || !userPwReg.MatchString(u.PassWord) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	emailReg := regexp.MustCompile("[0-9A-Za-z]+@qq.com")
	if !emailReg.MatchString(u.Email) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 创建用户
	err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		UserName: u.Name,
		PassWord: u.PassWord,
		Email:    u.Email,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}
