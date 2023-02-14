package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/grpc/userdemo"
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"html"
	"log"

	"github.com/gin-gonic/gin"
)

func PublishArtical(ctx *gin.Context) {
	var p PublishArticalParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 作者为空 标题 < 5 > 100 文本 > 50000 描述 <5 > 100 无封面
	if len(p.Author) == 0 || len(p.Text) > 50000 || len(p.Title) < 5 || len(p.Title) > 100 || len(p.Description) < 5 || len(p.Description) > 100 || len(p.Cover) == 0 {
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
	p.Description = html.EscapeString(p.Description)

	err = rpc.CreateArtical(context.Background(), &articaldemo.CreateArticalRequest{
		Author:      p.Author,
		Title:       p.Title,
		Text:        p.Text,
		Description: p.Description,
		Cover:       p.Cover,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 更新缓存
	err = rpc.RdbIncreaseItfUser(context.Background(), &userdemo.RdbIncreaseItfRequest{
		UserName: p.Author,
		Val:      1,
		Field:    constants.RdbUserFieldArtNum,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func DeleteArtical(ctx *gin.Context) {
	var p DeleteArticalParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// ID 不合法
	if p.ArticalID <= 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询是否存在该文章
	arts, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: []int32{p.ArticalID},
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

	// 检查删除者是否与文章作者相同
	err = pack.CheckAuthCookie(ctx, arts[0].Author)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	err = rpc.DeleteArtical(context.Background(), &articaldemo.DeleteArticalRequest{
		ID: p.ArticalID,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 删除文章缓存
	err = rpc.RdbDelArtical(context.Background(), &articaldemo.RdbDelArticalRequest{
		ID: p.ArticalID,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 更新缓存
	err = rpc.RdbIncreaseItfUser(context.Background(), &userdemo.RdbIncreaseItfRequest{
		UserName: arts[0].Author,
		Val:      -1,
		Field:    constants.RdbUserFieldArtNum,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendResponse(ctx, errno.Success)
}

func UpdateArtical(ctx *gin.Context) {
	var p UpdateArticalParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// ID 不合法 标题 < 5 > 100 文本 > 50000 无封面
	if p.ArticalID <= 0 || len(p.Text) > 50000 || len(p.Title) < 5 || len(p.Title) > 100 || len(p.Description) < 5 || len(p.Description) > 100 || len(p.Cover) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询是否存在该文章
	arts, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: []int32{p.ArticalID},
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

	// 目标账户必须与作者相匹配
	err = pack.CheckAuthCookie(ctx, arts[0].Author)
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 转义
	p.Text = html.EscapeString(p.Text)
	p.Title = html.EscapeString(p.Title)
	p.Description = html.EscapeString(p.Description)

	err = rpc.UpdateArtical(context.Background(), &articaldemo.UpdateArticalRequest{
		ArticalID:   p.ArticalID,
		Title:       p.Title,
		Text:        p.Text,
		Description: p.Description,
		Cover:       p.Cover,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 更新缓存
	err = rpc.RdbSetArtical(context.Background(), &articaldemo.RdbSetArticalRequest{
		RdbArtical: &articaldemo.RdbArtical{
			ID:          p.ArticalID,
			CreatedAt:   arts[0].CreatedAt,
			Title:       p.Title,
			Text:        p.Text,
			Description: p.Description,
			Cover:       p.Cover,
			Author:      arts[0].Author,
		},
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

	if len(p.IDs) <= 0 {
		log.Print(len(p.IDs))
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	rdbres, ungot, err := rpc.RdbGetArtical(context.Background(), &articaldemo.RdbGetArticalRequest{
		IDs: p.IDs,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	if len(ungot) == 0 {
		// 全部查询到了
		pack.SendData(ctx, errno.Success, rdbres)
		return
	}

	res, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
		IDs: ungot,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	// 将未查询到的文章全部缓存
	for _, art := range res {
		err := rpc.RdbSetArtical(context.Background(), &articaldemo.RdbSetArticalRequest{
			RdbArtical: ChangeArticalToRdbArtical([]*articaldemo.Artical{art})[0],
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
	}

	setAvator := func(art *articaldemo.RdbArtical) error {
		res, err := rpc.QueryAvator(context.Background(), &userdemo.QueryAvatorRequest{
			UserName: art.Author,
		})
		if err != nil {
			return errno.ConvertErr(err)
		}
		if len(res) == 0 {
			return errno.ServiceFault
		}
		art.AuthorAvator = res[0]
		return nil
	}

	rdbres = append(rdbres, ChangeArticalToRdbArtical(res)...)
	// 查询文章作者头像
	for _, art := range rdbres {
		if err := setAvator(art); err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
	}

	pack.SendData(ctx, errno.Success, rdbres)
}

func GetArticalIDsByAuthor(ctx *gin.Context) {
	author := ctx.Param("author")

	// 名称不合法
	if len(author) == 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	// 查询该作者是否存在
	_, err := rpc.QueryUserInfo(context.Background(), &userdemo.QueryUserInfoRequest{
		UserName: author,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	ids, err := rpc.QueryArticalByAuthor(context.Background(), &articaldemo.QueryArticalByAuthorRequest{
		Author: author,
	})

	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}

	pack.SendData(ctx, errno.Success, ids)
}
