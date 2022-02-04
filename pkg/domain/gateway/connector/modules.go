package connector

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/category"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/post"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/user"
)

type GatewayModules struct {
	CategoryModule category.CategoryModule
	PostModule     post.PostModule
	UserModule     user.UserModule
	LicenceModule  licence.LicenceModule
}
