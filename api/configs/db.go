package configs

import (
	"log"
	"net"

	"github.com/ilyakaznacheev/cleanenv"
)

type DbConfig struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	Username string `env:"DB_USERNAME" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:""`
	Database string `env:"DB_NAME" env-default:"events"`
	PoolSize int    `env:"DB_POOL_SIZE" env-default:"10"`
}

func NewDbConfig() *DbConfig {
	var config DbConfig

	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatal("Failed to initialize DB config: ", err)
	}

	return &config
}
func (rc *DbConfig) Address() string {
	return net.JoinHostPort(
		rc.Host,
		rc.Port,
	)
}
