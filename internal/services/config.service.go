package services

import "config_service/global"

type IConfigService interface {
	GetServiceConfig(serviceName string) interface{}
}

type configService struct {
}

func NewConfigService() IConfigService {
	return &configService{}
}

// GetServiceConfig implements IConfigService.
func (c *configService) GetServiceConfig(serviceName string) interface{} {
	switch serviceName {
	case "bmt_user":
		return global.Config.BMTUser
	default:
		return "invalid service name"
	}
}
