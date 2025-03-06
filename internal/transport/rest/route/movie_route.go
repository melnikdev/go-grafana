package apiroute

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/melnikdev/go-grafana/internal/repository"
	"github.com/melnikdev/go-grafana/internal/service"
	"github.com/melnikdev/go-grafana/internal/transport/rest"
)

func InitMovieRoutes(e *echo.Echo, db database.IdbService) {
	v := validator.New()
	r := repository.NewMovieRepository(db)
	s := service.NewMovieService(r, v)
	h := rest.NewMovieHandler(s)

	e.GET("/movie/:id", h.GetMovie)
	e.POST("/movie", h.CreateMovie)
	e.PUT("/movie/:id", h.UpdateMovie)
	e.DELETE("/movie/:id", h.DeleteMovie)

}
