package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (h *handler) ListMessage(c echo.Context) error {
	var page, limit int = 1, 10
	var err error

	if p := c.QueryParam("page"); p != "" {
		fmt.Println(p)
		page, err = strconv.Atoi(p)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid page parameter")
		}
	}
	if p := c.QueryParam("limit"); p != "" {
		limit, err = strconv.Atoi(p)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "invalid limit parameter")
		}
	}

	offset := (page - 1) * limit

	ws, err := h.wr.List(offset, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	resp := make([]CreateMessageResponse, 0, len(ws))
	for _, w := range ws {
		resp = append(resp, CreateMessageResponse{
			ID:      w.ID,
			Type:    w.Type,
			Message: w.Message,
		})
	}

	return c.JSON(http.StatusOK, resp)
}
