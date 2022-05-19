package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/PhanGiaPhat/gocrud/pkg/model"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func (h *handler) EditMessage(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 0, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid message id")
	}

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

	w, err := h.wr.Update(uint(id), wm)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := &CreateMessageResponse{
		ID:      w.ID,
		Type:    w.Type,
		Message: w.Message,
	}
	return c.JSON(http.StatusOK, resp)
}
