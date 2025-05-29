package server

import (
	"api/libs"
	"strings"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) middleware() {
	logger := libs.LoggerInstance()
	s.echo.Use(echozap.ZapLogger(logger))

	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	s.echo.Use(echoprometheus.NewMiddleware("queue-ws"))
}
