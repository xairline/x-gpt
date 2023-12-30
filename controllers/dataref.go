package controllers

import (
	"github.com/xairline/x-gpt/models"
	"github.com/xairline/x-gpt/services"
	"github.com/xairline/x-gpt/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DatarefController data type
type DatarefController struct {
	logger           utils.Logger
	datarefSvc       services.DatarefService
	webSocketService services.WebSocketService
}

// NewDatarefController creates new Dataref controller
func NewDatarefController(
	logger utils.Logger,
	datarefSvc services.DatarefService,
	webSocketService services.WebSocketService,
) DatarefController {
	return DatarefController{
		logger:           logger,
		datarefSvc:       datarefSvc,
		webSocketService: webSocketService,
	}
}

// GetDataref
// @Summary  Get Dataref
// @Tags     Dataref
// @Param    dataref_str  query string true "xplane dataref string"
// @Param    alias  query string false "alias name, if not set, dataref_str will be used"
// @Param    precision  query int8 true "-1: raw, 2: round up to two digits"
// @Param    is_byte_array query bool false "transform xplane byte array to string"
// @Accept   json
// @Produce  json
// @Success  200  {object}  models.DatarefValue
// @Failure  500  {object}  utils.ResponseError
// @Router   /xplm/dataref [get]
// @Security Oauth2Application[]
func (u DatarefController) GetDataref(c *gin.Context) {
	dataref, success := c.GetQuery("dataref_str")
	if !success {
		c.JSON(500, utils.ResponseError{Message: `missing "dataref_str"`})
	}

	var alias string
	alias, success = c.GetQuery("alias")
	if !success {
		alias = dataref
	}

	precisionInt, success := c.GetQuery("precision")
	if !success {
		precisionInt = "0"
	}
	precisionTmp, _ := strconv.Atoi(precisionInt)
	precision := int8(precisionTmp)

	// check hub if we have an active connection
	// Retrieve the value from the context
	clientId, exists := c.Get("clientId")
	if !exists {
		// Handle the case where "clientId" is not set
		c.JSON(http.StatusUnauthorized, gin.H{"error": "client ID not found"})
		return
	}
	if !u.webSocketService.IsClientExist(clientId.(string)) {
		c.JSON(http.StatusNotFound, gin.H{"error": "X Plane not connected"})
		return
	}

	res := u.datarefSvc.GetValueByDatarefName(clientId.(string), dataref, alias, &precision, c.GetBool("is_byte_array"))
	c.JSON(200, res)
}

// GetDatarefs
// @Summary  Get a list of Dataref
// @Tags     Dataref
// @Accept   json
// @Produce  json
// @Success  200  {object}  []models.DatarefValue
// @Failure  501  "Not Implemented"
// @Router   /xplm/datarefs [post]
// @Security Oauth2Application[]
func (u DatarefController) GetDatarefs(c *gin.Context) {
	c.JSON(501, "not implemented")
}

// SetDataref
// @Summary  Set Dataref
// @Tags     Dataref
// @Param    request body models.SetDatarefValue true "dataref and value"
// @Accept   json
// @Produce  json
// @Failure  501  "Not Implemented"
// @Router   /xplm/dataref [put]
// @Security Oauth2Application[]
func (u DatarefController) SetDataref(c *gin.Context) {
	// get dataref and value
	var data models.SetDatarefValue
	err := c.BindJSON(&data)
	u.logger.Infof("dataref: %+v", data)
	if err != nil {
		u.logger.Errorf("dataref: %+v", err)
		c.JSON(500, utils.ResponseError{Message: err.Error()})
		return
	}
	// check hub if we have an active connection
	// Retrieve the value from the context
	clientId, exists := c.Get("clientId")
	if !exists {
		// Handle the case where "clientId" is not set
		c.JSON(http.StatusUnauthorized, gin.H{"error": "client ID not found"})
		return
	}
	if !u.webSocketService.IsClientExist(clientId.(string)) {
		c.JSON(http.StatusNotFound, gin.H{"error": "X Plane not connected"})
		return
	}
	u.datarefSvc.SetValueByDatarefName(clientId.(string), data.Dataref, data.Value)
}

// SetDatarefs
// @Summary  Set a list of Dataref
// @Tags     Dataref
// @Accept   json
// @Produce  json
// @Failure  501  "Not Implemented"
// @Router   /xplm/datarefs [put]
// @Security Oauth2Application[]
func (u DatarefController) SetDatarefs(c *gin.Context) {
	c.JSON(501, "not implemented")
}

// SendCommand
// @Summary  Send command to X Plane
// @Tags     Dataref
// @Param    request body models.SendCommandReq true "xplane command"
// @Accept   json
// @Produce  json
// @Failure  501  "Not Implemented"
// @Router   /xplm/command [put]
// @Security Oauth2Application[]
func (u DatarefController) SendCommand(c *gin.Context) {
	// get dataref and value
	var data models.SendCommandReq
	err := c.BindJSON(&data)
	u.logger.Infof("command: %+v", data)
	if err != nil {
		u.logger.Errorf("command: %+v", err)
		c.JSON(500, utils.ResponseError{Message: err.Error()})
		return
	}
	// check hub if we have an active connection
	// Retrieve the value from the context
	clientId, exists := c.Get("clientId")
	if !exists {
		// Handle the case where "clientId" is not set
		c.JSON(http.StatusUnauthorized, gin.H{"error": "client ID not found"})
		return
	}
	if !u.webSocketService.IsClientExist(clientId.(string)) {
		c.JSON(http.StatusNotFound, gin.H{"error": "X Plane not connected"})
		return
	}
	u.datarefSvc.SendCommand(clientId.(string), data.Command)
}
