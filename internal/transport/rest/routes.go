package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/melnikdev/go-grafana/internal/repository"
	"github.com/melnikdev/go-grafana/internal/service"
)

func InitPublicRoutes(e *echo.Echo, db database.IdbService) {
	h := NewPublicHandler(db)
	e.GET("/", h.HelloWorld)
}

func InitMovieRoutes(e *echo.Echo, db database.IdbService) {
	rep := repository.NewMovieRepository(db)
	ser := service.NewMovieService(rep)
	h := NewMovieHandler(ser)

	e.GET("/movie/:id", h.GetMovie)
	e.POST("/movie", h.CreateMovie)
	e.PUT("/movie/:id", h.UpdateMovie)
	e.DELETE("/movie/:id", h.DeleteMovie)

}
