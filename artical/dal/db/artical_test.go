package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"
)

func TestArtical(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Artical{})
	DB.AutoMigrate(&Star{})
	DB.AutoMigrate(&Like{})
	DB.AutoMigrate(&Comment{})

	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}

	err := CreateArtical(cg, []*Artical{
		{Title: "title1", Author: "Grteen", Text: "this is a good text", Cover: "here", Description: "OK"},
	})

	if err != nil {
		t.Error(err)
	}

	res, err := QueryArtical(cg, []int32{10, 11, 12})
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("not found")
	}

	fmt.Println(res)

	err = DeleteArtical(cg, 28)
	if err != nil {
		t.Error(err)
	}
}
