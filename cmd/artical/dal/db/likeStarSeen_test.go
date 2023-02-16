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
	DB.AutoMigrate(&Seen{})

	ctx := context.WithValue(context.Background(), constants.LikeStarModel, &Star{})

	err := CreateLikeStar(ctx, []*LikeStar{
		{
			UserName:  "Grteen-test",
			ArticalID: 3,
		},
	}, &Like{})

	if err != nil {
		t.Error(err)
	}

	res, err := QueryLikeStar(ctx, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 3,
	}, &Like{})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res[0])

	err = DeleteLikeStar(ctx, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 3,
	}, &Like{})

	if err != nil {
		t.Error(err)
	}
}

func TestStar(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Like{})
	DB.AutoMigrate(&Star{})
	DB.AutoMigrate(&Seen{})
	DB.AutoMigrate(&StarFolder{})

	ctx := context.Background()
	err := CreateStarFolder(ctx, []*StarFolder{
		{
			UserName:   "Grteen-test",
			FolderName: "test o",
		},
	})
	if err != nil {
		t.Error(err)
	}
	err = CreateStar(ctx, []*Star{
		{
			UserName:  "Grteen-test",
			ArticalID: 16,
			FolderID:  1,
		},
	})
	if err != nil {
		t.Error(err)
	}

	res, err := QueryAllStarFolder(ctx, "Grteen-test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res[0])

	s, err := QueryAllStar(ctx, 1, 20, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(s)

	temp, err := QueryLikeStar(ctx, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 16,
	}, &Star{})
	fmt.Println(temp)

	err = DeleteLikeStar(ctx, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 16,
	}, &Star{})

	if err != nil {
		t.Error(err)
	}
}
