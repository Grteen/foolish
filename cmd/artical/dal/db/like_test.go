package db

import (
	"context"
	"testing"
)

func TestLike(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&Like{})

	err := CreateLike(context.Background(), []*Like{
		{UserName: "Grteen-test", ArticalID: 2},
	})

	if err != nil {
		t.Error(err)
	}

	err = DeleteLike(context.Background(), &Like{
		UserName:  "Grteen-test",
		ArticalID: 2,
	})
	if err != nil {
		t.Error(err)
	}
}
