package gateway

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/connector"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	echoHandler "github.com/kazmerdome/go-graphql-starter/pkg/module/provider/handler/echo"
)

type GatewayModule interface {
	GetEchoHttpHandler() echoHandler.EchoHandler
}

type gatewayModule struct {
	echoHandler echoHandler.EchoHandler
}

func NewGatewayModule(
	moduleConfig module.ModuleConfig,
	graphqlEndpoint string,
	playgroundPassword string,
	modules connector.GatewayModules,
) GatewayModule {
	m := new(gatewayModule)

	// Handlers
	m.echoHandler = newGatewayHandler(
		moduleConfig.GetProviderConfig(),
		graphqlEndpoint,
		playgroundPassword,
		modules,
	)

	return m
}

// Provider: Handler
func (r *gatewayModule) GetEchoHttpHandler() echoHandler.EchoHandler {
	return r.echoHandler
}
