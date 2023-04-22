package db

import (
	"be/pkg/config"
	"be/pkg/kafka"
	"context"
	"fmt"
	"testing"
)

func TestPubNotice(t *testing.T) {
	MySQLInit()
	kafka.Init()
	DB.AutoMigrate(&PubNotice{})
	cg := &config.Config{
		Ctx: context.Background(),
		Tx:  DB,
	}

	err := CreatePubNotice(cg, []*PubNotice{
		{
			UserName: "Grteen-test",
			Text:     "recovery",
		},
	})
	if err != nil {
		t.Error(err)
	}

	ids, err := QueryUserPubNotice(cg, "Grteen-test", 10, 0)
	if err != nil {
		t.Error(err)
	}

	pubs, err := QueryPubNotice(cg, ids)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(pubs[0])
}
