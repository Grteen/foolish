package service

import (
	"be/cmd/user/dal/db"
	"be/grpc/userdemo"
	"be/pkg/config"
	"be/pkg/errno"
)

func (s *UserService) UpdateUserInfo(req *userdemo.UpdateUserInfoRequest) error {
	// 查询 UserName 是否存在
	users, err := db.QueryUser(config.NewConfig(s.ctx, db.DB), req.UserName)
	if err != nil {
		return err
	}
	if len(users) == 0 {
		return errno.UserNotRegisterErr
	}

	return db.UpdateUserInfo(config.NewConfig(s.ctx, db.DB), &db.UserInfo{
		UserName:    req.UserName,
		UserAvator:  req.UserAvator,
		NickName:    req.NickName,
		Description: req.Description,
	})
}

func (s *UserService) QueryUserInfo(req *userdemo.QueryUserInfoRequest) ([]*db.UserInfo, error) {
	// 查询 UserName 是否存在
	users, err := db.QueryUser(config.NewConfig(s.ctx, db.DB), req.UserName)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserNotRegisterErr
	}

	return db.QueryUserInfo(config.NewConfig(s.ctx, db.DB), req.UserName)
}

func (s *UserService) QueryAvator(req *userdemo.QueryAvatorRequest) ([]string, error) {
	return db.QueryAvator(config.NewConfig(s.ctx, db.DB), req.UserName)
}
