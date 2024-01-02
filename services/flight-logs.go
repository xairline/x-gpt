package services

//go:generate mockgen -destination=../services/__mocks__/flight-logs.go -package=mocks -source=flight-logs.go

import "C"
import (
	_ "embed"
	"fmt"
	"github.com/xairline/x-gpt/models"
	"github.com/xairline/x-gpt/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

var flightLogsSvcLock = &sync.Mutex{}
var flightLogsSvc FlightLogsService

type FlightLogsService interface {
}

type flightLogsService struct {
	Logger utils.Logger
	db     *gorm.DB
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
