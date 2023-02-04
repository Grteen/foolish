package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"time"
)

type ArticalStar struct {
	ID        uint   `gorm:"primarykey"`
	UserName  string `gorm:"username"`
	ArticalID uint   `gorm:"articalID"`
}

func (as *ArticalStar) TableName() string {
	return constants.ArticalStarTableName
}

type Artical struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Title  string         `json:"title" gorm:"column:title; not null"`
	Author string         `json:"author" gorm:"column:author; not null"`
	Text   string         `json:"text" gorm:"column:text; type:text"`
	Stared []*ArticalStar `gorm:"foreignKey:ArticalID"`
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
func QueryArtical(ctx context.Context, id uint) (*Artical, error) {
	res := &Artical{}
	if err := DB.WithContext(ctx).Where("id = ?", id).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 ID 删除文章
func DeleteArtical(ctx context.Context, id uint) error {
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

// meaningless
// 根据 Title 查询文章 返回所有 Title 相同的文章
func QueryArticalByTitle(ctx context.Context, title string) ([]*Artical, error) {
	res := make([]*Artical, 0)
	if err := DB.WithContext(ctx).Where("title = ?", title).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}
