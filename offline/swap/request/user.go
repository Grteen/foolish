package request

import (
	"be/offline/swap/handlers"
	"be/offline/swap/service"
	"context"
)

func UserRequest() {
	var offset int32 = 0
	var limit int32 = 50
	for {
		users, _ := service.NewUserService(context.Background()).QueryUser(limit, offset)
		if len(users) == 0 {
			break
		}
		for _, user := range users {
			handlers.CreateUserHandler().Handle(user)
		}
		offset += limit
	}
}
