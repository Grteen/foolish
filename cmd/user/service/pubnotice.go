package service

import (
	"be/cmd/user/dal/db"
	"be/grpc/userdemo"
	"be/pkg/config"
)

// 创建公告
func (s *UserService) CreatePubNotice(req *userdemo.CreatePubNoticeRequest) error {
	return db.CreatePubNotice(config.NewConfig(s.ctx, db.DB), []*db.PubNotice{
		{
			UserName: req.UserName,
			Text:     req.Text,
		},
	})
}

// 删除公告
func (s *UserService) DeletePubNotice(req *userdemo.DeletePubNoticeRequest) error {
	return db.DeletePubNotice(config.NewConfig(s.ctx, db.DB), req.ID)
}

// 根据ID查询公告
func (s *UserService) QueryPubNotice(req *userdemo.QueryPubNoticeRequest) ([]*db.PubNotice, error) {
	return db.QueryPubNotice(config.NewConfig(s.ctx, db.DB), req.IDs)
}

// 查询某人的所有公告
func (s *UserService) QueryUserPubNotice(req *userdemo.QueryUserPubNoticeRequest) ([]int32, error) {
	return db.QueryUserPubNotice(config.NewConfig(s.ctx, db.DB), req.UserName, req.Limit, req.Offset)
}
