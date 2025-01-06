package initializations

import (
	"config_service/global"
	"fmt"
)

func Run() {
	loadConfigs()

	r := initRouter()

	r.Run(fmt.Sprintf("localhost:%s", global.Config.Server.Port))
}
