package models

import (
	"github.com/gorilla/websocket"
	"github.com/xairline/x-gpt/utils"
	"sync"
)

type Client struct {
	Id     string
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
	Logger utils.Logger
}

type Hub struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
	mu         sync.Mutex
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
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
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

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger.Errorf("error: %v", err)
			}
			break
		}
		c.Logger.Infof("Client %s received: %s", c.Id, message)
		c.Send <- message
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
			c.Logger.Infof("Client %s sent: %s", c.Id, message)
			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

// Lock locks the Hub. It should be used when accessing or modifying the Hub's data.
func (h *Hub) Lock() {
	h.mu.Lock()
}

// Unlock unlocks the Hub. It should be called after Lock when the Hub's data manipulation is done.
func (h *Hub) Unlock() {
	h.mu.Unlock()
}
