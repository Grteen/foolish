package db

import (
	"context"
	"fmt"
	"testing"
)

func TestArtical(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Star{})
	DB.AutoMigrate(&Like{})
	DB.AutoMigrate(&Comment{})
	DB.AutoMigrate(&Artical{})
	err := CreateArtical(context.Background(), []*Artical{
		{Title: "title1", Author: "Grteen", Text: "this is a good text"},
	})

	if err != nil {
		t.Error(err)
	}

	res, err := QueryArtical(context.Background(), []int32{10, 11, 12})
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("not found")
	}

	fmt.Println(res)

	err = DeleteArtical(context.Background(), 7)
	if err != nil {
		t.Error(err)
	}
}
