package user

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/service"
)

type UserModule interface {
	GetService() UserService
	// GetRepository() UserRepository
}

type userModule struct {
	userService    UserService
	userRepository UserRepository
}

func NewUserModule(moduleConfig module.ModuleConfig, mongodbAdapter mongodb.MongodbAdapter) UserModule {
	m := new(userModule)
	providerConfig := moduleConfig.GetProviderConfig()
	// Repository
	if moduleConfig.HasProviderOverwriter(provider.Repository) {
		m.userRepository = moduleConfig.GetProviderOverwriter(provider.Repository).(UserRepository)
	} else {
		m.userRepository = newUserRepository(repository.NewRepositoryConfig(providerConfig, mongodbAdapter))
	}
	// Service
	m.userService = newUserService(service.NewServiceConfig(providerConfig), m.userRepository)

	return m
}

func (r *userModule) GetService() UserService {
	return r.userService
}

// func (r *userModule) GetRepository() UserRepository {
// 	return r.userRepository
// }
