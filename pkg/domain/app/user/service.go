package user

// import (
// 	"context"
// 	"errors"
// 	"time"

// 	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/identity"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/repository"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/util/misc"
// )

// const (
// 	DB_COLLECTION_NAME = "User"
// 	ERR_ACCESS_DENIED  = "access denied"
// )

// type UserService interface {
// 	UserMe(ctx context.Context, bearerToken string) (*User, error)
// 	User(ctx context.Context, where *UserWhereDTO, search *string) (*User, error)
// 	Users(ctx context.Context, where *UserWhereDTO, orderBy *UserOrderByENUM, skip *int, limit *int, search *string) ([]*User, error)
// 	UserCount(ctx context.Context, where *UserWhereDTO, search *string) (*int, error)
// 	CreateUser(ctx context.Context, data UserCreateDTO) (*User, error)
// 	UpdateUser(ctx context.Context, where UserWhereUniqueDTO, data UserUpdateDTO) (*User, error)
// 	DeleteUser(ctx context.Context, where UserWhereUniqueDTO) (*User, error)
// }

// type userService struct {
// 	shared.SharedService
// 	identityService identity.IdentityService
// 	r               UserRepository
// }

// func NewUserService(s shared.SharedService, i identity.IdentityService, db repository.MongoDatabase) UserService {
// 	r := NewUserRepository(s, db.Collection(DB_COLLECTION_NAME))
// 	return &userService{SharedService: s, identityService: i, r: r}
// }

// // UserMe
// func (r *userService) UserMe(ctx context.Context, bearerToken string) (*User, error) {
// 	userID, err := r.identityService.GetIdGuard(bearerToken)
// 	if err != nil {
// 		return nil, err
// 	}
// 	where := UserWhereDTO{}
// 	where.ID = &userID

// 	data, err := r.r.One(&where)
// 	if err != nil || data.Email == "" {
// 		return nil, errors.New(ERR_ACCESS_DENIED)
// 	}
// 	// Update last-seen
// 	now := time.Now()
// 	seenWhere := UserUpdateDTO{}
// 	seenWhere.LastSeen = &now
// 	r.r.Update(data.ID, &seenWhere)

// 	return data, nil
// }

// // User
// func (r *userService) User(ctx context.Context, where *UserWhereDTO, search *string) (*User, error) {
// 	if search != nil {
// 		where.OR = misc.MongoSearchFieldParser(SEARCH_FILEDS, *search)
// 	}
// 	return r.r.One(where)
// }

// // Users
// func (r *userService) Users(ctx context.Context, where *UserWhereDTO, orderBy *UserOrderByENUM, skip *int, limit *int, search *string) ([]*User, error) {
// 	if search != nil {
// 		where.OR = misc.MongoSearchFieldParser(SEARCH_FILEDS, *search)
// 	}
// 	return r.r.List(where, orderBy, skip, limit, nil)
// }

// // UserCount
// func (r *userService) UserCount(ctx context.Context, where *UserWhereDTO, search *string) (*int, error) {
// 	if search != nil {
// 		where.OR = misc.MongoSearchFieldParser(SEARCH_FILEDS, *search)
// 	}
// 	return r.r.Count(where)
// }

// // CreateUser
// func (r *userService) CreateUser(ctx context.Context, data UserCreateDTO) (*User, error) {
// 	return r.r.Create(&data)
// }

// // UpdateUser
// func (r *userService) UpdateUser(ctx context.Context, where UserWhereUniqueDTO, data UserUpdateDTO) (*User, error) {
// 	return r.r.Update(where.ID, &data)
// }

// // DeleteUser
// func (r *userService) DeleteUser(ctx context.Context, where UserWhereUniqueDTO) (*User, error) {
// 	return r.r.Delete(where.ID)
// }
