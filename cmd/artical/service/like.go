package service

import (
	"be/cmd/artical/dal/db"
	"be/grpc/articaldemo"
	"be/pkg/errno"
)

func (s *ArticalService) CreateLike(req *articaldemo.CreateLikeRequest) error {
	return db.CreateLike(s.ctx, []*db.Like{
		{
			UserName:  req.UserName,
			ArticalID: uint(req.ArticalID),
		},
	})
}

func (s *ArticalService) DeleteLike(req *articaldemo.DeleteLikeRequest) error {
	return db.DeleteLike(s.ctx, &db.Like{
		UserName:  req.UserName,
		ArticalID: uint(req.ArticalID),
	})
}

func (s *ArticalService) QueryLike(req *articaldemo.QueryLikeRequest) ([]*db.Like, error) {
	res, err := db.QueryLike(s.ctx, &db.Like{
		UserName:  req.UserName,
		ArticalID: uint(req.ArticalID),
	})

	if err != nil {
		return nil, err
	}

	// 没有查询到
	if len(res) == 0 {
		return nil, errno.NoLikesErr
	}

	return res, nil
}
