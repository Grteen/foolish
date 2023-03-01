package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func TestLike(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Like{})
	ctx := context.Background()
	cg := config.NewConfig(ctx, DB)
	err := CreateActionLike(cg, []*Like{
		{
			UserName: "Grteen-test",
			ActionID: 5,
		},
	})
	if err != nil {
		t.Error(err)
	}

	res, err := QueryActionLike(cg, &Like{
		UserName: "Grteen-test",
		ActionID: 5,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)

	err = DeleteActionLike(cg, &Like{
		Model: gorm.Model{
			ID: 1,
		},
		UserName: "Grteen-test",
		ActionID: 5,
	})
	if err != nil {
		t.Error(err)
	}
}
