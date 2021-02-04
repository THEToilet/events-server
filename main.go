package main

import (
	"fmt"
	"github.com/THEToilet/events-server/pkg/usercase"

	"github.com/THEToilet/events-server/pkg/gateway/database"
	"github.com/THEToilet/events-server/pkg/domain/repository"
)

func main() {

	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		fmt.Print("unko")
	}


	userRepository := repository.NewUserRepository()

	userUseCase := usercase.NewUserUseCase(userRepository)

	s :=
}
