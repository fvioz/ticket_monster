package domain

import (
	"encoding/json"
	"fetcher/configs"
	"fetcher/libs"

	"github.com/gocraft/work"
)

type Repository struct {
	aggregator *Aggregator
}

func NewRepository() *Repository {
	return &Repository{
		aggregator: NewAggregator(),
	}
}

func (r *Repository) SaveEvent(event BasePlan) (*work.Job, error) {
	globalConfig := configs.GlobalConfigInstance()

	logger := libs.LoggerInstance()
	redis := libs.RedisInstance()

	jsonData, err := json.Marshal(event)
	if err != nil {
		logger.Error("Failed to marshal event to JSON", err)
		return nil, err
	}

	enqueuer := work.NewEnqueuer(redis.Config().ApplicationNamespace, redis.Pool())

	job, err := enqueuer.Enqueue(globalConfig.RedisNewEventName, work.Q{"json_event": jsonData})
	if err != nil {
		logger.Error("Failed to enqueue the event", err)
	}

	return job, nil
}
