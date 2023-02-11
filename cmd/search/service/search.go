package service

import (
	"be/cmd/search/dal/db"
	"be/grpc/searchdemo"
	"context"
)

type SearchService struct {
	ctx context.Context
}

func NewSearchService(ctx context.Context) *SearchService {
	return &SearchService{ctx: ctx}
}

func (s *SearchService) SearchArtical(req *searchdemo.SearchArticalRequest) ([]int32, error) {
	return db.Search(s.ctx, req.Keyword, req.Limit, req.Offset)
}
