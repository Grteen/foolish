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

// 查询所有通知的id 按照发布时间降序
func SearchAllNotify(cg *config.Config, username string, limit, offset int32) ([]*AllNotify, error) {
	res := make([]*AllNotify, 0)
	cg.Tx.WithContext(cg.Ctx).Raw("select * from (select `id`, `createdAt`, `isread`, 1 as `NotifyType` from `likeNotify`where `username` = ? and `isdelete` = false union all select `id`, `createdAt`, `isread`, 0 as `NotifyType` from `replyNotify` where `username` = ? and `isdelete` = false) as temp order by `isread` asc, `createdAt` desc limit ? offset ?", username, username, limit, offset).Scan(&res)
	return res, nil
}
