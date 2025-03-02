package apiroute

import (
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/melnikdev/go-grafana/internal/transport/rest"
)

func InitPublicRoutes(e *echo.Echo, db database.IdbService) {
	h := rest.NewPublicHandler(db)
	e.GET("/hello", h.HelloWorld)
}
