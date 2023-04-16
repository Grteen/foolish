package db

import (
	"be/cmd/action/pack"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"

	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	UserName string `gorm:"column:username; not null"`
	ActionID int32  `gorm:"column:actionID; not null"`
}

func (l *Like) TableName() string {
	return constants.ActionLikeTableName
}

// 点赞
func CreateActionLike(cg *config.Config, likes []*Like) error {
	return cg.Tx.Transaction(func(tx *gorm.DB) error {
		for _, l := range likes {
			err := tx.Model(&Action{}).Where("id = ?", l.ActionID).Update("likeNum", gorm.Expr("likeNum + ?", 1)).Error
			if err != nil {
				pack.EPrint(err.Error())
				return errno.ServiceFault
			}
		}
		if err := tx.WithContext(cg.Ctx).Create(likes).Error; err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}
		return nil
	})
}

// 取消点赞
func DeleteActionLike(cg *config.Config, like *Like) error {
	return cg.Tx.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&Action{}).Where("id = ?", like.ActionID).Update("likeNum", gorm.Expr("likeNum + ?", -1)).Error
		if err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}
		if err := tx.WithContext(cg.Ctx).Where("username = ?", like.UserName).Where("actionID = ?", like.ActionID).Delete(&Like{}).Error; err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}
		return nil
	})
}

// 查询 username 对于 actionID 的点赞
func QueryActionLike(cg *config.Config, like *Like) ([]*Like, error) {
	res := make([]*Like, 0)
	err := cg.Tx.Transaction(func(tx *gorm.DB) error {
		if err := cg.Tx.WithContext(cg.Ctx).Where("username = ?", like.UserName).Where("actionID = ?", like.ActionID).Find(&res).Error; err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}
		return nil
	})
	return res, err
}
