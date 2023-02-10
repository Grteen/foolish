package rdb

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestArtical(t *testing.T) {
	RedisInit()
	ctx := context.Background()
	err := SetArtical(ctx, []*RdbArtical{
		{
			ID:        114514,
			CreatedAt: time.Now().String(),
			Author:    "Grteen-test",
			Title:     "test for redis",
			Text:      "a good test for ",

			LikeNum: 11,
			StarNum: 45,
			SeenNum: 14,
		},
	})
	if err != nil {
		t.Error(err)
	}

	res, _, err := GetArtical(ctx, []int32{114514})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res[0])
}
