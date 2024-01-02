package routes

import (
	"github.com/xairline/x-gpt/controllers"
	"github.com/xairline/x-gpt/utils"
)

// FlightLogsRoutes struct
type FlightLogsRoutes struct {
	logger               utils.Logger
	handler              utils.RequestHandler
	flightLogsController controllers.FlightLogsController
}

// Setup FlightLogs routes
func (s FlightLogsRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/apis/flight-logs")
	{
		api.GET("", s.flightLogsController.GetFlightLogs)
		api.GET(":id", s.flightLogsController.GetFlightLog)

	}
}

// NewFlightLogsRoutes creates new FlightLogs controller
func NewFlightLogsRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	flightLogsController controllers.FlightLogsController,
) FlightLogsRoutes {
	return FlightLogsRoutes{
		logger:               logger,
		handler:              handler,
		flightLogsController: flightLogsController,
	}
}
