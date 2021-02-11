package main

import (
	"github.com/THEToilet/events-server/pkg/config"
	"github.com/THEToilet/events-server/pkg/gateway"
	"github.com/THEToilet/events-server/pkg/gateway/database"
	"github.com/THEToilet/events-server/pkg/log"
	"github.com/THEToilet/events-server/pkg/server"
	"github.com/THEToilet/events-server/pkg/usecase"
	"go.uber.org/zap"
)

func main() {

	logger := log.New()
	sqlDB, err := database.NewMySqlDB()
	if err != nil {
		logger.Fatal("mysql do not connect", zap.Error(err))
	}
	redisConn, err := database.NewRedis()
	if err != nil {
		logger.Fatal("mysql do not connect", zap.Error(err))
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

	s.Start("")//config.Port())
	logger.Info("Port: " + config.Port())
	logger.Info("server running......")
}
