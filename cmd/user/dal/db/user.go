package db

import (
	"be/cmd/user/dal/rdb"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName        string    `json:"username" gorm:"column:username; not null; unique"`
	Email           string    `json:"email" gorm:"column:email;not null; unique"`
	PassWord        string    `json:"password" gorm:"column:pw; not null"`
	IsAdministrator int32     `json:"isadministrator" gorm:"column:isadministrator; not null"`
	UserInfo        *UserInfo `gorm:"foreignKey:UserName; references:UserName"`

	SubNum    int32   `json:"subNum" gorm:"column:subNum; not null"`
	FanNum    int32   `json:"fanNum" gorm:"column:fanNum; not null"`
	ArtNum    int32   `json:"artNum" gorm:"column:artNum; not null"`
	ActNum    int32   `json:"actNum" gorm:"column:actNum; not null"`
	Subscribe []*User `gorm:"many2many:subscribe; foreignKey:UserName; joinForeignKey:User; References:UserName; joinReferences:Subscribe"`

	FanPublic int32 `gorm:"column:fanPublic; not null"`
	SubPublic int32 `gorm:"column:subPublic; not null"`
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

// 更新用户 粉丝关注列表权限
func UpdateUserPublic(cg *config.Config, userName string, fanPublic, subPublic int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Model(&User{}).Where("username = ?", userName).Updates(map[string]interface{}{"fanPublic": fanPublic, "subPublic": subPublic}).Error; err != nil {
		rdb.EPrint(err.Error())
		return errno.ServiceFault
	}
	return nil
}
