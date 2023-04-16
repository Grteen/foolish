package db

import (
	"be/cmd/action/pack"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"

	"gorm.io/gorm"
)

type Action struct {
	gorm.Model

	Text    string     `json:"text" gorm:"column:text; not null"`
	Author  string     `json:"author" gorm:"column:author; not null"`
	PicFile []*PicFile `json:"pictures" gorm:"foreignKey:ActionID"`

	LikeNum int32 `gorm:"column:likeNum; not null"`
}

func (act *Action) TableName() string {
	return constants.ActionTableName
}

type PicFile struct {
	File     string `gorm:"column:file"`
	ActionID int32  `gorm:"column:actionID"`
}

func (pic *PicFile) TableName() string {
	return constants.ActionPicFileTableName
}

type User struct{}

func (u *User) TableName() string {
	return constants.UserTableName
}

// 创建动态
func CreateAction(cg *config.Config, acts []*Action) error {
	return cg.Tx.Transaction(func(tx *gorm.DB) error {
		// 作者动态数增加
		if err := tx.WithContext(cg.Ctx).Model(&User{}).Where("username = ?", acts[0].Author).Update("actNum", gorm.Expr("actNum + ?", len(acts))).Error; err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}
		if err := tx.WithContext(cg.Ctx).Create(acts).Error; err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}
		return nil
	})
}

// 根据动态ID查询动态
func QueryAction(cg *config.Config, ids []int32) ([]*Action, error) {
	res := make([]*Action, 0)
	cg.Tx.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			act := &Action{}
			if err := tx.WithContext(cg.Ctx).Where("id = ?", id).Preload("PicFile").Find(act).Error; err != nil {
				pack.EPrint(err.Error())
				return errno.ServiceFault
			}
			// 未查询到该id
			if act.ID == 0 {
				continue
			}
			res = append(res, act)
		}
		return nil
	})

	return res, nil
}

// 删除动态
func DeleteAction(cg *config.Config, id int32) error {
	return cg.Tx.Transaction(func(tx *gorm.DB) error {
		// 查询动态作者
		var author string
		if err := tx.WithContext(cg.Ctx).Model(&Action{}).Select("author").Where("id = ?", id).Find(&author).Error; err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}
		// 作者动态数 - 1
		if err := tx.WithContext(cg.Ctx).Model(&User{}).Where("username = ?", author).Update("actNum", gorm.Expr("actNum - ?", 1)).Error; err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}

		if err := tx.WithContext(cg.Ctx).Where("id = ?", id).Delete(&Action{}).Error; err != nil {
			pack.EPrint(err.Error())
			return errno.ServiceFault
		}
		return nil
	})
}

// 根据 Author 查询动态
func QueryArticalByAuthor(cg *config.Config, author, field, order string) ([]int32, error) {
	res := make([]int32, 0)
	if err := cg.Tx.Model(&Action{}).WithContext(cg.Ctx).Select("id").Where("author = ?", author).Order(field + " " + order).Find(&res).Error; err != nil {
		pack.EPrint(err.Error())
		return nil, errno.ServiceFault
	}
	return res, nil
}
