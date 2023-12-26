package services

//go:generate mockgen -destination=../services/__mocks__/webSocket.go -package=mocks -source=webSocket.go

import "C"
import (
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
