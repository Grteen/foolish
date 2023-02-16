package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"time"

	"gorm.io/gorm"
)

type Artical struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt; not null"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updatedAt; not null"`

	Title       string `json:"title" gorm:"column:title; not null"`
	Author      string `json:"author" gorm:"column:author; not null"`
	Text        string `json:"text" gorm:"column:text; type:text"`
	Description string `json:"description" gorm:"column:description; not null"`
	Cover       string `json:"cover" gorm:"column:cover; not null"`

	LikeNum int32 `gorm:"column:likeNum; not null"`
	StarNum int32 `gorm:"column:starNum; not null"`
	SeenNum int32 `gorm:"column:seenNum; not null"`

	Liked   []*Like    `gorm:"foreignKey:ArticalID"`
	Stared  []*Star    `gorm:"foreignKey:ArticalID"`
	Comment []*Comment `gorm:"foreignKey:ArticalID"`
}

func (art *Artical) TableName() string {
	return constants.ArticalTableName
}

type User struct{}

func (u *User) TableName() string {
	return constants.UserTableName
}

// 创建文章
func CreateArtical(ctx context.Context, arts []*Artical) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		for _, art := range arts {
			// 作者文章数 + 1
			if err := tx.WithContext(ctx).Model(&User{}).Where("username = ?", art.Author).Update("artNum", gorm.Expr("artNum + ?", 1)).Error; err != nil {
				return errno.ServiceFault
			}
			if err := tx.WithContext(ctx).Create(arts).Error; err != nil {
				return errno.ServiceFault
			}
		}
		return nil
	})
}

// 根据 ID 查询文章
func QueryArtical(ctx context.Context, ids []int32) ([]*Artical, error) {
	res := make([]*Artical, 0)
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("id in ?", ids).Find(&res).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})

	return res, nil
}

// 根据 ID 删除文章 及其所有评论
func DeleteArtical(ctx context.Context, id int32) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		res := make([]int32, 0)
		// 查询该文章的所有 master 评论 并删除
		if err := tx.WithContext(ctx).Model(&Comment{}).Select("id").Where("articalID = ?", id).Where("master is null").Order("updatedAt DESC").Find(&res).Error; err != nil {
			return errno.ServiceFault
		}
		for _, cm := range res {
			if err := tx.WithContext(ctx).Where("master = ?", cm).Delete(&Comment{}).Error; err != nil {
				return errno.ServiceFault
			}
			if err := tx.WithContext(ctx).Where("id = ?", cm).Delete(&Comment{}).Error; err != nil {
				return errno.ServiceFault
			}
		}

		// 查询文章作者
		var author string
		if err := tx.WithContext(ctx).Model(&Artical{}).Select("author").Where("id = ?", id).Find(&author).Error; err != nil {
			return errno.ServiceFault
		}
		// 作者文章数 - 1
		if err := tx.WithContext(ctx).Model(&User{}).Where("username = ?", author).Update("artNum", gorm.Expr("artNum - ?", 1)).Error; err != nil {
			return errno.ServiceFault
		}
		if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&Artical{}).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 根据 ID 更新文章为 art
func UpdateArtical(ctx context.Context, art *Artical) error {
	if err := DB.WithContext(ctx).Where("id = ?", art.ID).Updates(*art).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 Author 查询文章
func QueryArticalByAuthor(ctx context.Context, author string) ([]int32, error) {
	res := make([]int32, 0)
	if err := DB.Model(&Artical{}).WithContext(ctx).Select("id").Where("author = ?", author).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// meaningless
// 根据 Title 查询文章 返回所有 Title 相同的文章
func QueryArticalByTitle(ctx context.Context, title string) ([]*Artical, error) {
	res := make([]*Artical, 0)
	if err := DB.WithContext(ctx).Where("title = ?", title).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}
