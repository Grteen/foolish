package rdb

import (
	"be/pkg/constants"
	"context"
	"log"
)

func EPrint(str string) {
	if err := RDB.LPush(context.Background(), constants.RdbErrorLogKey, str).Err(); err != nil {
		log.Print(err)
	}
}

func APrint(str string) {
	if err := RDB.LPush(context.Background(), constants.RdbAccessLogKey, str).Err(); err != nil {
		log.Print(err)
	}
}

func SPrint(str string) {
	if err := RDB.LPush(context.Background(), constants.RdbSwapLogKey, str).Err(); err != nil {
		log.Print(err)
	}
}
