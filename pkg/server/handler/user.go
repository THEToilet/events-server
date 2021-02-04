package handler

import "../../usercase"

type UserHandler struct {
	userUseCase *usercase.UserUseCase
}

func NewUserCase(userUseCase *usercase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetUser() error {

}