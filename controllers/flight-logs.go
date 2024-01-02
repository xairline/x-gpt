package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/xairline/x-gpt/services"
	"github.com/xairline/x-gpt/utils"
)

// FlightLogsController data type
type FlightLogsController struct {
	logger            utils.Logger
	flightLogsService services.FlightLogsService
}

// NewFlightLogsController creates new FlightLogs controller
func NewFlightLogsController(
	logger utils.Logger,
	flightLogsService services.FlightLogsService,
) FlightLogsController {
	return FlightLogsController{
		logger:            logger,
		flightLogsService: flightLogsService,
	}
}

// GetFlightLogs
// @Summary  Get a list of FlightLogs
// @Param    isOverview    query     string  false  "specify if it's overview"
// @Param    clientId    query     string  false  "specify clientId"
// @Param    departureAirportId query string false "departure airport"
// @Param    arrivalAirportId query string false "arrival airport"
// @Param    aircraftICAO query string false "aircraft ICAO"
// @Param    source query string false "xplane or xws"
// @Tags     Flight_Logs
// @Accept   json
// @Produce  json
// @Success  200  {object}  []models.FlightStatus
// @Failure  500  {object}  utils.ResponseError
// @Router   /flight-logs [get]
func (u FlightLogsController) GetFlightLogs(c *gin.Context) {
	u.flightLogsService.GetFlightLogs(c)
	return
}

// GetFlightLog
// @Summary  Get one FlightLog
// @Param    id  path  string  true  "id of a flight log item"
// @Tags     Flight_Logs
// @Accept   json
// @Produce  json
// @Success  200  {object}  models.FlightStatus
// @Failure  404  "Not Found"
// @Router   /flight-logs/{id} [get]
func (u FlightLogsController) GetFlightLog(c *gin.Context) {
	u.flightLogsService.GetFlightLog(c)
	return
}
