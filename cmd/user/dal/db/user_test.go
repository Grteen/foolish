package db

import (
	"context"
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&User{})
	err := CreateUser(context.Background(), []*User{
		{
			UserName: "Grteen1",
			PassWord: "123456",
			Email:    "temp.com",
		},
	})

	if err != nil {
		t.Error(err)
	}

	u, err := QueryUser(context.Background(), "Grteen1")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(u[0])

	err = DeleteUser(context.Background(), "Grteen1")
	if err != nil {
		t.Error(err)
	}
}

func TestSubscribe(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&User{})
	err := CreateUser(context.Background(), []*User{
		{
			UserName: "Grteen114514",
			PassWord: "123456",
			Email:    "not unique.com",

			Subscribe: []*User{
				{
					UserName: "Grteen1437",
					PassWord: "123456",
					Email:    "ok.com",
				},
			},
		},
	})

	if err != nil {
		t.Error(err)
	}
}
