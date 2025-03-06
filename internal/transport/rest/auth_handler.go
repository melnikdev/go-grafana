package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/internal/request"
	"github.com/melnikdev/go-grafana/internal/service"
)

type AuthHandler struct {
	service service.IAuthService
}

func NewAuthHandler(service service.IAuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h AuthHandler) RegisterUser(c echo.Context) error {

	var r request.RegisterUserRequest

	err := c.Bind(&r)
	if err != nil {
		return c.String(500, err.Error())
	}

	res, err := h.service.Register(r)

	if err != nil {
		return c.String(500, err.Error())
	}

	return c.String(200, res)
}

func (h AuthHandler) Login(c echo.Context) error {
	var r request.LoginUserRequest

	err := c.Bind(&r)
	if err != nil {
		return c.String(500, err.Error())
	}

	res, err := h.service.Login(r)

	if err != nil {
		return c.String(401, "Invalid credential")
	}

	return c.JSON(200, map[string]string{"token": res})
}
