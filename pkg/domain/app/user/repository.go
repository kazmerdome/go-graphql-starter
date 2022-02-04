package user

// import (
// 	"context"
// 	"errors"
// 	"time"

// 	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/identity"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/repository"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/util/misc"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/util/validator"
// 	"github.com/kazmerdome/go-graphql-starter/pkg/util/zeroval"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const (
// 	ERR_ALREADY_EXIST = "item is already exist"
// 	ERR_INTERNAL      = "internal server error"
// 	ERR_INVALID_ID    = "invalid id"
// 	ERR_NOT_FOUND     = "not found"
// )

// var SEARCH_FILEDS = []string{"email", "username", "firstname", "lastname"}

// type UserRepository interface {
// 	One(filter *UserWhereDTO) (*User, error)
// 	List(filter *UserWhereDTO, orderBy *UserOrderByENUM, skip *int, limit *int, customQuery *bson.M) ([]*User, error)
// 	Count(filter *UserWhereDTO) (*int, error)
// 	Create(data *UserCreateDTO) (*User, error)
// 	Update(where primitive.ObjectID, data *UserUpdateDTO) (*User, error)
// 	Delete(where primitive.ObjectID) (*User, error)
// }

// type userRepository struct {
// 	shared.SharedService
// 	c repository.MongoCollection
// }

// func NewUserRepository(s shared.SharedService, c repository.MongoCollection) UserRepository {
// 	return &userRepository{
// 		SharedService: s,
// 		c:             c,
// 	}
// }

// // One
// func (r *userRepository) One(filter *UserWhereDTO) (*User, error) {
// 	var item *User
// 	if filter == nil {
// 		filter = &UserWhereDTO{}
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	r.c.FindOne(ctx, &filter).Decode(&item)
// 	return item, nil
// }

// // List
// func (r *userRepository) List(filter *UserWhereDTO, orderBy *UserOrderByENUM, skip *int, limit *int, customQuery *bson.M) ([]*User, error) {
// 	var items []*User

// 	if filter == nil {
// 		filter = &UserWhereDTO{}
// 	}

// 	orderByKey := "created_at"
// 	orderByValue := -1

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	options := options.Find()
// 	if limit != nil {
// 		options.SetLimit(int64(*limit))
// 	}
// 	if skip != nil {
// 		options.SetSkip(int64(*skip))
// 	}
// 	if orderBy != nil {
// 		orderByKey, orderByValue = misc.GetOrderByKeyAndValue(string(*orderBy))
// 	}
// 	options.SetSort(map[string]int{orderByKey: orderByValue})

// 	var queryFilter interface{}
// 	if filter != nil {
// 		queryFilter = filter
// 	}
// 	if !zeroval.IsZeroVal(customQuery) {
// 		queryFilter = customQuery
// 	}

// 	cursor, err := r.c.Find(ctx, &queryFilter, options)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = cursor.All(ctx, &items); err != nil {
// 		return nil, err
// 	}

// 	return items, nil
// }

// // Count
// func (r *userRepository) Count(filter *UserWhereDTO) (*int, error) {
// 	var c *int

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if filter == nil {
// 		filter = &UserWhereDTO{}
// 	}

// 	count, err := r.c.CountDocuments(ctx, filter, nil)
// 	if err != nil {
// 		return c, err
// 	}

// 	ic := int(count)
// 	c = &ic

// 	return c, nil
// }

// // Create
// func (r *userRepository) Create(data *UserCreateDTO) (*User, error) {
// 	item := new(User)
// 	now := time.Now()
// 	data.CreatedAt = now
// 	data.UpdatedAt = now
// 	data.LastLogin = &now

// 	// validate
// 	if err := validator.Validate(data); err != nil {
// 		return nil, err
// 	}

// 	// ctx
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// check uniqe
// 	existedItem := new(User)
// 	f := bson.M{
// 		"$or": []bson.M{
// 			{"email": data.Email},
// 			{"username": data.Username},
// 		},
// 	}

// 	r.c.FindOne(ctx, f).Decode(&existedItem)
// 	if existedItem.Email != "" {
// 		return nil, errors.New(ERR_ALREADY_EXIST)
// 	}

// 	// operation
// 	res, err := r.c.InsertOne(ctx, data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	oid, ok := res.InsertedID.(primitive.ObjectID)
// 	if !ok {
// 		return nil, errors.New(ERR_INTERNAL)
// 	}

// 	// provide new item
// 	r.c.FindOne(ctx, bson.M{"_id": oid}).Decode(&item)

// 	return item, nil
// }

// // Update
// func (r *userRepository) Update(where primitive.ObjectID, data *UserUpdateDTO) (*User, error) {
// 	item := new(User)
// 	data.UpdatedAt = time.Now()

// 	// validate
// 	if zeroval.IsZeroVal(where) {
// 		return nil, errors.New(ERR_INTERNAL)
// 	}
// 	if err := validator.Validate(data); err != nil {
// 		return nil, err
// 	}

// 	// ctx
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	// check user is available
// 	eu := new(User)
// 	r.c.FindOne(ctx, bson.M{"_id": where}).Decode(&eu)
// 	if eu.Email == "" {
// 		return nil, errors.New(ERR_ALREADY_EXIST)
// 	}

// 	// check unique
// 	existedItem := new(User)
// 	f := bson.M{
// 		"$or": []bson.M{
// 			{"email": data.Email, "_id": bson.M{"$ne": where}},
// 			{"username": data.Username, "_id": bson.M{"$ne": where}},
// 		},
// 	}
// 	r.c.FindOne(ctx, f).Decode(&existedItem)
// 	if existedItem.Email != "" {
// 		return nil, errors.New(ERR_ALREADY_EXIST)
// 	}
// 	// if array items are not provided from dto keep old data
// 	// instead of bson omitempty
// 	if zeroval.IsZeroVal(data.AuthStrategy) || zeroval.IsZeroVal(data.Policy) {
// 		oldItem := new(User)
// 		r.c.FindOne(ctx, bson.M{"_id": where}).Decode(&oldItem)

// 		if zeroval.IsZeroVal(data.AuthStrategy) {
// 			var authStrategyArray []*identity.AuthStrategy
// 			for _, i := range oldItem.AuthStrategy {
// 				newItem := identity.AuthStrategy{
// 					Type:   i.Type,
// 					Secret: i.Secret,
// 				}
// 				authStrategyArray = append(authStrategyArray, &newItem)
// 			}
// 			data.AuthStrategy = authStrategyArray
// 		}
// 		if zeroval.IsZeroVal(data.Policy) {
// 			var policyArray []*identity.Policy
// 			for _, i := range oldItem.Policy {
// 				newItem := identity.Policy{
// 					Resource: i.Resource,
// 					Role:     i.Role,
// 				}
// 				policyArray = append(policyArray, &newItem)
// 			}
// 			data.Policy = policyArray
// 		}
// 	}
// 	// operation
// 	_, err := r.c.UpdateOne(ctx, bson.M{"_id": where}, bson.M{"$set": data})
// 	if err != nil {
// 		return nil, err
// 	}
// 	r.c.FindOne(ctx, bson.M{"_id": where}).Decode(&item)
// 	return item, nil
// }

// // Delete
// func (r *userRepository) Delete(where primitive.ObjectID) (*User, error) {
// 	item := new(User)

// 	// validate
// 	if zeroval.IsZeroVal(where) {
// 		return nil, errors.New(ERR_INVALID_ID)
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	r.c.FindOne(ctx, bson.M{"_id": where}).Decode(&item)
// 	if item.Email == "" {
// 		return nil, errors.New(ERR_NOT_FOUND)
// 	}
// 	_, err := r.c.DeleteOne(ctx, bson.M{"_id": item.ID})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return item, nil
// }
