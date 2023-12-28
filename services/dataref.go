package services

//go:generate mockgen -destination=../services/__mocks__/dataref.go -package=mocks -source=dataref.go

import "C"
import (
	_ "embed"
	"github.com/xairline/x-gpt/models"
	"github.com/xairline/x-gpt/utils"
	"sync"
)

var datarefSvcLock = &sync.Mutex{}
var datarefSvc DatarefService

type DatarefService interface {
	GetValueByDatarefName(clientId string, dataref, name string, precision *int8, isByteArray bool) models.DatarefValue
	SetValueByDatarefName(clientId string, dataref string, value interface{})
}

type datarefService struct {
	Logger           utils.Logger
	WebSocketService WebSocketService
}

func (d datarefService) SetValueByDatarefName(clientId string, dataref string, value interface{}) {
	// find websocket client and forward request
	return
}

func (d datarefService) GetValueByDatarefName(clientId, dataref, name string, precision *int8, isByteArray bool) models.DatarefValue {
	_ = d.WebSocketService.SendWsMsgByClientId(clientId, "action:GetDataref, dataref: "+dataref)
	return models.DatarefValue{}
}

func NewDatarefService(logger utils.Logger, webSocketService WebSocketService) DatarefService {
	if datarefSvc != nil {
		logger.Info("Dataref SVC has been initialized already")
		return datarefSvc
	} else {
		logger.Info("Dataref SVC: initializing")
		datarefSvcLock.Lock()
		defer datarefSvcLock.Unlock()

		datarefSvc = datarefService{
			Logger:           logger,
			WebSocketService: webSocketService,
		}
		return datarefSvc
	}
}
