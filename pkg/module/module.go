package module

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider"
	"github.com/kazmerdome/go-graphql-starter/pkg/observer/logger"
)

type ModuleConfig interface {
	GetProviderConfig() provider.ProviderConfig

	HasProviderOverwriter(providerType provider.ProviderType) bool
	GetProviderOverwriter(providerType provider.ProviderType) interface{}
	SetProviderOverwriter(providerType provider.ProviderType, value interface{})
}

type moduleConfig struct {
	logger              logger.Logger
	config              config.Config
	providerOverwriters map[provider.ProviderType]interface{}
}

func NewModuleConfig(l logger.Logger, c config.Config) ModuleConfig {
	return &moduleConfig{l, c, map[provider.ProviderType]interface{}{}}
}

func (r *moduleConfig) GetProviderConfig() provider.ProviderConfig {
	return provider.NewProviderConfig(r.logger, r.config)
}

func (r *moduleConfig) HasProviderOverwriter(providerType provider.ProviderType) bool {
	_, ok := r.providerOverwriters[providerType]
	return ok
}

func (r *moduleConfig) GetProviderOverwriter(providerType provider.ProviderType) interface{} {
	return r.providerOverwriters[providerType]
}

func (r *moduleConfig) SetProviderOverwriter(providerType provider.ProviderType, value interface{}) {
	r.providerOverwriters[providerType] = value
}
