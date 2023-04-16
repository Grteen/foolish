package service

import (
	"be/cmd/action/dal/db"
	"be/grpc/actiondemo"
	"be/pkg/config"
)

// 点赞动态
func (s *ActionService) CreateActionLike(req *actiondemo.CreateActionLikeRequest) error {
	return db.CreateActionLike(config.NewConfig(s.ctx, db.DB), []*db.Like{
		{
			UserName: req.Actionlike.Username,
			ActionID: req.Actionlike.ActionID,
		},
	})
}

// 取消点赞
func (s *ActionService) DeleteActionLike(req *actiondemo.DeleteActionLikeRequest) error {
	return db.DeleteActionLike(config.NewConfig(s.ctx, db.DB), &db.Like{
		UserName: req.Username,
		ActionID: req.ActionID,
	})
}

// 查询动态点赞
func (s *ActionService) QueryActionLike(req *actiondemo.QueryActionLikeRequest) ([]*db.Like, error) {
	return db.QueryActionLike(config.NewConfig(s.ctx, db.DB), &db.Like{
		UserName: req.Username,
		ActionID: req.ActionID,
	})
}
