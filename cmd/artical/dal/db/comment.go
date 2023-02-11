package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID          uint      `gorm:"primarykey"`
	CreatedAt   time.Time `gorm:"column:createdAt"`
	UpdatedAt   time.Time `gorm:"column:updatedAt"`
	UserName    string    `gorm:"column:username"`
	ArticalID   uint      `gorm:"column:articalID"`
	CommentText string    `gorm:"column:comment"`

	Master *uint      `gorm:"column:master"`
	Reply  []*Comment `gorm:"foreignkey:Master"`
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
	if err := DB.WithContext(ctx).Preload("Reply").Where("id in ?", id).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 ID 删除评论 并删除其所有 reply
func DeleteComment(ctx context.Context, id int32) error {
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("master = ?", id).Delete(&Comment{}).Error; err != nil {
			return errno.ServiceFault
		}

		if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&Comment{}).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
	return nil
}

// 暂时无用
// 根据 ID 更新评论
func UpdateComment(ctx context.Context, cm *Comment) error {
	if err := DB.WithContext(ctx).Where("id = ?", cm.ID).Updates(cm).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 ArticalID 查看 master 评论 并按照时间降序返回
func QueryCommentByArticalID(ctx context.Context, articalID int32) ([]int32, error) {
	res := make([]int32, 0)
	if err := DB.WithContext(ctx).Model(&Comment{}).Select("id").Where("articalID = ?", articalID).Where("master is null").Order("updatedAt DESC").Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}
