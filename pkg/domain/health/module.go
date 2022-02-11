package health

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	echoHandler "github.com/kazmerdome/go-graphql-starter/pkg/module/provider/handler/echo"
)

type HealthModule interface {
	GetService() HealthService
	GetEchoHttpHandler() echoHandler.EchoHandler
}

type healthModule struct {
	service     HealthService
	echoHandler echoHandler.EchoHandler
}

func NewHealthModule(moduleConfig module.ModuleConfig) HealthModule {
	m := new(healthModule)
	// Service
	m.service = newHealthService(moduleConfig.GetProviderConfig())
	// Handlers
	m.echoHandler = newHealthHandler(moduleConfig.GetProviderConfig(), m.service)
	return m
}

// Provider: Service
func (r *healthModule) GetService() HealthService {
	return r.service
}

func (r *healthModule) GetEchoHttpHandler() echoHandler.EchoHandler {
	return r.echoHandler
}
