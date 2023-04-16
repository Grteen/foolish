package db

import (
	"be/pkg/config"
	"time"
)

type ArtAct struct {
	ID        uint
	CreatedAt time.Time
	Type      int32
}

// 查询一个用户的所有文章和动态 按照发布时间降序
func SearchArtAct(cg *config.Config, username string, limit, offset int32) ([]*ArtAct, error) {
	res := make([]*ArtAct, 0)
	cg.Tx.WithContext(cg.Ctx).Raw("select * from ( select `id` , `created_at` , 0 as `type` from `artical` where `author` = ? and isnull(`deleted_at`) UNION ALL select `id` , `created_at` , 1 as `type` from `action` where `author` = ? and isnull(`deleted_at`) ) as temp order by `created_at` DESC limit ? offset ?", username, username, limit, offset).Scan(&res)
	return res, nil
}
