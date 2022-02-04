package handler

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
)

type HandlerConfig struct {
	provider.ProviderConfig
}

func NewHandlerConfig(pc provider.ProviderConfig) *HandlerConfig {
	return &HandlerConfig{ProviderConfig: pc}
}
