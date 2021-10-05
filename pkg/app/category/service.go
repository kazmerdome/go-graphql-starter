package category

import (
	"context"

	"github.com/kazmerdome/go-graphql-starter/pkg/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
)

const (
	DB_COLLECTION_NAME = "Category"
)

type CategoryService interface {
	Category(ctx context.Context, where *CategoryWhereDTO) (*Category, error)
	Categories(ctx context.Context, where *CategoryWhereDTO, orderBy *CategoryOrderByENUM, skip *int, limit *int) ([]*Category, error)
	CategoryCount(ctx context.Context, where *CategoryWhereDTO) (*int, error)
	CreateCategory(ctx context.Context, data Category) (*Category, error)
	UpdateCategory(ctx context.Context, where CategoryWhereUniqueDTO, data Category) (*Category, error)
	DeleteCategory(ctx context.Context, where CategoryWhereUniqueDTO) (*Category, error)
	GetRepository() CategoryRepository
}

type categoryService struct {
	shared.SharedService
	cr CategoryRepository
}

func NewCategoryService(s shared.SharedService, db repository.MongoDatabase) CategoryService {
	cr := NewCategoryRepository(s, db.Collection(DB_COLLECTION_NAME))
	return &categoryService{SharedService: s, cr: cr}
}

// Dataloader helper - repository export
func (r *categoryService) GetRepository() CategoryRepository {
	return r.cr
}

// Category
func (r *categoryService) Category(ctx context.Context, where *CategoryWhereDTO) (*Category, error) {
	return r.cr.One(where)
}

// Categories
func (r *categoryService) Categories(ctx context.Context, where *CategoryWhereDTO, orderBy *CategoryOrderByENUM, skip *int, limit *int) ([]*Category, error) {
	return r.cr.List(where, orderBy, skip, limit, nil)
}

// CategoryCount
func (r *categoryService) CategoryCount(ctx context.Context, where *CategoryWhereDTO) (*int, error) {
	return r.cr.Count(where)
}

// CreateCategory
func (r *categoryService) CreateCategory(ctx context.Context, data Category) (*Category, error) {
	return r.cr.Create(&data)
}

// UpdateCategory
func (r *categoryService) UpdateCategory(ctx context.Context, where CategoryWhereUniqueDTO, data Category) (*Category, error) {
	return r.cr.Update(where.ID, &data)
}

// DeleteCategory
func (r *categoryService) DeleteCategory(ctx context.Context, where CategoryWhereUniqueDTO) (*Category, error) {
	return r.cr.Delete(where.ID)
}
