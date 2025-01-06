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
		responses.Response(c, http.StatusBadRequest, "", nil)
		return
	}

	config := cc.configService.GetServiceConfig(req.ServiceName)

	responses.Response(c, http.StatusOK, "data returned successfully", config)
}
