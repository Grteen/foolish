package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"time"
)

// 显示在用户个人主页的基础信息
type UserInfo struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserName  string `json:"username" gorm:"column:username"`

	NickName    string `json:"nickname" gorm:"column:nickname"`
	Description string `json:"description" gorm:"column:description"`
	UserAvator  string `json:"avator" gorm:"column:avator"`
	// UserBackGround string
}

func (uf *UserInfo) TableName() string {
	return constants.UserInfoTableName
}

// 创建用户信息 uf
func CreateUserInfo(ctx context.Context, uf *UserInfo) error {
	if err := DB.WithContext(ctx).Create(uf).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 更新用户信息为 uf
func UpdateUserInfo(ctx context.Context, uf *UserInfo) error {
	if err := DB.WithContext(ctx).Where("username = ?", uf.UserName).Updates(*uf).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 userName 查询用户信息
func QueryUserInfo(ctx context.Context, userName string) ([]*UserInfo, error) {
	res := make([]*UserInfo, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 userName 查询用户头像
func QueryAvator(ctx context.Context, userName string) ([]string, error) {
	res := make([]string, 0)
	if err := DB.WithContext(ctx).Model(&UserInfo{}).Select("avator").Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}
