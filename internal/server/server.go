package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/melnikdev/go-grafana/internal/config"
	"github.com/melnikdev/go-grafana/internal/database"
	apiroute "github.com/melnikdev/go-grafana/internal/transport/rest/route"
	webroute "github.com/melnikdev/go-grafana/internal/transport/web/route"
)

type Server struct {
	port int

	db database.IdbService
}

func NewServer(db database.IdbService, config *config.Server) *http.Server {
	NewServer := &Server{
		port: config.Port,

		db: db,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.initServer(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

func (s *Server) initServer() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(echoprometheus.NewMiddleware("go_grafana"))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	apiroute.InitPublicRoutes(e, s.db)
	apiroute.InitMovieRoutes(e, s.db)
	webroute.InitWebRoutes(e, s.db)

	return e
}
