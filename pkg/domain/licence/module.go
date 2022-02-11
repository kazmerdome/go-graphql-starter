package licence

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/guard"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/service"
)

type LicenceModule interface {
	GetService() LicenceService
	GetRepository() LicenceRepository
	GetGuard() LicenceGuard
}

type licenceModule struct {
	licenceService    LicenceService
	licenceRepository LicenceRepository
	licenceGuard      LicenceGuard
}

func NewLicenceModule(moduleConfig module.ModuleConfig, adapters module.Adapters) LicenceModule {
	m := new(licenceModule)
	providerConfig := moduleConfig.GetProviderConfig()
	// Repository
	if moduleConfig.HasProviderOverwriter(provider.Repository) {
		m.licenceRepository = moduleConfig.GetProviderOverwriter(provider.Repository).(LicenceRepository)
	} else {
		m.licenceRepository = newLicenceRepository(repository.NewRepositoryConfig(providerConfig, adapters))
	}
	// Service
	m.licenceService = newLicenceService(service.NewServiceConfig(providerConfig), m.licenceRepository)
	// Guard
	m.licenceGuard = newLicenceGuard(guard.NewGuardConfig(providerConfig), m.licenceRepository)
	return m
}

// Provider: Service
func (r *licenceModule) GetService() LicenceService {
	return r.licenceService
}

// Provider: Repository
func (r *licenceModule) GetRepository() LicenceRepository {
	return r.licenceRepository
}

// Provider: Guard
func (r *licenceModule) GetGuard() LicenceGuard {
	return r.licenceGuard
}
