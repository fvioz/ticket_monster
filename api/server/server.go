package server

import (
	"api/configs"
	"context"
	"net"

	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func Init() *Server {
	return &Server{
		echo: echo.New(),
	}
}

func (s *Server) Start(ctx context.Context) error {
	config := configs.GlobalConfigInstance()

	s.echo.Debug = config.Debug()

	s.routes()
	s.middleware()

	s.echo.Start(
		net.JoinHostPort(
			config.Host,
			config.Port,
		),
	)

	return nil
}
