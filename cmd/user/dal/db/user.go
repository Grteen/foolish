package db

import (
	"be/pkg/constants"
	"be/pkg/errno"
	"context"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserName  string    `json:"username" gorm:"column:username; not null; unique"`
	Email     string    `json:"email" gorm:"column:email; unique"`
	PassWord  string    `json:"password" gorm:"column:pw"`
	UserInfo  *UserInfo `gorm:"foreignKey:UserName; references:UserName"`

	SubNum    int32   `json:"subNum" gorm:"column:subNum; not null"`
	FanNum    int32   `json:"fanNum" gorm:"column:fanNum; not null"`
	ArtNum    int32   `json:"artNum" gorm:"column:artNum; not null"`
	Subscribe []*User `gorm:"many2many:subscribe; foreignKey:UserName; joinForeignKey:User; References:UserName; joinReferences:Subscribe"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// 创建用户
func CreateUser(ctx context.Context, users []*User) error {
	if err := DB.WithContext(ctx).Create(users).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}

// 根据 userName 查询用户 返回所有 userName 相同的用户 （正常情况只返回一个）
func QueryUser(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("username = ?", userName).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 email 查询用户 返回所有 email 相同的用户 （正常情况只返回一个）
func QueryUserByEmail(ctx context.Context, email string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("email = ?", email).Find(&res).Error; err != nil {
		return nil, errno.ServiceFault
	}
	return res, nil
}

// 根据 userName 删除用户
func DeleteUser(ctx context.Context, userName string) error {
	if err := DB.WithContext(ctx).Where("username = ?", userName).Delete(&User{}).Error; err != nil {
		return errno.ServiceFault
	}
	return nil
}
