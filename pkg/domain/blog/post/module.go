package post

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/service"
)

type PostModule interface {
	GetService() PostService
	GetRepository() PostRepository
}

type postModule struct {
	service    PostService
	repository PostRepository
}

func NewPostModule(moduleConfig module.ModuleConfig, mongodbAdapter mongodb.MongodbAdapter) PostModule {
	m := new(postModule)
	providerConfig := moduleConfig.GetProviderConfig()
	// Repository
	if moduleConfig.HasProviderOverwriter(provider.Repository) {
		m.repository = moduleConfig.GetProviderOverwriter(provider.Repository).(PostRepository)
	} else {
		m.repository = newPostRepository(repository.NewRepositoryConfig(providerConfig, mongodbAdapter))
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
