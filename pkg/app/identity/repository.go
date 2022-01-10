package identity

// import (
// 	"context"
// 	"errors"
// 	"time"

// 	"github.com/kazmerdome/go-graphql-starter/pkg/repository"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
// )

// const (
// 	ERR_ALREADY_EXIST = "item is already exist"
// 	ERR_INTERNAL      = "internal server error"
// 	ERR_INVALID_ID    = "invalid id"
// 	ERR_NOT_FOUND     = "not found"
// )

// type IdentityRepository interface {
// 	OneById(filter *IdentityWhereUniqueDTO) (*Identity, error)
// }

// type identityRepository struct {
// 	shared.SharedService
// 	c repository.MongoCollection
// }

// func NewIdentityRepository(s shared.SharedService, c repository.MongoCollection) IdentityRepository {
// 	return &identityRepository{
// 		SharedService: s,
// 		c:             c,
// 	}
// }

// // OneById
// func (r *identityRepository) OneById(filter *IdentityWhereUniqueDTO) (*Identity, error) {
// 	if filter == nil {
// 		return nil, errors.New(ERR_INTERNAL)
// 	}
// 	if filter.ID.Hex() == "" {
// 		return nil, errors.New(ERR_INTERNAL)
// 	}

// 	var item *Identity

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	r.c.FindOne(ctx, &filter).Decode(&item)
// 	return item, nil
// }
