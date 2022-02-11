package gateway

import (
	"context"
	"errors"

	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/connector"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/dataloader"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/graph/generated"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/resolver"

	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"

	echoHandler "github.com/kazmerdome/go-graphql-starter/pkg/module/provider/handler/echo"

	httpUtil "github.com/kazmerdome/go-graphql-starter/pkg/util/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
)

const (
	AUTHORIZATION_HEADER_KEY  = "Authorization"
	PLAYGROUNDPASS_HEADER_KEY = "Playground-Password"
	PLAYGROUND_TITLE          = "GraphQL Playground"
)

type gatewayHandler struct {
	echoHandler.EchoHandlerConfig
	modules                  connector.GatewayModules
	graphqlEndpoint          string
	authToken                string
	playgroundPassword       string
	playgroundPasswordHeader string
}

func newGatewayHandler(
	c echoHandler.EchoHandlerConfig,
	graphqlEndpoint string,
	playgroundPassword string,
	modules connector.GatewayModules,
) echoHandler.EchoHandler {
	var authToken string
	var playgroundPasswordHeader string
	h := gatewayHandler{
		EchoHandlerConfig:        c,
		graphqlEndpoint:          graphqlEndpoint,
		authToken:                authToken,
		playgroundPassword:       playgroundPassword,
		playgroundPasswordHeader: playgroundPasswordHeader,
		modules:                  modules,
	}
	return &h
}

func (r *gatewayHandler) AddSubroute(e *echo.Echo) {
	/*
	 * Add Request.Header Reader Middleware
	 */
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			headers := ctx.Request().Header
			r.authToken = httpUtil.GetHeaderString(AUTHORIZATION_HEADER_KEY, headers)
			r.playgroundPasswordHeader = httpUtil.GetHeaderString(PLAYGROUNDPASS_HEADER_KEY, headers)
			return next(ctx)
		}
	})

	/*
	 * Init GQL: add Resolvers and Directives
	 */
	resolver := resolver.NewResolver(&r.authToken, r.modules)
	config := generated.Config{Resolvers: resolver}
	config.Directives.Auth = func(
		ctx context.Context,
		obj interface{},
		next graphql.Resolver,
		feature licence.Feature,
		permissions []licence.Permission,
	) (interface{}, error) {
		if err := r.modules.LicenceModule.GetGuard().AuthGuard(feature, permissions, *resolver.AuthToken); err != nil {
			return nil, errors.New(licence.ERR_UNAUTHORIZED)
		}
		return next(ctx)
	}

	// new custom handler based on gqlgen version <0.11.3
	queryHandler := handler.New(generated.NewExecutableSchema(config))
	queryHandler.AddTransport(transport.POST{})
	queryHandler.AddTransport(transport.MultipartForm{})
	queryHandler.SetQueryCache(lru.New(1000))
	queryHandler.Use(extension.AutomaticPersistedQuery{Cache: lru.New(100)})
	queryHandler.Use(apollotracing.Tracer{})
	queryHandler.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		rc := graphql.GetOperationContext(ctx)
		if r.playgroundPassword == r.playgroundPasswordHeader {
			rc.DisableIntrospection = false
		} else {
			rc.DisableIntrospection = true
		}
		return next(ctx)
	})

	e.GET("/", echo.WrapHandler(playground.Handler(PLAYGROUND_TITLE, r.graphqlEndpoint)))
	e.POST("/query", echo.WrapHandler(dataloader.DataLoaderMiddleware(r.modules, queryHandler)))
}
