package service

import (
	"be/cmd/artical/dal/db"
	"be/cmd/artical/dal/rdb"
	"be/grpc/articaldemo"
	"be/pkg/config"
	"be/pkg/errno"

	"gorm.io/gorm"
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
	return db.CreateLikeStar(config.NewConfig(s.ctx, db.DB), []*db.LikeStar{
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
	return db.DeleteLikeStar(config.NewConfig(s.ctx, db.DB), &db.LikeStar{
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

	return db.UpdateLikeStarTime(config.NewConfig(s.ctx, db.DB), &db.LikeStar{
		UserName:  req.Likestar.UserName,
		ArticalID: uint(req.Likestar.ArticalID),
	}, req.UpdateTime.AsTime(), itf)
}

// 查询点赞收藏 如果不存在则返回NoLikeStarErr
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
	res, err := db.QueryLikeStar(config.NewConfig(s.ctx, db.DB), &db.LikeStar{
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
	res, err := db.QueryAllLikeStar(config.NewConfig(s.ctx, db.DB), req.UserName, itf)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 创建收藏
func (s *ArticalService) CreateStar(req *articaldemo.CreateStarRequest) error {
	return db.CreateStar(config.NewConfig(s.ctx, db.DB), []*db.Star{
		{
			UserName:  req.Username,
			ArticalID: uint(req.ArticalID),
			FolderID:  uint(req.StarFolderID),
		},
	})
}

// 创建收藏夹
func (s *ArticalService) CreateStarFolder(req *articaldemo.CreateStarFolderRequest) error {
	return db.CreateStarFolder(config.NewConfig(s.ctx, db.DB), []*db.StarFolder{
		{
			UserName:   req.UserName,
			FolderName: req.FolderName,
			IsDefault:  req.IsDefault,
			Public:     req.Public,
		},
	})
}

// 查询收藏夹
func (s *ArticalService) QueryStarFolder(req *articaldemo.QueryStarFolderRequest) ([]*db.StarFolder, error) {
	return db.QueryStarFolder(config.NewConfig(s.ctx, db.DB), req.IDs)
}

// 查询所有的收藏夹
func (s *ArticalService) QueryAllStarFolder(req *articaldemo.QueryAllStarFolderRequest) ([]*db.StarFolder, error) {
	return db.QueryAllStarFolder(config.NewConfig(s.ctx, db.DB), req.UserName)
}

// 查询收藏夹的所有收藏
func (s *ArticalService) QueryAllStar(req *articaldemo.QueryAllStarRequest) ([]*db.Star, error) {
	return db.QueryAllStar(config.NewConfig(s.ctx, db.DB), req.StarFolderID, req.Limit, req.Offset)
}

// 删除收藏夹
func (s *ArticalService) DeleteStarFolder(req *articaldemo.DeleteStarFolderRequest) error {
	return db.DeleteStarFolder(config.NewConfig(s.ctx, db.DB), req.ID)
}

// 更新收藏夹
func (s *ArticalService) UpdateStarFolder(req *articaldemo.UpdateStarFolderRequest) error {
	return db.UpdateStarFolder(config.NewConfig(s.ctx, db.DB), req.StarFolder.ID, req.StarFolder.FolderName, req.StarFolder.Public)
}

// 删除收藏夹并将其中的收藏移动到默认收藏夹
func (s *ArticalService) DeleteStarFolderAndMove(req *articaldemo.DeleteStarFolderAndMoveRequest) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		// 查询原收藏夹的所有收藏
		stars, err := db.QueryAllStar(config.NewConfig(s.ctx, tx), req.StarFolderID, 0, 0)
		if err != nil {
			return errno.ServiceFault
		}
		// 查询某人收藏夹
		id, err := db.QueryDefaultFolder(config.NewConfig(s.ctx, tx), req.Username)
		if err != nil {
			return errno.ServiceFault
		}
		// 移动收藏
		for _, star := range stars {
			if err := db.UpdateStarOwner(config.NewConfig(s.ctx, tx), int32(star.ID), id); err != nil {
				return errno.ServiceFault
			}
		}
		// 删除收藏夹
		if err := db.DeleteStarFolder(config.NewConfig(s.ctx, tx), req.StarFolderID); err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 更改收藏所属的收藏夹
func (s *ArticalService) UpdateStarOwner(req *articaldemo.UpdateStarOwnerRequest) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {
		// 查询该收藏
		star, err := db.QueryLikeStar(config.NewConfig(s.ctx, tx), &db.LikeStar{
			UserName:  req.Username,
			ArticalID: uint(req.ArticalID),
		}, &db.Star{})
		if err != nil {
			return errno.ServiceFault
		}
		if len(star) == 0 {
			return errno.NoStarErr
		}

		err = db.UpdateStarOwner(config.NewConfig(s.ctx, tx), int32(star[0].ID), req.OwnerID)
		if err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

func (s *ArticalService) RdbIncreaseitf(req *articaldemo.RdbIncreaseitfRequest) error {
	return rdb.IncreaseLikeStar(s.ctx, req.ArticalID, req.Val, req.Field)
}

// 设置点赞收藏缓存
func (s *ArticalService) RdbSetLikeStar(req *articaldemo.RdbSetLikeStarRequest) error {
	return rdb.SetLikeStar(s.ctx, req.UserName, req.ArticalID, req.Type, req.UpdatedAt)
}

// 获取点赞收藏缓存
func (s *ArticalService) RdbGetLikeStar(req *articaldemo.RdbGetLikeStarRequest) (bool, string, error) {
	return rdb.GetLikeStar(s.ctx, req.UserName, req.ArticalID, req.Type)
}

// 删除点赞收藏缓存
func (s *ArticalService) RdbDelLikeStar(req *articaldemo.RdbDelLikeStarRequest) error {
	return rdb.DelLikeStar(s.ctx, req.UserName, req.ArticalID, req.Type)
}
