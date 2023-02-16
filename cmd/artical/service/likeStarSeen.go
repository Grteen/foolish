package service

import (
	"be/cmd/artical/dal/db"
	"be/cmd/artical/dal/rdb"
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

func (s *ArticalService) QueryAllLikeStar(req *articaldemo.QueryAllLikeStarRequest) ([]int32, error) {
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

func (s *ArticalService) RdbIncreaseitf(req *articaldemo.RdbIncreaseitfRequest) error {
	return rdb.IncreaseLikeStar(s.ctx, req.ArticalID, req.Val, req.Field)
}

// 创建收藏
func (s *ArticalService) CreateStar(req *articaldemo.CreateStarRequest) error {
	return db.CreateStar(s.ctx, []*db.Star{
		{
			UserName:  req.Username,
			ArticalID: uint(req.ArticalID),
			FolderID:  uint(req.StarFolderID),
		},
	})
}

// 创建收藏夹
func (s *ArticalService) CreateStarFolder(req *articaldemo.CreateStarFolderRequest) error {
	return db.CreateStarFolder(s.ctx, []*db.StarFolder{
		{
			UserName:   req.UserName,
			FolderName: req.FolderName,
			IsDefault:  req.IsDefault,
		},
	})
}

// 查询收藏夹
func (s *ArticalService) QueryStarFolder(req *articaldemo.QueryStarFolderRequest) ([]*db.StarFolder, error) {
	return db.QueryStarFolder(s.ctx, req.IDs)
}

// 查询所有的收藏夹
func (s *ArticalService) QueryAllStarFolder(req *articaldemo.QueryAllStarFolderRequest) ([]*db.StarFolder, error) {
	return db.QueryAllStarFolder(s.ctx, req.UserName)
}

// 查询收藏夹的所有收藏
func (s *ArticalService) QueryAllStar(req *articaldemo.QueryAllStarRequest) ([]*db.Star, error) {
	return db.QueryAllStar(s.ctx, req.StarFolderID, req.Limit, req.Offset)
}

// 删除收藏夹
func (s *ArticalService) DeleteStarFolder(req *articaldemo.DeleteStarFolderRequest) error {
	return db.DeleteStarFolder(s.ctx, req.ID)
}

// 更新收藏夹
func (s *ArticalService) UpdateStarFolder(req *articaldemo.UpdateStarFolderRequest) error {
	return db.UpdateStarFolder(s.ctx, req.StarFolder.ID, req.StarFolder.FolderName)
}
