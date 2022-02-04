package repository

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
)

type RepositoryConfig struct {
	*provider.ProviderConfig
	Adapters struct {
		MongodbAdapter mongodb.MongodbAdapter
	}
}

func NewRepositoryConfig(pc *provider.ProviderConfig, mongodbAdapter mongodb.MongodbAdapter) *RepositoryConfig {
	return &RepositoryConfig{
		ProviderConfig: pc,
		Adapters:       struct{ MongodbAdapter mongodb.MongodbAdapter }{mongodbAdapter},
	}
}
