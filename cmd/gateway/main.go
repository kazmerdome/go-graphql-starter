package main

import (
	"os"
	"os/signal"

	"github.com/kazmerdome/go-graphql-starter/pkg/adapter"
	"github.com/kazmerdome/go-graphql-starter/pkg/adapter/repository/mongodb"
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/exposure"
	echoExposure "github.com/kazmerdome/go-graphql-starter/pkg/exposure/http/echo"
	"github.com/kazmerdome/go-graphql-starter/pkg/module"
	observerLogger "github.com/kazmerdome/go-graphql-starter/pkg/observer/logger"
	"github.com/labstack/echo"

	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/category"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/post"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/gateway/connector"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/health"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/user"

	"syscall"
)

const (
	DEFAULT_PORT = "9090"
)

func main() {
	// Load Config
	c := config.NewConfig(config.MODE_GLOBALENV)
	APP_PORT := c.Get(config.ENV_PORT)
	if APP_PORT == "" {
		APP_PORT = DEFAULT_PORT
	}

	// Load Observers
	logger := observerLogger.NewStandardLogger()

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
	adapters := module.NewAdapters(mongodbAdapter)

	var (
		healthModule  = health.NewHealthModule(moduleConfig)
		gatewayModule = gateway.NewGatewayModule(
			moduleConfig,
			c.Get(config.ENV_GRAPHQL_ENDPOINT),
			c.Get(config.ENV_GRAPHQL_PLAYGROUND_PASS),
			connector.GatewayModules{
				CategoryModule: category.NewCategoryModule(moduleConfig, adapters),
				UserModule:     user.NewUserModule(moduleConfig, adapters),
				LicenceModule:  licence.NewLicenceModule(moduleConfig, adapters),
				PostModule:     post.NewPostModule(moduleConfig, adapters),
			},
		)
	)

	// Load Exposers
	echoExposure := echoExposure.NewEchoExposure(
		exposure.NewExposureConfig(logger, c),
		[]echo.MiddlewareFunc{},
		[]echoExposure.Handler{
			healthModule.GetEchoHttpHandler(),
			gatewayModule.GetEchoHttpHandler(),
		},
		APP_PORT,
		true,
	)

	echoExposure.Start()
	defer echoExposure.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
