package controllers

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/xairline/x-gpt/middlewares"
	"github.com/xairline/x-gpt/utils"
	"net/http"
	"net/http/httptest"
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
		dc := NewDatarefController(utils.NewLogger(), mockService)
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
		mockService := mocks.NewMockDatarefService(mockCtrl)

		// Set up Gin
		router := gin.Default()
		env := utils.NewEnv(utils.NewLogger())
		dc := NewDatarefController(utils.NewLogger(), mockService)
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
		req.Header.Add("Authorization", "Bearer test-token")
		// Add headers or other request setup for valid token simulation

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		// Assert status code and response
		assert.Equal(t, http.StatusUnauthorized, w.Code)
		// Additional assertions as needed
	})
}
