package services

//go:generate mockgen -destination=../services/__mocks__/dataref.go -package=mocks -source=dataref.go

import "C"
import (
	_ "embed"
	"encoding/json"
	"github.com/xairline/x-gpt/models"
	"github.com/xairline/x-gpt/utils"
	"sync"
)

var datarefSvcLock = &sync.Mutex{}
var datarefSvc DatarefService

type DatarefService interface {
	GetValueByDatarefName(clientId string, dataref, name string, precision *int8, isByteArray bool) models.DatarefValue
	SetValueByDatarefName(clientId string, dataref string, value interface{})
	SendCommand(clientId string, command string)
}

type datarefService struct {
	Logger           utils.Logger
	WebSocketService WebSocketService
}

func (d datarefService) SetValueByDatarefName(clientId string, dataref string, value interface{}) {
	datarefObj := models.SetDatarefValue{
		Dataref: dataref,
		Value:   value,
	}
	datarefBytes, _ := json.Marshal(datarefObj)
	d.WebSocketService.SendWsMsgByClientId(clientId, "SetDataref|"+string(datarefBytes))
	return
}

func (d datarefService) SendCommand(clientId string, cmd string) {
	d.WebSocketService.SendWsMsgByClientId(clientId, "SendCommand|"+cmd)
	return
}

func (d datarefService) GetValueByDatarefName(clientId, dataref, name string, precision *int8, isByteArray bool) models.DatarefValue {
	datarefObj := models.Dataref{
		Name:         dataref,
		Precision:    *precision,
		IsBytesArray: isByteArray,
	}
	datarefBytes, _ := json.Marshal(datarefObj)
	message, _ := d.WebSocketService.SendWsMsgByClientId(clientId, "GetDataref|"+string(datarefBytes))
	res := models.DatarefValue{}
	json.Unmarshal([]byte(message), &res)
	return res
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
