package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/commentdemo"
	"be/grpc/notifydemo"
	"be/grpc/userdemo"
	"be/pkg/check"
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
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	type temp struct {
		Type   int32       `json:"type"`
		Notify interface{} `json:"notify"`
	}
	res := make([]*temp, 0)

	// 查询sender头像 和被评论的文本
	for _, ntf := range ntfs {
		// 如果通知所属人不属于当前用户 则不返回该通知
		err := pack.CheckAuthCookie(ctx, ntf.UserName)
		if err != nil {
			continue
		}

		avator, err := rpc.QueryAvator(context.Background(), &userdemo.QueryAvatorRequest{
			UserName: ntf.Sender,
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		ntf.Avator = avator[0]
		if ntf.Master != 0 {
			cmtext, err := rpc.QueryComment(context.Background(), &commentdemo.QueryCommentRequest{
				CommentID: []int32{ntf.Master},
			})
			if err != nil {
				pack.SendResponse(ctx, errno.ConvertErr(err))
				return
			}
			// 没有该评论
			if len(cmtext) == 0 {
				pack.SendResponse(ctx, errno.NoSuchCommentErr)
				return
			}
			ntf.MasterText = cmtext[0].CommentText
		}
		res = append(res, &temp{
			Type:   0,
			Notify: ntf,
		})
	}

	pack.SendData(ctx, errno.Success, res)
}

// 将回复通知设定为已读
func ReadReplyNotify(ctx *gin.Context) {
	var p ReadNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	ReadNotify(ctx, p.ID, 0)
}

// 将回复通知设置为已删除
func DeleteReplyNotify(ctx *gin.Context) {
	var p DeleteNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	DeleteNotify(ctx, p.ID, 0)
}

func QueryAllLikeNotify(ctx *gin.Context) {
	var p QueryAllLikeNotifyParma
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

	IDs, err := rpc.QueryAllLikeNotify(context.Background(), &notifydemo.QueryAllLikeNotifyRequest{
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

func QueryLikeNotify(ctx *gin.Context) {
	var p QueryLikeNotifyParma
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

	ltfs, err := rpc.QueryLikeNotify(context.Background(), &notifydemo.QueryLikeNotifyRequest{
		IDs: p.IDs,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	type temp struct {
		Type   int32       `json:"type"`
		Notify interface{} `json:"notify"`
	}
	res := make([]*temp, 0)

	// 查询sender头像
	for _, ltf := range ltfs {
		// 如果通知所属人不属于当前用户 则不返回该通知
		err := pack.CheckAuthCookie(ctx, ltf.UserName)
		if err != nil {
			continue
		}
		avator, err := rpc.QueryAvator(context.Background(), &userdemo.QueryAvatorRequest{
			UserName: ltf.Sender,
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		ltf.Avator = avator[0]
		res = append(res, &temp{
			Type:   1,
			Notify: ltf,
		})
	}

	pack.SendData(ctx, errno.Success, res)
}

// 讲点赞通知设定为已读
func ReadLikeNotify(ctx *gin.Context) {
	var p ReadNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	ReadNotify(ctx, p.ID, 1)
}

// 将点赞通知设置为已删除
func DeleteLikeNotify(ctx *gin.Context) {
	var p DeleteNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	DeleteNotify(ctx, p.ID, 1)
}

// 将通知设定为已读
func ReadNotify(ctx *gin.Context, ID int32, tp int32) {
	if ID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 通知所属人
	var owner string
	// 查询通知是否存在
	if tp == 0 {
		// reply notify
		ntfs, err := rpc.QueryReplyNotify(context.Background(), &notifydemo.QueryReplyNotifyRequest{
			IDs: []int32{ID},
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		if len(ntfs) == 0 {
			pack.SendResponse(ctx, errno.NoNotifyErr)
			return
		}
		owner = ntfs[0].UserName
	} else if tp == 1 {
		// like notify
		ltfs, err := rpc.QueryLikeNotify(context.Background(), &notifydemo.QueryLikeNotifyRequest{
			IDs: []int32{ID},
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		if len(ltfs) == 0 {
			pack.SendResponse(ctx, errno.NoNotifyErr)
			return
		}
		owner = ltfs[0].UserName
	} else {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 目标账户必须与username相匹配
	err := pack.CheckAuthCookie(ctx, owner)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.ReadNotify(context.Background(), &notifydemo.ReadNotifyRequest{
		ID:   ID,
		Type: tp,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

// 将通知设定为已删除
func DeleteNotify(ctx *gin.Context, ID int32, tp int32) {
	if ID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	var owner string
	// 查询通知是否存在
	if tp == 0 {
		// reply notify
		ntfs, err := rpc.QueryReplyNotify(context.Background(), &notifydemo.QueryReplyNotifyRequest{
			IDs: []int32{ID},
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		if len(ntfs) == 0 {
			pack.SendResponse(ctx, errno.NoNotifyErr)
			return
		}
		owner = ntfs[0].UserName
	} else if tp == 1 {
		// like notify
		ltfs, err := rpc.QueryLikeNotify(context.Background(), &notifydemo.QueryLikeNotifyRequest{
			IDs: []int32{ID},
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		if len(ltfs) == 0 {
			pack.SendResponse(ctx, errno.NoNotifyErr)
			return
		}
		owner = ltfs[0].UserName
	} else {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 目标账户必须与username相匹配
	err := pack.CheckAuthCookie(ctx, owner)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.DeleteNotify(context.Background(), &notifydemo.DeleteNotifyRequest{
		ID:   ID,
		Type: tp,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

// 查询所有通知的id 并按照时间降序返回
func SearchAllNotify(ctx *gin.Context) {
	var p SearchAllNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if p.Limit < 0 || p.Offset < 0 || len(p.UserName) <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与username相匹配
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	ntfs, err := rpc.SearchAllNotify(context.Background(), &notifydemo.SearchAllNotifyRequest{
		UserName: p.UserName,
		Limit:    p.Limit,
		Offset:   p.Offset,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	type temp struct {
		Type   int32       `json:"type"`
		Notify interface{} `json:"notify"`
	}
	res := make([]*temp, 0)

	// 查询具体的通知信息
	for _, ntf := range ntfs {
		// replyNotify
		if ntf.NotifyType == 0 {
			rtf, err := rpc.QueryReplyNotify(context.Background(), &notifydemo.QueryReplyNotifyRequest{
				IDs: []int32{ntf.ID},
			})
			if err != nil {
				pack.SendResponse(ctx, errno.ConvertErr(err))
				return
			}
			if rtf[0].Master != 0 {
				cmtext, err := rpc.QueryComment(context.Background(), &commentdemo.QueryCommentRequest{
					CommentID: []int32{rtf[0].Master},
				})
				if err != nil {
					pack.SendResponse(ctx, errno.ConvertErr(err))
					return
				}
				// 没有该评论
				if len(cmtext) == 0 {
					pack.SendResponse(ctx, errno.NoSuchCommentErr)
					return
				}
				rtf[0].MasterText = cmtext[0].CommentText
			}
			res = append(res, &temp{
				Type:   ntf.NotifyType,
				Notify: rtf[0],
			})
		} else if ntf.NotifyType == 1 {
			ltf, err := rpc.QueryLikeNotify(context.Background(), &notifydemo.QueryLikeNotifyRequest{
				IDs: []int32{ntf.ID},
			})
			if err != nil {
				pack.SendResponse(ctx, errno.ConvertErr(err))
				return
			}
			res = append(res, &temp{
				Type:   ntf.NotifyType,
				Notify: ltf[0],
			})
		}
	}

	pack.SendData(ctx, errno.Success, res)
}

// 创建系统消息
func CreateSystemNotify(ctx *gin.Context) {
	var p CreateSystemNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 目标账户必须与username相匹配
	err := pack.CheckAuthCookie(ctx, "")
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 检测参数
	if !check.CheckNotifyText(p.Text) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	err = rpc.CreateSystemNotify(context.Background(), &notifydemo.CreateSystemNotifyRequest{
		Text: p.Text,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

// 查询所有系统消息
func QueryAllSystemNotify(ctx *gin.Context) {
	var p QueryAllSystemNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if !check.CheckOffset(p.Offset) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	res, err := rpc.QueryAllSystemNotify(context.Background(), &notifydemo.QueryAllSystemNotifyRequest{
		Limit:  p.Limit,
		Offset: p.Offset,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, res)
}

// 根据ID查询指定系统消息
func QuerySystemNotify(ctx *gin.Context) {
	var p QuerySystemNotifyParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if !check.CheckPostiveArray(p.IDs) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	res, err := rpc.QuerySystemNotify(context.Background(), &notifydemo.QuerySystemNotifyRequest{
		IDs: p.IDs,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, res)
}
