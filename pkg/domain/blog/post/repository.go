package post

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
	ERR_ALREADY_EXIST = "item is already exist"
	ERR_INTERNAL      = "internal server error"
	ERR_INVALID_ID    = "invalid id"
	ERR_NOT_FOUND     = "not found"
)

type PostRepository interface {
	One(filter *PostWhereDTO) (*Post, error)
	List(filter *PostWhereDTO, orderBy *PostOrderByENUM, skip *int, limit *int, customQuery *bson.M) ([]*Post, error)
	Count(filter *PostWhereDTO) (*int, error)
	Create(data *Post) (*Post, error)
	Update(where primitive.ObjectID, data *Post) (*Post, error)
	Delete(where primitive.ObjectID) (*Post, error)
}

type postRepository struct {
	*repository.RepositoryConfig
	postCollection mongodb.MongoCollection
}

func newPostRepository(c *repository.RepositoryConfig) PostRepository {
	cc := c.Adapters.MongodbAdapter.Collection(DB_COLLECTION_NAME)
	return &postRepository{
		RepositoryConfig: c,
		postCollection:   cc,
	}
}

// One
func (r *postRepository) One(filter *PostWhereDTO) (*Post, error) {
	var item *Post
	if filter == nil {
		filter = &PostWhereDTO{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r.postCollection.FindOne(ctx, &filter).Decode(&item)
	return item, nil
}

// List
func (r *postRepository) List(filter *PostWhereDTO, orderBy *PostOrderByENUM, skip *int, limit *int, customQuery *bson.M) ([]*Post, error) {
	var items []*Post

	if filter == nil {
		filter = &PostWhereDTO{}
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

	cursor, err := r.postCollection.Find(ctx, &queryFilter, options)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

// Count
func (r *postRepository) Count(filter *PostWhereDTO) (*int, error) {
	var c *int

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if filter == nil {
		filter = &PostWhereDTO{}
	}

	count, err := r.postCollection.CountDocuments(ctx, filter, nil)
	if err != nil {
		return c, err
	}

	ic := int(count)
	c = &ic

	return c, nil
}

// Create
func (r *postRepository) Create(data *Post) (*Post, error) {
	item := new(Post)
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
	existedItem := new(Post)
	f := bson.M{
		"$or": []bson.M{
			{"title": data.Title},
			{"slug": data.Slug},
		},
	}

	r.postCollection.FindOne(ctx, f).Decode(&existedItem)

	if existedItem.Slug != "" {
		return nil, errors.New(ERR_ALREADY_EXIST)
	}

	// operation
	res, err := r.postCollection.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New(ERR_INTERNAL)
	}

	// provie new item
	r.postCollection.FindOne(ctx, bson.M{"_id": oid}).Decode(&item)

	return item, nil
}

// Update
func (r *postRepository) Update(where primitive.ObjectID, data *Post) (*Post, error) {
	item := new(Post)
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
	existedItem := new(Post)
	// f := bson.M{"locale": data.Locale, "_id": bson.M{"$ne": where}}
	f := bson.M{
		"_id": bson.M{"$ne": where},
		"$or": []bson.M{
			{"title": data.Title},
			{"slug": data.Slug},
		},
	}
	r.postCollection.FindOne(ctx, f).Decode(&existedItem)
	if existedItem.Slug != "" {
		return nil, errors.New(ERR_ALREADY_EXIST)
	}

	// operation
	_, err := r.postCollection.UpdateOne(ctx, bson.M{"_id": where}, bson.M{"$set": data})
	if err != nil {
		return nil, err
	}
	r.postCollection.FindOne(ctx, bson.M{"_id": where}).Decode(&item)
	return item, nil
}

// Delete
func (r *postRepository) Delete(where primitive.ObjectID) (*Post, error) {
	item := new(Post)

	// validate
	if zeroval.IsZeroVal(where) {
		return nil, errors.New(ERR_INVALID_ID)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r.postCollection.FindOne(ctx, bson.M{"_id": where}).Decode(&item)
	if item.Slug == "" {
		return nil, errors.New(ERR_NOT_FOUND)
	}
	_, err := r.postCollection.DeleteOne(ctx, bson.M{"_id": item.ID})
	if err != nil {
		return nil, err
	}
	return item, nil
}
