package handlers

import (
	"log"
	"queue-ws/storage"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 30 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024
)

type positionMessage struct {
	Position int64 `json:"pos"`
}

type WebSocketHandler struct {
	// Key is the unique identifier for the client.
	key string

	// Handler is the main handler for the websocket connection.
	conn *websocket.Conn

	// Storage is the storage instance that manages the clients.
	storage *storage.Storage
}

func (h *Handler) WebSocket(key string, conn *websocket.Conn, storage *storage.Storage) error {

	w := &WebSocketHandler{
		key:     key,
		conn:    conn,
		storage: storage,
	}

	go w.StartClient()

	return nil
}

func (w *WebSocketHandler) StartClient() {
	client := w.storage.SubscribeClient(w.key, w.conn)

	defer func() {
		w.storage.UnSubscribeClient(client)
		w.conn.Close()
	}()

	go w.SendPosition()

	w.conn.SetReadLimit(maxMessageSize)
	w.conn.SetReadDeadline(time.Now().Add(pongWait))
	w.conn.SetPongHandler(func(string) error {
		w.conn.SetReadDeadline(time.Now().Add(pongWait))

		return nil
	})

	for {
		_, _, err := w.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
				w.storage.UnSubscribeClient(client)
				w.conn.Close()
			}
			break
		}
	}
}

func (w *WebSocketHandler) SendPosition() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Reset(pingPeriod)
		ticker.Stop()
	}()

	for range ticker.C {
		msg := positionMessage{
			Position: 0,
		}

		w.conn.SetWriteDeadline(time.Now().Add(writeWait))
		err := w.conn.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
