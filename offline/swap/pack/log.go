package pack

import (
	"be/offline/swap/dal/rdb"
	"strconv"
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

// 在log中添加ID
func SuffixID(id int32) string {
	return " ID = " + strconv.Itoa(int(id))
}
