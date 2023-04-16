package service

import (
	"be/cmd/comment/dal/db"
	"be/cmd/comment/dal/other"
	"be/cmd/comment/dal/rdb"
	"be/grpc/commentdemo"
	"be/pkg/config"
	"be/pkg/errno"
	"context"
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

// 根据 ID 查询评论 已缓存
func (s *CommentService) QueryComment(req *commentdemo.QueryCommentRequest) ([]*db.Comment, error) {
	res := make([]*db.Comment, 0)
	// 查询缓存
	rcms, ungot, err := rdb.GetRdbComment(s.ctx, req.CommentID)
	if err != nil {
		return nil, errno.ServiceFault
	}
	if len(ungot) != 0 {
		cm, err := db.QueryComment(config.NewConfig(s.ctx, db.DB), ungot)
		if err != nil {
			return nil, err
		}
		// 缓存数据
		rcm, err := other.ChangeCommentToRdb(cm)
		if err != nil {
			rdb.EPrint(err.Error())
			return nil, errno.ServiceFault
		}
		if err := rdb.SetRdbComment(s.ctx, rcm); err != nil {
			return nil, errno.ServiceFault
		}

		cms, err := other.ChangeRdbToComment(rcms)
		if err != nil {
			rdb.EPrint(err.Error())
			return nil, errno.ServiceFault
		}

		res = append(res, cms...)
		res = append(res, cm...)
		return res, nil
	}

	cms, err := other.ChangeRdbToComment(rcms)
	if err != nil {
		rdb.EPrint(err.Error())
		return nil, errno.ServiceFault
	}
	res = append(res, cms...)

	return res, nil
}

// 查询一篇文章的所有 评论 ID 只会返回 没有 master 字段的 评论
func (s *CommentService) QueryCommentByTargetID(req *commentdemo.QueryCommentByTargetIDRequest) ([]int32, error) {
	return db.QueryCommentByTargetID(config.NewConfig(s.ctx, db.DB), req.TargetID, req.Type)
}

// 根据 ID 删除评论及其所有回复
func (s *CommentService) DeleteComment(req *commentdemo.DeleteCommentRequest) error {
	if err := db.DeleteComment(config.NewConfig(s.ctx, db.DB), req.CommentID); err != nil {
		return err
	}
	return rdb.DelRdbComment(s.ctx, req.CommentID)
}

// // 暂时无用
// func (s *CommentService) UpdateComment(req *commentdemo.UpdateCommentRequest) error {
// 	return db.UpdateComment(config.NewConfig(s.ctx, db.DB), &db.Comment{
// 		Model: gorm.Model{
// 			ID: uint(req.CommentID),
// 		},
// 		CommentText: req.CommentText,
// 	})
// }
