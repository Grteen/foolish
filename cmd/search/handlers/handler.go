package handlers

import (
	"be/cmd/search/pack"
	"be/cmd/search/service"
	"be/grpc/searchdemo"
	"be/pkg/check"
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
	if !check.CheckKeyWord(req.Keyword) || !check.CheckOffset(req.Offset) || !check.CheckSearchLimit(req.Limit) {
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

// 个人空间内搜索
func (s *SearchServiceImpl) SearchUserZoom(ctx context.Context, req *searchdemo.SearchUserZoomRequest) (*searchdemo.SearchUserZoomResponse, error) {
	resp := new(searchdemo.SearchUserZoomResponse)

	// 检测参数
	if !check.CheckKeyWord(req.Keyword) || !check.CheckOffset(req.Offset) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	res, err := service.NewSearchService(ctx).SearchUserZoom(req)
	if err != nil {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	for _, t := range res {
		resp.Targets = append(resp.Targets, &searchdemo.Target{
			TargetID: t.TargetID,
			Type:     t.Type,
		})
	}
	return resp, nil
}
