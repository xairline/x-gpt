package utils

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xairline/x-gpt/docs"
)

// RequestHandler function
type RequestHandler struct {
	Gin *gin.Engine
}

type ResponseError struct {
	Message string `json:"message"`
} //@name ResponseError

type ResponseOk struct {
	Message string `json:"message"`
} //@name ResponseOk

// NewRequestHandler creates a new request handler
func NewRequestHandler(logger Logger, env Env) RequestHandler {
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

	return RequestHandler{Gin: engine}
}
