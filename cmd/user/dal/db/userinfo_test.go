package db

import (
	"context"
	"fmt"
	"testing"
)

func TestUserInfo(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&UserInfo{})
	DB.AutoMigrate(&User{})

	err := CreateUserInfo(context.Background(), &UserInfo{
		UserName: "Grteen",
	})

	if err != nil {
		t.Error(err)
	}

	uf := &UserInfo{
		UserName:    "Grteen",
		NickName:    "Good Name",
		Description: "this is a desc",
		UserAvator:  "127.0.0.1/here",
	}

	err = UpdateUserInfo(context.Background(), uf)
	if err != nil {
		t.Error(err)
	}

	res, err := QueryUserInfo(context.Background(), uf.UserName)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
