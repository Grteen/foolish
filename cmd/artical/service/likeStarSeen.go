package service

import (
	"be/cmd/artical/dal/db"
	"be/grpc/articaldemo"
	"be/pkg/errno"
)

func (s *ArticalService) CreateLikeStar(req *articaldemo.CreateLikeStarRequest) error {
	var itf db.LikeStarInterface
	if req.Type == 0 {
		// Like
		itf = &db.Like{}
	} else if req.Type == 1 {
		// Star
		itf = &db.Star{}
	} else if req.Type == 2 {
		// Seen
		itf = &db.Seen{}
	} else {
		return errno.ServiceFault
	}
	return db.CreateLikeStar(s.ctx, []*db.LikeStar{
		{
			UserName:  req.UserName,
			ArticalID: uint(req.ArticalID),
		},
	}, itf)
}

func (s *ArticalService) DeleteLikeStar(req *articaldemo.DeleteLikeStarRequest) error {
	var itf db.LikeStarInterface
	if req.Type == 0 {
		// Like
		itf = &db.Like{}
	} else if req.Type == 1 {
		// Star
		itf = &db.Star{}
	} else if req.Type == 2 {
		// Seen
		itf = &db.Seen{}
	} else {
		return errno.ServiceFault
	}
	return db.DeleteLikeStar(s.ctx, &db.LikeStar{
		UserName:  req.UserName,
		ArticalID: uint(req.ArticalID),
	}, itf)
}

func (s *ArticalService) UpdateLikeStarTime(req *articaldemo.UpdateLikeStarTimeRequest) error {
	var itf db.LikeStarInterface
	if req.Type == 0 {
		// Like
		itf = &db.Like{}
	} else if req.Type == 1 {
		// Star
		itf = &db.Star{}
	} else if req.Type == 2 {
		// Seen
		itf = &db.Seen{}
	} else {
		return errno.ServiceFault
	}

	return db.UpdateLikeStarTime(s.ctx, &db.LikeStar{
		UserName:  req.Likestar.UserName,
		ArticalID: uint(req.Likestar.ArticalID),
	}, req.UpdateTime.AsTime(), itf)
}

func (s *ArticalService) QueryLikeStar(req *articaldemo.QueryLikeStarRequest) ([]*db.LikeStar, error) {
	var itf db.LikeStarInterface
	if req.Type == 0 {
		// Like
		itf = &db.Like{}
	} else if req.Type == 1 {
		// Star
		itf = &db.Star{}
	} else if req.Type == 2 {
		// Seen
		itf = &db.Seen{}
	} else {
		return nil, errno.ServiceFault
	}
	res, err := db.QueryLikeStar(s.ctx, &db.LikeStar{
		UserName:  req.UserName,
		ArticalID: uint(req.ArticalID),
	}, itf)

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
	var itf db.LikeStarInterface
	if req.Type == 0 {
		// Like
		itf = &db.Like{}
	} else if req.Type == 1 {
		// Star
		itf = &db.Star{}
	} else if req.Type == 2 {
		// Seen
		itf = &db.Seen{}
	} else {
		return nil, errno.ServiceFault
	}
	res, err := db.QueryAllLikeStar(s.ctx, req.UserName, itf)
	if err != nil {
		return nil, err
	}

	return res, nil
}
