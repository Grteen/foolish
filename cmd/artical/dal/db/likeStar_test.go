package db

import (
	"be/pkg/constants"
	"context"
	"fmt"
	"testing"
)

func TestLike(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Like{})
	DB.AutoMigrate(&Star{})

	ctx := context.WithValue(context.Background(), constants.LikeStarModel, &Star{})
	err := CreateLikeStar(ctx, []*LikeStar{
		{
			UserName:  "Grteen-test",
			ArticalID: 2,
		},
	})

	if err != nil {
		t.Error(err)
	}

	res, err := QueryLikeStar(ctx, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 2,
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res[0])

	tmp, err := QueryAllLikeStar(ctx, "Grteen-test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tmp)

	err = DeleteLikeStar(ctx, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 2,
	})

	if err != nil {
		t.Error(err)
	}

	ctx = context.WithValue(context.Background(), constants.LikeStarModel, &Like{})
	err = CreateLikeStar(ctx, []*LikeStar{
		{
			UserName:  "Grteen-test",
			ArticalID: 3,
		},
	})

	if err != nil {
		t.Error(err)
	}

	res, err = QueryLikeStar(ctx, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 3,
	})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res[0])

	err = DeleteLikeStar(ctx, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 3,
	})

	if err != nil {
		t.Error(err)
	}
}
