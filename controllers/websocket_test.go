package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/xairline/x-gpt/services"
	mocks "github.com/xairline/x-gpt/services/__mocks__"
	"github.com/xairline/x-gpt/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestGetConnectionToken tests the GetConnectionToken method
func TestGetConnectionToken(tt *testing.T) {
	// Set up Gin
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Create a mock controller
	mockCtrl := gomock.NewController(tt)
	defer mockCtrl.Finish()

	// Set up your environment and dependencies for WebSocketController
	mockEnv := utils.Env{JwtSecret: "your-secret-key"}
	mockLogger := utils.Logger{} // Assume this is your actual logger or a mock
	mockWebSocketService := mocks.NewMockWebSocketService(mockCtrl)

	// Initialize your WebSocketController
	wsController := NewWebSocketController(mockEnv, mockLogger, mockWebSocketService)

	// Set up the route
	router.GET("/apis/ws/token", wsController.GetConnectionToken)

	tt.Run("Success", func(t *testing.T) {
		// Create a request with a Bearer token in the Authorization header
		req, _ := http.NewRequest("GET", "/apis/ws/token", nil)
		req.Header.Add("Authorization", "Bearer test-token")

		// Record the response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusOK, w.Code)
		// Parse the response body to get the token
		var response utils.ResponseOk
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)

		// Extract the token from the response
		tokenString := response.Message
		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Ensure the token's algorithm matches what you expect
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(mockEnv.JwtSecret), nil
		})

		assert.NoError(t, err)
		assert.NotNil(t, token)
		assert.True(t, token.Valid)

		// Optionally check the claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check for specific claims you expect
			// For example, assert the 'token' claim matches 'test-token'
			assert.Equal(t, "test-token", claims["token"])
		} else {
			t.Errorf("Claims are not of type jwt.MapClaims")
		}
	})

	tt.Run("Invalid Authorization header", func(t *testing.T) {
		// Create a request without the Authorization header
		req, _ := http.NewRequest("GET", "/apis/ws/token", nil)

		// Record the response
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assertions
		assert.Equal(t, http.StatusBadRequest, w.Code) // Assuming 400 is the response for missing Authorization

		// Optionally, you can also check the error message in the response
		var response utils.ResponseError
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Invalid authorization header", response.Message)
	})
}

// TestUpgrade tests the Upgrade method for WebSocket connections
func TestUpgrade(tt *testing.T) {
	// Setup for the token request
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Mocks and WebSocketController setup
	mockCtrl := gomock.NewController(tt)
	defer mockCtrl.Finish()
	mockEnv := utils.Env{JwtSecret: "your-secret-key"}
	mockLogger := utils.NewLogger()
	webSocketSvc := services.NewWebSocketService(mockLogger)
	wsController := NewWebSocketController(mockEnv, mockLogger, webSocketSvc)
	router.GET("/ws/token", wsController.GetConnectionToken)
	router.GET("/ws/upgrade", wsController.Upgrade)

	// Start a local HTTP server
	server := httptest.NewServer(router)
	defer server.Close()

	tt.Run("Success", func(t *testing.T) {
		// Request a token
		req, _ := http.NewRequest("GET", server.URL+"/ws/token", nil)
		req.Header.Add("Authorization", "Bearer test-token")
		tokenResp := httptest.NewRecorder()
		router.ServeHTTP(tokenResp, req)

		// Parse the token response
		var response utils.ResponseOk
		json.Unmarshal(tokenResp.Body.Bytes(), &response)
		token := response.Message

		// Use the token to connect to the WebSocket endpoint
		wsUrl := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws/upgrade?auth=" + token
		ws, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
		if err != nil {
			tt.Fatalf("Could not open a ws connection on %s %v", wsUrl, err)
		}
		defer ws.Close()

		// Test sending a message over the WebSocket
		testMessage := "hello"
		if err := ws.WriteMessage(websocket.TextMessage, []byte(testMessage)); err != nil {
			tt.Fatalf("Could not send message over ws connection %v", err)
		}

		// Read the response from the WebSocket server
		messageType, responseMessage, err := ws.ReadMessage()
		if err != nil {
			tt.Fatalf("Could not read message from ws connection %v", err)
		}

		// Assertions and additional tests
		assert.NotNil(tt, ws)
		assert.Equal(tt, websocket.TextMessage, messageType)
		assert.GreaterOrEqual(tt, len(string(responseMessage)), 0) // Replace with your expected response
	})

	tt.Run("InvalidToken", func(t *testing.T) {
		// Use an invalid token to connect to the WebSocket endpoint
		invalidToken := "invalid-token"
		wsUrl := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws/upgrade?auth=" + invalidToken
		ws, resp, err := websocket.DefaultDialer.Dial(wsUrl, nil)

		// Expect an error, as the token is invalid
		assert.NotNil(t, err)
		if err == nil {
			ws.Close()
		}
		if resp != nil {
			// Optionally check the response status code if your implementation sends a specific response
			assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
		}
	})
}
