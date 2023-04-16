package service

import (
	"be/cmd/artical/dal/db"
	"be/cmd/artical/dal/rdb"
	"be/grpc/articaldemo"
	"be/pkg/config"
	"context"

	"gorm.io/gorm"
)

type ArticalService struct {
	ctx context.Context
}

func NewArticalService(ctx context.Context) *ArticalService {
	return &ArticalService{ctx: ctx}
}

func (s *ArticalService) CreateArtical(req *articaldemo.CreateArticalRequest) error {
	return db.CreateArtical(config.NewConfig(s.ctx, db.DB), []*db.Artical{
		{
			Author:      req.Author,
			Title:       req.Title,
			Text:        req.Text,
			Description: req.Description,
			Cover:       req.Cover,
			// CreatedAt:   time.Now(),
			// LikeNum:     0,
			// StarNum:     0,
			// SeenNum:     0,
		},
	})
}

func (s *ArticalService) QueryArtical(req *articaldemo.QueryArticalRequest) ([]*db.Artical, error) {
	return db.QueryArtical(config.NewConfig(s.ctx, db.DB), req.IDs)
}

func (s *ArticalService) DeleteArtical(req *articaldemo.DeleteArticalRequest) error {
	return db.DeleteArtical(config.NewConfig(s.ctx, db.DB), req.ID)
}

// 不更新作者
func (s *ArticalService) UpdateArtical(req *articaldemo.UpdateArticalRequest) error {
	return db.UpdateArtical(config.NewConfig(s.ctx, db.DB), &db.Artical{
		Model: gorm.Model{
			ID: uint(req.ArticalID),
		},
		Title:       req.Title,
		Text:        req.Text,
		Description: req.Description,
		Cover:       req.Cover,
	})
}

func (s *ArticalService) QueryArticalByAuthor(req *articaldemo.QueryArticalByAuthorRequest) ([]int32, error) {
	return db.QueryArticalByAuthor(config.NewConfig(s.ctx, db.DB), req.Author, req.Field, req.Order)
}

func (s *ArticalService) RdbSetArtical(req *articaldemo.RdbSetArticalRequest) error {
	return rdb.SetArtical(s.ctx, []*rdb.RdbArtical{
		{
			ID:          uint(req.RdbArtical.ID),
			CreatedAt:   req.RdbArtical.CreatedAt,
			Title:       req.RdbArtical.Title,
			Author:      req.RdbArtical.Author,
			Text:        req.RdbArtical.Text,
			Description: req.RdbArtical.Description,

			LikeNum:      req.RdbArtical.LikeNum,
			StarNum:      req.RdbArtical.StarNum,
			SeenNum:      req.RdbArtical.SeenNum,
			Cover:        req.RdbArtical.Cover,
			AuthorAvator: req.RdbArtical.AuthorAvator,
		},
	})
}

func (s *ArticalService) RdbDelArtical(req *articaldemo.RdbDelArticalRequest) error {
	return rdb.DelArtical(s.ctx, req.ID)
}

func (s *ArticalService) RdbGetArtical(req *articaldemo.RdbGetArticalRequest) ([]*rdb.RdbArtical, []int32, error) {
	return rdb.GetArtical(s.ctx, req.IDs)
}

// 不获取 Text 的版本
func (s *ArticalService) RdbGetArticalEx(req *articaldemo.RdbGetArticalRequest) ([]*rdb.RdbArtical, []int32, error) {
	return rdb.GetArticalInfo(s.ctx, req.IDs)
}
