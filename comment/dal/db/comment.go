package db

import (
	"be/cmd/action/pack"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserName    string `gorm:"column:username"`
	TargetID    uint   `gorm:"column:targetID"`
	CommentText string `gorm:"column:comment"`
	Type        int32  `gorm:"column:type"`

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
		pack.EPrint(err.Error())
		return nil, errno.ServiceFault
	}
	for _, cm := range cms {
		res = append(res, int32(cm.ID))
	}
	return res, nil
}

// 根据 ID 查询评论
func QueryComment(cg *config.Config, ids []int32) ([]*Comment, error) {
	res := make([]*Comment, 0)
	for _, id := range ids {
		temp := &Comment{}
		if err := cg.Tx.WithContext(cg.Ctx).Preload("Reply").Where("id = ?", id).Find(temp).Error; err != nil {
			pack.EPrint(err.Error())
			return nil, errno.ServiceFault
		}
		if temp.ID == 0 {
			continue
		}
		res = append(res, temp)
	}
	return res, nil
}

// 根据 ID 删除评论 并删除其所有 reply
func DeleteComment(cg *config.Config, id int32) error {
	cg.Tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(cg.Ctx).Where("master = ?", id).Delete(&Comment{}).Error; err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}

		if err := tx.WithContext(cg.Ctx).Where("id = ?", id).Delete(&Comment{}).Error; err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}
		return nil
	})
	return nil
}

// 根据 TargetID 和 Type 查看 所有 master 评论
func QueryCommentByTargetID(cg *config.Config, targetID, tp int32) ([]int32, error) {
	res := make([]int32, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Model(&Comment{}).Select("id").Where("targetID = ?", targetID).Where("type = ?", tp).Where("master is null").Find(&res).Error; err != nil {
		pack.EPrint(err.Error())
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 暂时无用
// 根据 ID 更新评论
func UpdateComment(cg *config.Config, cm *Comment) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("id = ?", cm.ID).Updates(cm).Error; err != nil {
		pack.EPrint(err.Error())
		return errno.ServiceFault
	}
	return nil
}
