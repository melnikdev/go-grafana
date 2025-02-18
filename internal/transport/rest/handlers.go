package rest

import (
	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
