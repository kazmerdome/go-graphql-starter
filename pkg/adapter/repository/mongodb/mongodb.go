package mongodb

import (
	"context"

	"errors"
	"fmt"

	"github.com/kazmerdome/go-graphql-starter/pkg/shared"

	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongodbAdapter interface {
	Collection(name string, opts ...*options.CollectionOptions) MongoCollection
	Disconnect()
}

type MongoCollection interface {
	Drop(ctx context.Context) error
	Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult
}

type mongodbAdapter struct {
	shared.SharedService
	client   *mongo.Client
	database *mongo.Database
}

func NewMongodbAdapter(c shared.SharedService, uri string, name string, retrywrites bool) MongodbAdapter {
	if uri == "" {
		panic(errors.New("uri is required"))
	}

	if name == "" {
		panic(errors.New("database name is required"))
	}

	c.Logger.Info("connecting " + name + " db...")
	connectionURI := uri + "/" + name

	if retrywrites {
		connectionURI = connectionURI + "?retryWrites=true"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		fmt.Println(err)
		c.Logger.Fatal("mongo connection error!")
	}

	// Check the connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		c.Logger.Fatal(err.Error())
	}

	database := client.Database(name)
	c.Logger.Info(name + " db is connected successfully!")

	return &mongodbAdapter{
		database:      database,
		client:        client,
		SharedService: c,
	}
}

func (d *mongodbAdapter) Disconnect() {
	d.Logger.Info("disconnection " + d.database.Name() + " db...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	d.database.Client().Disconnect(ctx)
	d.Logger.Info(d.database.Name() + " is disconnected  successfully")
}

func (d *mongodbAdapter) Collection(name string, opts ...*options.CollectionOptions) MongoCollection {
	return d.database.Collection(name, opts...)
}
