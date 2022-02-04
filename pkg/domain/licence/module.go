package licence

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/guard"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/service"
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

func NewLicenceModule(moduleConfig module.ModuleConfig, mongodbAdapter mongodb.MongodbAdapter) LicenceModule {
	m := new(licenceModule)
	providerConfig := moduleConfig.GetProviderConfig()
	// Repository
	if moduleConfig.HasProviderOverwriter(provider.Repository) {
		m.licenceRepository = moduleConfig.GetProviderOverwriter(provider.Repository).(LicenceRepository)
	} else {
		m.licenceRepository = newLicenceRepository(repository.NewRepositoryConfig(providerConfig, mongodbAdapter))
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
