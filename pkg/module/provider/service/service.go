package service

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider"
)

type ServiceConfig interface {
	provider.ProviderConfig
}
type serviceConfig struct {
	provider.ProviderConfig
}

func NewServiceConfig(c provider.ProviderConfig) ServiceConfig {
	return &serviceConfig{ProviderConfig: c}
}
