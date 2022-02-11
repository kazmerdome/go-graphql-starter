package gateway

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/connector"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/handler"
)

type GatewayModule interface {
	GetEchoHttpHandler() GatewayHandler
}

type gatewayModule struct {
	handler GatewayHandler
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
		handler.NewHandlerConfig(providerConfig),
		graphqlEndpoint,
		playgroundPassword,
		modules,
	)

	return m
}

// Provider: Handler
func (r *gatewayModule) GetEchoHttpHandler() GatewayHandler {
	return r.handler
}
