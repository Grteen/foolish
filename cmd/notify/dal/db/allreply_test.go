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
	res, _ := SearchAllNotify(cg, 20, 0)
	fmt.Println(res)
}
