package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
)

type Like struct {
	ID        uint   `gorm:"primarykey"`
	UserName  string `gorm:"column:username; not null"`
	ArticalID uint   `gorm:"column:articalID; not null"`
}

func (l *Like) TableName() string {
	return constants.LikeTableName
}

// 点赞
func CreateLike(ctx context.Context, likes []*Like) error {
	if err := DB.WithContext(ctx).Create(likes).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 UserName 和 ArticalID 取消点赞
func DeleteLike(ctx context.Context, like *Like) error {
	if err := DB.WithContext(ctx).Where("username = ?", like.UserName).Where("articalID = ?", like.ArticalID).Delete(&Like{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 查询 Username 对于 ArticalID 的点赞 （正常情况只有一个）
func QueryLike(ctx context.Context, like *Like) ([]*Like, error) {
	res := make([]*Like, 0)
	if err := DB.WithContext(ctx).Where("username = ?", like.UserName).Where("articalID = ?", like.ArticalID).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 删除文章时可用
// 根据 ArticalID 批量删除点赞
func DeleteLikeByArticalID(ctx context.Context, articalID uint) error {
	if err := DB.WithContext(ctx).Where("articalID = ?", articalID).Delete(&Like{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}
