package db

import "testing"

func TestSystemNotify(t *testing.T) {
	MySQLInit()
	DB.AutoMigrate(&SystemNotify{})
}
