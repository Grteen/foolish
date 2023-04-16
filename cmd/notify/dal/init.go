package dal

import (
	"be/cmd/notify/dal/db"
	"be/cmd/notify/dal/rdb"
)

// init MySQL and Redis
func Init() {
	db.MySQLInit()
	rdb.RedisInit()
}
