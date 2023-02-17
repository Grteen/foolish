package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"
)

func TestComment(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Comment{})

	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}
	_, err := CreateComment(cg, []*Comment{
		{
			UserName:    "Grteen-test",
			ArticalID:   3,
			CommentText: "First-Comment",
		},
		{
			UserName:    "Grteen-test",
			ArticalID:   3,
			CommentText: "Second-Comment",
		},
	})

	if err != nil {
		t.Error(err)
	}

	cm, err := QueryComment(cg, []int32{5, 6, 7})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(cm)

	err = UpdateComment(cg, &Comment{
		ID:          2,
		UserName:    "Grteen-test",
		ArticalID:   3,
		CommentText: "New-Comment",
	})

	if err != nil {
		t.Error(err)
	}

	cms, err := QueryCommentByArticalID(cg, 3)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(cms[0], cms[1])

	if err = DeleteComment(cg, 1); err != nil {
		t.Error(err)
	}
	if err = DeleteComment(cg, 2); err != nil {
		t.Error(err)
	}

}

func TestReply(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Comment{})

	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}
	// var m uint = 1
	// err := CreateComment(cg, []*Comment{
	// 	{
	// 		UserName:    "Grteen-test",
	// 		ArticalID:   3,
	// 		CommentText: "First-Comment",
	// 	},
	// 	{
	// 		UserName:    "Grteen-test",
	// 		ArticalID:   3,
	// 		CommentText: "Reply",
	// 		Master:      &m,
	// 	},
	// })

	// if err != nil {
	// 	t.Error(err)
	// }

	res, err := QueryComment(cg, []int32{1})
	if err != nil {
		t.Error(err)
	}
	if len(res) != 0 {
		fmt.Println(res[0].Reply)
	}

	// err = DeleteComment(cg, 1)
	// if err != nil {
	// 	t.Error(err)
	// }
}
