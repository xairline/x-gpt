package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/xairline/x-gpt/services"
	"github.com/xairline/x-gpt/utils"
	"strings"
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
	mySigningKey := []byte(ws.env.JwtSecret)
	auth, success := c.GetQuery("auth")
	if !success {
		c.JSON(500, utils.ResponseError{Message: `missing "auth"`})
		return
	}
	// Parse the token. Make sure to validate the alg is what you expect.
	token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})
	if err != nil {
		c.JSON(401, utils.ResponseError{Message: err.Error()})
		return
	}
	ws.webSocketSvc.Upgrade(c, token.Claims.(jwt.MapClaims)["token"].(string))
}

// GetConnectionToken
//
//	@Summary	Get Token for WebSocket connection
//	@Tags		Web Socket
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	utils.ResponseOk
//	@Failure	500	{object}	utils.ResponseError
//	@Router		/ws/token [get]
//
// @Security Oauth2Application[]
func (ws WebSocketController) GetConnectionToken(c *gin.Context) {
	// Extract the Authorization header from the request
	authHeader := c.GetHeader("Authorization")

	// Check if the Authorization header is in the correct format
	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(400, utils.ResponseError{Message: "Invalid authorization header"})
		return
	}

	// Extract the token from the Authorization header
	bearerToken := strings.TrimPrefix(authHeader, "Bearer ")

	// Create a new JWT token with the extracted bearer token as a claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token": bearerToken,
	})

	// Sign the new JWT token with the secret
	signedToken, err := token.SignedString([]byte(ws.env.JwtSecret))
	if err != nil {
		// Return an error response if there's an issue with signing the token
		c.JSON(500, utils.ResponseError{Message: err.Error()})
		return
	}

	// Return the signed JWT token in the response
	c.JSON(200, utils.ResponseOk{Message: signedToken})
}
