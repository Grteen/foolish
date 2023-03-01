package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"
)

func TestLikeNotify(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&LikeNotify{})
	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}
	err := CreateLikeNotify(cg, []*LikeNotify{
		{
			Notify: Notify{
				UserName: "Grteen-test",
				Title:    "a first test",
				Sender:   "Grteen-test",
				Text:     "very very good good",
				IsRead:   false,
			},
			TargetID: 16,
		},
	})
	if err != nil {
		t.Error(err)
	}

	res, err := QueryLikeNotify(cg, []int32{1})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
