package main

import (
	"context"
	"os"
	"os/signal"
	"queue-ws/handlers"
	"queue-ws/libs"
	"queue-ws/server"
	"queue-ws/storage"
	"syscall"
)

func main() {

	logger := libs.LoggerInstance()
	logger.Info("Starting service ...")

	// Create a new context
	ctx := context.Background()

	// Initialize the handlers
	handlers := handlers.New()

	// Initialize the hub to store web socket clients
	storage := storage.NewStorage()
	storage.Start(ctx)

	// Initialize the server with the storage, handlers and system logger
	server := server.NewServer(storage, handlers)
	server.Start(ctx)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	logger.Info("Service shutting down gracefully")
}
