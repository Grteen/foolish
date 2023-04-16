package db

import (
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"

	"gorm.io/gorm"
)

// 显示在用户个人主页的基础信息
type UserInfo struct {
	gorm.Model
	UserName string `json:"username" gorm:"column:username"`

	NickName    string `json:"nickname" gorm:"column:nickname"`
	Description string `json:"description" gorm:"column:description"`
	UserAvator  string `json:"avator" gorm:"column:avator"`
}

func (uf *UserInfo) TableName() string {
	return constants.UserInfoTableName
}

// 创建用户信息 uf
func CreateUserInfo(cg *config.Config, uf *UserInfo) error {
	if err := cg.Tx.WithContext(cg.Ctx).Create(uf).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 更新用户信息为 uf
func UpdateUserInfo(cg *config.Config, uf *UserInfo) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("username = ?", uf.UserName).Updates(*uf).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 userName 查询用户信息
func QueryUserInfo(cg *config.Config, userName string) ([]*UserInfo, error) {
	res := make([]*UserInfo, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 userName 查询用户头像
func QueryAvator(cg *config.Config, userName string) ([]string, error) {
	res := make([]string, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Model(&UserInfo{}).Select("avator").Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}
