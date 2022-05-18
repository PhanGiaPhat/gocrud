package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/PhanGiaPhat/gocrud/pkg/handler"
	"github.com/PhanGiaPhat/gocrud/pkg/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server interface {
	Start()
}

type svr struct {
	server *echo.Echo
}

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if cv.validator.Struct(i) != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, cv.validator.Struct(i).Error())
	}
	return nil
}

func (s *svr) Start() {
	go func() {
		if err := s.server.Start(":8080"); err != nil && err != http.ErrServerClosed {
			s.server.Logger.Fatal("shutting down the server, ", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		s.server.Logger.Fatal(err)
	}
}

func NewServer(wr repository.MessageRepository) Server {
	e := echo.New()
	validate := validator.New()

	e.Validator = &CustomValidator{validator: validate}
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := handler.NewHandler(wr)
	e.POST("/messages", h.CreateMessage)
	e.GET("/messages/:id", h.GetMessage)
	e.GET("/messages", h.ListMessage)
	return &svr{server: e}
}
