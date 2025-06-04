package storage

type MemoryStorage struct {
	// Registered clients.
	clients map[string]map[*Client]int64

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		clients:    make(map[string]map[*Client]int64),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (memoryStorage *MemoryStorage) Start() {
	for {
		select {
		case client := <-memoryStorage.register:
			memoryStorage.clients[client.key][client] = client.id
		case client := <-memoryStorage.unregister:
			delete(memoryStorage.clients[client.key], client)
		}
	}
}

func (storage *MemoryStorage) clientExist(client *Client) bool {
	if _, ok := storage.clients[client.key][client]; ok {
		return true
	}

	return false
}
