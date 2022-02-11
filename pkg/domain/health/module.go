package health

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/handler"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/service"
)

type HealthModule interface {
	GetService() HealthService
	GetEchoHttpHandler() HealthHandler
}

type healthModule struct {
	service HealthService
	handler HealthHandler
}

func NewHealthModule(moduleConfig module.ModuleConfig) HealthModule {
	m := new(healthModule)
	providerConfig := moduleConfig.GetProviderConfig()
	// Service
	m.service = newHealthService(service.NewServiceConfig(providerConfig))
	// Handler
	m.handler = newHealthHandler(handler.NewHandlerConfig(providerConfig), m.service)
	return m
}

// Provider: Service
func (r *healthModule) GetService() HealthService {
	return r.service
}

func (r *healthModule) GetEchoHttpHandler() HealthHandler {
	return r.handler
}
