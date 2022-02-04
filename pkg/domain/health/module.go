package health

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/handler"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/service"
	"github.com/kazmerdome/go-graphql-starter/pkg/server"
)

type HealthModule interface {
	GetService() HealthService
	GetHandler() server.Handler
}

type healthModule struct {
	service HealthService
	handler server.Handler
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

// Provider: Handler
func (r *healthModule) GetHandler() server.Handler {
	return r.handler
}
