package workers

import (
	"context"
	"encoding/json"
	"processors/configs"
	"processors/events/domain"
	"processors/events/persistence"
	"processors/libs"

	"github.com/gocraft/work"
	"go.uber.org/zap"
)

type EventWorker struct {
	pool          *work.WorkerPool
	workerContext WorkerContext
}

func (w *Worker) InitEventWorkers() *EventWorker {
	redis := libs.RedisInstance()
	config := redis.Config()

	return &EventWorker{
		pool: work.NewWorkerPool(
			w.context,
			config.Concurrency,
			config.ApplicationNamespace,
			redis.Pool(),
		),
		workerContext: w.context,
	}
}

func (w *EventWorker) Start(ctx context.Context) error {
	logger := libs.LoggerInstance()
	errorChan := make(chan error, 1)

	go w.run()

	select {
	case <-ctx.Done():
		logger.Info("Shutting down event worker")
		w.pool.Stop()
	case err := <-errorChan:
		logger.Error("Failed to start the event worker", zap.Error(err))
		return err
	}

	return nil
}

func (w *EventWorker) run() {
	logger := libs.LoggerInstance()
	logger.Info("Starting events worker")

	globalConfig := configs.GlobalConfigInstance()

	w.pool.Job(globalConfig.RedisNewEventName, (*WorkerContext).FetchNewEvent)
	w.pool.Start()
}

func (w *WorkerContext) FetchNewEvent(job *work.Job) error {
	logger := libs.LoggerInstance()
	logger.Info("Starting to process new event")

	jsonEvent := job.ArgString("json_event")
	if err := job.ArgError(); err != nil {
		return err
	}

	logger.Info("Parsed event from JSON",
		zap.String("raw", jsonEvent),
	)

	var persistenceEvent persistence.BasePlan
	json.Unmarshal([]byte(jsonEvent), &persistenceEvent)

	eventsAggregator := domain.NewAggregator()
	eventsRepository := domain.NewRepository()

	event := eventsAggregator.FromPersistenceBasePlanToDomain(persistenceEvent)

	eventsRepository.SaveEvent(event)

	logger.Info("Finished to process new event")

	return nil
}
