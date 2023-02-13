package service

import (
	"be/cmd/user/dal/db"
	"be/grpc/userdemo"
)

func (s *UserService) CreateSubscribe(req *userdemo.CreateSubscribeRequest) error {
	return db.CreateSubscribe(s.ctx, []*db.UserSub{
		{
			User: req.User,
			Sub:  req.Sub,
		},
	})
}

func (s *UserService) DeleteSubscribe(req *userdemo.DeleteSubscribeRequest) error {
	return db.DeleteSubscribe(s.ctx, &db.UserSub{
		User: req.User,
		Sub:  req.Sub,
	})
}

func (s *UserService) QuerySubscribe(req *userdemo.QuerySubscribeRequest) ([]*db.UserSub, error) {
	return db.QuerySubscribe(s.ctx, &db.UserSub{
		User: req.User,
		Sub:  req.Sub,
	})
}

func (s *UserService) QueryAllSubscribe(req *userdemo.QueryAllSubscribeRequest) ([]string, error) {
	return db.QueryAllSubscribe(s.ctx, req.User)
}

func (s *UserService) QueryALLFans(req *userdemo.QueryAllFansRequest) ([]string, error) {
	return db.QueryAllFans(s.ctx, req.User)
}
