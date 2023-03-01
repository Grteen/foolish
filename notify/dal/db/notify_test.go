package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"
)

func TestReplyNotify(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&ReplyNotify{})
	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}
	err := CreateReplyNotify(cg, []*ReplyNotify{
		{
			Notify: Notify{
				UserName: "Grteen-test",
				Title:    "a first test",
				Sender:   "Grteen-test",
				Text:     "very very good good",
				IsRead:   false,
			},
			TargetID:  16,
			CommentID: 1,
		},
	})
	if err != nil {
		t.Error(err)
	}

	res, err := QueryReplyNotify(cg, []int32{1})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
