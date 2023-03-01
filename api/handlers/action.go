package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/actiondemo"
	"be/grpc/notifydemo"
	"be/grpc/userdemo"
	"be/pkg/check"
	"be/pkg/errno"
	"context"
	"html"

	"github.com/gin-gonic/gin"
)

func PublishAction(ctx *gin.Context) {
	var p PublishActionParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if !check.CheckUserName(p.Author) || (!check.CheckActionText(p.Text) && !check.CheckStringArray(p.PicFiles)) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与作者相匹配
	err := pack.CheckAuthCookie(ctx, p.Author)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 转义
	p.Text = html.EscapeString(p.Text)

	err = rpc.CreateAction(context.Background(), &actiondemo.CreateActionRequest{
		Author:   p.Author,
		Text:     p.Text,
		Picfiles: p.PicFiles,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func DeleteAction(ctx *gin.Context) {
	var p DeleteActionParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if !check.CheckPostiveNumber(p.ID) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询是否存在该动态
	acts, err := rpc.QueryAction(context.Background(), &actiondemo.QueryActionRequest{
		IDs: []int32{p.ID},
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

	// 检查删除者是否与动态作者相同
	err = pack.CheckAuthCookie(ctx, acts[0].Author)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.DeleteAction(context.Background(), &actiondemo.DeleteActionRequest{
		ID: p.ID,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func GetAction(ctx *gin.Context) {
	var p GetActionParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// 检测参数
	if !check.CheckPostiveArray(p.IDs) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	acts, err := rpc.QueryAction(context.Background(), &actiondemo.QueryActionRequest{
		IDs: p.IDs,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, acts)
}

func GetActionByAuthor(ctx *gin.Context) {
	var p GetActionByAuthorParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// 检测参数
	if !check.CheckUserName(p.Author) || !check.CheckStringLength(p.Field) || !check.CheckStringLength(p.Order) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询该作者是否存在
	_, err := rpc.QueryUserInfo(context.Background(), &userdemo.QueryUserInfoRequest{
		UserName: p.Author,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	var fieldTemp = map[string]string{
		"time": "updated_at",
		"like": "likeNum",
	}

	var orderTemp = map[string]string{
		"ASC":  "ASC",
		"DESC": "DESC",
	}

	field := fieldTemp[p.Field]
	order := orderTemp[p.Order]

	ids, err := rpc.QueryActionByAuthor(context.Background(), &actiondemo.QueryActionByAuthorRequest{
		Author: p.Author,
		Field:  field,
		Order:  order,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, ids)
}

func CreateActionLike(ctx *gin.Context) {
	var p CreateActionLikeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// 检测参数
	if !check.CheckUserName(p.UserName) || !check.CheckPostiveNumber(p.ActionID) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与 username 相同
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 检查是否存在动态
	acts, err := rpc.QueryAction(context.Background(), &actiondemo.QueryActionRequest{
		IDs: []int32{p.ActionID},
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

	err = rpc.CreateActionLike(context.Background(), &actiondemo.CreateActionLikeRequest{
		Actionlike: &actiondemo.ActionLike{
			Username: p.UserName,
			ActionID: p.ActionID,
		},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 创建点赞通知
	err = rpc.CreateLikeNotify(context.Background(), &notifydemo.CreateLikeNotifyRequest{
		Likentf: &notifydemo.LikeNotify{
			UserName: acts[0].Author,
			Sender:   p.UserName,
			Title:    "收到了点赞",
			Text:     p.UserName + "点赞了你的动态",
			Target: &notifydemo.Target{
				TargetID: p.ActionID,
				Type:     1,
			},
		},
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func DeleteActionLike(ctx *gin.Context) {
	var p DeleteActionLikeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// 检测参数
	if !check.CheckUserName(p.UserName) || !check.CheckPostiveNumber(p.ActionID) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与 username 相同
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 检查是否存在动态
	acts, err := rpc.QueryAction(context.Background(), &actiondemo.QueryActionRequest{
		IDs: []int32{p.ActionID},
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

	err = rpc.DeleteActionLike(context.Background(), &actiondemo.DeleteActionLikeRequest{
		Username: p.UserName,
		ActionID: p.ActionID,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func QueryActionLike(ctx *gin.Context) {
	var p QueryActionLikeParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	// 检测参数
	if !check.CheckUserName(p.UserName) || !check.CheckPostiveNumber(p.ActionID) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 目标账户必须与 username 相同
	err := pack.CheckAuthCookie(ctx, p.UserName)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 检查是否存在动态
	acts, err := rpc.QueryAction(context.Background(), &actiondemo.QueryActionRequest{
		IDs: []int32{p.ActionID},
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

	likes, err := rpc.QueryActionLike(context.Background(), &actiondemo.QueryActionLikeRequest{
		Username: p.UserName,
		ActionID: p.ActionID,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	// 没有点赞
	if len(likes) == 0 {
		pack.SendData(ctx, errno.Success, false)
		return
	}

	// 有点赞收藏
	pack.SendData(ctx, errno.Success, true)
}
