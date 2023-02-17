package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"
)

func TestUserInfo(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&UserInfo{})
	DB.AutoMigrate(&User{})
	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}
	err := CreateUserInfo(cg, &UserInfo{
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

	err = UpdateUserInfo(cg, uf)
	if err != nil {
		t.Error(err)
	}

	res, err := QueryUserInfo(cg, uf.UserName)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
