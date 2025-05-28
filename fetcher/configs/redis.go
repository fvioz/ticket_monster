package configs

import (
	"log"
	"net"

	"github.com/ilyakaznacheev/cleanenv"
)

type RedisConfig struct {
	Host                 string `env:"REDIS_HOST" env-default:"localhost"`
	Port                 string `env:"REDIS_PORT" env-default:"6379"`
	Username             string `env:"REDIS_USERNAME" env-default:""`
	Password             string `env:"REDIS_PASSWORD" env-default:""`
	Concurrency          uint   `env:"REDIS_CONCURRENCY" env-default:"10"`
	MaxActive            int    `env:"REDIS_MAX_ACTIVE" env-default:"5"`
	MaxIdle              int    `env:"REDIS_MAX_IDLE" env-default:"5"`
	Wait                 bool   `env:"REDIS_WAIT" env-default:"true"`
	ApplicationNamespace string `env:"REDIS_APPLICATION_NAMESPACE" env-default:"fetcher"`
	EventCrontab         string `env:"REDIS_EVENT_CRONTAB" env-default:"*/30 * * * *"`
}

func NewRedisConfig() *RedisConfig {
	var config RedisConfig

	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatal("Failed to initialize global config: ", err)
	}

	return &config
}
func (rc *RedisConfig) Address() string {
	return net.JoinHostPort(
		rc.Host,
		rc.Port,
	)
}
