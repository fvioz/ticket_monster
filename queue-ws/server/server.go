package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"queue-ws/configs"
	"queue-ws/handlers"
	"queue-ws/libs"
	"queue-ws/storage"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type Server struct {
	echo     *echo.Echo
	handlers *handlers.Handler
	storage  *storage.Storage
}

func NewServer(storage *storage.Storage, handlers *handlers.Handler) *Server {
	return &Server{
		echo:     echo.New(),
		handlers: handlers,
		storage:  storage,
	}
}

func (s *Server) Start(ctx context.Context) error {
	logger := libs.LoggerInstance()
	errorChan := make(chan error, 1)

	s.routes()
	s.middleware()

	go s.run(errorChan)

	select {
	case <-ctx.Done():
		logger.Info("Shutting down the server")
		if shutdownErr := s.echo.Shutdown(ctx); shutdownErr != nil {
			logger.Error("Error shutting down the server", zap.Error(shutdownErr))
			return shutdownErr
		}
	case err := <-errorChan:
		logger.Fatal("Failed to start HTTP server", zap.Error(err))
		return err
	}

	return nil
}

func (s *Server) run(errorChan chan<- error) {
	defer close(errorChan)

	config := configs.GlobalConfigInstance()

	if err := s.echo.Start(
		net.JoinHostPort(
			config.Host,
			config.Port,
		),
	); err != nil && !errors.Is(err, http.ErrServerClosed) {
		errorChan <- err
	}
}
