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

type Target struct {
	TargetID int32 `gorm:"column:id"`
	Type     int32 // 0 is artical 1 is action
}

// 搜索某一个特定用户的文章和动态
func SearchUserZoom(cg *config.Config, username, keyword string, limit, offset int32) ([]*Target, error) {
	key := "%" + keyword + "%"
	res := make([]*Target, 0)
	if len(keyword) == 0 {
		if err := cg.Tx.Raw("select * from ( select `id` , `created_at` , 0 as `type` from `artical` where `author` = ? and isnull(`deleted_at`) UNION ALL select `id` , `created_at` , 1 as `type` from `action` where `author` = ? and isnull(`deleted_at`) ) as temp order by `created_at` DESC limit ? offset ?", username, username, limit, offset).Scan(&res).Error; err != nil {
			return nil, errno.ServiceFault
		}
	} else {
		if err := cg.Tx.Raw("select * from ( select `id` , `created_at` , 0 as `type` from `artical` where `author` = ? and isnull(`deleted_at`) and ( `title` like ? or `description` like ? or `text` like ? ) UNION ALL select `id` , `created_at` , 1 as `type` from `action` where `author` = ? and isnull(`deleted_at`) and `text` like ? ) as temp order by `created_at` DESC limit ? offset ?", username, key, key, key, username, key, limit, offset).Scan(&res).Error; err != nil {
			return nil, errno.ServiceFault
		}
	}

	return res, nil
}
