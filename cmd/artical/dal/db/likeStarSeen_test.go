package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"
)

func TestLike(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Like{})
	DB.AutoMigrate(&Star{})
	DB.AutoMigrate(&Seen{})

	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}

	err := CreateLikeStar(cg, []*LikeStar{
		{
			UserName:  "Grteen-test",
			ArticalID: 3,
		},
	}, &Like{})

	if err != nil {
		t.Error(err)
	}

	res, err := QueryLikeStar(cg, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 3,
	}, &Like{})

	if err != nil {
		t.Error(err)
	}

	fmt.Println(res[0])

	err = DeleteLikeStar(cg, &LikeStar{
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

	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}
	err := CreateStarFolder(cg, []*StarFolder{
		{
			UserName:   "Grteen-test",
			FolderName: "test o",
			Public:     0,
		},
	})
	if err != nil {
		t.Error(err)
	}
	err = CreateStar(cg, []*Star{
		{
			UserName:  "Grteen-test",
			ArticalID: 16,
			FolderID:  1,
		},
	})
	if err != nil {
		t.Error(err)
	}

	res, err := QueryAllStarFolder(cg, "Grteen-test")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res[0])

	s, err := QueryAllStar(cg, 1, 20, 0)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(s)

	temp, err := QueryLikeStar(cg, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 16,
	}, &Star{})
	fmt.Println(temp)

	err = DeleteLikeStar(cg, &LikeStar{
		UserName:  "Grteen-test",
		ArticalID: 16,
	}, &Star{})

	if err != nil {
		t.Error(err)
	}
}
