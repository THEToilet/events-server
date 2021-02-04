package handler

import (
	"github.com/THEToilet/events-server/pkg/usercase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	userUseCase *usercase.UserUseCase
}

func NewUserHandler(userUseCase *usercase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetUser(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := h.userUseCase.GetUser(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &userResponse{
		Mail: user.Mail,
	})
}

type userResponse struct {
	Mail string `json:"main"`
}
