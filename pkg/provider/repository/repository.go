package repository

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider"
)

type RepositoryConfig interface {
	provider.ProviderConfig
	GetMongoAdapter() mongodb.MongodbAdapter
}
type repositoryConfig struct {
	provider.ProviderConfig
	adapters struct {
		mongodbAdapter mongodb.MongodbAdapter
	}
}

func NewRepositoryConfig(c provider.ProviderConfig, mongodbAdapter mongodb.MongodbAdapter) RepositoryConfig {
	return &repositoryConfig{
		ProviderConfig: c,
		adapters:       struct{ mongodbAdapter mongodb.MongodbAdapter }{mongodbAdapter},
	}
}

func (r *repositoryConfig) GetMongoAdapter() mongodb.MongodbAdapter {
	return r.adapters.mongodbAdapter
}
