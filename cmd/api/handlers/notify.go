package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/notifydemo"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

func QueryAllReplyNotify(ctx *gin.Context) {
	var p QueryAllReplyNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if len(p.UserName) == 0 || p.Limit <= 0 || p.Offset < 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与username相匹配
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	IDs, err := rpc.QueryAllReplyNotify(context.Background(), &notifydemo.QueryAllReplyNotifyRequest{
		UserName: p.UserName,
		Limit:    p.Limit,
		Offset:   p.Offset,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, IDs)
}

func QueryReplyNotify(ctx *gin.Context) {
	var p QueryReplyNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if len(p.IDs) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}
	for _, id := range p.IDs {
		if id <= 0 {
			pack.SendResponse(ctx, errno.ParamErr)
			return
		}
	}

	ntfs, err := rpc.QueryReplyNotify(context.Background(), &notifydemo.QueryReplyNotifyRequest{
		IDs: p.IDs,
	})

	// 查询sender头像
	for _, ntf := range ntfs {
		avator, err := rpc.QueryAvator(context.Background(), &userdemo.QueryAvatorRequest{
			UserName: ntf.Sender,
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
		}
		ntf.Avator = avator[0]
	}
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, ntfs)
}

// 将回复通知设定为已读
func ReadReplyNotify(ctx *gin.Context) {
	var p ReadReplyNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if p.ID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	err := rpc.ReadReplyNotify(context.Background(), &notifydemo.ReadReplyNotifyRequest{
		ID: p.ID,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}
