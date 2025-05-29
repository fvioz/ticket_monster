package main

import (
	"os"
	"os/signal"
	"processors/events/persistence"
	"processors/libs"
	"processors/workers"
	"syscall"

	"github.com/go-pg/pg/v10/orm"
)

func createSchema() error {
	db := libs.DBInstance()

	models := []interface{}{
		(*persistence.BasePlan)(nil),
		(*persistence.Plan)(nil),
		(*persistence.Zone)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	logger := libs.LoggerInstance()
	logger.Info("Starting service ...")

	createSchema()

	w := workers.InitWorkers()
	w.StartWorkers()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan

	logger.Info("Service shutting down gracefully")
}
