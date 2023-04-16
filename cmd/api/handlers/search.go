package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/actiondemo"
	"be/grpc/articaldemo"
	"be/grpc/searchdemo"
	"be/pkg/check"
	"be/pkg/errno"
	"context"

	"github.com/gin-gonic/gin"
)

func SearchArtical(ctx *gin.Context) {
	var p SearchArticalParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if len(p.KeyWord) >= 30 || p.Limit <= 0 || p.Limit > 20 || p.Offset < 0 {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	IDs, err := rpc.SearchArtical(context.Background(), &searchdemo.SearchArticalRequest{
		Keyword: p.KeyWord,
		Limit:   p.Limit,
		Offset:  p.Offset,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
	}
	if len(IDs) == 0 {
		pack.SendData(ctx, errno.Success, nil)
		return
	}

	artinfos := make([]*ArticalInfo, 0)
	// 查询缓存
	rdbarts, ungot, err := rpc.RdbGetArticalEx(context.Background(), &articaldemo.RdbGetArticalRequest{
		IDs: IDs,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
		return
	}
	if len(ungot) != 0 {
		// 没有查到
		arts, err := rpc.QueryArtical(context.Background(), &articaldemo.QueryArticalRequest{
			IDs: ungot,
		})
		if err != nil {
			pack.SendResponse(ctx, errno.ConvertErr(err))
			return
		}
		rdbarts = append(rdbarts, ChangeArticalToRdbArtical(arts)...)
		// 缓存
		for _, art := range arts {
			rpc.RdbSetArtical(context.Background(), &articaldemo.RdbSetArticalRequest{
				RdbArtical: &articaldemo.RdbArtical{
					ID:        art.ID,
					CreatedAt: art.CreatedAt,
					Title:     art.Title,
					Author:    art.Author,
					// Text: art.Text,
					Description: art.Description,
					LikeNum:     art.LikeNum,
					StarNum:     art.StarNum,
					SeenNum:     art.SeenNum,
					Cover:       art.Cover,
				},
			})
		}
	}

	artinfos = append(artinfos, ChangeRdbArticalToArticalInfo(rdbarts)...)

	pack.SendData(ctx, errno.Success, artinfos)
}

func SearchUserZoom(ctx *gin.Context) {
	var p SearchUserZoomParma
	if err := ctx.ShouldBind(&p); err != nil {
		pack.SendResponse(ctx, errno.ServiceFault)
		return
	}

	// 检测参数
	if !check.CheckUserName(p.UserName) || !check.CheckKeyWord(p.KeyWord) || !check.CheckOffset(p.Offset) {
		pack.SendResponse(ctx, errno.ParamErr)
		return
	}

	targets, err := rpc.SearchUserZoom(context.Background(), &searchdemo.SearchUserZoomRequest{
		Author:  p.UserName,
		Keyword: p.KeyWord,
		Limit:   p.Limit,
		Offset:  p.Offset,
	})
	if err != nil {
		pack.SendResponse(ctx, errno.ConvertErr(err))
	}
	if len(targets) == 0 {
		pack.SendData(ctx, errno.Success, nil)
		return
	}

	type Temp struct {
		Type   int32       `json:"type"`
		Target interface{} `json:"artact"`
	}
	res := make([]*Temp, 0)

	for _, t := range targets {
		// 文章
		if t.Type == 0 {
			artinfo, err := rpc.QueryArticalEx(context.Background(), &articaldemo.QueryArticalRequest{
				IDs: []int32{t.TargetID},
			})
			if err != nil {
				pack.SendResponse(ctx, errno.ConvertErr(err))
			}
			res = append(res, &Temp{
				Type:   t.Type,
				Target: artinfo[0],
			})
		} else if t.Type == 1 {
			// action
			act, err := rpc.QueryAction(context.Background(), &actiondemo.QueryActionRequest{
				IDs: []int32{t.TargetID},
			})
			if err != nil {
				pack.SendResponse(ctx, errno.ConvertErr(err))
			}
			res = append(res, &Temp{
				Type:   t.Type,
				Target: act[0],
			})
		}
	}

	pack.SendData(ctx, errno.Success, res)
}
