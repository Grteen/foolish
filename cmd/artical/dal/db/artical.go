package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"time"

	"gorm.io/gorm"
)

type Artical struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Title       string `json:"title" gorm:"column:title; not null"`
	Author      string `json:"author" gorm:"column:author; not null"`
	Text        string `json:"text" gorm:"column:text; type:text"`
	Description string `json:"description" gorm:"column:description; not null"`

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

// 创建文章
func CreateArtical(ctx context.Context, arts []*Artical) error {
	if err := DB.WithContext(ctx).Create(arts).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 ID 查询文章
func QueryArtical(ctx context.Context, ids []int32) ([]*Artical, error) {
	res := make([]*Artical, 0)
	DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Where("ID in ?", ids).Find(&res).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})

	return res, nil
}

// 根据 ID 删除文章
func DeleteArtical(ctx context.Context, id int32) error {
	if err := DB.WithContext(ctx).Where("id = ?", id).Delete(&Artical{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
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
