package services

type IConfigService interface {
	GetServiceConfig(serviceName string) (interface{}, error)
}
