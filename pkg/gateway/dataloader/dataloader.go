package dataloader

import (
	"context"
	"net/http"

	"github.com/kazmerdome/go-graphql-starter/pkg/gateway/connector"
)

// ContextKey ...
type loaderCTX string

const ContextKey loaderCTX = "DATALOADER"

// Loaders ...
type Loaders struct {
	CategoryLoader *CategoryLoader
}

// DataLoaderMiddleware ...
// services: load all of the services for db operations
func DataLoaderMiddleware(services connector.GatewayServices, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loaders := Loaders{
			CategoryLoader: getCategoryLoader(services.CategoryService),
		}
		ctx := context.WithValue(r.Context(), ContextKey, loaders)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetContextLoaders ...
func GetContextLoaders(ctx context.Context) Loaders {
	return ctx.Value(ContextKey).(Loaders)
}
