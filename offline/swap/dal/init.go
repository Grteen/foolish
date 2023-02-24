package dal

import (
	"be/offline/swap/dal/db"
	"be/offline/swap/dal/rdb"
)

// Init mysql
func Init() {
	db.MySQLInit()
	rdb.RedisInit()
}
