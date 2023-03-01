package service

import (
	"be/cmd/user/dal/db"
	"be/grpc/userdemo"
	"be/pkg/config"
)

func (s *UserService) CreateSubscribe(req *userdemo.CreateSubscribeRequest) error {
	return db.CreateSubscribe(config.NewConfig(s.ctx, db.DB), []*db.UserSub{
		{
			User: req.User,
			Sub:  req.Sub,
		},
	})
}

func (s *UserService) DeleteSubscribe(req *userdemo.DeleteSubscribeRequest) error {
	return db.DeleteSubscribe(config.NewConfig(s.ctx, db.DB), &db.UserSub{
		User: req.User,
		Sub:  req.Sub,
	})
}

func (s *UserService) QuerySubscribe(req *userdemo.QuerySubscribeRequest) ([]*db.UserSub, error) {
	return db.QuerySubscribe(config.NewConfig(s.ctx, db.DB), &db.UserSub{
		User: req.User,
		Sub:  req.Sub,
	})
}

func (s *UserService) QueryAllSubscribe(req *userdemo.QueryAllSubscribeRequest) ([]string, error) {
	return db.QueryAllSubscribe(config.NewConfig(s.ctx, db.DB), req.User)
}

func (s *UserService) QueryALLFans(req *userdemo.QueryAllFansRequest) ([]string, error) {
	return db.QueryAllFans(config.NewConfig(s.ctx, db.DB), req.User)
}
