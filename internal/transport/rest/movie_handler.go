package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/service"
)

type MovieHandler struct {
	service service.IMovieService
}

func NewMovieHandler(service service.IMovieService) *MovieHandler {
	return &MovieHandler{
		service: service,
	}
}

func (h MovieHandler) GetMovie(c echo.Context) error {
	movie, err := h.service.FindById()

	if err != nil {
		return c.String(200, err.Error())
	}
	return c.String(200, movie.Title)
}
