package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"
)

func TestAction(t *testing.T) {
	MySQLInit()
	// DB.AutoMigrate(&Comment{})
	DB.AutoMigrate(&PicFile{})
	DB.AutoMigrate(&Action{})
	ctx := context.Background()
	err := CreateAction(config.NewConfig(ctx, DB), []*Action{
		{
			Text:   "first test of action",
			Author: "Grteen-test",
			PicFile: []*PicFile{
				{
					File: "http://test of a file",
				},
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	res, err := QueryAction(config.NewConfig(ctx, DB), []int32{
		4,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)

	err = DeleteAction(config.NewConfig(ctx, DB), 4)
	if err != nil {
		t.Error(err)
	}
}
