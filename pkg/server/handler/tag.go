package handler

import (
	"github.com/THEToilet/events-server/pkg/log"
	"github.com/THEToilet/events-server/pkg/usecase"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type TagHandler struct {
	tagUseCase *usecase.TagUseCase
}

func NewTagHandler(tagUseCase *usecase.TagUseCase) *TagHandler {
	return &TagHandler{
		tagUseCase: tagUseCase,
	}
}

//GetTag は　GET /events/tags に対応するハンドラーです
func (h *TagHandler) GetTag(c echo.Context) error {
	ctx := c.Request().Context()
	tags, err := h.tagUseCase.GetTag(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	res := make([]tagResponse, len(tags))
	for i, v := range tags {
		res[i] = tagResponse{
			TagId:     v.TagID,
			TagName:   v.TagName,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}
	return c.JSON(http.StatusOK, res)
}

type tagResponse struct {
	TagId     string    `json:"tag_id"`
	TagName   string    `json:"tag_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//PostTag は POST /events/tags に対応するハンドラーです
func (h *TagHandler) PostTag(c echo.Context) error {
	logger := log.New()
	logger.Info("POST /events/tags")
	ctx := c.Request().Context()
	param := new(tagRequest)
	if err := c.Bind(param); err != nil {
		logger.Error("request body can not unmarshal", zap.Error(err))
		return c.NoContent(http.StatusBadRequest)
	}
	name := param.TagName
	logger.Info("name = " + name)
	_, err := h.tagUseCase.PostTag(ctx, name)
	if err != nil {
		logger.Error("post tag failed", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

type tagRequest struct {
	TagName string `json:"tag_name"`
}
