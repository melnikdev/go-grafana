package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/database"
)

type PublicHandler struct {
	db database.Service
}

func NewPublicHandler(db database.Service) *PublicHandler {

	return &PublicHandler{
		db: db,
	}
}

func (h PublicHandler) HelloWorld(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
