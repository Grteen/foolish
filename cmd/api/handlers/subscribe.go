package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/userdemo"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

func Subscribe(ctx *gin.Context) {
	var p SubscribeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// userName空
	if len(p.User) == 0 || len(p.Sub) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与 user 匹配
	err := pack.CheckAuthCookie(ctx, p.User)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 查询 sub 是否存在 不存在返回未注册错误
	_, err = rpc.QueryUserInfo(context.Background(), &userdemo.QueryUserInfoRequest{
		UserName: p.Sub,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.CreateSubscribe(context.Background(), &userdemo.CreateSubscribeRequest{
		User: p.User,
		Sub:  p.Sub,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 更新缓存
	err = rpc.RdbIncreaseItfUser(context.Background(), &userdemo.RdbIncreaseItfRequest{
		UserName: p.User,
		Val:      1,
		Field:    constants.RdbUserFieldSubNum,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.RdbIncreaseItfUser(context.Background(), &userdemo.RdbIncreaseItfRequest{
		UserName: p.Sub,
		Val:      1,
		Field:    constants.RdbUserFieldFanNum,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func UnSubscribe(ctx *gin.Context) {
	var p SubscribeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// userName空
	if len(p.User) == 0 || len(p.Sub) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与 user 匹配
	err := pack.CheckAuthCookie(ctx, p.User)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 查询 sub 是否存在 不存在返回未注册错误
	_, err = rpc.QueryUserInfo(context.Background(), &userdemo.QueryUserInfoRequest{
		UserName: p.Sub,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.DeleteSubscribe(context.Background(), &userdemo.DeleteSubscribeRequest{
		User: p.User,
		Sub:  p.Sub,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 更新缓存
	err = rpc.RdbIncreaseItfUser(context.Background(), &userdemo.RdbIncreaseItfRequest{
		UserName: p.User,
		Val:      -1,
		Field:    constants.RdbUserFieldSubNum,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.RdbIncreaseItfUser(context.Background(), &userdemo.RdbIncreaseItfRequest{
		UserName: p.Sub,
		Val:      -1,
		Field:    constants.RdbUserFieldFanNum,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func QuerySubscribe(ctx *gin.Context) {
	var p SubscribeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// userName空
	if len(p.User) == 0 || len(p.Sub) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与 user 匹配
	err := pack.CheckAuthCookie(ctx, p.User)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 查询 sub 是否存在 不存在返回未注册错误
	_, err = rpc.QueryUserInfo(context.Background(), &userdemo.QueryUserInfoRequest{
		UserName: p.Sub,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	_, err = rpc.QuerySubscribe(context.Background(), &userdemo.QuerySubscribeRequest{
		User: p.User,
		Sub:  p.Sub,
	})
	if err != nil && err != errno.NoSubscribeErr {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 没有 p.User 对于 p.Sub 的订阅
	if err == errno.NoSubscribeErr {
		pack.SendData(ctx, errno.Success, false)
		return
	}

	pack.SendData(ctx, errno.Success, true)
}

func QueryAllSubscribe(ctx *gin.Context) {
	var p QueryAllSubscribeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// userName空
	if len(p.User) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询 user 是否存在 不存在返回未注册错误
	u, err := rpc.QueryUser(context.Background(), &userdemo.QueryUserRequest{
		User: p.User,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 检测权限
	if u[0].SubPublic == 0 {
		// 不公开
		// 目标账户必须与 user 匹配
		err := pack.CheckAuthCookie(ctx, p.User)
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
	}

	uss, err := rpc.QueryAllSubscribe(context.Background(), &userdemo.QueryAllSubscribeRequest{
		User: p.User,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, uss)
}

func QueryALLFans(ctx *gin.Context) {
	var p QueryAllFansParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// userName空
	if len(p.User) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询 user 是否存在 不存在返回未注册错误
	u, err := rpc.QueryUser(context.Background(), &userdemo.QueryUserRequest{
		User: p.User,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 检测权限
	if u[0].FanPublic == 0 {
		// 不公开
		// 目标账户必须与 user 匹配
		err := pack.CheckAuthCookie(ctx, p.User)
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
	}

	uss, err := rpc.QueryALLFans(context.Background(), &userdemo.QueryAllFansRequest{
		User: p.User,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, uss)
}
