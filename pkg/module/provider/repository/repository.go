package repository

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider"
)

type RepositoryConfig interface {
	provider.ProviderConfig
	module.Adapters
}

type repositoryConfig struct {
	provider.ProviderConfig
	module.Adapters
}

func NewRepositoryConfig(config provider.ProviderConfig, adapters module.Adapters) RepositoryConfig {
	return &repositoryConfig{
		ProviderConfig: config,
		Adapters:       adapters,
	}
}
