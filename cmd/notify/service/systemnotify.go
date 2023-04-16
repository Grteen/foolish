package service

import (
	"be/cmd/notify/dal/db"
	"be/grpc/notifydemo"
	"be/pkg/config"
)

// 创建系统消息
func (s *NotifyService) CreateSystemNotify(req *notifydemo.CreateSystemNotifyRequest) error {
	return db.CreateSystemNotify(config.NewConfig(s.ctx, db.DB), []*db.SystemNotify{
		{
			Text: req.Text,
		},
	})
}

// 查询某个ID的系统消息
func (s *NotifyService) QuerySystemNotify(req *notifydemo.QuerySystemNotifyRequest) ([]*db.SystemNotify, error) {
	return db.QuerySystemNotify(config.NewConfig(s.ctx, db.DB), req.IDs)
}

// 查询所有系统消息的ID
func (s *NotifyService) QueryAllSystemNotify(req *notifydemo.QueryAllSystemNotifyRequest) ([]int32, error) {
	return db.QueryAllSystemNotify(config.NewConfig(s.ctx, db.DB), req.Limit, req.Offset)
}
