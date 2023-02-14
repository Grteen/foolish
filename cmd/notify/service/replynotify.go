package service

import (
	"be/cmd/notify/dal/db"
	"be/grpc/notifydemo"
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
	return db.CreateReplyNotify(s.ctx, []*db.ReplyNotify{
		{
			Notify: db.Notify{
				UserName: req.Replyntf.UserName,
				Title:    req.Replyntf.Title,
				Sender:   req.Replyntf.Sender,
				Text:     req.Replyntf.Text,
				IsRead:   false,
			},
		},
	})
}
