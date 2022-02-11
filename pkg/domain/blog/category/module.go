package category

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/service"
)

type CategoryModule interface {
	GetService() CategoryService
	GetRepository() CategoryRepository
}

type categoryModule struct {
	service    CategoryService
	repository CategoryRepository
}

func NewCategoryModule(moduleConfig module.ModuleConfig, adapters module.Adapters) CategoryModule {
	m := new(categoryModule)
	providerConfig := moduleConfig.GetProviderConfig()
	// Repository
	if moduleConfig.HasProviderOverwriter(provider.Repository) {
		m.repository = moduleConfig.GetProviderOverwriter(provider.Repository).(CategoryRepository)
	} else {
		m.repository = newCategoryRepository(repository.NewRepositoryConfig(providerConfig, adapters))
	}
	// Service
	m.service = newCategoryService(service.NewServiceConfig(providerConfig), m.repository)
	return m
}

func (r *categoryModule) GetService() CategoryService {
	return r.service
}

func (r *categoryModule) GetRepository() CategoryRepository {
	return r.repository
}
