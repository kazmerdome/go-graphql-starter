package licence

import (
	"context"
	"errors"
	"time"

	"github.com/kazmerdome/go-graphql-starter/pkg/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/misc"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/validator"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/zeroval"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB_COLLECTION_NAME = "Licence"
)
const (
	ERR_ALREADY_EXIST = "item is already exist"
	ERR_INTERNAL      = "internal server error"
	ERR_INVALID_ID    = "invalid id"
	ERR_NOT_FOUND     = "not found"
)

var SEARCH_FILEDS = []string{"id"}

type LicenceRepository interface {
	One(filter *LicenceWhereDTO) (*Licence, error)
	List(filter *LicenceWhereDTO, orderBy *LicenceOrderByENUM, skip *int, limit *int, customQuery *bson.M) ([]*Licence, error)
	Count(filter *LicenceWhereDTO) (*int, error)
	Create(data *LicenceCreateDTO) (*Licence, error)
	Update(where primitive.ObjectID, data *LicenceUpdateDTO) (*Licence, error)
	Delete(where primitive.ObjectID) (*Licence, error)
}

type licenceRepository struct {
	shared.SharedService
	c repository.MongoCollection
}

func NewLicenceRepository(s shared.SharedService, db repository.MongoDatabase) LicenceRepository {
	c := db.Collection(DB_COLLECTION_NAME)
	return &licenceRepository{
		SharedService: s,
		c:             c,
	}
}

// One
func (r *licenceRepository) One(filter *LicenceWhereDTO) (*Licence, error) {
	var item *Licence
	if filter == nil {
		filter = &LicenceWhereDTO{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r.c.FindOne(ctx, &filter).Decode(&item)
	return item, nil
}

// List
func (r *licenceRepository) List(filter *LicenceWhereDTO, orderBy *LicenceOrderByENUM, skip *int, limit *int, customQuery *bson.M) ([]*Licence, error) {
	var items []*Licence

	if filter == nil {
		filter = &LicenceWhereDTO{}
	}

	orderByKey := "created_at"
	orderByValue := -1

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	options := options.Find()
	if limit != nil {
		options.SetLimit(int64(*limit))
	}
	if skip != nil {
		options.SetSkip(int64(*skip))
	}
	if orderBy != nil {
		orderByKey, orderByValue = misc.GetOrderByKeyAndValue(string(*orderBy))
	}
	options.SetSort(map[string]int{orderByKey: orderByValue})

	var queryFilter interface{}
	if filter != nil {
		queryFilter = filter
	}
	if !zeroval.IsZeroVal(customQuery) {
		queryFilter = customQuery
	}

	cursor, err := r.c.Find(ctx, &queryFilter, options)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// Count
func (r *licenceRepository) Count(filter *LicenceWhereDTO) (*int, error) {
	var c *int

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if filter == nil {
		filter = &LicenceWhereDTO{}
	}

	count, err := r.c.CountDocuments(ctx, filter, nil)
	if err != nil {
		return c, err
	}

	ic := int(count)
	c = &ic

	return c, nil
}

// Create
func (r *licenceRepository) Create(data *LicenceCreateDTO) (*Licence, error) {
	item := new(Licence)
	now := time.Now()
	data.CreatedAt = now
	data.UpdatedAt = now

	// validate
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// operation
	res, err := r.c.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New(ERR_INTERNAL)
	}

	// provide new item
	r.c.FindOne(ctx, bson.M{"_id": oid}).Decode(&item)

	return item, nil
}

// Update
func (r *licenceRepository) Update(where primitive.ObjectID, data *LicenceUpdateDTO) (*Licence, error) {
	item := new(Licence)
	data.UpdatedAt = time.Now()

	// validate
	if zeroval.IsZeroVal(where) {
		return nil, errors.New(ERR_INTERNAL)
	}
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	// ctx
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check user is available
	eu := new(Licence)
	r.c.FindOne(ctx, bson.M{"_id": where}).Decode(&eu)
	if eu.ID.Hex() == "" {
		return nil, errors.New(ERR_ALREADY_EXIST)
	}

	// if array items are not provided from dto keep old data
	// instead of bson omitempty
	if zeroval.IsZeroVal(data.Grants) {
		oldItem := new(Licence)
		r.c.FindOne(ctx, bson.M{"_id": where}).Decode(&oldItem)

		var grantArray []*Grant
		for _, i := range oldItem.Grants {
			newItem := Grant{
				Feature:     i.Feature,
				Permissions: i.Permissions,
				Version:     i.Version,
			}
			grantArray = append(grantArray, &newItem)
		}
		data.Grants = grantArray
	}

	// operation
	_, err := r.c.UpdateOne(ctx, bson.M{"_id": where}, bson.M{"$set": data})
	if err != nil {
		return nil, err
	}
	r.c.FindOne(ctx, bson.M{"_id": where}).Decode(&item)
	return item, nil
}

// Delete
func (r *licenceRepository) Delete(where primitive.ObjectID) (*Licence, error) {
	item := new(Licence)

	// validate
	if zeroval.IsZeroVal(where) {
		return nil, errors.New(ERR_INVALID_ID)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r.c.FindOne(ctx, bson.M{"_id": where}).Decode(&item)
	if item.ID.Hex() == "" {
		return nil, errors.New(ERR_NOT_FOUND)
	}
	_, err := r.c.DeleteOne(ctx, bson.M{"_id": item.ID})
	if err != nil {
		return nil, err
	}
	return item, nil
}
