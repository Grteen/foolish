package handlers

import (
	"be/cmd/artical/pack"
	"be/cmd/artical/service"
	"be/grpc/articaldemo"
	"be/pkg/errno"
	"context"
)

// implements the service interface defined in IDL
type ArticalServiceImpl struct {
	articaldemo.UnimplementedArticalServiceServer
}

func (s *ArticalServiceImpl) CreateArtical(ctx context.Context, req *articaldemo.CreateArticalRequest) (*articaldemo.CreateArticalResponse, error) {
	resp := new(articaldemo.CreateArticalResponse)

	// 作者为空 标题 < 5 && > 100 文本 > 50000
	if len(req.Author) == 0 || len(req.Text) > 50000 || (len(req.Title) < 5 && len(req.Title) > 100) {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	err := service.NewArticalService(ctx).CreateArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}
