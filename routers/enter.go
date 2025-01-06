package routers

import "config_service/routers/config"

type RouterGroup struct {
	Config config.ConfigRouter
}

var RouterGroupApp = new(RouterGroup)
