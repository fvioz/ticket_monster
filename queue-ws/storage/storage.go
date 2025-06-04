package storage

import (
	"context"
	"queue-ws/libs"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Storage struct {
	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Memory storage
	memoryStorage *MemoryStorage

	// Redis storage
	redisStorage *RedisStorage
}

func NewStorage() *Storage {
	return &Storage{
		register:      make(chan *Client),
		unregister:    make(chan *Client),
		memoryStorage: NewMemoryStorage(context.Background()),
		redisStorage:  NewRedisStorage(context.Background()),
	}
}

func (storage *Storage) Start(ctx context.Context) error {
	logger := libs.LoggerInstance()
	errorChan := make(chan error, 1)

	go storage.run(errorChan)

	select {
	case <-ctx.Done():
		logger.Info("Shutting down the server")
	case err := <-errorChan:
		logger.Fatal("Failed to start the storage", zap.Error(err))
		return err
	}

	return nil

}

func (storage *Storage) ClientRank(client *Client) (int64, error) {
	return storage.redisStorage.ClientRank(client)
}

func (storage *Storage) SubscribeClient(key string, conn *websocket.Conn) *Client {
	client := &Client{
		id:   time.Now().UnixNano(),
		key:  key,
		conn: conn,
	}

	storage.register <- client

	return client
}

func (storage *Storage) UnSubscribeClient(client *Client) {
	storage.unregister <- client
}

func (storage *Storage) run(errorChan chan<- error) {
	defer close(errorChan)

	go storage.memoryStorage.Start()

	for {
		select {
		case client := <-storage.register:
			// If the client is not already registered in memory, register it.
			if ok := storage.memoryStorage.clientExist(client); !ok {
				if rank, _ := storage.redisStorage.SubscribeClient(client); rank > 0 {
					storage.memoryStorage.register <- client
				}
			}
		case client := <-storage.unregister:
			if ok := storage.memoryStorage.clientExist(client); ok {
				if err := storage.redisStorage.UnSubscribeClient(client); err != nil {
					storage.memoryStorage.unregister <- client
				}
			}
		}
	}
}
