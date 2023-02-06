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

func (s *ArticalServiceImpl) QueryArtical(ctx context.Context, req *articaldemo.QueryArticalRequest) (*articaldemo.QueryArticalResponse, error) {
	resp := new(articaldemo.QueryArticalResponse)

	art, err := service.NewArticalService(ctx).QueryArtical(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	// 没有查询到文章
	if art.ID == 0 {
		resp.Resp = pack.BuildResp(errno.NoSuchArticalErr)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Artical = &articaldemo.Artical{
		ID:     int64(art.ID),
		Author: art.Author,
		Title:  art.Title,
		Text:   art.Text,
	}

	return resp, nil
}

func (s *ArticalServiceImpl) CreateLike(ctx context.Context, req *articaldemo.CreateLikeRequest) (*articaldemo.CreateLikeResponse, error) {
	resp := new(articaldemo.CreateLikeResponse)

	// username 为空 ID 不合法
	if len(req.UserName) == 0 || req.ArticalID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询是否有该用户对于该文章的点赞
	res, err := service.NewArticalService(ctx).QueryLike(&articaldemo.QueryLikeRequest{
		UserName:  req.UserName,
		ArticalID: req.ArticalID,
	})

	if err != nil && err != errno.NoLikesErr {
		resp.Resp = pack.BuildResp(errno.ConvertErr(err))
		return resp, nil
	}

	// 已经有点赞了
	if len(res) != 0 {
		resp.Resp = pack.BuildResp(errno.AlreadyLikesErr)
		return resp, nil
	}

	err = service.NewArticalService(ctx).CreateLike(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

func (s *ArticalServiceImpl) DeleteLike(ctx context.Context, req *articaldemo.DeleteLikeRequest) (*articaldemo.DeleteLikeResponse, error) {
	resp := new(articaldemo.DeleteLikeResponse)

	// username 为空 ID 不合法
	if len(req.UserName) == 0 || req.ArticalID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	// 查询是否有该用户对于该文章的点赞
	_, err := service.NewArticalService(ctx).QueryLike(&articaldemo.QueryLikeRequest{
		UserName:  req.UserName,
		ArticalID: req.ArticalID,
	})

	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	err = service.NewArticalService(ctx).DeleteLike(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	return resp, nil
}

func (s *ArticalServiceImpl) QueryLike(ctx context.Context, req *articaldemo.QueryLikeRequest) (*articaldemo.QueryLikeResponse, error) {
	resp := new(articaldemo.QueryLikeResponse)

	// username 为空 ID 不合法
	if len(req.UserName) == 0 || req.ArticalID <= 0 {
		resp.Resp = pack.BuildResp(errno.ParamErr)
		return resp, nil
	}

	res, err := service.NewArticalService(ctx).QueryLike(req)
	if err != nil {
		resp.Resp = pack.BuildResp(err)
		return resp, nil
	}

	resp.Resp = pack.BuildResp(errno.Success)
	resp.Like = &articaldemo.Like{
		UserName:  res[0].UserName,
		ArticalID: int64(res[0].ArticalID),
	}

	return resp, nil
}
