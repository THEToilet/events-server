package handler

import (
	"github.com/THEToilet/events-server/pkg/domain/model"
	"github.com/THEToilet/events-server/pkg/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
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
	res := make([]eventResponse, len(events))
	for i, v := range events {
		res[i] = eventResponse{
			ID:          v.ID,
			PostedUser:  v.PostedUser,
			EventURL:    v.EventURL,
			DeadLine:    v.DeadLine,
			Description: v.Description,
			Tag:         v.Tag,
			CreatedAt:   v.UpdatedAt,
			UpdatedAt:   v.UpdatedAt,
		}
	}
	return c.JSON(http.StatusOK, res)
}

type eventResponse struct {
	ID          string       `json:"event_id"`
	PostedUser  string       `json:"posted_user"`
	EventURL    string       `json:"event_url"`
	DeadLine    time.Time    `json:"dead_line"`
	Description string       `json:"description"`
	Tag         []*model.Tag `json:"tags"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

//PutEvent は PUT /events/:id に対応するハンドラーです
func (h *EventHandler) PutEvent(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	event, err := h.eventUseCase.PutEvent(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, &eventResponse{
		ID:          event.ID,
		PostedUser:  event.PostedUser,
		EventURL:    event.EventURL,
		DeadLine:    event.DeadLine,
		Description: event.Description,
		Tag:         event.Tag,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	},
	)
}

//PostEvent は POST /events に対応するハンドラーです
func (h *EventHandler) PostEvent(c echo.Context) error {
	ctx := c.Request().Context()
	id := "uu" //c.Request().Body.Read("id").(string)
	_, err := h.eventUseCase.PostEvent(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

//DeleteEvent は DELETE /events/:id　に対応するハンドラーです
func (h *EventHandler) DeleteEvent(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	err := h.eventUseCase.DeleteEvent(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
