package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/pkg/errno"
	"context"
	"html"

	"github.com/gin-gonic/gin"
)

func PublishArtical(ctx *gin.Context) {
	var p PublishArticalParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 作者为空 标题 < 5 && > 100 文本 > 50000
	if len(p.Author) == 0 || len(p.Text) > 50000 || (len(p.Title) < 5 && len(p.Title) > 100) {
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
	p.Title = html.EscapeString(p.Title)

	err = rpc.CreateArtical(context.Background(), &articaldemo.CreateArticalRequest{
		Author: p.Author,
		Title:  p.Title,
		Text:   p.Text,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func GetArtical(ctx *gin.Context) {
	var p GetArticalParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, err)
		return
	}

	if p.ID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	res, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		ID: p.ID,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, res)
}
