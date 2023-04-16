package service

import (
	"be/cmd/search/dal/db"
	"be/grpc/searchdemo"
	"be/pkg/config"
	"context"
)

type SearchService struct {
	ctx context.Context
}

func NewSearchService(ctx context.Context) *SearchService {
	return &SearchService{ctx: ctx}
}

func (s *SearchService) SearchArtical(req *searchdemo.SearchArticalRequest) ([]int32, error) {
	return db.Search(config.NewConfig(s.ctx, db.DB), req.Keyword, req.Limit, req.Offset)
}

func (s *SearchService) SearchUserZoom(req *searchdemo.SearchUserZoomRequest) ([]*db.Target, error) {
	return db.SearchUserZoom(config.NewConfig(s.ctx, db.DB), req.Author, req.Keyword, req.Limit, req.Offset)
}
