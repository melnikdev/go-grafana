package rest

import (
	"github.com/labstack/echo/v4"
)

func InitPublicRoutes(e *echo.Echo) {
	e.GET("/", HelloWorld)
}
