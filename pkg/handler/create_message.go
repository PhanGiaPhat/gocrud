package handler

import (
	"net/http"

	"github.com/PhanGiaPhat/gocrud/pkg/model"
	"github.com/labstack/echo"
)

type (
	CreateMessageRequest struct {
		Type    string `json:"type" validate:"required"`
		Message string `json:"message" validate:"required"`
	}

	CreateMessageResponse struct {
		ID      uint   `json:"id"`
		Type    string `json:"type"`
		Message string `json:"message"`
	}

	Error struct {
		Message string `json:"message"`
	}
)

func (h *handler) CreateMessage(c echo.Context) error {
	u := new(CreateMessageRequest)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	wm := model.Message{
		Type:    u.Type,
		Message: u.Message,
	}
	w, err := h.wr.Create(wm)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	resp := CreateMessageResponse{
		ID:      w.ID,
		Type:    w.Type,
		Message: w.Message,
	}
	return c.JSON(http.StatusOK, resp)
}
