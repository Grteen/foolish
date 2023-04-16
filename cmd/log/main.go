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
	go rdb.WriteSwapLog()

	// 阻塞主线程
	select {}
}
