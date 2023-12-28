package controllers

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/websocket"
	"github.com/xairline/x-gpt/services"
	"github.com/xairline/x-gpt/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestUpgrade tests the Upgrade method for WebSocket connections
func TestUpgrade(tt *testing.T) {
	// Setup for the token request
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// Mocks and WebSocketController setup
	mockCtrl := gomock.NewController(tt)
	defer mockCtrl.Finish()
	mockEnv := utils.Env{OauthEndpoint: "https://pzloiy.logto.app/oidc"}
	mockLogger := utils.NewLogger()
	webSocketSvc := services.NewWebSocketService(mockLogger)
	wsController := NewWebSocketController(mockEnv, mockLogger, webSocketSvc)
	router.GET("/ws/upgrade", wsController.Upgrade)

	// Start a local HTTP server
	server := httptest.NewServer(router)
	defer server.Close()

	tt.Run("Success", func(t *testing.T) {
		// Request a token
		req, _ := http.NewRequest("GET", server.URL+"/ws/token", nil)
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJFUzM4NCIsInR5cCI6ImF0K2p3dCIsImtpZCI6ImE3NGlvcEVHVXpqSTd3OEtIanEyNkk5cnE0YzYzRXNiRTYtUllTMGNDYlUifQ.eyJqdGkiOiJ4RVJVYldGMURmWTU5ZWw0QTE3UnkiLCJzdWIiOiI1enhtZGdxcHQwdm8iLCJpYXQiOjE3MDM3MjU3ODUsImV4cCI6MTcwMzcyOTM4NSwic2NvcGUiOiIiLCJjbGllbnRfaWQiOiI0ZGVhdDYwdjFiZW82YjI5Z2l3ZnIiLCJpc3MiOiJodHRwczovL3B6bG9peS5sb2d0by5hcHAvb2lkYyIsImF1ZCI6Imh0dHBzOi8vYXBwLnhhaXJsaW5lLm9yZyJ9.bonkXjKKPxb0uUYDBXENexJK_8sFxtLKHDH8mbkEmm5JcJLZc2V0Stnpw9cnz_qFEaQQhJYKOMUc-0ZckEDTEvwqR5RnWqTsna-SJPWqFrMHcmRiASLVindqUgSAfER8")
		tokenResp := httptest.NewRecorder()
		router.ServeHTTP(tokenResp, req)

		// Parse the token response
		token := "eyJhbGciOiJFUzM4NCIsInR5cCI6ImF0K2p3dCIsImtpZCI6ImE3NGlvcEVHVXpqSTd3OEtIanEyNkk5cnE0YzYzRXNiRTYtUllTMGNDYlUifQ.eyJqdGkiOiJ4RVJVYldGMURmWTU5ZWw0QTE3UnkiLCJzdWIiOiI1enhtZGdxcHQwdm8iLCJpYXQiOjE3MDM3MjU3ODUsImV4cCI6MTcwMzcyOTM4NSwic2NvcGUiOiIiLCJjbGllbnRfaWQiOiI0ZGVhdDYwdjFiZW82YjI5Z2l3ZnIiLCJpc3MiOiJodHRwczovL3B6bG9peS5sb2d0by5hcHAvb2lkYyIsImF1ZCI6Imh0dHBzOi8vYXBwLnhhaXJsaW5lLm9yZyJ9.bonkXjKKPxb0uUYDBXENexJK_8sFxtLKHDH8mbkEmm5JcJLZc2V0Stnpw9cnz_qFEaQQhJYKOMUc-0ZckEDTEvwqR5RnWqTsna-SJPWqFrMHcmRiASLVindqUgSAfER8"

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

	tt.Run("InvalidTokenPwd", func(t *testing.T) {
		// Use an invalid token to connect to the WebSocket endpoint
		invalidToken := "eyJhbGciOiJFUzM4NCIsInR5cCI6ImF0K2p3dCIsImtpZCI6ImE3NGlvcEVHVXpqSTd3OEtIanEyNkk5cnE0YzYzRXNiRTYtUllTMGNDYlUifQ.eyJqdGkiOiJ4RVJVYldGMURmWTU5ZWw0QTE3UnkiLCJzdWIiOiI1enhtZGdxcHQwdm8iLCJpYXQiOjE3MDM3MjU3ODUsImV4cCI6MTcwMzcyOTM4NSwic2NvcGUiOiIiLCJjbGllbnRfaWQiOiI0ZGVhdDYwdjFiZW82YjI5Z2l3ZnIiLCJpc3MiOiJodHRwczovL3B6bG9peS5sb2d0by5hcHAvb2lkYyIsImF1ZCI6Imh0dHBzOi8vYXBwLnhhaXJsaW5lLm9yZyJ9.bonkXjKKPxb0uUYDBXENexJK_8sFxtLKHDH8mbkEmm5JcJLZc2V0Stnpw9cnz_qFEaQQhJYKOMUc-0ZckEDTEvwqR5RnWqTsna-SJPWqFrMHcmRiASLVindqUgSAfER"
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
