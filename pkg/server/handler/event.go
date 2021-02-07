package handler

import (
	"github.com/THEToilet/events-server/pkg/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type EventHandler struct {
	eventUseCase *usecase.EventUseCase
}

func NewEventHandler(eventUseCase *usecase.EventUseCase) *EventHandler {
	return &EventHandler{
		eventUseCase: eventUseCase,
	}
}

func (h *EventHandler) GetEvent(c echo.Context) error {
	ctx := c.Request().Context()
	event, err := h.eventUseCase.GetEvent(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &userResponse{
		Mail: user.Mail,
	})
}
func (h *EventHandler) PutEvent(c echo.Context) error {
	ctx := c.Request().Context()
	event, err := h.eventUseCase.GetEvent(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, 9)
}
func (h *EventHandler) PostEvent(c echo.Context) error {
	ctx := c.Request().Context()
	event, err := h.eventUseCase.GetEvent(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, 9)
}

func (h *EventHandler) DeleteEvent(c echo.Context) error {
	ctx := c.Request().Context()
	event, err := h.eventUseCase.GetEvent(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, 9)
}

type eventResponse struct {
	Mail string `json:"main"`
}
