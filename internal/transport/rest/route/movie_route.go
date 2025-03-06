package apiroute

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/melnikdev/go-grafana/internal/middleware"
	"github.com/melnikdev/go-grafana/internal/repository"
	"github.com/melnikdev/go-grafana/internal/service"
	"github.com/melnikdev/go-grafana/internal/transport/rest"
)

func InitMovieRoutes(e *echo.Echo, db database.IdbService) {
	v := validator.New()
	r := repository.NewMovieRepository(db)
	s := service.NewMovieService(r, v)
	h := rest.NewMovieHandler(s)

	private := e.Group("/movie", middleware.AuthMiddleware)

	private.GET("/:id", h.GetMovie)
	private.POST("/", h.CreateMovie)
	private.PUT("/:id", h.UpdateMovie)
	private.DELETE("/:id", h.DeleteMovie)

}
