package db

import (
	"be/pkg/config"
	"time"
)

type AllNotify struct {
	ID         uint
	CreatedAt  time.Time
	NotifyType int32
}

// 查询所有通知的id 按照更新时间降序
func SearchAllNotify(cg *config.Config, limit, offset int32) ([]*AllNotify, error) {
	res := make([]*AllNotify, 0)
	cg.Tx.WithContext(cg.Ctx).Raw("select * from (select `id`, `createdAt`, 1 as `NotifyType` from `likeNotify` union all select `id`, `createdAt`, 0 as `NotifyType` from `replyNotify`) as temp order by `createdAt` desc limit ? offset ?", limit, offset).Scan(&res)
	return res, nil
}
