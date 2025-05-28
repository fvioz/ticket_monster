package configs

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpClientConfig struct {
	TransportTimeout             uint `env:"HTTP_CLIENT_TRANSPORT_TIMEOUT" env-default:"5"`
	TransportKeepAlive           uint `env:"HTTP_CLIENT_TRANSPORT_KEEP_ALIVE" env-default:"30"`
	TransportTLSHandshakeTimeout uint `env:"HTTP_CLIENT_TRANSPORT_TLS_HAND_SHAKE_TIMEOUT" env-default:"5"`
	ClientTimeout                uint `env:"HTTP_CLIENT_CLIENT_TIMEOUT" env-default:"10"`
}

func NewHttpClientConfig() *HttpClientConfig {
	var config HttpClientConfig

	if err := cleanenv.ReadEnv(&config); err != nil {
		log.Fatal("Failed to initialize Http config: ", err)
	}

	return &config
}
