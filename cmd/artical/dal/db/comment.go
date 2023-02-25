package db

import (
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserName    string `gorm:"column:username"`
	ArticalID   uint   `gorm:"column:articalID"`
	CommentText string `gorm:"column:comment"`

	Master *uint      `gorm:"column:master"`
	Reply  []*Comment `gorm:"foreignkey:Master"`
}

func (c *Comment) TableName() string {
	return constants.CommentTableName
}

// 创建评论 并返回创建的评论的 ID
func CreateComment(cg *config.Config, cms []*Comment) ([]int32, error) {
	res := make([]int32, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Create(cms).Error; err != nil {
		return nil, errno.ServiceFault
	}
	for _, cm := range cms {
		res = append(res, int32(cm.ID))
	}
	return res, nil
}

// 根据 ID 查询评论
func QueryComment(cg *config.Config, id []int32) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Preload("Reply").Where("id in ?", id).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 ID 删除评论 并删除其所有 reply
func DeleteComment(cg *config.Config, id int32) error {
	cg.Tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(cg.Ctx).Where("master = ?", id).Delete(&Comment{}).Error; err != nil {
			return errno.ServiceFault
		}

		if err := tx.WithContext(cg.Ctx).Where("id = ?", id).Delete(&Comment{}).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
	return nil
}

// 暂时无用
// 根据 ID 更新评论
func UpdateComment(cg *config.Config, cm *Comment) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("id = ?", cm.ID).Updates(cm).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 ArticalID 查看 master
func QueryCommentByArticalID(cg *config.Config, articalID int32) ([]int32, error) {
	res := make([]int32, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Model(&Comment{}).Select("id").Where("articalID = ?", articalID).Where("master is null").Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}
