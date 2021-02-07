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

//GetEvent は GET /events に対応するハンドラーです
func (h *EventHandler) GetEvent(c echo.Context) error {
	ctx := c.Request().Context()
	events, err := h.eventUseCase.GetEvent(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	for k := events {

	}

	return c.JSON(http.StatusOK, &userResponse{
		Mail: user.Mail,
	})
}

//PutEvent は PUT /events/:id に対応するハンドラーです
func (h *EventHandler) PutEvent(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	event, err := h.eventUseCase.GetEvent(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

//PostEvent は POST /events に対応するハンドラーです
func (h *EventHandler) PostEvent(c echo.Context) error {
	ctx := c.Request().Context()
	event, err := h.eventUseCase.GetEvent(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

//DeleteEvent は DELETE /events/:id　に対応するハンドラーです
func (h *EventHandler) DeleteEvent(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	event, err := h.eventUseCase.DeleteEvent(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

