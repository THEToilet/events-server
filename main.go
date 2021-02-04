package main

import (
	"fmt"
	"github.com/THEToilet/events-server/pkg/gateway"
	"github.com/THEToilet/events-server/pkg/gateway/database"
	"github.com/THEToilet/events-server/pkg/usercase"
)

func main() {

	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		fmt.Print("unko")
	}


	userRepository := gateway.NewUserRepository(sqlDB)

	userUseCase := usercase.NewUserUseCase(userRepository)

	s :=
}
