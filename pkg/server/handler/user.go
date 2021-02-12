package handler

import (
	"github.com/THEToilet/events-server/pkg/log"
	"github.com/THEToilet/events-server/pkg/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

//GetUser は GET /users に対応するハンドラーです
func (h *UserHandler) GetUser(c echo.Context) error {
	logger := log.New()
	logger.Info("ACCESS GET /users")
	ctx := c.Request().Context()
	user, err := h.userUseCase.GetUser(ctx)
	if err != nil {
		logger.Error("user not found", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &userResponse{
		Mail: user.UserMail,
	})
}

type userResponse struct {
	Mail string `json:"mail"`
}

//UserLogin は POST /users/login に対応するハンドラーです
func (h *UserHandler) UserLogin(c echo.Context) error {
	logger := log.New()
	ctx := c.Request().Context()
	_, err := h.userUseCase.UserLogin(ctx)
	if err != nil {
		logger.Error("login failed", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	sessionID := uuid.New().String()
	c.SetCookie(&http.Cookie{
		Name:  "session",
		Value: sessionID,
	})
	return c.NoContent(http.StatusOK)
}

//UserEntry は POST /users/entry に対応するハンドラーです
func (h *UserHandler) UserEntry(c echo.Context) error {
	logger := log.New()
	ctx := c.Request().Context()
	_, err := h.userUseCase.UserLogin(ctx)
	if err != nil {
		logger.Error("login failed", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	sessionID := uuid.New().String()
	c.SetCookie(&http.Cookie{
		Name:  "session",
		Value: sessionID,
	})
	return c.NoContent(http.StatusOK)
}
