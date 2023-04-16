package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/actiondemo"
	"be/grpc/articaldemo"
	"be/grpc/commentdemo"
	"be/grpc/notifydemo"
	"be/grpc/userdemo"
	"be/pkg/check"
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

	// 检测参数
	if !check.CheckUserName(p.UserName) || !check.CheckCommentText(p.CommentText) || !check.CheckPostiveNumber(p.TargetID) || p.Master < 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与 username 相同
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 被回复的人
	var replyed string
	// 被回复的文章名称
	var replyedArt string

	// 查看是否存在目标
	if p.Type == 0 {
		// 文章回复
		res, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
			IDs: []int32{p.TargetID},
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		if len(res) == 0 {
			pack.SendResponse(ctx, errno.NoSuchArticalErr)
			return
		}
		replyed = res[0].Author
		replyedArt = res[0].Title
	} else if p.Type == 1 {
		// 动态
		res, err := rpc.QueryAction(context.Background(), &actiondemo.QueryActionRequest{
			IDs: []int32{p.TargetID},
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		if len(res) == 0 {
			pack.SendResponse(ctx, errno.NoActionErr)
			return
		}
		replyed = res[0].Author
	} else {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	master := p.Master

	if p.Master != 0 {
		// 查询是否存在评论
		res, err := rpc.QueryComment(context.Background(), &commentdemo.QueryCommentRequest{
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

		// master 评论的 TargetID 必须与 回复 相同
		if res[0].TargetID != p.TargetID {
			pack.SendResponse(ctx, errno.ParamErr)
			return
		}

		// 如果回复的评论也是 reply 则更改 master 为回复的评论的master
		if res[0].Master != 0 {
			p.Master = res[0].Master
		}

		replyed = res[0].UserName
	}

	ids, err := rpc.CreateComment(context.Background(), &commentdemo.CreateCommentRequest{
		UserName:    p.UserName,
		TargetID:    p.TargetID,
		CommentText: p.CommentText,
		Master:      p.Master,
		Type:        p.Type,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 创建回复消息
	if p.UserName != replyed {
		var title string
		if p.Type == 0 {
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
					CommentID: ids[0],
					Master:    master,
					Target: &notifydemo.Target{
						TargetID: p.TargetID,
						Type:     p.Type,
					},
				},
			})
		} else if p.Type == 1 {
			if p.Master != 0 {
				// 回复
				title = "有人回复了你"
			} else {
				// 评论
				title = "有人评论了你的动态"
			}
			err = rpc.CreateReplyNotify(context.Background(), &notifydemo.CreateReplyNotifyRequest{
				Replyntf: &notifydemo.ReplyNotify{
					Sender:    p.UserName,
					UserName:  replyed,
					Title:     title,
					Text:      p.CommentText,
					CommentID: ids[0],
					Master:    master,
					Target: &notifydemo.Target{
						TargetID: p.TargetID,
						Type:     p.Type,
					},
				},
			})
		}
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

	//检测参数
	if !check.CheckPostiveArray(p.ComentIDs) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	cms, err := rpc.QueryComment(context.Background(), &commentdemo.QueryCommentRequest{
		CommentID: p.ComentIDs,
	})

	setAvator := func(cm *commentdemo.Comment) error {
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

func QueryCommentByTargetID(ctx *gin.Context) {
	var p QueryCommentByTargetIDParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// 检测参数
	if !check.CheckPostiveNumber(p.TargetID) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	if p.Type == 0 {
		// 查询是否存在该文章
		arts, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
			IDs: []int32{int32(p.TargetID)},
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
	} else if p.Type == 1 {
		// 查询是否存在该动态
		acts, err := rpc.QueryAction(context.Background(), &actiondemo.QueryActionRequest{
			IDs: []int32{int32(p.TargetID)},
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		// 没有查到动态
		if len(acts) == 0 {
			pack.SendResponse(ctx, errno.NoActionErr)
			return
		}
	} else {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	ids, err := rpc.QueryCommentByTargetID(context.Background(), &commentdemo.QueryCommentByTargetIDRequest{
		TargetID: p.TargetID,
		Type:     p.Type,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, ids)
}

func DeleteComment(ctx *gin.Context) {
	var p DeleteCommentParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// 检测参数
	if !check.CheckPostiveNumber(p.CommentID) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询该评论是否存在
	res, err := rpc.QueryComment(context.Background(), &commentdemo.QueryCommentRequest{
		CommentID: []int32{p.CommentID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(res) == 0 {
		// 没有查询到
		pack.SendResponse(ctx, errno.NoSuchCommentErr)
		return
	}

	// 检测评论用户是否为当前用户
	err = pack.CheckAuthCookie(ctx, res[0].UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.DeleteComment(context.Background(), &commentdemo.DeleteCommentRequest{
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

	// 检测参数
	if !check.CheckPostiveNumber(p.CommentID) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询该评论是否存在
	res, err := rpc.QueryComment(context.Background(), &commentdemo.QueryCommentRequest{
		CommentID: []int32{p.CommentID},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(res) == 0 {
		// 没有查询到
		pack.SendResponse(ctx, errno.NoSuchCommentErr)
		return
	}

	// 检测评论用户是否为当前用户
	err = pack.CheckAuthCookie(ctx, res[0].UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.UpdateComment(context.Background(), &commentdemo.UpdateCommentRequest{
		CommentID:   p.CommentID,
		CommentText: p.CommentText,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}
