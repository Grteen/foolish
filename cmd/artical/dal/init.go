package dal

import "be/cmd/user/dal/db"

// init MySQL and Redis
func Init() {
	db.MySQLInit()
	db.RedisInit()
}
