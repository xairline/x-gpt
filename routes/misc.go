package routes

import (
	"github.com/xairline/x-gpt/controllers"
	"github.com/xairline/x-gpt/utils"
)

// MiscRoutes struct
type MiscRoutes struct {
	logger         utils.Logger
	handler        utils.RequestHandler
	miscController controllers.MiscController
}

// Setup Misc routes
func (s MiscRoutes) Setup() {
	s.logger.Info("Setting up routes")
	api := s.handler.Gin.Group("/apis")
	{
		api.GET("/liveness", s.miscController.GetLiveness)
		api.GET("/readiness", s.miscController.GetReadiness)
		api.GET("/version", s.miscController.GetVersion)

	}
}

// NewMiscRoutes creates new Misc controller
func NewMiscRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	miscController controllers.MiscController,
) MiscRoutes {
	return MiscRoutes{
		handler:        handler,
		logger:         logger,
		miscController: miscController,
	}
}
