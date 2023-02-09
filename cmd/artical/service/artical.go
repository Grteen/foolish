package service

import (
	"be/cmd/artical/dal/db"
	"be/cmd/artical/dal/rdb"
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
			Author:      req.Author,
			Title:       req.Title,
			Text:        req.Text,
			Description: req.Description,
		},
	})
}

func (s *ArticalService) QueryArtical(req *articaldemo.QueryArticalRequest) ([]*db.Artical, error) {
	return db.QueryArtical(s.ctx, req.IDs)
}

func (s *ArticalService) DeleteArtical(req *articaldemo.DeleteArticalRequest) error {
	return db.DeleteArtical(s.ctx, req.ID)
}

// 不更新作者
func (s *ArticalService) UpdateArtical(req *articaldemo.UpdateArticalRequest) error {
	return db.UpdateArtical(s.ctx, &db.Artical{
		ID:          uint(req.ArticalID),
		Title:       req.Title,
		Text:        req.Text,
		Description: req.Description,
	})
}

func (s *ArticalService) QueryArticalByAuthor(req *articaldemo.QueryArticalByAuthorRequest) ([]int32, error) {
	return db.QueryArticalByAuthor(s.ctx, req.Author)
}

func (s *ArticalService) RdbSetArtical(req *articaldemo.RdbSetArticalRequest) error {
	return rdb.SetArtical(s.ctx, []*rdb.RdbArtical{
		{
			ID:        uint(req.RdbArtical.ID),
			CreatedAt: req.RdbArtical.CreateTime.AsTime(),
			Title:     req.RdbArtical.Title,
			Author:    req.RdbArtical.Author,
			Text:      req.RdbArtical.Text,

			LikeNum: req.RdbArtical.LikeNum,
			StarNum: req.RdbArtical.StarNum,
			SeenNum: req.RdbArtical.SeenNum,
		},
	})
}

func (s *ArticalService) RdbGetArtical(req *articaldemo.RdbGetArticalRequest) ([]*rdb.RdbArtical, []int32, error) {
	return rdb.GetArtical(s.ctx, req.IDs)
}
