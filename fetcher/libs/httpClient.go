package libs

import (
	"fetcher/configs"
	"net"
	"net/http"
	"time"
)

func NewHttpClient() http.Client {
	config := configs.NewHttpClientConfig()

	netTransport := http.Transport{
		Dial: (&net.Dialer{
			Timeout:   time.Duration(config.TransportTimeout) * time.Second,
			KeepAlive: time.Duration(config.TransportKeepAlive) * time.Second,
		}).Dial,
		TLSHandshakeTimeout: time.Duration(config.TransportTLSHandshakeTimeout) * time.Second,
	}

	return http.Client{
		Transport: &netTransport,
		Timeout:   time.Duration(config.ClientTimeout) * time.Second,
	}
}
