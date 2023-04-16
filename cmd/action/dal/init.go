package dal

import (
	"be/cmd/action/dal/db"
	"be/cmd/action/dal/rdb"
)

func Init() {
	db.MySQLInit()
	rdb.RedisInit()
}
