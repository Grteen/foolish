package pack

import (
	"be/cmd/action/dal/rdb"
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
