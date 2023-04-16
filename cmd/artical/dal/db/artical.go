package db

import (
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"

	"gorm.io/gorm"
)

type Artical struct {
	gorm.Model

	Title       string `json:"title" gorm:"column:title; not null"`
	Author      string `json:"author" gorm:"column:author; not null"`
	Text        string `json:"text" gorm:"column:text; type:text"`
	Description string `json:"description" gorm:"column:description; not null"`
	Cover       string `json:"cover" gorm:"column:cover; not null"`

	LikeNum int32 `gorm:"column:likeNum; not null"`
	StarNum int32 `gorm:"column:starNum; not null"`
	SeenNum int32 `gorm:"column:seenNum; not null"`

	Liked  []*Like `gorm:"foreignKey:ArticalID"`
	Stared []*Star `gorm:"foreignKey:ArticalID"`
}

func (art *Artical) TableName() string {
	return constants.ArticalTableName
}

type User struct{}

func (u *User) TableName() string {
	return constants.UserTableName
}

type Comment struct{}

func (cm *Comment) TableName() string {
	return constants.CommentTableName
}

// 创建文章
func CreateArtical(cg *config.Config, arts []*Artical) error {
	return cg.Tx.Transaction(func(tx *gorm.DB) error {
		// 作者文章数增加
		if err := tx.WithContext(cg.Ctx).Model(&User{}).Where("username = ?", arts[0].Author).Update("artNum", gorm.Expr("artNum + ?", len(arts))).Error; err != nil {
			return errno.ServiceFault
		}
		if err := tx.WithContext(cg.Ctx).Create(arts).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 根据 ID 查询文章
func QueryArtical(cg *config.Config, ids []int32) ([]*Artical, error) {
	res := make([]*Artical, 0)
	cg.Tx.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(cg.Ctx).Where("id in ?", ids).Find(&res).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})

	return res, nil
}

// 根据 ID 删除文章 及其所有评论
func DeleteArtical(cg *config.Config, id int32) error {
	return cg.Tx.Transaction(func(tx *gorm.DB) error {
		res := make([]int32, 0)
		// 查询该文章的所有 master 评论 并删除
		if err := tx.WithContext(cg.Ctx).Model(&Comment{}).Select("id").Where("articalID = ?", id).Where("master is null").Order("updatedAt DESC").Find(&res).Error; err != nil {
			return err
		}
		for _, cm := range res {
			if err := tx.WithContext(cg.Ctx).Where("master = ?", cm).Delete(&Comment{}).Error; err != nil {
				return err
			}
			if err := tx.WithContext(cg.Ctx).Where("id = ?", cm).Delete(&Comment{}).Error; err != nil {
				return err
			}
		}

		// 查询文章作者
		var author string
		if err := tx.WithContext(cg.Ctx).Model(&Artical{}).Select("author").Where("id = ?", id).Find(&author).Error; err != nil {
			return err
		}
		// 作者文章数 - 1
		if err := tx.WithContext(cg.Ctx).Model(&User{}).Where("username = ?", author).Update("artNum", gorm.Expr("artNum - ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.WithContext(cg.Ctx).Where("id = ?", id).Delete(&Artical{}).Error; err != nil {
			return err
		}
		return nil
	})
}

// 根据 ID 更新文章为 art
func UpdateArtical(cg *config.Config, art *Artical) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("id = ?", art.ID).Updates(*art).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 Author 查询文章
func QueryArticalByAuthor(cg *config.Config, author, field, order string) ([]int32, error) {
	res := make([]int32, 0)
	if err := cg.Tx.Model(&Artical{}).WithContext(cg.Ctx).Select("id").Where("author = ?", author).Order(field + " " + order).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// meaningless
// 根据 Title 查询文章 返回所有 Title 相同的文章
func QueryArticalByTitle(cg *config.Config, title string) ([]*Artical, error) {
	res := make([]*Artical, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Where("title = ?", title).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}
