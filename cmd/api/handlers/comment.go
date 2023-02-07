package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	var p CommentParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// 评论者为空 ArticalID 不合法 文本 > 500
	if len(p.UserName) == 0 || len(p.CommentText) > 500 || p.ArticalID <= 0 {
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
	}

	// 目标账户必须与 username 相同
	err = pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.CreateComment(context.Background(), &articaldemo.CreateCommentRequest{
		UserName:    p.UserName,
		ArticalID:   p.ArticalID,
		CommentText: p.CommentText,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func QueryComment(ctx *gin.Context) {
	var p QueryCommentParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// CommentIDs 为空 或 ID 不合法
	if len(p.ComentIDs) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	for _, i := range p.ComentIDs {
		if i <= 0 {
			pack.SendResponse(ctx, errno.ParamErr)
			return
		}
	}

	cms, err := rpc.QueryComment(context.Background(), &articaldemo.QueryCommentRequest{
		CommentID: p.ComentIDs,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, cms)
}

func DeleteComment(ctx *gin.Context) {
	var p DeleteCommentParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// id 非法
	if p.CommentID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询该评论是否存在
	res, err := rpc.QueryComment(context.Background(), &articaldemo.QueryCommentRequest{
		CommentID: []int32{p.CommentID},
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	if len(res) == 0 {
		// 没有查询到
		pack.SendResponse(ctx, errno.NoSuchArticalErr)
		return
	}

	// 检测评论用户是否为当前用户
	err = pack.CheckAuthCookie(ctx, res[0].UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.DeleteComment(context.Background(), &articaldemo.DeleteCommentRequest{
		CommentID: p.CommentID,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
	}

	pack.SendResponse(ctx, errno.Success)
}
