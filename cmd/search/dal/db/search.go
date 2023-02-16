package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
)

type Artical struct {
}

func (art *Artical) TableName() string {
	return constants.ArticalTableName
}

func Search(ctx context.Context, keyword string, limit int32, offset int32) ([]int32, error) {
	key := "%" + keyword + "%"
	var temp int32
	if limit >= 20 || temp < 0 {
		temp = 20
	} else {
		temp = limit
	}
	res := make([]int32, temp)
	if err := DB.WithContext(ctx).Model(&Artical{}).Select("ID").Where("title like ?", key).Or("description like ?", key).
		Or("author like ?", key).Or("text like ?", key).
		Limit(int(limit)).Offset(int(offset)).Order("updatedAt DESC").Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}

	return res, nil
}
