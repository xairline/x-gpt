package utils

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xairline/x-gpt/docs"
	"github.com/xairline/x-gpt/middlewares"
)

// RequestHandler function
type RequestHandler struct {
	Gin      *gin.Engine
	Verifier *oidc.IDTokenVerifier
	Context  context.Context
}

type ResponseError struct {
	Message string `json:"message"`
} //@name ResponseError

type ResponseOk struct {
	Message string `json:"message"`
} //@name ResponseOk

// NewRequestHandler creates a new request handler
func NewRequestHandler(logger Logger, env Env) RequestHandler {
	// Initialize OIDC provider
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, env.OauthEndpoint)
	if err != nil {
		panic(err)
	}
	oidcVerifier := provider.Verifier(&oidc.Config{ClientID: env.OauthClientID, SkipClientIDCheck: true})

	gin.DefaultWriter = logger.GetGinLogger()
	engine := gin.New()

	docs.SwaggerInfo.Title = "X-Airline GPT"
	engine.GET(
		"/docs/*any",
		ginSwagger.WrapHandler(
			swaggerFiles.Handler,
			ginSwagger.DocExpansion("none"),
		),
	)
	engine.Use(middlewares.CORSMiddleware())

	return RequestHandler{
		Gin:      engine,
		Verifier: oidcVerifier,
		Context:  ctx,
	}
}
