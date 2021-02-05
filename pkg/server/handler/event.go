package handler

import (
	"github.com/THEToilet/events-server/pkg/usercase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EventHandler struct {
	eventUseCase *usercase.EventUseCase
}

func NewEventHandler(eventUseCase *usercase.EventUseCase) *EventHandler {
	return &EventHandler{
		eventUseCase: eventUseCase,
	}
}

func (h *EventHandler) GetEvent(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := h.eventUseCase.GetUser(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &userResponse{
		Mail: user.Mail,
	})
}
func (h *EventHandler) PutEvent(c echo.Context) error    {}
func (h *EventHandler) PostEvent(c echo.Context) error   {}
func (h *EventHandler) DeleteEvent(c echo.Context) error {}

type eventResponse struct {
	Mail string `json:"main"`
}
