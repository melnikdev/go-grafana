package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/database"
)

type PublicHandler struct {
	db database.IdbService
}

func NewPublicHandler(db database.IdbService) *PublicHandler {

	return &PublicHandler{
		db: db,
	}
}

func (h PublicHandler) HelloWorld(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
