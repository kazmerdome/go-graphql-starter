package services

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/app/category"
	"github.com/kazmerdome/go-graphql-starter/pkg/app/post"
)

type GatewayServices struct {
	CategoryService category.CategoryService
	PostService     post.PostService
}
