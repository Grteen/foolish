package dal

import (
	"be/cmd/log/dal/rdb"
)

func Init() {
	rdb.RedisInit()
}
