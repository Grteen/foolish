package dal

import "be/cmd/user/dal/db"

// Init mysql and redis
func Init() {
	db.MySQLInit()
	db.RedisInit()
}
