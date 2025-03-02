package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/cmd/web/views"
	"github.com/melnikdev/go-grafana/internal/service"
)

type WebHandler struct {
	service service.IMovieService
}

func NewWebHandler(service service.IMovieService) *WebHandler {
	return &WebHandler{
		service: service,
	}
}

func (h WebHandler) GetTop5Movie(c echo.Context) error {
	movies, err := h.service.GetTop5Movie()

	if err != nil {
		return c.String(404, err.Error())
	}

	return render(c, http.StatusOK, views.MoviesList(movies))
}

func render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
