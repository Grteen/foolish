package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
	"be/grpc/articaldemo"
	"be/grpc/searchdemo"
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
			IDs: IDs,
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
