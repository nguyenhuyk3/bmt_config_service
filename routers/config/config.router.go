package config

import (
	"config_service/internal/controllers"
	"config_service/internal/services"

	"github.com/gin-gonic/gin"
)

type ConfigRouter struct {
}

func (cr *ConfigRouter) InitConfigRouter(rg *gin.RouterGroup) {
	configService := services.NewConfigService()
	configController := controllers.NewConfigController(configService)
	configRouterPublic := rg.Group("/config")
	{
		configRouterPublic.GET("/get_service_config", configController.GetServiceConfig)
	}
}
