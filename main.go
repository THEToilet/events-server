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
	redisConn, err := database.NewRedis()
	if err != nil {
		fmt.Println("error")
	}

	userRepository := gateway.NewUserRepository(sqlDB)
	eventRepository := gateway.NewEventRepository(sqlDB)
	tagRepository := gateway.NewTagRepository(sqlDB)
	sessionRepository := gateway.NewSessionRepository(redisConn)

	userUseCase := usecase.NewUserUseCase(userRepository)
	eventUseCase := usecase.NewEventUseCase(eventRepository)
	tagUseCase := usecase.NewTagUseCase(tagRepository)
	authUseCase := usecase.NewAuthUseCase(sessionRepository)

	s := server.NewServer(userUseCase, eventUseCase, tagUseCase, authUseCase)

	s.Start(config.Port())
}
