package service

import (
	"be/cmd/artical/dal/db"
	"be/grpc/articaldemo"
	"be/pkg/errno"
)

func (s *ArticalService) CreateComment(req *articaldemo.CreateCommentRequest) error {
	return db.CreateComment(s.ctx, []*db.Comment{
		{
			UserName:    req.UserName,
			ArticalID:   uint(req.ArticalID),
			CommentText: req.CommentText,
		},
	})
}

func (s *ArticalService) QueryComment(req *articaldemo.QueryCommentRequest) ([]*db.Comment, error) {
	cm, err := db.QueryComment(s.ctx, req.CommentID)
	if err != nil {
		return nil, err
	}

	// 没有查询到该评论
	if len(cm) == 0 {
		return nil, errno.NoSuchCommentErr
	}

	return cm, nil
}

func (s *ArticalService) DeleteComment(req *articaldemo.DeleteCommentRequest) error {
	return db.DeleteComment(s.ctx, int32(req.CommentID))
}
