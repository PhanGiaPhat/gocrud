package handler

import (
	"github.com/PhanGiaPhat/gocrud/pkg/repository"
	"github.com/labstack/echo"
)

type Handler interface {
	CreateMessage(c echo.Context) error
	ListMessage(c echo.Context) error
	GetMessage(c echo.Context) error
	EditMessage(c echo.Context) error
	DeleteMessage(c echo.Context) error
}

type handler struct {
	wr repository.MessageRepository
}

func NewHandler(wr repository.MessageRepository) Handler {
	return &handler{wr: wr}
}
