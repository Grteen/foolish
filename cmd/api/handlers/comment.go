package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/grpc/notifydemo"
	"be/grpc/userdemo"
	"be/pkg/errno"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(ctx *gin.Context) {
	var p CommentParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// 评论者为空 ArticalID 不合法 文本 > 500 master < 0
	if len(p.UserName) == 0 || len(p.CommentText) > 500 || p.ArticalID <= 0 || p.Master < 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与 username 相同
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
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

	// 被回复的人
	var replyed string
	// 被回复的文章名称
	var replyedArt string

	if p.Master != 0 {
		// 查询是否存在评论
		res, err := rpc.QueryComment(context.Background(), &articaldemo.QueryCommentRequest{
			CommentID: []int32{p.Master},
		})

		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}

		if len(res) == 0 {
			pack.SendResponse(ctx, errno.NoSuchCommentErr)
			return
		}

		// master 评论的 articalID 必须与 回复 相同
		if res[0].ArticalID != p.ArticalID {
			pack.SendResponse(ctx, errno.ParamErr)
			return
		}

		// 如果回复的评论也是 reply 则更改 master 为回复的评论的master
		if res[0].Master != 0 {
			p.Master = res[0].Master
		}

		replyed = res[0].UserName
	} else {
		replyed = res[0].Author
	}

	replyedArt = res[0].Title

	ids, err := rpc.CreateComment(context.Background(), &articaldemo.CreateCommentRequest{
		UserName:    p.UserName,
		ArticalID:   p.ArticalID,
		CommentText: p.CommentText,
		Master:      p.Master,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 创建回复消息
	if p.UserName != replyed {
		var title string
		if p.Master != 0 {
			// 回复
			title = "有人回复了你 : " + replyedArt
		} else {
			// 评论
			title = "有人评论了你的文章 : " + replyedArt
		}
		err = rpc.CreateReplyNotify(context.Background(), &notifydemo.CreateReplyNotifyRequest{
			Replyntf: &notifydemo.ReplyNotify{
				Sender:    p.UserName,
				UserName:  replyed,
				Title:     title,
				Text:      p.CommentText,
				ArticalID: p.ArticalID,
				CommentID: ids[0],
			},
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
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

	setAvator := func(cm *articaldemo.Comment) error {
		res, err := rpc.QueryAvator(context.Background(), &userdemo.QueryAvatorRequest{
			UserName: cm.UserName,
		})
		if err != nil {
			return errno.ConvertErr(err)
		}
		if len(res) == 0 {
			return errno.ServiceFault
		}
		cm.Avator = res[0]
		return nil
	}

	// 查询用户头像
	for _, cm := range cms {
		if len(cm.Reply) != 0 {
			for _, rp := range cm.Reply {
				if err := setAvator(rp); err != nil {
					pack.SendResponse(ctx, errno.ConvertErr(err))
					return
				}
			}
		}
		if err := setAvator(cm); err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
	}

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, cms)
}

func QueryCommentByArticalID(ctx *gin.Context) {
	art := ctx.Param("articalID")
	articalID, err := strconv.ParseInt(art, 10, 32)
	if err != nil {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// ArticalID 非法
	if articalID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询是否存在该文章
	arts, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: []int32{int32(articalID)},
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 没有查到文章
	if len(arts) == 0 {
		pack.SendResponse(ctx, errno.NoSuchArticalErr)
		return
	}

	cms, err := rpc.QueryCommentByArticalID(context.Background(), &articaldemo.QueryCommentByArticalIDRequest{
		ArticalID: int32(articalID),
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
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func UpdateComment(ctx *gin.Context) {
	var p UpdateCommentParma
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

	err = rpc.UpdateComment(context.Background(), &articaldemo.UpdateCommentRequest{
		CommentID:   p.CommentID,
		CommentText: p.CommentText,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}
