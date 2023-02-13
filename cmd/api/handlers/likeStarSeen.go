package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/grpc/userdemo"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

// tp = 0  Like 请求
// tp = 1  Star 请求
// tp = 2  Seen 请求
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

	// 更新缓存
	var field string
	if tp == 0 {
		// Like
		field = constants.RdbArticalFieldLikeNum
	} else if tp == 1 {
		// Star
		field = constants.RdbArticalFieldStarNum
	} else if tp == 2 {
		// Seen
		// Seen 请求 不更新缓存
		pack.SendResponse(ctx, errno.Success)
		return
	} else {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	err = rpc.RdbIncreaseitf(context.Background(), &articaldemo.RdbIncreaseitfRequest{
		ArticalID: p.ArticalID,
		Val:       1,
		Field:     field,
	})
	pack.SendResponse(ctx, errno.Success)
}

// tp = 0  Like 请求
// tp = 1  Star 请求
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

	// 更新缓存
	var field string
	if tp == 0 {
		// Like
		field = constants.RdbArticalFieldLikeNum
	} else if tp == 1 {
		// Star
		field = constants.RdbArticalFieldStarNum
	} else if tp == 2 {
		// Seen
		field = constants.RdbArticalFieldSeenNum
	} else {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	err = rpc.RdbIncreaseitf(context.Background(), &articaldemo.RdbIncreaseitfRequest{
		ArticalID: p.ArticalID,
		Val:       -1,
		Field:     field,
	})

	pack.SendResponse(ctx, errno.Success)
}

// tp = 0 Like 请求
// tp = 1 Star 请求
func QueryLikeStar(ctx *gin.Context, tp int32) {
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

	_, err = rpc.QueryLikeStar(context.Background(), &articaldemo.QueryLikeStarRequest{
		UserName:  p.UserName,
		ArticalID: p.ArticalID,
		Type:      tp,
	})

	if err != nil && err != errno.NoLikeStarErr {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 没有点赞收藏
	if err == errno.NoLikeStarErr {
		pack.SendData(ctx, errno.Success, false)
		return
	}

	// 有点赞收藏
	pack.SendData(ctx, errno.Success, true)
}

// tp = 0 Like 请求
// tp = 1 Star 请求
// tp = 2 Seen 请求
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

func GiveSeen(ctx *gin.Context) {
	GiveLikeStar(ctx, 2)
}

func QueryAllSeen(ctx *gin.Context) {
	QueryAllLikeStar(ctx, 2)
}

func HasLike(ctx *gin.Context) {
	QueryLikeStar(ctx, 0)
}

func HasStar(ctx *gin.Context) {
	QueryLikeStar(ctx, 1)
}
