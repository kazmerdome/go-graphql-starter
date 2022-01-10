package dataloader

import (
	"time"

	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/category"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate go run github.com/vektah/dataloaden CategoryLoader string *github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/category.Category

func getCategoryLoader(s category.CategoryService) *CategoryLoader {
	maxLimit := 150

	return NewCategoryLoader(
		CategoryLoaderConfig{
			MaxBatch: maxLimit,
			Wait:     1 * time.Millisecond,
			Fetch: func(keys []string) ([]*category.Category, []error) {
				var filter category.CategoryWhereDTO

				objectIds := make([]primitive.ObjectID, len(keys))
				for i, k := range keys {
					oid, _ := primitive.ObjectIDFromHex(k)
					objectIds[i] = oid
				}
				customQuery := bson.M{"_id": bson.M{"$in": objectIds}}

				r := s.GetRepository()
				items, err := r.List(&filter, nil, nil, &maxLimit, &customQuery)

				if err != nil {
					return nil, []error{err}
				}

				w := make(map[string]*category.Category, len(items))
				if len(items) > 0 {
					for _, item := range items {
						if item != nil {
							w[item.ID.Hex()] = item
						}
					}
				}

				result := make([]*category.Category, len(keys))
				for i, key := range keys {
					result[i] = w[key]
				}

				return result, nil
			},
		},
	)
}
