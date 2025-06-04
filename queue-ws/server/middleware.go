package server

import (
	"queue-ws/libs"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) middleware() {
	logger := libs.LoggerInstance()
	s.echo.Use(echozap.ZapLogger(logger))

	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	s.echo.Use(middleware.CORS())
	s.echo.Use(echoprometheus.NewMiddleware("queue-ws"))
}
