package db

import (
	"be/cmd/notify/log"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"

	"gorm.io/gorm"
)

// 系统通知消息  用于公告栏
type SystemNotify struct {
	gorm.Model
	Text string `gorm:"column:text"`
}

func (s *SystemNotify) TableName() string {
	return constants.SystemNotifyTableName
}

// 创建系统消息
func CreateSystemNotify(cg *config.Config, stfs []*SystemNotify) error {
	if err := cg.Tx.WithContext(cg.Ctx).Create(stfs).Error; err != nil {
		log.EPrint(err.Error())
		return errno.ServiceFault
	}
	return nil
}

// 查询所有系统消息的ID 按照时间降序排序
func QueryAllSystemNotify(cg *config.Config, limit, offset int32) ([]int32, error) {
	res := make([]int32, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Model(&SystemNotify{}).Select("id").Order("created_at DESC").Limit(int(limit)).Offset(int(offset)).Find(&res).Error; err != nil {
		log.EPrint(err.Error())
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 查询某个ID的系统消息
func QuerySystemNotify(cg *config.Config, ids []int32) ([]*SystemNotify, error) {
	res := make([]*SystemNotify, 0)
	for _, id := range ids {
		stf := &SystemNotify{}
		if err := cg.Tx.WithContext(cg.Ctx).Where("id = ?", id).Find(stf).Error; err != nil {
			log.EPrint(err.Error())
			return nil, errno.ServiceFault
		}
		if stf.ID == 0 {
			continue
		}
		res = append(res, stf)
	}
	return res, nil
}
