package initializations

import (
	"config_service/global"
	"config_service/routers"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		// Force color for console
		gin.ForceConsoleColor()
		// Create a router with default middleware (Logger and Recovery)
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		// Create a new router without default middleware
		r = gin.New()
	}

	configRouter := routers.RouterGroupApp.Config

	mainGroup := r.Group("v1")
	{
		configRouter.InitConfigRouter(mainGroup)
	}

	return r
}
