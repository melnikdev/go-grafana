package apiroute

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/database"
	"github.com/melnikdev/go-grafana/internal/repository"
	"github.com/melnikdev/go-grafana/internal/service"
	"github.com/melnikdev/go-grafana/internal/transport/rest"
)

func InitAuthRoutes(e *echo.Echo, db database.IdbService) {
	v := validator.New()
	r := repository.NewUserRepository(db)
	s := service.NewAuthService(r, v)
	h := rest.NewAuthHandler(s)

	e.POST("/register", h.RegisterUser)

}
