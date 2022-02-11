package adapter

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/observer/logger"
)

type AdapterType string

type AdapterConfig interface {
	GetConfig() config.Config
	GetLogger() logger.Logger
}

type adapterConfig struct {
	logger logger.Logger
	config config.Config
}

func NewAdapterConfig(l logger.Logger, c config.Config) AdapterConfig {
	return &adapterConfig{l, c}
}

func (r *adapterConfig) GetConfig() config.Config {
	return r.config
}

func (r *adapterConfig) GetLogger() logger.Logger {
	return r.logger
}
