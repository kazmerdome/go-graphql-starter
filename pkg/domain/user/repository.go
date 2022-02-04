package user

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/repository"
)

const (
	DB_COLLECTION_NAME = "User"
)
const (
	ERR_ALREADY_EXIST = "item is already exist"
	ERR_INTERNAL      = "internal server error"
	ERR_INVALID_ID    = "invalid id"
	ERR_NOT_FOUND     = "not found"
)

var SEARCH_FILEDS = []string{"id"}

type UserRepository interface {
	// One(filter *LicenceWhereDTO) (*Licence, error)
	// List(filter *LicenceWhereDTO, orderBy *LicenceOrderByENUM, skip *int, limit *int, customQuery *bson.M) ([]*Licence, error)
	// Count(filter *LicenceWhereDTO) (*int, error)
	// Create(data *LicenceCreateDTO) (*Licence, error)
	// Update(where primitive.ObjectID, data *LicenceUpdateDTO) (*Licence, error)
	// Delete(where primitive.ObjectID) (*Licence, error)
}

type userRepository struct {
	*repository.RepositoryConfig
	userCollection mongodb.MongoCollection
}

func newUserRepository(c *repository.RepositoryConfig) UserRepository {
	userCollection := c.Adapters.MongodbAdapter.Collection(DB_COLLECTION_NAME)
	return &userRepository{
		RepositoryConfig: c,
		userCollection:   userCollection,
	}
}
