package module

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/observe/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
)

type ModuleConfig interface {
	GetProviderConfig() provider.ProviderConfig

	HasProviderOverwriter(providerType provider.ProviderType) bool
	GetProviderOverwriter(providerType provider.ProviderType) interface{}
	SetProviderOverwriter(providerType provider.ProviderType, value interface{})
}

type moduleConfig struct {
	Logger              logger.Logger
	Config              config.Config
	providerOverwriters map[provider.ProviderType]interface{}
}

func NewModuleConfig(l logger.Logger, c config.Config) ModuleConfig {
	return &moduleConfig{l, c, map[provider.ProviderType]interface{}{}}
}

func (r *moduleConfig) GetProviderConfig() provider.ProviderConfig {
	return provider.NewProviderConfig(r.Logger, r.Config)
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
