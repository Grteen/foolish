package service

import (
	"be/cmd/artical/dal/db"
	"be/grpc/articaldemo"
	"context"
)

type ArticalService struct {
	ctx context.Context
}

func NewArticalService(ctx context.Context) *ArticalService {
	return &ArticalService{ctx: ctx}
}

func (s *ArticalService) CreateArtical(req *articaldemo.CreateArticalRequest) error {
	return db.CreateArtical(s.ctx, []*db.Artical{
		{
			Author: req.Author,
			Title:  req.Title,
			Text:   req.Text,
		},
	})
}
