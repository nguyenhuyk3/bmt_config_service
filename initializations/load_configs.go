package initializations

import (
	"config_service/global"
	"fmt"

	"github.com/spf13/viper"
)

func loadConfigs() {
	viper := viper.New()
	// File name
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	// Path to config
	viper.AddConfigPath("configs")
	// Read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config: %w", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode configuration %v", err)
	}
	fmt.Println(global.Config)
}
