package post

import (
	"context"

	"github.com/kazmerdome/go-graphql-starter/pkg/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"
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
	shared.SharedService
	r PostRepository
}

func NewPostService(s shared.SharedService, db repository.MongoDatabase) PostService {
	r := NewPostRepository(s, db.Collection(DB_COLLECTION_NAME))
	return &postService{SharedService: s, r: r}
}

// Post
func (r *postService) Post(ctx context.Context, where *PostWhereDTO) (*Post, error) {
	return r.r.One(where)
}

// Posts
func (r *postService) Posts(ctx context.Context, where *PostWhereDTO, orderBy *PostOrderByENUM, skip *int, limit *int) ([]*Post, error) {
	return r.r.List(where, orderBy, skip, limit, nil)
}

// PostCount
func (r *postService) PostCount(ctx context.Context, where *PostWhereDTO) (*int, error) {
	return r.r.Count(where)
}

// CreatePost
func (r *postService) CreatePost(ctx context.Context, data Post) (*Post, error) {
	return r.r.Create(&data)
}

// UpdatePost
func (r *postService) UpdatePost(ctx context.Context, where PostWhereUniqueDTO, data Post) (*Post, error) {
	return r.r.Update(where.ID, &data)
}

// DeletePost
func (r *postService) DeletePost(ctx context.Context, where PostWhereUniqueDTO) (*Post, error) {
	return r.r.Delete(where.ID)
}
