package storage

import (
	"context"
	"queue-ws/libs"

	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

type RedisStorage struct {
	// Redis connection
	conn redis.Conn

	// Logger instance for logging operations
	logger *zap.Logger

	// Context for managing the lifecycle of the Redis instance
	ctx context.Context
}

func NewRedisStorage(ctx context.Context) *RedisStorage {
	client := libs.RedisInstance()
	logger := libs.LoggerInstance()

	conn := client.Pool().Get()
	defer conn.Close()

	return &RedisStorage{
		conn:   conn,
		logger: logger,
		ctx:    ctx,
	}
}

func (r *RedisStorage) ClientRank(client *Client) (int64, error) {
	rank, err := redis.Int64(r.conn.Do("ZRANK", client.redisKey(), client.redisID()))

	if err != nil {
		r.logger.Error("Failed to get client rank", zap.Error(err))
		return 0, err
	}

	r.logger.Info("Client rank",
		zap.String("client_id", client.redisID()),
		zap.String("key", client.redisKey()),
		zap.Int64("rank", rank),
	)

	return rank, nil
}

func (r *RedisStorage) SubscribeClient(client *Client) (int64, error) {
	rank, err := redis.Int64(r.conn.Do("ZADD", client.redisKey(), client.redisScore(), client.redisID()))

	if err != nil {
		r.logger.Error("Failed to subscribe the client", zap.Error(err))
		return 0, err
	}

	r.logger.Info("Client subscription",
		zap.String("client_id", client.redisID()),
		zap.String("key", client.redisKey()),
		zap.Int64("rank", rank),
	)

	return rank, nil
}

func (r *RedisStorage) UnSubscribeClient(client *Client) error {
	_, err := r.conn.Do("ZREM", client.redisKey(), client.redisID())

	if err != nil {
		r.logger.Error("Failed to unsubscribe the client", zap.Error(err))
		return err
	}

	r.logger.Info("Client unsubscription",
		zap.String("client_id", client.redisID()),
		zap.String("key", client.redisKey()),
	)

	return nil
}
