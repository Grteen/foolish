package rdb

import (
	"be/pkg/constants"
	"context"
	"testing"
)

func TestLikeStar(t *testing.T) {
	RedisInit()
	IncreaseLikeStar(context.Background(), 1, 1, constants.RdbArticalFieldLikeNum)
	IncreaseLikeStar(context.Background(), 1, 1, constants.RdbArticalFieldStarNum)
	IncreaseLikeStar(context.Background(), 1, 1, constants.RdbArticalFieldSeenNum)
	IncreaseLikeStar(context.Background(), 1, -1, constants.RdbArticalFieldLikeNum)
	IncreaseLikeStar(context.Background(), 1, -1, constants.RdbArticalFieldStarNum)
	IncreaseLikeStar(context.Background(), 1, -1, constants.RdbArticalFieldSeenNum)
}
