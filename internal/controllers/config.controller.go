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

type GetServiceConfigReq struct {
	ServiceName string `json:"service_name"`
}

func (cc *ConfigController) GetServiceConfig(c *gin.Context) {
	var req GetServiceConfigReq
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.FailureResponse(c, http.StatusBadRequest, "invalid request")
		return
	}

	config, err := cc.configService.GetServiceConfig(req.ServiceName)
	if err != nil {
		responses.FailureResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	responses.SuccessResponse(c, http.StatusOK, "data returned successfully", config)
}
