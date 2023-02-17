package service

import (
	"be/cmd/artical/dal/db"
	"be/grpc/articaldemo"
	"be/pkg/config"
	"be/pkg/errno"
)

// 创建评论或是 reply
// 如果 master 不为 0 则为 reply
func (s *ArticalService) CreateComment(req *articaldemo.CreateCommentRequest) ([]int32, error) {
	if req.Master != 0 {
		m := uint(req.Master)
		return db.CreateComment(config.NewConfig(s.ctx, db.DB), []*db.Comment{
			{
				UserName:    req.UserName,
				ArticalID:   uint(req.ArticalID),
				CommentText: req.CommentText,

				Master: &m,
			},
		})
	} else {
		return db.CreateComment(config.NewConfig(s.ctx, db.DB), []*db.Comment{
			{
				UserName:    req.UserName,
				ArticalID:   uint(req.ArticalID),
				CommentText: req.CommentText,
			},
		})
	}
}

// 根据 ID 查询评论
func (s *ArticalService) QueryComment(req *articaldemo.QueryCommentRequest) ([]*db.Comment, error) {
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
func (s *ArticalService) QueryCommentByArticalID(req *articaldemo.QueryCommentByArticalIDRequest) ([]int32, error) {
	return db.QueryCommentByArticalID(config.NewConfig(s.ctx, db.DB), req.ArticalID)
}

// 暂时无用
func (s *ArticalService) UpdateComment(req *articaldemo.UpdateCommentRequest) error {
	return db.UpdateComment(config.NewConfig(s.ctx, db.DB), &db.Comment{
		ID:          uint(req.CommentID),
		CommentText: req.CommentText,
	})
}

// 根据 ID 删除评论及其所有回复
func (s *ArticalService) DeleteComment(req *articaldemo.DeleteCommentRequest) error {
	return db.DeleteComment(config.NewConfig(s.ctx, db.DB), req.CommentID)
}
