package controllers

import (
	"config_service/internal/services"
	"config_service/pkgs/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConfigController struct {
	configService services.IConfigService
}

func NewConfigController(configService services.IConfigService) *ConfigController {
	return &ConfigController{
		configService: configService,
	}
}

func (cc *ConfigController) GetServiceConfig(c *gin.Context) {
	serviceName := c.Query("service_name")
	config, err := cc.configService.GetServiceConfig(serviceName)
	if err != nil {
		responses.FailureResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	responses.SuccessResponse(c, http.StatusOK, "data returned successfully", config)
}
