package routes

import (
	"github.com/xairline/x-gpt/controllers"
	"github.com/xairline/x-gpt/middlewares"
	"github.com/xairline/x-gpt/utils"
)

// WebSocketRoutes struct
type WebSocketRoutes struct {
	logger              utils.Logger
	handler             utils.RequestHandler
	webSocketController controllers.WebSocketController
}

// Setup WebSocket routes
func (s WebSocketRoutes) Setup() {
	s.logger.Info("Setting up WEBSOCKET routes")

	api := s.handler.Gin.Group("/apis")
	{
		api.GET("/ws", s.webSocketController.Upgrade)
		api.GET("/ws/token", middlewares.OIDCMiddleware(s.handler.Context, s.handler.Verifier), s.webSocketController.GetConnectionToken)
	}
}

// NewWebSocketRoutes creates new WebSocket controller
func NewWebSocketRoutes(
	logger utils.Logger,
	handler utils.RequestHandler,
	webSocketController controllers.WebSocketController,
) WebSocketRoutes {
	return WebSocketRoutes{
		handler:             handler,
		logger:              logger,
		webSocketController: webSocketController,
	}
}
