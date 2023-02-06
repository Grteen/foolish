package dal

import "be/cmd/artical/dal/db"

// init MySQL and Redis
func Init() {
	db.MySQLInit()
	db.RedisInit()
}
