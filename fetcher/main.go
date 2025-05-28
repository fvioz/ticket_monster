package main

import (
	"fetcher/libs"
	"fetcher/workers"
	"os"
	"os/signal"
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
