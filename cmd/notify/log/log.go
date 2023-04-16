package log

import (
	"be/cmd/notify/dal/rdb"
)

func EPrint(str string) {
	rdb.EPrint(str)
}

func APrint(str string) {
	rdb.APrint(str)
}

func SPrint(str string) {
	rdb.SPrint(str)
}
