package db

import (
	"be/pkg/config"
	"context"
	"fmt"
	"testing"
)

func TestAllReply(t *testing.T) {
	MySQLInit()
	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}
	username := "Grteen-test"
	res, _ := SearchAllNotify(cg, username, 20, 0)
	fmt.Println(res)
}
