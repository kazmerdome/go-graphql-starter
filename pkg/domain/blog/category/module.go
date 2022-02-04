package category

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/service"
)

type CategoryModule interface {
	GetService() CategoryService
	GetRepository() CategoryRepository
}

type categoryModule struct {
	service    CategoryService
	repository CategoryRepository
}

func NewCategoryModule(moduleConfig module.ModuleConfig, mongodbAdapter mongodb.MongodbAdapter) CategoryModule {
	m := new(categoryModule)
	providerConfig := moduleConfig.GetProviderConfig()
	// Repository
	if moduleConfig.HasProviderOverwriter(provider.Repository) {
		m.repository = moduleConfig.GetProviderOverwriter(provider.Repository).(CategoryRepository)
	} else {
		m.repository = newCategoryRepository(repository.NewRepositoryConfig(providerConfig, mongodbAdapter))
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
