package handler

import (
	"github.com/THEToilet/events-server/pkg/usercase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TagHandler struct {
	tagUseCase *usercase.TagUseCase
}

func NewTagHandler(tagUseCase *usercase.TagUseCase) *TagHandler {
	return &TagHandler{
		tagUseCase: tagUseCase,
	}
}

func (h *TagHandler) GetTagList(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := h.userUseCase.GetUser(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &userResponse{
		Mail: user.Mail,
	})
}

type tagResponse struct {
}


func (h *TagHandler) PostTagList(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := h.userUseCase.GetUser(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, &userResponse{
		Mail: user.Mail,
	})
}
