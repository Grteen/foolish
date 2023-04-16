package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/grpc/userdemo"
	"be/pkg/check"
	"be/pkg/errno"
	"context"

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

	// 检测参数
	if !check.CheckUserPassWord(u.PassWord) || !check.CheckUserName(u.Name) || !check.CheckUserEmail(u.Email) {
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

	// 创建用户的默认收藏夹
	err = rpc.CreateStarFolder(context.Background(), &articaldemo.CreateStarFolderRequest{
		UserName:   u.Name,
		FolderName: "默认",
		IsDefault:  true,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}
