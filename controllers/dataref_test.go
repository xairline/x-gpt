package controllers

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/websocket"
	"github.com/xairline/x-gpt/middlewares"
	"github.com/xairline/x-gpt/services"
	"github.com/xairline/x-gpt/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/xairline/x-gpt/services/__mocks__"
	// other necessary imports
)

func TestGetDataref(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("TestGetDataref_NoToken", func(t *testing.T) {
		// Setup mock controller and service
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockDatarefService(mockCtrl)

		// Set up Gin
		router := gin.Default()
		env := utils.NewEnv(utils.NewLogger())
		dc := NewDatarefController(utils.NewLogger(), mockService, services.NewWebSocketService(utils.NewLogger()))
		ctx := context.Background()
		provider, err := oidc.NewProvider(ctx, env.OauthEndpoint)
		if err != nil {
			panic(err)
		}
		oidcVerifier := provider.Verifier(&oidc.Config{ClientID: env.OauthClientID})
		router.Use(middlewares.OIDCMiddleware(context.Background(), oidcVerifier))
		router.GET("/xplm/dataref", dc.GetDataref)

		// Simulate a request with a valid token
		req, _ := http.NewRequest("GET", "/xplm/dataref?dataref_str=test&precision=2", nil)
		// Add headers or other request setup for valid token simulation

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert status code and response
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		// Additional assertions as needed
	})

	t.Run("TestGetDataref_ValidTokenWithWebSocket", func(t *testing.T) {
		// Setup mock controller and service
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		// Set up Gin
		router := gin.Default()
		env := utils.NewEnv(utils.NewLogger())
		dc := NewDatarefController(
			utils.NewLogger(),
			services.NewDatarefService(utils.NewLogger()),
			services.NewWebSocketService(utils.NewLogger()),
		)

		// setup token verification
		ctx := context.Background()
		provider, err := oidc.NewProvider(ctx, env.OauthEndpoint)
		if err != nil {
			panic(err)
		}
		oidcVerifier := provider.Verifier(&oidc.Config{ClientID: env.OauthClientID, SkipClientIDCheck: true})

		router.GET("/xplm/dataref", middlewares.OIDCMiddleware(context.Background(), oidcVerifier), dc.GetDataref)

		// Simulate a request with a valid token
		mockEnv := utils.Env{OauthEndpoint: "https://pzloiy.logto.app/oidc"}
		mockLogger := utils.NewLogger()
		webSocketSvc := services.NewWebSocketService(mockLogger)
		wsController := NewWebSocketController(mockEnv, mockLogger, webSocketSvc)
		router.GET("/ws/upgrade", wsController.Upgrade)
		// fake a websocket connection
		token := "eyJhbGciOiJFUzM4NCIsInR5cCI6ImF0K2p3dCIsImtpZCI6ImE3NGlvcEVHVXpqSTd3OEtIanEyNkk5cnE0YzYzRXNiRTYtUllTMGNDYlUifQ.eyJqdGkiOiJ4RVJVYldGMURmWTU5ZWw0QTE3UnkiLCJzdWIiOiI1enhtZGdxcHQwdm8iLCJpYXQiOjE3MDM3MjU3ODUsImV4cCI6MTcwMzcyOTM4NSwic2NvcGUiOiIiLCJjbGllbnRfaWQiOiI0ZGVhdDYwdjFiZW82YjI5Z2l3ZnIiLCJpc3MiOiJodHRwczovL3B6bG9peS5sb2d0by5hcHAvb2lkYyIsImF1ZCI6Imh0dHBzOi8vYXBwLnhhaXJsaW5lLm9yZyJ9.bonkXjKKPxb0uUYDBXENexJK_8sFxtLKHDH8mbkEmm5JcJLZc2V0Stnpw9cnz_qFEaQQhJYKOMUc-0ZckEDTEvwqR5RnWqTsna-SJPWqFrMHcmRiASLVindqUgSAfER8"

		// Simulate a request with a valid token
		req, _ := http.NewRequest("GET", "/xplm/dataref?dataref_str=test&precision=2", nil)
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJFUzM4NCIsInR5cCI6ImF0K2p3dCIsImtpZCI6ImE3NGlvcEVHVXpqSTd3OEtIanEyNkk5cnE0YzYzRXNiRTYtUllTMGNDYlUifQ.eyJqdGkiOiJ4RVJVYldGMURmWTU5ZWw0QTE3UnkiLCJzdWIiOiI1enhtZGdxcHQwdm8iLCJpYXQiOjE3MDM3MjU3ODUsImV4cCI6MTcwMzcyOTM4NSwic2NvcGUiOiIiLCJjbGllbnRfaWQiOiI0ZGVhdDYwdjFiZW82YjI5Z2l3ZnIiLCJpc3MiOiJodHRwczovL3B6bG9peS5sb2d0by5hcHAvb2lkYyIsImF1ZCI6Imh0dHBzOi8vYXBwLnhhaXJsaW5lLm9yZyJ9.bonkXjKKPxb0uUYDBXENexJK_8sFxtLKHDH8mbkEmm5JcJLZc2V0Stnpw9cnz_qFEaQQhJYKOMUc-0ZckEDTEvwqR5RnWqTsna-SJPWqFrMHcmRiASLVindqUgSAfER8")

		// Start a local HTTP server
		server := httptest.NewServer(router)
		defer server.Close()

		// Use the token to connect to the WebSocket endpoint
		wsUrl := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws/upgrade?auth=" + token
		ws, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
		if err != nil {
			t.Fatalf("Could not open a ws connection on %s %v", wsUrl, err)
		}
		defer ws.Close()

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert status code and response
		assert.Equal(t, http.StatusOK, w.Code)
		// Receive a message from the WebSocket connection
		_, message, err := ws.ReadMessage()
		if err != nil {
			t.Fatalf("Error reading from WebSocket: %v", err)
		}
		assert.Equal(t, `{"dataref_str":"test","value":0,"alias":"test","precision":2,"is_byte_array":false}`, string(message))
	})

	t.Run("TestGetDataref_ValidTokenWithoutWebSocket", func(t *testing.T) {
		// Setup mock controller and service
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockService := mocks.NewMockDatarefService(mockCtrl)

		// Set up Gin
		router := gin.Default()
		env := utils.NewEnv(utils.NewLogger())
		dc := NewDatarefController(utils.NewLogger(), mockService, services.NewWebSocketService(utils.NewLogger()))

		// setup token verification
		ctx := context.Background()
		provider, err := oidc.NewProvider(ctx, env.OauthEndpoint)
		if err != nil {
			panic(err)
		}
		oidcVerifier := provider.Verifier(&oidc.Config{ClientID: env.OauthClientID, SkipClientIDCheck: true})

		router.GET("/xplm/dataref", middlewares.OIDCMiddleware(context.Background(), oidcVerifier), dc.GetDataref)

		// Simulate a request with a valid token
		req, _ := http.NewRequest("GET", "/xplm/dataref?dataref_str=test&precision=2", nil)
		req.Header.Add("Authorization", "Bearer eyJhbGciOiJFUzM4NCIsInR5cCI6ImF0K2p3dCIsImtpZCI6ImE3NGlvcEVHVXpqSTd3OEtIanEyNkk5cnE0YzYzRXNiRTYtUllTMGNDYlUifQ.eyJqdGkiOiJ4RVJVYldGMURmWTU5ZWw0QTE3UnkiLCJzdWIiOiI1enhtZGdxcHQwdm8iLCJpYXQiOjE3MDM3MjU3ODUsImV4cCI6MTcwMzcyOTM4NSwic2NvcGUiOiIiLCJjbGllbnRfaWQiOiI0ZGVhdDYwdjFiZW82YjI5Z2l3ZnIiLCJpc3MiOiJodHRwczovL3B6bG9peS5sb2d0by5hcHAvb2lkYyIsImF1ZCI6Imh0dHBzOi8vYXBwLnhhaXJsaW5lLm9yZyJ9.bonkXjKKPxb0uUYDBXENexJK_8sFxtLKHDH8mbkEmm5JcJLZc2V0Stnpw9cnz_qFEaQQhJYKOMUc-0ZckEDTEvwqR5RnWqTsna-SJPWqFrMHcmRiASLVindqUgSAfER8")

		// Start a local HTTP server
		server := httptest.NewServer(router)
		defer server.Close()

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert status code and response
		assert.Equal(t, http.StatusNotFound, w.Code)
		// Additional assertions as needed
	})
}
