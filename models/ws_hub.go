package models

import (
	"github.com/gorilla/websocket"
	"github.com/xairline/x-gpt/utils"
	"sync"
	"time"
)

type Client struct {
	Id           string
	Hub          *Hub
	Conn         *websocket.Conn
	Send         chan []byte
	Logger       utils.Logger
	mu           sync.Mutex
	locked       bool
	LastActivity time.Time
}

type Hub struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	go h.Cleanup()
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
				_ = client.Conn.Close()
			}
		case message := <-h.Broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:

					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

func (h *Hub) Cleanup() {
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			counter := 0
			logger := utils.NewLogger()
			logger.Infof("Active clients before: %d", counter)
			for client := range h.Clients {
				if time.Since(client.LastActivity) > 20*time.Minute {
					client.Logger.Infof("Client: %s inactive for 20 minutes, closing connection", client.Id)
					close(client.Send) // This will cause WritePump to close the connection
				} else {
					counter++
				}
			}
			logger.Infof("Active clients after: %d", counter)
		}
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			c.Logger.Infof("Client: %s sent: %s", c.Id, message)
			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// Lock locks the Hub. It should be used when accessing or modifying the Hub's data.
func (c *Client) Lock() {
	c.mu.Lock()
	c.locked = true
	c.Logger.Infof("Client: %s locked", c.Id)
	// Schedule an automatic unlock after the timeout.
	time.AfterFunc(5*time.Second, func() {
		if c.locked {
			c.Logger.Infof("Client: %s auto unlock", c.Id)
			c.Unlock()
		}
	})
}

// Unlock unlocks the Hub. It should be called after Lock when the Hub's data manipulation is done.
func (c *Client) Unlock() {
	c.locked = false
	c.mu.Unlock()
	c.Logger.Infof("Client: %s unlocked", c.Id)
}
