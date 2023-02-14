package db

import (
	"context"
	"fmt"
	"testing"
)

func TestReplyNotify(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&ReplyNotify{})
	ctx := context.Background()
	err := CreateReplyNotify(ctx, []*ReplyNotify{
		{
			Notify: Notify{
				UserName: "Grteen-test",
				Title:    "a first test",
				Sender:   "Grteen-test",
				Text:     "very good",
				IsRead:   false,
			},
		},
	})
	if err != nil {
		t.Error(err)
	}

	res, err := QueryReplyNotify(ctx, []int32{1})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
