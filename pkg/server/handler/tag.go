package handler

import (
	"github.com/THEToilet/events-server/pkg/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TagHandler struct {
	tagUseCase *usecase.TagUseCase
}

func NewTagHandler(tagUseCase *usecase.TagUseCase) *TagHandler {
	return &TagHandler{
		tagUseCase: tagUseCase,
	}
}

func (h *TagHandler) GetTagList(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := h.tagUseCase.GetTagList(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &userResponse{
		Mail: user.Mail,
	})
}

type tagResponse struct {
}


func (h *TagHandler) PostTag(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := h.tagUseCase.PostTag(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &userResponse{
		Mail: user.Mail,
	})
}
