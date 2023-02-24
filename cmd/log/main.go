package main

import (
	"be/cmd/log/dal"
	rdb "be/cmd/log/dal/rdb"
	"be/cmd/log/pack"
)

func Init() {
	dal.Init()
	pack.LogInit()
}

func main() {
	Init()

	go rdb.WriteALog()
	go rdb.WriteELog()

	// 阻塞主线程
	select {}
}
