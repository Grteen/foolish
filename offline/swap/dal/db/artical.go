package db

import (
	"be/offline/swap/pack"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"
	"time"
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
}

func (art *Artical) TableName() string {
	return constants.ArticalTableName
}

// 查询某个用户的所有文章数量
func QueryArtNum(cg *config.Config, username string) (int32, error) {
	var count int64
	if err := cg.Tx.WithContext(cg.Ctx).Model(&Artical{}).Where("author = ?", username).Count(&count).Error; err != nil {
		pack.EPrint(err.Error())
		return 0, errno.ServiceFault
	}
	return int32(count), nil
}

// 删除某个用户的所有文章
func DeleteArtical(cg *config.Config, username string) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("author = ?", username).Delete(&Artical{}).Error; err != nil {
		pack.EPrint(err.Error())
		return errno.ServiceFault
	}
	pack.SPrint("Delete all Articals of " + username)
	return nil
}
