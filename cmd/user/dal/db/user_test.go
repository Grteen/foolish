package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&User{})
	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}
	err := CreateUser(cg, []*User{
		{
			UserName: "Grteen33",
			PassWord: "123456",
			Email:    "temp333.com",
		},
	})

	if err != nil {
		t.Error(err)
	}

	u, err := QueryUser(cg, "Grteen33")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u[0])

	err = DeleteUser(cg, "Grteen33")
	if err != nil {
		t.Error(err)
	}
}

func TestSubscribe(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&User{})
	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}
	err := CreateUser(cg, []*User{
		{
			UserName: "Grteen777",
			PassWord: "123456",
			Email:    "not good  unique.com",

			Subscribe: []*User{
				{
					UserName: "Grteen1437666",
					PassWord: "123456",
					Email:    "ok of unique.com",
				},
			},
		},
	})

	if err != nil {
		t.Error(err)
	}
}
