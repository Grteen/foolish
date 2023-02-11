package handlers

import (
	"be/cmd/api/pack"
	"be/cmd/api/rpc"
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
	if len(p.KeyWord) == 0 || len(p.KeyWord) >= 30 || p.Limit <= 0 || p.Limit > 20 || p.Offset < 0 {
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
		return
	}

	pack.SendData(ctx, errno.Success, IDs)
}
