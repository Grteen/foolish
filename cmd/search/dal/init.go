package dal

import "be/cmd/search/dal/db"

func Init() {
	db.MySQLInit()
}
