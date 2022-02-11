package provider

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/observer/logger"
)

type ProviderType string

const (
	Service    ProviderType = "SERVICE"
	Repository ProviderType = "REPOSITORY"
	Guard      ProviderType = "GUARD"
)

type ProviderConfig interface {
	GetConfig() config.Config
	GetLogger() logger.Logger
}

type providerConfig struct {
	logger logger.Logger
	config config.Config
}

func NewProviderConfig(l logger.Logger, c config.Config) ProviderConfig {
	return &providerConfig{l, c}
}

func (r *providerConfig) GetConfig() config.Config {
	return r.config
}

func (r *providerConfig) GetLogger() logger.Logger {
	return r.logger
}
