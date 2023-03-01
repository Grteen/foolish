package dal

import (
	"be/cmd/comment/dal/db"
	"be/cmd/comment/dal/rdb"
)

func Init() {
	db.MySQLInit()
	rdb.RedisInit()
}
