package category

import (
	"context"

	"github.com/kazmerdome/go-graphql-starter/pkg/provider/service"
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
	*service.ServiceConfig
	categoryRepository CategoryRepository
}

func newCategoryService(c *service.ServiceConfig, r CategoryRepository) CategoryService {
	return &categoryService{ServiceConfig: c, categoryRepository: r}
}

// Dataloader helper - repository export
func (r *categoryService) GetRepository() CategoryRepository {
	return r.categoryRepository
}

// Category
func (r *categoryService) Category(ctx context.Context, where *CategoryWhereDTO) (*Category, error) {
	return r.categoryRepository.One(where)
}

// Categories
func (r *categoryService) Categories(ctx context.Context, where *CategoryWhereDTO, orderBy *CategoryOrderByENUM, skip *int, limit *int) ([]*Category, error) {
	return r.categoryRepository.List(where, orderBy, skip, limit, nil)
}

// CategoryCount
func (r *categoryService) CategoryCount(ctx context.Context, where *CategoryWhereDTO) (*int, error) {
	return r.categoryRepository.Count(where)
}

// CreateCategory
func (r *categoryService) CreateCategory(ctx context.Context, data Category) (*Category, error) {
	return r.categoryRepository.Create(&data)
}

// UpdateCategory
func (r *categoryService) UpdateCategory(ctx context.Context, where CategoryWhereUniqueDTO, data Category) (*Category, error) {
	return r.categoryRepository.Update(where.ID, &data)
}

// DeleteCategory
func (r *categoryService) DeleteCategory(ctx context.Context, where CategoryWhereUniqueDTO) (*Category, error) {
	return r.categoryRepository.Delete(where.ID)
}
