package config

import (
	"config_service/internal/controllers"
	"config_service/internal/implementations"

	"github.com/gin-gonic/gin"
)

type ConfigRouter struct {
}

func (cr *ConfigRouter) InitConfigRouter(rg *gin.RouterGroup) {
	configService := implementations.NewConfigService()
	configController := controllers.NewConfigController(configService)
	configRouterPublic := rg.Group("/config")
	{
		configRouterPublic.GET("/get_service_config", configController.GetServiceConfig)
	}
}
