package server

import (
	"context"
	"fmt"
	"net/http"
	"pocsrv/internal/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	ctx  context.Context
	cfg  config.Config
	echo *echo.Echo
}

func New(cfg config.Config) (*Server, error) {
	return &Server{
		ctx:  context.Background(),
		cfg:  cfg,
		echo: echo.New(),
	}, nil
}

func (s *Server) onStart() {
	s.echo.HideBanner = true
	s.echo.Debug = s.cfg.Debug
}

func (s *Server) ListenAndServe() error {

	s.onStart()

	s.echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	s.echo.GET("/ping", s.pingHandler)
	s.echo.POST("/exec", s.execHandler)

	pocs := s.echo.Group("/s")
	{
		pocs.Static("/", s.cfg.StaticPath)
	}

	if s.cfg.UseSSL {
		return s.echo.StartTLS(fmt.Sprintf(":%d", s.cfg.HTTPPort), s.cfg.SSLCertPath, s.cfg.SSLKeyPath)
	}

	return s.echo.Start(fmt.Sprintf(":%d", s.cfg.HTTPPort))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
