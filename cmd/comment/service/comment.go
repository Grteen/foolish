package service

import (
	"be/cmd/comment/dal/db"
	"be/grpc/commentdemo"
	"be/pkg/config"
	"be/pkg/errno"
	"context"

	"gorm.io/gorm"
)

type CommentService struct {
	ctx context.Context
}

func NewCommentService(ctx context.Context) *CommentService {
	return &CommentService{ctx: ctx}
}

// 创建评论或是 reply
// 如果 master 不为 0 则为 reply
func (s *CommentService) CreateComment(req *commentdemo.CreateCommentRequest) ([]int32, error) {
	if req.Master != 0 {
		m := uint(req.Master)
		return db.CreateComment(config.NewConfig(s.ctx, db.DB), []*db.Comment{
			{
				UserName:    req.UserName,
				TargetID:    uint(req.TargetID),
				CommentText: req.CommentText,

				Master: &m,
				Type:   req.Type,
			},
		})
	} else {
		return db.CreateComment(config.NewConfig(s.ctx, db.DB), []*db.Comment{
			{
				UserName:    req.UserName,
				TargetID:    uint(req.TargetID),
				CommentText: req.CommentText,

				Type: req.Type,
			},
		})
	}
}

// 根据 ID 查询评论
func (s *CommentService) QueryComment(req *commentdemo.QueryCommentRequest) ([]*db.Comment, error) {
	cm, err := db.QueryComment(config.NewConfig(s.ctx, db.DB), req.CommentID)
	if err != nil {
		return nil, err
	}

	// 没有查询到该评论
	if len(cm) == 0 {
		return nil, errno.NoSuchCommentErr
	}

	return cm, nil
}

// 查询一篇文章的所有 评论 ID 只会返回 没有 master 字段的 评论
func (s *CommentService) QueryCommentByTargetID(req *commentdemo.QueryCommentByTargetIDRequest) ([]int32, error) {
	return db.QueryCommentByTargetID(config.NewConfig(s.ctx, db.DB), req.TargetID, req.Type)
}

// 根据 ID 删除评论及其所有回复
func (s *CommentService) DeleteComment(req *commentdemo.DeleteCommentRequest) error {
	return db.DeleteComment(config.NewConfig(s.ctx, db.DB), req.CommentID)
}

// 暂时无用
func (s *CommentService) UpdateComment(req *commentdemo.UpdateCommentRequest) error {
	return db.UpdateComment(config.NewConfig(s.ctx, db.DB), &db.Comment{
		Model: gorm.Model{
			ID: uint(req.CommentID),
		},
		CommentText: req.CommentText,
	})
}
