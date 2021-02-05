package main

import (
	"fmt"
	"github.com/THEToilet/events-server/pkg/config"
	"github.com/THEToilet/events-server/pkg/gateway"
	"github.com/THEToilet/events-server/pkg/gateway/database"
	"github.com/THEToilet/events-server/pkg/server"
	"github.com/THEToilet/events-server/pkg/usercase"
)

func main() {

	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		fmt.Print("unko")
	}

	userRepository := gateway.NewUserRepository(sqlDB)
	eventRepository := gateway.NewEventRepository(sqlDB)
	tagRepository := gateway.NewTagRepository(sqlDB)

	userUseCase := usercase.NewUserUseCase(userRepository)
	eventUseCase := usercase.NewEventUseCase(eventRepository)
	tagUseCase := usercase.NewTagUseCase(tagRepository)

	s := server.NewServer(userUseCase, eventUseCase, tagUseCase)

	s.Start(config.Port())
}
