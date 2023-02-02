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
