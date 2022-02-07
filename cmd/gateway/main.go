package main

import (
	"os"
	"os/signal"

	"github.com/kazmerdome/go-graphql-starter/pkg/adapter"
	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/health"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	observeLogger "github.com/kazmerdome/go-graphql-starter/pkg/observe/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/server"

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
	// Load Config
	c := config.NewConfig(config.MODE_GLOBALENV)

	// Load Observes
	logger := observeLogger.NewStandardLogger()

	// Load Adapters
	adapterConfig := adapter.NewAdapterConfig(logger, c)
	mongodbAdapter := mongodb.NewMongodbAdapter(
		adapterConfig,
		c.Get(config.ENV_MONGO_URI),
		c.Get(config.ENV_MONGO_DATABASE),
		true,
	)
	defer mongodbAdapter.Disconnect()

	// Load Modules
	moduleConfig := module.NewModuleConfig(logger, c)
	var (
		healthModule   = health.NewHealthModule(moduleConfig)
		licenceModule  = licence.NewLicenceModule(moduleConfig, mongodbAdapter)
		userModule     = user.NewUserModule(moduleConfig, mongodbAdapter)
		categoryModule = category.NewCategoryModule(moduleConfig, mongodbAdapter)
		postModule     = post.NewPostModule(moduleConfig, mongodbAdapter)
		gatewayModule  = gateway.NewGatewayModule(
			moduleConfig,
			c.Get(config.ENV_GRAPHQL_ENDPOINT),
			c.Get(config.ENV_GRAPHQL_PLAYGROUND_PASS),
			connector.GatewayModules{
				CategoryModule: categoryModule,
				UserModule:     userModule,
				LicenceModule:  licenceModule,
				PostModule:     postModule,
			},
		)
	)

	// Init server
	APP_PORT := c.Get(config.ENV_PORT)
	if APP_PORT == "" {
		APP_PORT = DEFAULT_PORT
	}

	server := server.NewServer(
		server.NewServerConfig(logger, c),
		APP_PORT,
		[]server.Handler{
			healthModule.GetHandler(),
			gatewayModule.GetHandler(),
		},
		[]echo.MiddlewareFunc{
			// server.ShowReqHeadersMiddleware,
		},
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
