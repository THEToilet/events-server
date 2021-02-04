package handler

import (
	"github.com/THEToilet/events-server/pkg/usercase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase *usercase.UserUseCase
}

func NewUserHandler(userUseCase *usercase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetUser(echo.Context) error {

}
