package handler

import (
	"github.com/THEToilet/events-server/pkg/usecase"
	"github.com/labstack/echo/v4"
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

//GetTagList は　GET /events/tags に対応するハンドラーです
func (h *TagHandler) GetTagList(c echo.Context) error {
	ctx := c.Request().Context()
	tags, err := h.tagUseCase.GetTagList(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	res := make([]tagResponse, len(tags))
	for i, v := range tags {
		res[i] = tagResponse{
			TagId:     v.ID,
			TagName:   v.Name,
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
	ctx := c.Request().Context()
	name := "uu"
	_, err := h.tagUseCase.PostTag(ctx, name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
