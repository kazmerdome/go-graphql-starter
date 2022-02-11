package post

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/service"
)

type PostModule interface {
	GetService() PostService
	GetRepository() PostRepository
}

type postModule struct {
	service    PostService
	repository PostRepository
}

func NewPostModule(moduleConfig module.ModuleConfig, adapters module.Adapters) PostModule {
	m := new(postModule)
	providerConfig := moduleConfig.GetProviderConfig()
	// Repository
	if moduleConfig.HasProviderOverwriter(provider.Repository) {
		m.repository = moduleConfig.GetProviderOverwriter(provider.Repository).(PostRepository)
	} else {
		m.repository = newPostRepository(repository.NewRepositoryConfig(providerConfig, adapters))
	}
	// Service
	m.service = newPostService(service.NewServiceConfig(providerConfig), m.repository)
	return m
}

func (r *postModule) GetService() PostService {
	return r.service
}

func (r *postModule) GetRepository() PostRepository {
	return r.repository
}
