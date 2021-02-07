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

	res := make([]tagListResponse, len(tags))
	for i, v := range tags {
		res[i] = tagListResponse{
			TagId:     v.ID,
			TagName:   v.Name,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}
	return c.JSON(http.StatusOK, res)
}

type tagListResponse struct {
	TagId     string    `json:"tagId"`
	TagName   string    `json:"tagName"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//PostTag は POST /events/tags に対応するハンドラーです
func (h *TagHandler) PostTag(c echo.Context) error {
	ctx := c.Request().Context()
	err := h.tagUseCase.PostTag(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}
