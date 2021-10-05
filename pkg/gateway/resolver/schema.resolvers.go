package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kazmerdome/go-graphql-starter/pkg/app/category"
	"github.com/kazmerdome/go-graphql-starter/pkg/app/post"
	"github.com/kazmerdome/go-graphql-starter/pkg/gateway/dataloader"
	"github.com/kazmerdome/go-graphql-starter/pkg/gateway/graph/generated"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, data category.Category) (*category.Category, error) {
	return r.services.CategoryService.CreateCategory(ctx, data)
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, where category.CategoryWhereUniqueDTO, data category.Category) (*category.Category, error) {
	return r.services.CategoryService.UpdateCategory(ctx, where, data)
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, where category.CategoryWhereUniqueDTO) (*category.Category, error) {
	return r.services.CategoryService.DeleteCategory(ctx, where)
}

func (r *mutationResolver) CreatePost(ctx context.Context, data post.Post) (*post.Post, error) {
	return r.services.PostService.CreatePost(ctx, data)
}

func (r *mutationResolver) UpdatePost(ctx context.Context, where post.PostWhereUniqueDTO, data post.Post) (*post.Post, error) {
	return r.services.PostService.UpdatePost(ctx, where, data)
}

func (r *mutationResolver) DeletePost(ctx context.Context, where post.PostWhereUniqueDTO) (*post.Post, error) {
	return r.services.PostService.DeletePost(ctx, where)
}

func (r *postResolver) Category(ctx context.Context, obj *post.Post) (*category.Category, error) {
	return dataloader.GetContextLoaders(ctx).CategoryLoader.Load(obj.Category.Hex())
}

func (r *queryResolver) Category(ctx context.Context, where *category.CategoryWhereDTO) (*category.Category, error) {
	return r.services.CategoryService.Category(ctx, where)
}

func (r *queryResolver) Categories(ctx context.Context, where *category.CategoryWhereDTO, orderBy *category.CategoryOrderByENUM, skip *int, limit *int) ([]*category.Category, error) {
	return r.services.CategoryService.Categories(ctx, where, orderBy, skip, limit)
}

func (r *queryResolver) CategoryCount(ctx context.Context, where *category.CategoryWhereDTO) (*int, error) {
	return r.services.CategoryService.CategoryCount(ctx, where)
}

func (r *queryResolver) Post(ctx context.Context, where *post.PostWhereDTO) (*post.Post, error) {
	return r.services.PostService.Post(ctx, where)
}

func (r *queryResolver) Posts(ctx context.Context, where *post.PostWhereDTO, orderBy *post.PostOrderByENUM, skip *int, limit *int) ([]*post.Post, error) {
	return r.services.PostService.Posts(ctx, where, orderBy, skip, limit)
}

func (r *queryResolver) PostCount(ctx context.Context, where *post.PostWhereDTO) (*int, error) {
	return r.services.PostService.PostCount(ctx, where)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
