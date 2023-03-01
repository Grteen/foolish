package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	MySQLInit()
	cg := &config.Config{Ctx: context.Background(), Tx: DB}
	res, err := Search(cg, "Grteen", 15, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res)
}
