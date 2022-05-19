package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func (h *handler) DeleteMessage(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 0, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid message id")
	}
	_, err = h.wr.Delete(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Deleted Succed !")
}
