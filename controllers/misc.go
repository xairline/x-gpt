package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xairline/x-gpt/utils"
)

// use ldflags to replace this value during build:
//
//	https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications
const VERSION string = "development"

// MiscController data type
type MiscController struct {
	logger utils.Logger
}

// NewMiscController creates new Misc controller
func NewMiscController(logger utils.Logger) MiscController {
	return MiscController{
		logger: logger,
	}
}

// GetVersion
//
//	@Summary	Get version of GPT X-Plane plugin
//	@Tags		Misc
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	utils.ResponseOk
//	@Failure	500	{object}	utils.ResponseError
//	@Router		/version [get]
func (u MiscController) GetVersion(c *gin.Context) {
	// log headers
	u.logger.Info("Headers: %v", c.Request.Header)
	c.JSON(200, utils.ResponseOk{Message: VERSION})
}

// GetReadiness
//
//	@Summary	K8s Readiness endpoint
//	@Tags		Misc
//	@Accept		json
//	@Produce	json
//	@Success	204
//	@Failure	500	{object}	utils.ResponseError
//	@Router		/readiness [get]
func (u MiscController) GetReadiness(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

// GetLiveness
//
//	@Summary	K8s Liveness endpoint
//	@Tags		Misc
//	@Accept		json
//	@Produce	json
//	@Success	204
//	@Router		/liveness [get]
func (u MiscController) GetLiveness(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
