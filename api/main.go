package main

import (
	"api/libs"
	"api/server"
	"context"

	"go.uber.org/zap"
)

func main() {
	logger := libs.LoggerInstance()

	server := server.Init()

	if err := server.Start(context.Background()); err != nil {
		logger.Fatal("Failed to start server:", zap.Error(err))
	}
}
