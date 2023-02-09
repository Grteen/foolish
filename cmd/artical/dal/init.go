package dal

import (
	"be/cmd/artical/dal/db"
	"be/cmd/artical/dal/rdb"
)

// init MySQL and Redis
func Init() {
	db.MySQLInit()
	rdb.RedisInit()
}
