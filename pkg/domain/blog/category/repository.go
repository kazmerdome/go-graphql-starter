package category

import (
	"context"
	"errors"
	"time"

	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/misc"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/validator"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/zeroval"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB_COLLECTION_NAME = "Category"
)
const (
	ERR_ALREADY_EXIST = "item is already exist"
	ERR_INTERNAL      = "internal server error"
	ERR_INVALID_ID    = "invalid id"
	ERR_NOT_FOUND     = "not found"
)

type CategoryRepository interface {
	One(filter *CategoryWhereDTO) (*Category, error)
	List(filter *CategoryWhereDTO, orderBy *CategoryOrderByENUM, skip *int, limit *int, customQuery *bson.M) ([]*Category, error)
	Count(filter *CategoryWhereDTO) (*int, error)
	Create(data *Category) (*Category, error)
	Update(where primitive.ObjectID, data *Category) (*Category, error)
	Delete(where primitive.ObjectID) (*Category, error)
}

type categoryRepository struct {
	repository.RepositoryConfig
	categoryCollection mongodb.MongoCollection
}

func newCategoryRepository(c repository.RepositoryConfig) CategoryRepository {
	cc := c.GetMongoAdapter().Collection(DB_COLLECTION_NAME)
	return &categoryRepository{
		RepositoryConfig:   c,
		categoryCollection: cc,
	}
}

// One
func (r *categoryRepository) One(filter *CategoryWhereDTO) (*Category, error) {
	var item *Category
	if filter == nil {
		filter = &CategoryWhereDTO{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r.categoryCollection.FindOne(ctx, &filter).Decode(&item)
	return item, nil
}

// List
func (r *categoryRepository) List(filter *CategoryWhereDTO, orderBy *CategoryOrderByENUM, skip *int, limit *int, customQuery *bson.M) ([]*Category, error) {
	var items []*Category
	if filter == nil {
		filter = &CategoryWhereDTO{}
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

	cursor, err := r.categoryCollection.Find(ctx, &queryFilter, options)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// Count
func (r *categoryRepository) Count(filter *CategoryWhereDTO) (*int, error) {
	var c *int

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if filter == nil {
		filter = &CategoryWhereDTO{}
	}

	count, err := r.categoryCollection.CountDocuments(ctx, filter, nil)
	if err != nil {
		return c, err
	}

	ic := int(count)
	c = &ic

	return c, nil
}

// Create
func (r *categoryRepository) Create(data *Category) (*Category, error) {
	item := new(Category)
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()

	// validate
	if err := validator.Validate(data); err != nil {
		return nil, err
	}

	// collection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// check uniqe
	existedItem := new(Category)
	f := bson.M{
		"$or": []bson.M{
			{"title": data.Title},
			{"slug": data.Slug},
		},
	}

	r.categoryCollection.FindOne(ctx, f).Decode(&existedItem)

	if existedItem.Slug != "" {
		return nil, errors.New(ERR_ALREADY_EXIST)
	}

	// operation
	res, err := r.categoryCollection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New(ERR_INTERNAL)
	}

	// provie new item
	r.categoryCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&item)

	return item, nil
}

// Update
func (r *categoryRepository) Update(where primitive.ObjectID, data *Category) (*Category, error) {
	item := new(Category)
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

	// check unique
	existedItem := new(Category)
	// f := bson.M{"locale": data.Locale, "_id": bson.M{"$ne": where}}
	f := bson.M{
		"_id": bson.M{"$ne": where},
		"$or": []bson.M{
			{"title": data.Title},
			{"slug": data.Slug},
		},
	}
	r.categoryCollection.FindOne(ctx, f).Decode(&existedItem)
	if existedItem.Slug != "" {
		return nil, errors.New(ERR_ALREADY_EXIST)
	}

	// operation
	_, err := r.categoryCollection.UpdateOne(ctx, bson.M{"_id": where}, bson.M{"$set": data})
	if err != nil {
		return nil, err
	}
	r.categoryCollection.FindOne(ctx, bson.M{"_id": where}).Decode(&item)
	return item, nil
}

// Delete
func (r *categoryRepository) Delete(where primitive.ObjectID) (*Category, error) {
	item := new(Category)

	// validate
	if zeroval.IsZeroVal(where) {
		return nil, errors.New(ERR_INVALID_ID)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r.categoryCollection.FindOne(ctx, bson.M{"_id": where}).Decode(&item)
	if item.Slug == "" {
		return nil, errors.New(ERR_NOT_FOUND)
	}
	_, err := r.categoryCollection.DeleteOne(ctx, bson.M{"_id": item.ID})
	if err != nil {
		return nil, err
	}
	return item, nil
}
