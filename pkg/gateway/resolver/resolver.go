package resolver

import "github.com/kazmerdome/go-graphql-starter/pkg/gateway/services"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	services  services.GatewayServices
	AuthToken *string
}

func NewResolver(authToken *string, services services.GatewayServices) *Resolver {
	return &Resolver{
		services:  services,
		AuthToken: authToken,
	}
}
