package handlers

import (
	"be/cmd/search/pack"
	"be/cmd/search/service"
	"be/grpc/searchdemo"
	"be/pkg/errno"
	"context"
)

// implements the service interface defined in IDL
type SearchServiceImpl struct {
	searchdemo.UnimplementedSearchServiceServer
}

// 根据关键字查询文章
func (s *SearchServiceImpl) SearchArtical(ctx context.Context, req *searchdemo.SearchArticalRequest) (*searchdemo.SearchArticalResponse, error) {
	resp := new(searchdemo.SearchArticalResponse)

	// 检测参数
	if len(req.Keyword) == 0 || len(req.Keyword) >= 30 || req.Limit <= 0 || req.Limit > 20 || req.Offset < 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	res, err := service.NewSearchService(ctx).SearchArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.ArticalIDs = res
	return resp, nil
}
