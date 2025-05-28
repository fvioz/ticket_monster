package workers

import (
	"context"
	"fetcher/configs"
	"fetcher/events/api"
	"fetcher/events/domain"
	"fetcher/libs"

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
	redis := libs.RedisInstance()

	globalConfig := configs.GlobalConfigInstance()
	redisConfig := redis.Config()

	logger.Info("Setting up fetch_events worker pool",
		zap.String("namespace", redisConfig.ApplicationNamespace),
		zap.String("crontab", redisConfig.EventCrontab),
	)

	jobOptions := work.JobOptions{
		MaxConcurrency: 1,
	}

	w.pool.PeriodicallyEnqueue(
		redisConfig.EventCrontab,
		globalConfig.RedisEventFetcherName,
	)

	w.pool.JobWithOptions(
		globalConfig.RedisEventFetcherName,
		jobOptions,
		(*WorkerContext).FetchEvents,
	)

	w.pool.Start()
}

func (w *WorkerContext) FetchEvents(job *work.Job) error {
	logger := libs.LoggerInstance()
	logger.Info("Starting to fetch events from API")

	apiRepository := api.NewRepository()
	domainRepository := domain.NewRepository()

	xmlEvents, err := apiRepository.GetEvents()
	if err != nil {
		logger.Error("Failed to fetch events from API", err)
		return err
	}

	aggregator := domain.NewAggregator()
	events := aggregator.FromApiToDomain(xmlEvents)

	for _, event := range events {
		go domainRepository.SaveEvent(event)
	}

	logger.Info("Finished fetching events from API")

	return nil
}
