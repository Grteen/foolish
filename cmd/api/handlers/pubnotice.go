package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

func CreatePubNotice(ctx *gin.Context) {
	var p CreatePubNoticeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 目标账户必须与username相匹配
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.CreatePubNotice(context.Background(), &userdemo.CreatePubNoticeRequest{
		UserName: p.UserName,
		Text:     p.Text,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func DeletePubNotice(ctx *gin.Context) {
	var p DeletePubNoticeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 查询公告是否存在
	pubs, err := rpc.QueryPubNotice(context.Background(), &userdemo.QueryPubNoticeRequest{
		IDs: []int32{p.ID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 不存在
	if len(pubs) == 0 {
		pack.SendResponse(ctx, errno.NoPubNoticeErr)
		return
	}

	// 目标账户必须与username相匹配
	err = pack.CheckAuthCookie(ctx, pubs[0].UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.DeletePubNotice(context.Background(), &userdemo.DeletePubNoticeRequest{
		ID: p.ID,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func QueryPubNotice(ctx *gin.Context) {
	var p QueryPubNoticeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	pubs, err := rpc.QueryPubNotice(context.Background(), &userdemo.QueryPubNoticeRequest{
		IDs: p.IDs,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, pubs)
}

func QueryUserPubNotice(ctx *gin.Context) {
	var p QueryUserPubNoticeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	ids, err := rpc.QueryUserPubNotice(context.Background(), &userdemo.QueryUserPubNoticeRequest{
		UserName: p.UserName,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, ids)
}
