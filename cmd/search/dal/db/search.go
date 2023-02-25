package db

import (
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"
)

type Artical struct {
}

func (art *Artical) TableName() string {
	return constants.ArticalTableName
}

func Search(cg *config.Config, keyword string, limit int32, offset int32) ([]int32, error) {
	key := "%" + keyword + "%"
	res := make([]int32, 0)
	if len(keyword) == 0 {
		if err := cg.Tx.WithContext(cg.Ctx).Model(&Artical{}).Select("ID").Order("updated_at DESC").Limit(int(limit)).Offset(int(offset)).Find(&res).Error; err != nil {
			return nil, errno.ServiceFault
		}
	} else {
		if err := cg.Tx.WithContext(cg.Ctx).Model(&Artical{}).Select("ID").Where("title like ?", key).Or("description like ?", key).
			Or("author like ?", key).Or("text like ?", key).
			Limit(int(limit)).Offset(int(offset)).Order("updated_at DESC").Find(&res).Error; err != nil {
			return nil, errno.ServiceFault
		}
	}

	return res, nil
}
