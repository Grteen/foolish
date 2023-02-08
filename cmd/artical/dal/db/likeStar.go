package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"

	"gorm.io/gorm"
)

type LikeStar struct {
	ID        uint   `gorm:"primarykey"`
	UserName  string `gorm:"column:username; not null"`
	ArticalID uint   `gorm:"column:articalID; not null"`
}

type Like struct {
	ID        uint   `gorm:"primarykey"`
	UserName  string `gorm:"column:username; not null"`
	ArticalID uint   `gorm:"column:articalID; not null"`
}

func (l *Like) TableName() string {
	return constants.LikeTableName
}

func (l *Like) ColumnForArtical() string {
	return "likeNum"
}

type Star struct {
	ID        uint   `gorm:"primarykey"`
	UserName  string `gorm:"column:username; not null"`
	ArticalID uint   `gorm:"column:articalID; not null"`
}

func (s *Star) TableName() string {
	return constants.StarTableName
}

func (s *Star) ColumnForArtical() string {
	return "StarNum"
}

type LikeStarInterface interface {
	ColumnForArtical() string
}

// 点赞 （收藏）
func CreateLikeStar(ctx context.Context, likeStars []*LikeStar, itf LikeStarInterface) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		for _, ls := range likeStars {
			err := DB.Model(&Artical{}).Where("ID = ?", ls.ArticalID).Update(itf.ColumnForArtical(), gorm.Expr(itf.ColumnForArtical()+" + ?", 1)).Error
			if err != nil {
				return errno.ServiceFault
			}
		}
		if err := DB.WithContext(ctx).Model(itf).Create(likeStars).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 根据 UserName 和 ArticalID 取消点赞  （收藏）
func DeleteLikeStar(ctx context.Context, likeStar *LikeStar, itf LikeStarInterface) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		err := DB.Model(&Artical{}).Where("ID = ?", likeStar.ArticalID).Update(itf.ColumnForArtical(), gorm.Expr(itf.ColumnForArtical()+" - ?", 1)).Error
		if err != nil {
			return errno.ServiceFault
		}
		if err := DB.WithContext(ctx).Model(itf).Where("username = ?", likeStar.UserName).Where("articalID = ?", likeStar.ArticalID).Delete(itf).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 查询 Username 对于 ArticalID 的点赞 (收藏) （正常情况只有一个）
func QueryLikeStar(ctx context.Context, likeStar *LikeStar, itf LikeStarInterface) ([]*LikeStar, error) {
	res := make([]*LikeStar, 0)
	if err := DB.WithContext(ctx).Model(itf).Where("username = ?", likeStar.UserName).Where("articalID = ?", likeStar.ArticalID).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 查询 UserName 所有的 收藏 (点赞) 返回文章ID
func QueryAllLikeStar(ctx context.Context, userName string) ([]uint32, error) {
	res := make([]uint32, 0)
	if err := DB.WithContext(ctx).Model(ctx.Value(constants.LikeStarModel)).Select("ArticalID").Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 删除文章时可用
// 根据 ArticalID 批量删除点赞 （收藏）
func DeleteLikeStarByArticalID(ctx context.Context, articalID int32) error {
	if err := DB.WithContext(ctx).Model(ctx.Value(constants.LikeStarModel)).Where("articalID = ?", articalID).Delete(&Like{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}
