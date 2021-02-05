package main

import (
	"fmt"
	"github.com/THEToilet/events-server/pkg/config"
	"github.com/THEToilet/events-server/pkg/gateway"
	"github.com/THEToilet/events-server/pkg/gateway/database"
	"github.com/THEToilet/events-server/pkg/server"
	"github.com/THEToilet/events-server/pkg/usecase"
)

func main() {

	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		fmt.Print("unko")
	}

	userRepository := gateway.NewUserRepository(sqlDB)
	eventRepository := gateway.NewEventRepository(sqlDB)
	tagRepository := gateway.NewTagRepository(sqlDB)

	userUseCase := usecase.NewUserUseCase(userRepository)
	eventUseCase := usecase.NewEventUseCase(eventRepository)
	tagUseCase := usecase.NewTagUseCase(tagRepository)
	authUseCase := usecase.NewAuthUseCase()

	s := server.NewServer(userUseCase, eventUseCase, tagUseCase, authUseCase)

	s.Start(config.Port())
}
