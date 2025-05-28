package libs

import (
	"fetcher/configs"
	"sync"

	"github.com/gomodule/redigo/redis"
)

type RedisClient struct {
	config *configs.RedisConfig
	pool   *redis.Pool
}

var redisClient *RedisClient
var redisClientOnce sync.Once

func RedisInstance() *RedisClient {
	redisClientOnce.Do(func() {
		config := configs.NewRedisConfig()

		var pool = &redis.Pool{
			MaxActive: config.MaxActive,
			MaxIdle:   config.MaxIdle,
			Wait:      config.Wait,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", config.Address(),
					redis.DialUsername(config.Username),
					redis.DialPassword(config.Password),
				)
			},
		}

		redisClient = &RedisClient{
			config: config,
			pool:   pool,
		}
	})

	return redisClient
}

func (rc *RedisClient) Config() *configs.RedisConfig {
	return rc.config
}

func (rc *RedisClient) Pool() *redis.Pool {
	return rc.pool
}
