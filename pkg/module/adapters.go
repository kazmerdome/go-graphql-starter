package module

import "github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"

type Repository struct {
	MongodbAdapter mongodb.MongodbAdapter
}

type Adapters interface {
	GetRepositoryMongodbAdapter() mongodb.MongodbAdapter
}

type adapters struct {
	repository struct {
		mongodb mongodb.MongodbAdapter
	}
}

func NewAdapters(mongodb mongodb.MongodbAdapter) Adapters {
	e := new(adapters)
	e.repository.mongodb = mongodb
	return e
}

func (r *adapters) GetRepositoryMongodbAdapter() mongodb.MongodbAdapter {
	return r.repository.mongodb
}
