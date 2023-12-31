package routes

import (
	"github.com/xairline/x-gpt/controllers"
	"github.com/xairline/x-gpt/middlewares"
	"github.com/xairline/x-gpt/utils"
)

// DatarefRoutes struct
type DatarefRoutes struct {
	logger            utils.Logger
	handler           utils.RequestHandler
	datarefController controllers.DatarefController
}

// Setup Dataref routes
func (s DatarefRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/apis/xplm")
	api.Use(middlewares.OIDCMiddleware(s.handler.Context, s.handler.Verifier))
	{
		api.GET("/dataref", s.datarefController.GetDataref)
		api.PUT("/dataref", s.datarefController.SetDataref)
		api.PUT("/command", s.datarefController.SendCommand)
	}
}

// NewDatarefRoutes creates new Dataref controller
func NewDatarefRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	datarefController controllers.DatarefController,
) DatarefRoutes {
	return DatarefRoutes{
		logger:            logger,
		handler:           handler,
		datarefController: datarefController,
	}
}
