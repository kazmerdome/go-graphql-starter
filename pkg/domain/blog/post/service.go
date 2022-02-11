package post

import (
	"context"

	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/service"
)

const (
	DB_COLLECTION_NAME = "Post"
)

type PostService interface {
	Post(ctx context.Context, where *PostWhereDTO) (*Post, error)
	Posts(ctx context.Context, where *PostWhereDTO, orderBy *PostOrderByENUM, skip *int, limit *int) ([]*Post, error)
	PostCount(ctx context.Context, where *PostWhereDTO) (*int, error)
	CreatePost(ctx context.Context, data Post) (*Post, error)
	UpdatePost(ctx context.Context, where PostWhereUniqueDTO, data Post) (*Post, error)
	DeletePost(ctx context.Context, where PostWhereUniqueDTO) (*Post, error)
}

type postService struct {
	service.ServiceConfig
	postRepository PostRepository
}

func newPostService(c service.ServiceConfig, r PostRepository) PostService {
	return &postService{ServiceConfig: c, postRepository: r}
}

// Post
func (r *postService) Post(ctx context.Context, where *PostWhereDTO) (*Post, error) {
	return r.postRepository.One(where)
}

// Posts
func (r *postService) Posts(ctx context.Context, where *PostWhereDTO, orderBy *PostOrderByENUM, skip *int, limit *int) ([]*Post, error) {
	return r.postRepository.List(where, orderBy, skip, limit, nil)
}

// PostCount
func (r *postService) PostCount(ctx context.Context, where *PostWhereDTO) (*int, error) {
	return r.postRepository.Count(where)
}

// CreatePost
func (r *postService) CreatePost(ctx context.Context, data Post) (*Post, error) {
	return r.postRepository.Create(&data)
}

// UpdatePost
func (r *postService) UpdatePost(ctx context.Context, where PostWhereUniqueDTO, data Post) (*Post, error) {
	return r.postRepository.Update(where.ID, &data)
}

// DeletePost
func (r *postService) DeletePost(ctx context.Context, where PostWhereUniqueDTO) (*Post, error) {
	return r.postRepository.Delete(where.ID)
}
