package db

import (
	"be/offline/swap/pack"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"
)

type UserSub struct {
	User string `gorm:"column:user"`
	Sub  string `gorm:"column:subscribe"`
}

func (*UserSub) TableName() string {
	return constants.UserSubTableName
}

// 查询一个用户的所有订阅数
func QuerySubNum(cg *config.Config, username string) (int32, error) {
	var count int64
	if err := cg.Tx.WithContext(cg.Ctx).Model(&UserSub{}).Where("user = ?", username).Count(&count).Error; err != nil {
		pack.EPrint(err.Error())
		return 0, errno.ServiceFault
	}
	return int32(count), nil
}

// 查询一个用户的所有粉丝数
func QueryFanNum(cg *config.Config, username string) (int32, error) {
	var count int64
	if err := cg.Tx.WithContext(cg.Ctx).Model(&UserSub{}).Where("subscribe = ?", username).Count(&count).Error; err != nil {
		pack.EPrint(err.Error())
		return 0, errno.ServiceFault
	}
	return int32(count), nil
}

// 删除对于一个用户的所有订阅 以及该用户的订阅
func DeleteSub(cg *config.Config, username string) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("user = ?", username).Or("subscribe = ?", username).Delete(&UserSub{}).Error; err != nil {
		pack.EPrint(err.Error())
		return errno.ServiceFault
	}
	pack.SPrint("Delete all subscribe of " + username)
	return nil
}
