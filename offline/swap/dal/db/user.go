package db

import (
	"be/offline/swap/dal/rdb"
	"be/pkg/config"
	"be/pkg/constants"
	"be/pkg/errno"
	"strconv"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserName  string `json:"username" gorm:"column:username; not null; unique"`
	Email     string `json:"email" gorm:"column:email; unique"`
	PassWord  string `json:"password" gorm:"column:pw"`
	// UserInfo  *UserInfo `gorm:"foreignKey:UserName; references:UserName"`

	SubNum    int32   `json:"subNum" gorm:"column:subNum; not null"`
	FanNum    int32   `json:"fanNum" gorm:"column:fanNum; not null"`
	ArtNum    int32   `json:"artNum" gorm:"column:artNum; not null"`
	Subscribe []*User `gorm:"many2many:subscribe; foreignKey:UserName; joinForeignKey:User; References:UserName; joinReferences:Subscribe"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// 更新用户
func UpdateUser(cg *config.Config, user *User) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		rdb.SPrint(err.Error())
		return errno.ServiceFault
	}
	rdb.SPrint("Update a User ID = " + strconv.Itoa(int(user.ID)))
	return nil
}

// 删除用户
func DeleteUser(cg *config.Config, id int32) error {
	if err := cg.Tx.WithContext(cg.Ctx).Where("id = ?", id).Delete(&User{}).Error; err != nil {
		rdb.SPrint(err.Error())
		return errno.ServiceFault
	}
	rdb.SPrint("Delete a User ID = " + strconv.Itoa(int(id)))
	return nil
}
