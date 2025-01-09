package implementations

import (
	"config_service/global"
	"config_service/internal/services"
	"fmt"
)

type configService struct {
}

func NewConfigService() services.IConfigService {
	return &configService{}
}

// GetServiceConfig implements IConfigService.
func (c *configService) GetServiceConfig(serviceName string) (interface{}, error) {
	switch serviceName {
	case "bmt_user":
		return global.Config.BMTUser, nil
	default:
		return nil, fmt.Errorf("service %s not found", serviceName)
	}
}
