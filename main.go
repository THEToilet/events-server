package main

import (
	"./pkg/usercase"
	"./pkg/domain/repository" )

func main() {

	userRepository := repository.NewUserRepository()

	userUseCase := usercase.NewUserUseCase(userRepository)

	s :=
}
