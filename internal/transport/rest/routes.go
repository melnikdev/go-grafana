package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/database"
)

func InitPublicRoutes(e *echo.Echo, db database.Service) {
	h := NewPublicHandler(db)
	e.GET("/", h.HelloWorld)
}
