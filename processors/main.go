package main

import (
	"os"
	"os/signal"
	"processors/libs"
	"processors/workers"
	"syscall"
)

func main() {
	logger := libs.LoggerInstance()
	logger.Info("Starting service ...")

	w := workers.InitWorkers()
	w.StartWorkers()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	logger.Info("Service shutting down gracefully")
}
