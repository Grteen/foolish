package service

import (
	"be/cmd/action/dal/db"
	"be/grpc/actiondemo"
	"be/pkg/config"
	"context"
)

type ActionService struct {
	ctx context.Context
}

func NewActionService(ctx context.Context) *ActionService {
	return &ActionService{ctx: ctx}
}

// 创建动态
func (s *ActionService) CreateAction(req *actiondemo.CreateActionRequest) error {
	pics := make([]*db.PicFile, 0)
	for _, pic := range req.Picfiles {
		pics = append(pics, &db.PicFile{
			File: pic,
		})
	}
	return db.CreateAction(config.NewConfig(s.ctx, db.DB), []*db.Action{
		{
			Author:  req.Author,
			Text:    req.Text,
			PicFile: pics,
		},
	})
}

// 根据动态ID查询动态
func (s *ActionService) QueryAction(req *actiondemo.QueryActionRequest) ([]*db.Action, error) {
	return db.QueryAction(config.NewConfig(s.ctx, db.DB), req.IDs)
}

// 查询某个人的所有动态
func (s *ActionService) QueryActionByAuthor(req *actiondemo.QueryActionByAuthorRequest) ([]int32, error) {
	return db.QueryArticalByAuthor(config.NewConfig(s.ctx, db.DB), req.Author, req.Field, req.Order)
}

// 根据动态ID删除动态
func (s *ActionService) DeleteAction(req *actiondemo.DeleteActionRequest) error {
	return db.DeleteAction(config.NewConfig(s.ctx, db.DB), req.ID)
}
