package service

import (
	"be/cmd/notify/dal/db"
	"be/grpc/notifydemo"
	"be/pkg/config"
	"context"
)

type NotifyService struct {
	ctx context.Context
}

func NewNotifyService(ctx context.Context) *NotifyService {
	return &NotifyService{ctx: ctx}
}

// 创建回复通知
func (s *NotifyService) CreateReplyNotify(req *notifydemo.CreateReplyNotifyRequest) error {
	return db.CreateReplyNotify(config.NewConfig(s.ctx, db.DB), []*db.ReplyNotify{
		{
			Notify: db.Notify{
				UserName: req.Replyntf.UserName,
				Title:    req.Replyntf.Title,
				Sender:   req.Replyntf.Sender,
				Text:     req.Replyntf.Text,

				IsRead: false,
			},
			ArticalID: req.Replyntf.ArticalID,
			CommentID: req.Replyntf.CommentID,
		},
	})
}

// 根据 ID 查询回复消息
func (s *NotifyService) QueryReplyNotify(req *notifydemo.QueryReplyNotifyRequest) ([]*db.ReplyNotify, error) {
	return db.QueryReplyNotify(config.NewConfig(s.ctx, db.DB), req.IDs)
}

// 查询某人的 回复消息id
func (s *NotifyService) QueryAllReplyNotify(req *notifydemo.QueryAllReplyNotifyRequest) ([]int32, error) {
	return db.QueryAllReplyNotify(config.NewConfig(s.ctx, db.DB), req.UserName, req.Limit, req.Offset)
}

// 根据ID 更新回复通知为已阅读
func (s *NotifyService) ReadReplyNotify(req *notifydemo.ReadReplyNotifyRequest) error {
	return db.UpdateReplyNotify(config.NewConfig(s.ctx, db.DB), req.ID)
}
