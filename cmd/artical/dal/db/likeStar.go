package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
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

type Star struct {
	ID        uint   `gorm:"primarykey"`
	UserName  string `gorm:"column:username; not null"`
	ArticalID uint   `gorm:"column:articalID; not null"`
}

func (s *Star) TableName() string {
	return constants.StarTableName
}

// 点赞 （收藏）
func CreateLikeStar(ctx context.Context, likeStars []*LikeStar) error {
	if err := DB.WithContext(ctx).Model(ctx.Value(constants.LikeStarModel)).Create(likeStars).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 UserName 和 ArticalID 取消点赞  （收藏）
func DeleteLikeStar(ctx context.Context, likeStar *LikeStar) error {
	if err := DB.WithContext(ctx).Model(ctx.Value(constants.LikeStarModel)).Where("username = ?", likeStar.UserName).Where("articalID = ?", likeStar.ArticalID).Delete(&LikeStar{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 查询 Username 对于 ArticalID 的点赞 (收藏) （正常情况只有一个）
func QueryLikeStar(ctx context.Context, likeStar *LikeStar) ([]*LikeStar, error) {
	res := make([]*LikeStar, 0)
	if err := DB.WithContext(ctx).Model(ctx.Value(constants.LikeStarModel)).Where("username = ?", likeStar.UserName).Where("articalID = ?", likeStar.ArticalID).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 删除文章时可用
// 根据 ArticalID 批量删除点赞 （收藏）
func DeleteLikeStarByArticalID(ctx context.Context, articalID uint) error {
	if err := DB.WithContext(ctx).Model(ctx.Value(constants.LikeStarModel)).Where("articalID = ?", articalID).Delete(&Like{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}
