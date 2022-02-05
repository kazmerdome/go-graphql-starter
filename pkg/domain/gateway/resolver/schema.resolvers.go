package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/category"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/post"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/dataloader"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/graph/generated"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, data category.Category) (*category.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, where category.CategoryWhereUniqueDTO, data category.Category) (*category.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, where category.CategoryWhereUniqueDTO) (*category.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePost(ctx context.Context, data post.Post) (*post.Post, error) {
	return r.modules.PostModule.GetService().CreatePost(ctx, data)
}

func (r *mutationResolver) UpdatePost(ctx context.Context, where post.PostWhereUniqueDTO, data post.Post) (*post.Post, error) {
	return r.modules.PostModule.GetService().UpdatePost(ctx, where, data)
}

func (r *mutationResolver) DeletePost(ctx context.Context, where post.PostWhereUniqueDTO) (*post.Post, error) {
	return r.modules.PostModule.GetService().DeletePost(ctx, where)
}

func (r *postResolver) Category(ctx context.Context, obj *post.Post) (*category.Category, error) {
	return dataloader.GetContextLoaders(ctx).CategoryLoader.Load(obj.Category.Hex())
}

func (r *queryResolver) Category(ctx context.Context, where *category.CategoryWhereDTO) (*category.Category, error) {
	return r.modules.CategoryModule.GetService().Category(ctx, where)
}

func (r *queryResolver) Categories(ctx context.Context, where *category.CategoryWhereDTO, orderBy *category.CategoryOrderByENUM, skip *int, limit *int) ([]*category.Category, error) {
	return r.modules.CategoryModule.GetService().Categories(ctx, where, orderBy, skip, limit)
}

func (r *queryResolver) CategoryCount(ctx context.Context, where *category.CategoryWhereDTO) (*int, error) {
	return r.modules.CategoryModule.GetService().CategoryCount(ctx, where)
}

func (r *queryResolver) Post(ctx context.Context, where *post.PostWhereDTO) (*post.Post, error) {
	return r.modules.PostModule.GetService().Post(ctx, where)
}

func (r *queryResolver) Posts(ctx context.Context, where *post.PostWhereDTO, orderBy *post.PostOrderByENUM, skip *int, limit *int) ([]*post.Post, error) {
	return r.modules.PostModule.GetService().Posts(ctx, where, orderBy, skip, limit)
}

func (r *queryResolver) PostCount(ctx context.Context, where *post.PostWhereDTO) (*int, error) {
	return r.modules.PostModule.GetService().PostCount(ctx, where)
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
