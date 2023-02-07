package service

import (
	"be/cmd/artical/dal/db"
	"be/grpc/articaldemo"
	"be/pkg/errno"
)

func (s *ArticalService) CreateLikeStar(req *articaldemo.CreateLikeStarRequest) error {
	return db.CreateLikeStar(s.ctx, []*db.LikeStar{
		{
			UserName:  req.UserName,
			ArticalID: uint(req.ArticalID),
		},
	})
}

func (s *ArticalService) DeleteLikeStar(req *articaldemo.DeleteLikeStarRequest) error {
	return db.DeleteLikeStar(s.ctx, &db.LikeStar{
		UserName:  req.UserName,
		ArticalID: uint(req.ArticalID),
	})
}

func (s *ArticalService) QueryLikeStar(req *articaldemo.QueryLikeStarRequest) ([]*db.LikeStar, error) {
	res, err := db.QueryLikeStar(s.ctx, &db.LikeStar{
		UserName:  req.UserName,
		ArticalID: uint(req.ArticalID),
	})

	if err != nil {
		return nil, err
	}

	// 没有查询到
	if len(res) == 0 {
		return nil, errno.NoLikeStarErr
	}

	return res, nil
}

func (s *ArticalService) QueryAllLikeStar(req *articaldemo.QueryAllLikeStarRequest) ([]uint32, error) {
	res, err := db.QueryAllLikeStar(s.ctx, req.UserName)
	if err != nil {
		return nil, err
	}

	return res, nil
}
