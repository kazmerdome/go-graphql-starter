package gateway

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/connector"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/handler"
	"github.com/kazmerdome/go-graphql-starter/pkg/server"
)

type GatewayModule interface {
	GetHandler() server.Handler
}

type gatewayModule struct {
	handler server.Handler
}

func NewGatewayModule(
	moduleConfig module.ModuleConfig,
	graphqlEndpoint string,
	playgroundPassword string,
	modules connector.GatewayModules,
) GatewayModule {
	m := new(gatewayModule)
	providerConfig := moduleConfig.GetProviderConfig()

	// Handler
	m.handler = newGatewayHandler(
		handler.NewHandlerConfig(*providerConfig),
		graphqlEndpoint,
		playgroundPassword,
		modules,
	)

	return m
}

// Provider: Handler
func (r *gatewayModule) GetHandler() server.Handler {
	return r.handler
}
