package service

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
)

type ServiceConfig struct {
	*provider.ProviderConfig
}

func NewServiceConfig(pc *provider.ProviderConfig) *ServiceConfig {
	return &ServiceConfig{ProviderConfig: pc}
}
