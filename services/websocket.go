package services

//go:generate mockgen -destination=../services/__mocks__/webSocket.go -package=mocks -source=webSocket.go

import "C"
import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/xairline/x-gpt/models"
	"github.com/xairline/x-gpt/utils"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var webSocketSvcLock = &sync.Mutex{}
var webSocketSvc WebSocketService

type WebSocketService interface {
	Upgrade(c *gin.Context, clientId string)
	IsClientExist(clientId string) bool
	SendWsMsgByClientId(clientId string, message string) (string, error)
}

type webSocketService struct {
	Logger            utils.Logger
	Hub               *models.Hub
	FlightLogsService FlightLogsService
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
		Id:           clientId,
		Hub:          ws.Hub,
		Conn:         conn,
		Send:         make(chan []byte, 256),
		Logger:       ws.Logger,
		LastActivity: time.Now(),
	}
	client.Hub.Register <- client
	go client.WritePump()

	go func() {
		for {
			syncedLocalId, _ := ws.FlightLogsService.GetLastSyncedLocalIdByUsername(clientId)
			response, err := ws.SendWsMsgByClientId(clientId, "SyncFlightLogs|"+strconv.Itoa(syncedLocalId))
			if err != nil {
				ws.Logger.Errorf("Failed to sync flight logs: %+v", err)
				break
			}
			if response == "SyncFlightLogs|Done" {
				ws.Logger.Infof("Client: %s, Flight logs synced", clientId)
				break
			}
			var flightStatuses, syncedFlightStatuses []models.FlightStatus
			json.Unmarshal([]byte(response), &flightStatuses)
			for _, flightStatus := range flightStatuses {
				flightStatus.Username = clientId
				flightStatus.LocalId = flightStatus.ID
				flightStatus.ID = 0
				syncedFlightStatuses = append(syncedFlightStatuses, flightStatus)
			}
			ws.FlightLogsService.SaveFlightStatuses(syncedFlightStatuses)
			ws.Logger.Infof("Client: %s, Flight logs synced", clientId)
		}
	}()
	//go client.ReadPump()
}

func (ws webSocketService) IsClientExist(clientId string) bool {
	for client := range ws.Hub.Clients {
		if client.Id == clientId {
			return true
		}
	}
	return false
}

func NewWebSocketService(logger utils.Logger, flightLogsService FlightLogsService) WebSocketService {
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
			Logger:            logger,
			Hub:               hub,
			FlightLogsService: flightLogsSvc,
		}
		return webSocketSvc
	}
}

func (ws webSocketService) SendWsMsgByClientId(clientId string, message string) (string, error) {
	// Iterate over all clients in the Hub
	for client := range ws.Hub.Clients {
		if client.Id == clientId {
			// Found the client, send the message
			ws.Logger.Infof("Client: %s, sending: %s - wait for lock", clientId, message)
			client.Lock()
			select {
			case client.Send <- []byte(message):
				ws.Logger.Infof("Client: %s, sending: %s - waiting for response", clientId, message)
				for {
					_, message, err := client.Conn.ReadMessage()
					if err != nil {
						if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
							client.Logger.Errorf("error: %v", err)
						}
						break
					}
					if len(message) > 0 {
						ws.Logger.Infof("Client: %s, received: %s", clientId, message)
						client.LastActivity = time.Now()
						client.Unlock()
						return string(message), nil
					}
				}
				client.Unlock()
				break
			default:
				client.Unlock()
				return "", errors.New("failed to send message: channel is full")
			}
		}
	}

	// Client not found
	return "", errors.New("client not found")
}
