package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/request"
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
	id := c.Param("id")
	movie, err := h.service.FindById(id)

	if err != nil {
		return c.String(404, err.Error())
	}
	return c.String(200, movie.Title)
}

func (h MovieHandler) CreateMovie(c echo.Context) error {
	req := request.CreateMovieRequest{}
	c.Bind(&req)

	id, err := h.service.Create(req)

	if err != nil {
		return c.String(500, err.Error())
	}
	return c.String(200, "CreateMovie id = "+id)
}

func (h MovieHandler) UpdateMovie(c echo.Context) error {
	id := c.Param("id")

	req := request.UpdateMovieRequest{}
	err := h.service.Update(id, req)

	if err != nil {
		return c.String(500, err.Error())
	}

	return c.String(200, "UpdateMovie id = "+id)
}

func (h MovieHandler) DeleteMovie(c echo.Context) error {
	id := c.Param("id")

	err := h.service.Delete(id)

	if err != nil {
		return c.String(500, err.Error())
	}

	return c.String(200, "DeleteMovie id = "+id)
}
