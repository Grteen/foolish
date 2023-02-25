package db

import (
	"be/offline/swap/pack"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"
	"strconv"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"username" gorm:"column:username; not null; unique"`
	Email    string `json:"email" gorm:"column:email; unique"`
	PassWord string `json:"password" gorm:"column:pw"`
	// UserInfo  *UserInfo `gorm:"foreignKey:UserName; references:UserName"`

	SubNum    int32   `json:"subNum" gorm:"column:subNum; not null"`
	FanNum    int32   `json:"fanNum" gorm:"column:fanNum; not null"`
	ArtNum    int32   `json:"artNum" gorm:"column:artNum; not null"`
	Subscribe []*User `gorm:"many2many:subscribe; foreignKey:UserName; joinForeignKey:User; References:UserName; joinReferences:Subscribe"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// 查询所有用户
func QueryUser(cg *config.Config, limit, offset int32) ([]*User, error) {
	res := make([]*User, 0)
	if err := cg.Tx.WithContext(cg.Ctx).Model(&User{}).Limit(int(limit)).Offset(int(offset)).Find(&res).Error; err != nil {
		pack.EPrint(err.Error())
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 更新用户
func UpdateUser(cg *config.Config, user *User) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		pack.EPrint(err.Error())
		return errno.ServiceFault
	}
	pack.SPrint("Update a User ID = " + strconv.Itoa(int(user.ID)))
	return nil
}

// 删除用户
func DeleteUser(cg *config.Config, id int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("id = ?", id).Delete(&User{}).Error; err != nil {
		pack.EPrint(err.Error())
		return errno.ServiceFault
	}
	pack.SPrint("Delete a User ID = " + strconv.Itoa(int(id)))
	return nil
}
