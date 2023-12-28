package controllers

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/xairline/x-gpt/services"
	"github.com/xairline/x-gpt/utils"
	"net/http"
)

// WebSocketController data type
type WebSocketController struct {
	logger       utils.Logger
	env          utils.Env
	webSocketSvc services.WebSocketService
}

// NewWebSocketController creates new WebSocket controller
func NewWebSocketController(env utils.Env, logger utils.Logger, webSocketSvc services.WebSocketService) WebSocketController {
	return WebSocketController{
		logger:       logger,
		env:          env,
		webSocketSvc: webSocketSvc,
	}
}

func (ws WebSocketController) Upgrade(c *gin.Context) {
	auth, success := c.GetQuery("auth")
	if !success {
		c.JSON(http.StatusUnauthorized, utils.ResponseError{Message: `missing "auth"`})
		return
	}
	// Initialize OIDC provider
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, ws.env.OauthEndpoint)
	if err != nil {
		panic(err)
	}
	oidcVerifier := provider.Verifier(&oidc.Config{ClientID: ws.env.OauthClientID, SkipClientIDCheck: true, SkipExpiryCheck: true})
	token, err := oidcVerifier.Verify(ctx, auth)

	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.ResponseError{Message: err.Error()})
		return
	}
	ws.webSocketSvc.Upgrade(c, token.Subject)
}
