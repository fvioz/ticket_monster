package storage

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Client is a middleman between the websocket connection and the storage.
type Client struct {
	// Unix timestamp ID of the client.
	id int64

	// The key used to identify the client in the storage.
	key string

	// The websocket connection.
	conn *websocket.Conn
}

func (c *Client) redisID() string {
	return fmt.Sprint(c.id)
}

func (c *Client) redisScore() float64 {
	return float64(c.id)
}

func (c *Client) redisKey() string {
	return c.key
}

func (c *Client) getConn() *websocket.Conn {
	return c.conn
}
