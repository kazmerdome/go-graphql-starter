package main

import (
	"os"
	"os/signal"

	"github.com/kazmerdome/go-graphql-starter/pkg/auth/authorization/guards"
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/category"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/blog/post"
	"github.com/kazmerdome/go-graphql-starter/pkg/domain/licence"
	"github.com/kazmerdome/go-graphql-starter/pkg/gateway"
	"github.com/kazmerdome/go-graphql-starter/pkg/gateway/connector"
	"github.com/kazmerdome/go-graphql-starter/pkg/health"
	"github.com/kazmerdome/go-graphql-starter/pkg/logger"
	"github.com/kazmerdome/go-graphql-starter/pkg/repository"
	"github.com/kazmerdome/go-graphql-starter/pkg/server"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"

	"syscall"

	"github.com/labstack/echo"
)

const (
	DEFAULT_PORT = "9090"
)

func main() {
	/*
	 * Load config
	 */
	c := config.NewConfigService(config.MODE_GLOBALENV)

	/*
	 * Logger
	 */
	l := logger.NewStandardLogger()

	/*
	 * Init Shared Service for logging and config as base dependency
	 */
	s := *shared.NewSharedService(l, c)

	/*
	 * Database init
	 */
	db := repository.NewMongoDB(
		s, s.Config.Get(config.ENV_MONGO_URI),
		s.Config.Get(config.ENV_MONGO_DATABASE),
		true,
	)
	defer db.Disconnect()

	/*
	 * Init Modules
	 */
	// Health
	healthService := health.NewHealthService(s)
	healthHandler := health.NewHealthHandler(s, healthService)

	// Licence
	licenceRepository := licence.NewLicenceRepository(s, db)
	licenceService := licence.NewLicenceService(s, licenceRepository)
	licenceGuard := guards.NewLicenceGuard(s, licenceRepository)

	// Post
	postService := post.NewPostService(s, db)

	// Category
	categoryService := category.NewCategoryService(s, db)

	/*
	 * Load services and guards for gatewayService
	 */
	services := connector.GatewayServices{
		CategoryService: categoryService,
		PostService:     postService,
		LicenceService:  licenceService,
	}

	guards := connector.GatewayGuards{
		LicenceGuard: licenceGuard,
	}

	/*
	 * Gateway Module Init
	 */
	gatewayHandler := gateway.NewGatewayHandler(s, s.Config.Get(config.ENV_GRAPHQL_ENDPOINT),
		s.Config.Get(config.ENV_GRAPHQL_PLAYGROUND_PASS), services, guards)

	/*
	 * Collect Handlers
	 */
	handlers := []func(e *echo.Echo){
		healthHandler.GetRoutes,
		gatewayHandler.GetRoutes,
	}

	/*
	 * Server init
	 */
	middlewares := []echo.MiddlewareFunc{
		// server.ShowReqHeadersMiddleware,
	}

	APP_PORT := s.Config.Get(config.ENV_PORT)
	if APP_PORT == "" {
		APP_PORT = DEFAULT_PORT
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
