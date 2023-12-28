package services

//go:generate mockgen -destination=../services/__mocks__/webSocket.go -package=mocks -source=webSocket.go

import "C"
import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/xairline/x-gpt/models"
	"github.com/xairline/x-gpt/utils"
	"log"
	"net/http"
	"sync"
)

var webSocketSvcLock = &sync.Mutex{}
var webSocketSvc WebSocketService

type WebSocketService interface {
	Upgrade(c *gin.Context, clientId string)
	IsClientExist(clientId string) bool
	SendWsMsgByClientId(clientId string, message string) error
}

type webSocketService struct {
	Logger utils.Logger
	Hub    *models.Hub
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (ws webSocketService) Upgrade(c *gin.Context, clientId string) {
	w := c.Writer
	r := c.Request
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &models.Client{
		Id:     clientId,
		Hub:    ws.Hub,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		Logger: ws.Logger,
	}
	client.Hub.Register <- client

	go client.WritePump()
	go client.ReadPump()
}

func (ws webSocketService) IsClientExist(clientId string) bool {
	for client := range ws.Hub.Clients {
		if client.Id == clientId {
			return true
		}
	}
	return false
}

func NewWebSocketService(logger utils.Logger) WebSocketService {
	if webSocketSvc != nil {
		logger.Info("WebSocket SVC has been initialized already")
		return webSocketSvc
	} else {
		logger.Info("WebSocket SVC: initializing")
		webSocketSvcLock.Lock()
		defer webSocketSvcLock.Unlock()
		hub := models.NewHub()
		go hub.Run()
		webSocketSvc = webSocketService{
			Logger: logger,
			Hub:    hub,
		}
		return webSocketSvc
	}
}

func (ws webSocketService) SendWsMsgByClientId(clientId string, message string) error {
	// Lock the Hub for safe concurrent access
	ws.Hub.Lock()
	defer ws.Hub.Unlock()

	// Iterate over all clients in the Hub
	for client := range ws.Hub.Clients {
		if client.Id == clientId {
			// Found the client, send the message
			select {
			case client.Send <- []byte(message):
				return nil // Message queued successfully
			default:
				return errors.New("failed to send message: channel is full")
			}
		}
	}

	// Client not found
	return errors.New("client not found")
}
