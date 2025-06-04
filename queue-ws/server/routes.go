package server

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  128,
	WriteBufferSize: 128,
}

func (s *Server) routes() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	s.echo.GET("/metrics", echoprometheus.NewHandler())
}
