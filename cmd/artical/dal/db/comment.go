package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"time"
)

type Comment struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserName    string `gorm:"column:username"`
	ArticalID   uint   `gorm:"column:articalID"`
	CommentText string `gorm:"column:comment"`
}

func (c *Comment) TableName() string {
	return constants.CommentTableName
}

// 创建评论
func CreateComment(ctx context.Context, cms []*Comment) error {
	if err := DB.WithContext(ctx).Create(cms).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 ID 查询评论
func QueryComment(ctx context.Context, id []int32) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Where("id in ?", id).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 ID 删除评论
func DeleteComment(ctx context.Context, id int32) error {
	if err := DB.WithContext(ctx).Where("id = ?", id).Delete(&Comment{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 ID 更新评论
func UpdateComment(ctx context.Context, cm *Comment) error {
	if err := DB.WithContext(ctx).Where("id = ?", cm.ID).Updates(cm).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 ArticalID 查看评论
func QueryCommentByArticalID(ctx context.Context, articalID int32) ([]int32, error) {
	res := make([]int32, 0)
	if err := DB.WithContext(ctx).Model(&Comment{}).Select("id").Where("articalID = ?", articalID).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}
