package handler

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider"
)

type HandlerConfig interface {
	provider.ProviderConfig
}

type handlerConfig struct {
	provider.ProviderConfig
}

func NewHandlerConfig(config provider.ProviderConfig) HandlerConfig {
	return &handlerConfig{ProviderConfig: config}
}
