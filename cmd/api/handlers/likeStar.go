package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

// tp = 0  Like 请求
// tp = 1  Star 请求
func GiveLikeStar(ctx *gin.Context, tp int32) {
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
	res, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: []int32{p.ArticalID},
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	if len(res) == 0 {
		pack.SendResponse(ctx, errno.NoSuchArticalErr)
		return
	}

	// 目标账户必须与 username 相同
	err = pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.CreateLikeStar(context.Background(), &articaldemo.CreateLikeStarRequest{
		UserName:  p.UserName,
		ArticalID: p.ArticalID,
		Type:      tp,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func DeleteLikeStar(ctx *gin.Context, tp int32) {
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
	res, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: []int32{p.ArticalID},
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	if len(res) == 0 {
		pack.SendResponse(ctx, errno.NoSuchArticalErr)
		return
	}

	// 目标账户必须与 username 相同
	err = pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.DeleteLikeStar(context.Background(), &articaldemo.DeleteLikeStarRequest{
		ArticalID: p.ArticalID,
		UserName:  p.UserName,
		Type:      tp,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

// tp = 0 Like 请求
// tp = 1 Star 请求
func QueryAllLikeStar(ctx *gin.Context, tp int32) {
	userName := ctx.Query("username")

	// 查看是否存在该用户
	// 如果不存在则返回 10006 错误
	_, err := rpc.QueryUserInfo(ctx, &userdemo.QueryUserInfoRequest{
		UserName: userName,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	ids, err := rpc.QueryAllLikeStar(context.Background(), &articaldemo.QueryAllLikeStarRequest{
		UserName: userName,
		Type:     tp,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, map[string][]uint32{
		"ArticalIDs": ids,
	})
}

func GiveLike(ctx *gin.Context) {
	GiveLikeStar(ctx, 0)
}

func DeleteLike(ctx *gin.Context) {
	DeleteLikeStar(ctx, 0)
}

func GiveStar(ctx *gin.Context) {
	GiveLikeStar(ctx, 1)
}

func DeleteStar(ctx *gin.Context) {
	DeleteLikeStar(ctx, 1)
}

func QueryAllStar(ctx *gin.Context) {
	QueryAllLikeStar(ctx, 1)
}
