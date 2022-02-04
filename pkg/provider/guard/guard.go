package guard

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
)

type GuardConfig struct {
	*provider.ProviderConfig
}

func NewGuardConfig(pc *provider.ProviderConfig) *GuardConfig {
	return &GuardConfig{ProviderConfig: pc}
}
