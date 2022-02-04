package provider

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/observe/logger"
)

type ProviderType string

const (
	Service    ProviderType = "SERVICE"
	Repository ProviderType = "REPOSITORY"
	Guard      ProviderType = "GUARD"
)

type ProviderConfig struct {
	Logger logger.Logger
	Config config.Config
}

func NewProviderConfig(l logger.Logger, c config.Config) *ProviderConfig {
	return &ProviderConfig{l, c}
}
