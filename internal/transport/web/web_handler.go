package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/melnikdev/go-grafana/cmd/web/views/home"
	"github.com/melnikdev/go-grafana/cmd/web/views/layout"
	"github.com/melnikdev/go-grafana/cmd/web/views/movie"
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

func (h WebHandler) GetTopMovies(c echo.Context) error {
	movies, err := h.service.GetTopMovies(20)

	if err != nil {
		return c.String(404, err.Error())
	}

	return render(c, http.StatusOK, layout.Base(movie.List(movies), "Movies", "Movies Page"))
}

func (h WebHandler) GetHome(c echo.Context) error {
	return render(c, http.StatusOK, layout.Base(home.Home(), "Home", "Home Page"))
}

func render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
