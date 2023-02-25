package main

import (
	"be/offline/swap/dal"
	"be/offline/swap/request"
)

func main() {
	dal.Init()
	request.UserRequest()
}
