package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

func GiveLike(ctx *gin.Context) {
	var p LikeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	if p.ArticalID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查看是否存在文章
	_, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		ID: p.ArticalID,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 目标账户必须与 username 相同
	err = pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.CreateLike(context.Background(), &articaldemo.CreateLikeRequest{
		UserName:  p.UserName,
		ArticalID: p.ArticalID,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func DeleteLike(ctx *gin.Context) {
	var p LikeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	if p.ArticalID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查看是否存在文章
	_, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		ID: p.ArticalID,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 目标账户必须与 username 相同
	err = pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.DeleteLike(context.Background(), &articaldemo.DeleteLikeRequest{
		ArticalID: p.ArticalID,
		UserName:  p.UserName,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}
