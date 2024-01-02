package services

//go:generate mockgen -destination=../services/__mocks__/flight-logs.go -package=mocks -source=flight-logs.go

import "C"
import (
	_ "embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xairline/x-gpt/models"
	"github.com/xairline/x-gpt/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
)

var flightLogsSvcLock = &sync.Mutex{}
var flightLogsSvc FlightLogsService

type FlightLogsService interface {
	GetLastSyncedLocalIdByUsername(username string) (int, error)
	SaveFlightStatuses(statuses []models.FlightStatus)
	GetFlightLogs(c *gin.Context)
	GetFlightLog(c *gin.Context)
}

type flightLogsService struct {
	Logger utils.Logger
	db     *gorm.DB
}

func (f flightLogsService) GetFlightLog(c *gin.Context) {
	var res models.FlightStatus
	id, _ := strconv.Atoi(c.Param("id"))
	result := f.db.
		Model(&models.FlightStatus{}).
		Preload("Locations").
		Preload("Events").
		First(&res, id)
	if result.Error == nil {
		c.JSON(200, res)
		return
	} else {
		f.Logger.Infof("%+v", result.Error)
		c.JSON(404, "not found")
		return
	}
}

func (f flightLogsService) GetFlightLogs(c *gin.Context) {
	var res []models.FlightStatus
	isOverview := c.Request.URL.Query().Get("isOverview")
	var result *gorm.DB
	if isOverview == "true" {
		result = f.db.
			Preload("Locations" /*, "event_type = (?)", models.StateEvent*/).
			Model(&models.FlightStatus{})

	} else {
		result = f.db.Model(&models.FlightStatus{})
	}
	// departureAirportId
	departureAirportId := c.Request.URL.Query().Get("departureAirportId")
	if len(departureAirportId) > 0 {
		result = result.Where("departure_airport_id = ?", departureAirportId)
	}
	// arrivalAirportId
	arrivalAirportId := c.Request.URL.Query().Get("arrivalAirportId")
	if len(arrivalAirportId) > 0 {
		result = result.Where("arrival_airport_id = ?", arrivalAirportId)
	}
	// aircraftICAO
	aircraftICAO := c.Request.URL.Query().Get("aircraftICAO")
	if len(aircraftICAO) > 0 {
		result = result.Where("aircraft_icao = ?", aircraftICAO)
	}
	result = result.Where("LENGTH(arrival_airport_id) > 0")
	result = result.Order("id DESC")
	// only load a login user's own data
	clientId := c.Request.URL.Query().Get("clientId")
	if clientId != "" {
		result = result.Where("username = ?", clientId)
	}

	result = result.Find(&res)
	if result.Error != nil {
		c.JSON(500, utils.ResponseError{Message: fmt.Sprintf("Failed to get flight logs: %+v", result.Error)})
		return
	}
	c.JSON(200, res)
	return
}

func (f flightLogsService) SaveFlightStatuses(statuses []models.FlightStatus) {
	result := f.db.Save(statuses)
	if result.Error != nil {
		f.Logger.Errorf("Failed to save flight logs: %+v", result.Error)
	}
}

func (f flightLogsService) GetLastSyncedLocalIdByUsername(username string) (int, error) {
	var res models.FlightStatus
	result := f.db.Model(&models.FlightStatus{}).
		Where("username = ?", username).
		Order("local_id desc").
		Limit(1).
		Find(&res)
	if result.Error != nil {
		f.Logger.Errorf("Failed to get flight logs: %+v", result.Error)
	}

	return int(res.LocalId), nil
}

func NewFlightLogsService(logger utils.Logger, env utils.Env) FlightLogsService {
	if flightLogsSvc != nil {
		logger.Info("FlightLogs SVC has been initialized already")
		return flightLogsSvc
	} else {
		logger.Info("FlightLogs SVC: initializing")
		flightLogsSvcLock.Lock()
		defer flightLogsSvcLock.Unlock()

		// create database connection pool - postgres
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s password=%s sslmode=require",
			env.DbHost,
			env.DbPort,
			env.DbUser,
			env.DbName,
			env.DbPass)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Errorf("Failed to connect to database: %+v", err)
			panic(err)
		}

		pgDB, _ := db.DB()
		pgDB.SetMaxOpenConns(20)
		pgDB.SetMaxIdleConns(10)
		pgDB.SetConnMaxLifetime(time.Hour)

		// Migrate the schema
		err = db.AutoMigrate(
			&models.FlightStatus{},
			&models.FlightStatusLocation{},
			&models.FlightStatusEvent{},
		)
		if err != nil {
			logger.Errorf("%+v", err)
			panic(err)
		}

		flightLogsSvc = flightLogsService{
			Logger: logger,
			db:     db,
		}
		return flightLogsSvc
	}
}
