package rdb

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestComment(t *testing.T) {
	RedisInit()
	ctx := context.Background()
	err := SetRdbComment(ctx, []*RdbComment{
		{
			ID:          114514,
			CreatedAt:   time.Now().String(),
			UserName:    "Grteen-test",
			TargetID:    1,
			CommentText: "test of redis",
			Type:        0,

			Master: 15,
			Reply:  []int32{16, 17},
		},
	})
	if err != nil {
		t.Error(err)
	}

	res, _, err := GetRdbComment(ctx, []int32{114514})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res[0])
}
