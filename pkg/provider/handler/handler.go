package handler

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
)

type HandlerConfig interface {
	provider.ProviderConfig
}

type handlerConfig struct {
	provider.ProviderConfig
}

func NewHandlerConfig(c provider.ProviderConfig) HandlerConfig {
	return &handlerConfig{ProviderConfig: c}
}
