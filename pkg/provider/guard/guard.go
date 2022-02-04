package guard

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
)

type GuardConfig interface {
	provider.ProviderConfig
}
type guardConfig struct {
	provider.ProviderConfig
}

func NewGuardConfig(c provider.ProviderConfig) GuardConfig {
	return &guardConfig{ProviderConfig: c}
}
