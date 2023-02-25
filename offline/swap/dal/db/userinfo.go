package db

import (
	"be/offline/swap/pack"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"

	"gorm.io/gorm"
)

type UserInfo struct {
	gorm.Model
	UserName string `json:"username" gorm:"column:username"`

	NickName    string `json:"nickname" gorm:"column:nickname"`
	Description string `json:"description" gorm:"column:description"`
	UserAvator  string `json:"avator" gorm:"column:avator"`
	// UserBackGround string
}

func (uf *UserInfo) TableName() string {
	return constants.UserInfoTableName
}

// 软删某个用户的信息
func DeleteUserInfo(cg *config.Config, username string) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("username = ?", username).Delete(&UserInfo{}).Error; err != nil {
		pack.EPrint(err.Error())
		return errno.ServiceFault
	}
	pack.SPrint("Delete userinfo of " + username)
	return nil
}
