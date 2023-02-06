package db

import (
	"context"
	"fmt"
	"testing"
)

func TestArtical(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&ArticalStar{})
	DB.AutoMigrate(&Like{})
	DB.AutoMigrate(&Comment{})
	DB.AutoMigrate(&Artical{})
	err := CreateArtical(context.Background(), []*Artical{
		{Title: "title1", Author: "Grteen", Text: "this is a good text"},
	})

	if err != nil {
		t.Error(err)
	}

	res, err := QueryArtical(context.Background(), 1)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error("not found")
	}

	fmt.Println(res)

	err = DeleteArtical(context.Background(), res.ID)
	if err != nil {
		t.Error(err)
	}
}
