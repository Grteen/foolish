package service

import (
	"be/cmd/notify/dal/db"
	"be/grpc/notifydemo"
	"be/pkg/config"
	"be/pkg/errno"
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

				IsRead:   false,
				IsDelete: false,
			},
			TargetID:  req.Replyntf.Target.TargetID,
			CommentID: req.Replyntf.CommentID,
			Master:    req.Replyntf.Master,
		},
	})
}

// 根据 ID 查询回复通知
func (s *NotifyService) QueryReplyNotify(req *notifydemo.QueryReplyNotifyRequest) ([]*db.ReplyNotify, error) {
	return db.QueryReplyNotify(config.NewConfig(s.ctx, db.DB), req.IDs)
}

// 查询某人的 回复通知id
func (s *NotifyService) QueryAllReplyNotify(req *notifydemo.QueryAllReplyNotifyRequest) ([]int32, error) {
	return db.QueryAllReplyNotify(config.NewConfig(s.ctx, db.DB), req.UserName, req.Limit, req.Offset)
}

// 创建点赞通知
func (s *NotifyService) CreateLikeNotify(req *notifydemo.CreateLikeNotifyRequest) error {
	return db.CreateLikeNotify(config.NewConfig(s.ctx, db.DB), []*db.LikeNotify{
		{
			Notify: db.Notify{
				UserName: req.Likentf.UserName,
				Title:    req.Likentf.Title,
				Sender:   req.Likentf.Sender,
				Text:     req.Likentf.Text,

				IsRead:   false,
				IsDelete: false,
			},
			TargetID: req.Likentf.Target.TargetID,
			Type:     req.Likentf.Target.Type,
		},
	})
}

// 根据 ID 查询点赞通知
func (s *NotifyService) QueryLikeNotify(req *notifydemo.QueryLikeNotifyRequest) ([]*db.LikeNotify, error) {
	return db.QueryLikeNotify(config.NewConfig(s.ctx, db.DB), req.IDs)
}

// 查询某人的 点赞通知id
func (s *NotifyService) QueryAllLikeNotify(req *notifydemo.QueryAllLikeNotifyRequest) ([]int32, error) {
	return db.QueryAllLikeNotify(config.NewConfig(s.ctx, db.DB), req.UserName, req.Limit, req.Offset)
}

// 根据ID 更新通知为已阅读
// type = 0 为 replynotify
// type = 1 为 likenotify
func (s *NotifyService) ReadNotify(req *notifydemo.ReadNotifyRequest) error {
	if req.Type == 0 {
		return db.UpdateReplyNotify(config.NewConfig(s.ctx, db.DB), req.ID)
	} else if req.Type == 1 {
		return db.UpdateLikeNotify(config.NewConfig(s.ctx, db.DB), req.ID)
	} else {
		return errno.ServiceFault
	}
}

// 根据ID 更新通知为已删除
// type = 0 为 replynotify
// type = 1 为 likenotify
func (s *NotifyService) DeleteNotify(req *notifydemo.DeleteNotifyRequest) error {
	if req.Type == 0 {
		return db.DeleteReplyNotify(config.NewConfig(s.ctx, db.DB), req.ID)
	} else if req.Type == 1 {
		return db.DeleteLikeNotify(config.NewConfig(s.ctx, db.DB), req.ID)
	} else {
		return errno.ServiceFault
	}
}

// 查询所有通知的id 并按照时间降序返回
func (s *NotifyService) SearchAllNotify(req *notifydemo.SearchAllNotifyRequest) ([]*db.AllNotify, error) {
	ntfs, err := db.SearchAllNotify(config.NewConfig(s.ctx, db.DB), req.UserName, req.Limit, req.Offset)
	if err != nil {
		return nil, errno.ServiceFault
	}
	return ntfs, nil
}
