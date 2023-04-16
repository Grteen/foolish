package rdb

import (
	"be/cmd/log/pack"
	"be/pkg/constants"
	"context"
)

// 将redis中的log写出来
func WriteELog() {
	var err error
	var res []string
	for {
		res, err = RDB.BRPop(context.Background(), 0, constants.RdbErrorLogKey).Result()
		if err != nil {
			pack.ELoger.Print(err)
		}
		pack.ELoger.Print(res[1])
	}
}

func WriteALog() {
	var err error
	var res []string
	for {
		res, err = RDB.BRPop(context.Background(), 0, constants.RdbAccessLogKey).Result()
		if err != nil {
			pack.ELoger.Print(err)
		}
		pack.ALoger.Print(res[1])
	}
}

func WriteSwapLog() {
	var err error
	var res []string
	for {
		res, err = RDB.BRPop(context.Background(), 0, constants.RdbSwapLogKey).Result()
		if err != nil {
			pack.ELoger.Print(err)
		}
		pack.SLoger.Print(res[1])
	}
}
