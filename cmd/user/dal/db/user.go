package db

import (
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string    `json:"username" gorm:"column:username; not null; unique"`
	Email    string    `json:"email" gorm:"column:email; unique"`
	PassWord string    `json:"password" gorm:"column:pw"`
	UserInfo *UserInfo `gorm:"foreignKey:UserName; references:UserName"`

	SubNum    int32   `json:"subNum" gorm:"column:subNum; not null"`
	FanNum    int32   `json:"fanNum" gorm:"column:fanNum; not null"`
	ArtNum    int32   `json:"artNum" gorm:"column:artNum; not null"`
	Subscribe []*User `gorm:"many2many:subscribe; foreignKey:UserName; joinForeignKey:User; References:UserName; joinReferences:Subscribe"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// 创建用户
func CreateUser(cg *config.Config, users []*User) error {
	if err := cg.Tx.WithContext(cg.Ctx).Create(users).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 userName 查询用户 返回所有 userName 相同的用户 （正常情况只返回一个）
func QueryUser(cg *config.Config, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 email 查询用户 返回所有 email 相同的用户 （正常情况只返回一个）
func QueryUserByEmail(cg *config.Config, email string) ([]*User, error) {
	res := make([]*User, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Where("email = ?", email).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 userName 删除用户
func DeleteUser(cg *config.Config, userName string) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("username = ?", userName).Delete(&User{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}
