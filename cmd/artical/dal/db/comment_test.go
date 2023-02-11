package db

import (
	"context"
	"fmt"
	"testing"
)

func TestComment(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Comment{})

	ctx := context.Background()
	err := CreateComment(ctx, []*Comment{
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

	cm, err := QueryComment(ctx, []int32{5, 6, 7})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(cm)

	err = UpdateComment(ctx, &Comment{
		ID:          2,
		UserName:    "Grteen-test",
		ArticalID:   3,
		CommentText: "New-Comment",
	})

	if err != nil {
		t.Error(err)
	}

	cms, err := QueryCommentByArticalID(ctx, 3)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(cms[0], cms[1])

	if err = DeleteComment(ctx, 1); err != nil {
		t.Error(err)
	}
	if err = DeleteComment(ctx, 2); err != nil {
		t.Error(err)
	}

}

func TestReply(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Comment{})

	ctx := context.Background()
	// var m uint = 1
	// err := CreateComment(ctx, []*Comment{
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

	res, err := QueryComment(ctx, []int32{1})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res[0].Reply)

	// err = DeleteComment(ctx, 1)
	// if err != nil {
	// 	t.Error(err)
	// }
}
