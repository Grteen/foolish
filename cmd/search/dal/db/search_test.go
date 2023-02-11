package db

import (
	"context"
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	MySQLInit()
	ctx := context.Background()
	res, err := Search(ctx, "Grteen", 15, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
