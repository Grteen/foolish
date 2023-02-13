package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"

	"gorm.io/gorm"
)

// User 与 User 之间的连接表
type UserSub struct {
	User string `gorm:"column:user"`
	Sub  string `gorm:"column:subscribe"`
}

func (*UserSub) TableName() string {
	return constants.UserSubTableName
}

// 创建订阅 （user 订阅 sub）
func CreateSubscribe(ctx context.Context, subs []*UserSub) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		for _, sub := range subs {
			// 订阅者 订阅数量 + 1
			if err := tx.WithContext(ctx).Model(&User{}).Where("username = ?", sub.User).Update("subNum", gorm.Expr("subNum + ?", 1)).Error; err != nil {
				return errno.ServiceFault
			}
			// 被订阅者 粉丝数 + 1
			if err := tx.WithContext(ctx).Model(&User{}).Where("username = ?", sub.Sub).Update("fanNum", gorm.Expr("fanNum + ?", 1)).Error; err != nil {
				return errno.ServiceFault
			}
			if err := tx.WithContext(ctx).Create(subs).Error; err != nil {
				return errno.ServiceFault
			}
		}
		return nil
	})
}

// 删除订阅  (user 取消订阅 sub)
func DeleteSubscribe(ctx context.Context, sub *UserSub) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 订阅者 订阅数量 - 1
		if err := tx.WithContext(ctx).Model(&User{}).Where("username = ?", sub.User).Update("subNum", gorm.Expr("subNum - ?", 1)).Error; err != nil {
			return errno.ServiceFault
		}
		// 被订阅者 粉丝数 - 1
		if err := tx.WithContext(ctx).Model(&User{}).Where("username = ?", sub.Sub).Update("fanNum", gorm.Expr("fanNum - ?", 1)).Error; err != nil {
			return errno.ServiceFault
		}
		if err := tx.WithContext(ctx).Where("user = ?", sub.User).Where("subscribe = ?", sub.Sub).Delete(&UserSub{}).Error; err != nil {
			return errno.ServiceFault
		}
		return nil
	})
}

// 查询是否有 user 对于 sub 的订阅
func QuerySubscribe(ctx context.Context, sub *UserSub) ([]*UserSub, error) {
	res := make([]*UserSub, 0)
	if err := DB.WithContext(ctx).Where("user = ?", sub.User).Where("subscribe = ?", sub.Sub).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 查询一个用户的所有订阅 返回订阅的用户名称
func QueryAllSubscribe(ctx context.Context, userName string) ([]string, error) {
	res := make([]string, 0)
	if err := DB.WithContext(ctx).Model(&UserSub{}).Select("subscribe").Where("user = ?", userName).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 查询一个用户的所有粉丝 返回粉丝的用户名称
func QueryAllFans(ctx context.Context, userName string) ([]string, error) {
	res := make([]string, 0)
	if err := DB.WithContext(ctx).Model(&UserSub{}).Select("user").Where("subscribe = ?", userName).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}
