package dal

import (
	"be/cmd/smtp/dal/rdb"
)

func Init() {
	rdb.RedisInit()
}
