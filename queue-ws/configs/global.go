package configs

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host string `env:"HOST" env-default:"127.0.0.1"`
	Port string `env:"PORT" env-default:"8080"`
	Env  string `env:"ENV" env-default:"development"`
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
