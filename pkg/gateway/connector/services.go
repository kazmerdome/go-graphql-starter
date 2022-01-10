package connector

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/category"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/post"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
)

type GatewayServices struct {
	CategoryService category.CategoryService
	PostService     post.PostService
	// UserService     user.UserService
	LicenceService licence.LicenceService
}
