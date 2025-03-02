package webroute

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/melnikdev/go-grafana/internal/repository"
	"github.com/melnikdev/go-grafana/internal/service"
	"github.com/melnikdev/go-grafana/internal/transport/web"
)

func InitWebRoutes(e *echo.Echo, db database.IdbService) {
	v := validator.New()
	r := repository.NewMovieRepository(db)
	s := service.NewMovieService(r, v)
	h := web.NewWebHandler(s)

	e.GET("/", h.GetHome)
	e.GET("/movies", h.GetTop5Movie)

}
