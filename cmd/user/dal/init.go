package dal

import (
	"be/cmd/user/dal/db"
	"be/cmd/user/dal/rdb"
)

// Init mysql and redis
func Init() {
	db.MySQLInit()
	rdb.RedisInit()
}
