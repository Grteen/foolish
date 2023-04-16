package rdb

import (
	"context"
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	RedisInit()
	ctx := context.Background()
	SetUser(ctx, []*RdbUser{
		{
			UserName:    "Grteen-test",
			NickName:    "Grteen-Nick",
			Description: "good Description",
			UserAvator:  "http://127.0.0.1/here",
			SubNum:      11,
			FanNum:      45,
			ArtNum:      14,
		},
	})

	res, ungot, err := GetUser(ctx, []string{"Grteen-test", "ungot"})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res[0])
	fmt.Println(ungot)

	err = DelUser(ctx, "Grteen-test")
	if err != nil {
		t.Error(err)
	}
}
