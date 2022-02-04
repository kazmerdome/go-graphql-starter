package main

import (
	"os"
	"os/signal"

	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/health"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	"github.com/kazmerdome/go-graphql-starter/pkg/observe/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/server"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"

	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/category"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/post"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/connector"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/user"

	"syscall"

	"github.com/labstack/echo"
)

const (
	DEFAULT_PORT = "9090"
)

func main() {
	/*
	 * Load Config
	 */
	c := config.NewConfig(config.MODE_GLOBALENV)

	/*
	 * Load Observes
	 */
	l := logger.NewStandardLogger()

	/*
	 * Init providerConfig for logging and config as base dependency
	 */
	s := *shared.NewSharedService(l, c)

	/*
	 * Adapters init
	 */
	mongodbAdapter := mongodb.NewMongodbAdapter(
		s,
		c.Get(config.ENV_MONGO_URI),
		c.Get(config.ENV_MONGO_DATABASE),
		true,
	)
	defer mongodbAdapter.Disconnect()

	/*
	 * Modules Init
	 */
	var (
		moduleConfig   = module.NewModuleConfig(l, c)
		healthModule   = health.NewHealthModule(moduleConfig)
		licenceModule  = licence.NewLicenceModule(moduleConfig, mongodbAdapter)
		userModule     = user.NewUserModule(moduleConfig, mongodbAdapter)
		categoryModule = category.NewCategoryModule(moduleConfig, mongodbAdapter)
		postModule     = post.NewPostModule(moduleConfig, mongodbAdapter)
		gatewayModule  = gateway.NewGatewayModule(
			moduleConfig,
			s.Config.Get(config.ENV_GRAPHQL_ENDPOINT),
			s.Config.Get(config.ENV_GRAPHQL_PLAYGROUND_PASS),
			connector.GatewayModules{
				CategoryModule: categoryModule,
				UserModule:     userModule,
				LicenceModule:  licenceModule,
				PostModule:     postModule,
			},
		)
	)

	/*
	 * Server init
	 */
	APP_PORT := s.Config.Get(config.ENV_PORT)
	if APP_PORT == "" {
		APP_PORT = DEFAULT_PORT
	}

	middlewares := []echo.MiddlewareFunc{
		// server.ShowReqHeadersMiddleware,
	}
	handlers := []func(e *echo.Echo){
		healthModule.GetHandler().GetRoutes,
		gatewayModule.GetHandler().GetRoutes,
	}

	server := server.NewServer(
		s,
		APP_PORT,
		&handlers,
		&middlewares,
		true,
	)

	go server.Start()
	defer server.Stop()

	// Stop server gracefully
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
