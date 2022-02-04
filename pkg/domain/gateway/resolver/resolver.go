package resolver

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/connector"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	modules   connector.GatewayModules
	AuthToken *string
}

func NewResolver(authToken *string, modules connector.GatewayModules) *Resolver {
	return &Resolver{
		modules:   modules,
		AuthToken: authToken,
	}
}
