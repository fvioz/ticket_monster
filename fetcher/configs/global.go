package configs

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env                   string `env:"ENV" env-default:"development"`
	RedisEventFetcherName string `env:"ENV" env-default:"fetch_events"`
	RedisNewEventName     string `env:"ENV" env-default:"new_event"`
}

var config *Config
var configOnce sync.Once

func GlobalConfigInstance() *Config {
	configOnce.Do(func() {
		var c Config

		if err := cleanenv.ReadEnv(&c); err != nil {
			log.Fatal("Failed to initialize global config: ", err)
		}

		config = &c
	})

	return config
}

func (sc *Config) Debug() bool {
	return sc.Env == "development"
}
