package rdb

import "testing"

func TestLog(t *testing.T) {
	RedisInit()
	APrint("test of access again")
	EPrint("test of error again")
}
