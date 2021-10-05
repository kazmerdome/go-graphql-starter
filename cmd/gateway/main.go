package main

import (
	"os"
	"os/signal"

	"github.com/kazmerdome/go-graphql-starter/pkg/app/category"
	"github.com/kazmerdome/go-graphql-starter/pkg/app/post"
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/gateway"
	"github.com/kazmerdome/go-graphql-starter/pkg/gateway/services"
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
	db := repository.NewMongoDB(s, s.Config.Get("MONGO_URI"), s.Config.Get("MONGO_DATABASE"), true)
	defer db.Disconnect()

	/*
	 * Health Module Init
	 */
	healthService := health.NewHealthService(s)
	healthHandler := health.NewHealthHandler(s, healthService)

	/*
	 * Load services for gatewayService
	 */
	services := services.GatewayServices{
		CategoryService: category.NewCategoryService(s, db),
		PostService:     post.NewPostService(s, db),
	}

	/*
	 * Gateway Module Init
	 */
	gatewayHandler := gateway.NewGatewayHandler(s, s.Config.Get("GRAPHQL_ENDPOINT"), s.Config.Get("GRAPHQL_PLAYGROUND_PASS"), services)

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

	APP_PORT := s.Config.Get("PORT")
	if APP_PORT == "" {
		APP_PORT = DEFAULT_PORT
	}

	identityServer := server.NewServer(
		s,
		APP_PORT,
		&handlers,
		&middlewares,
		true,
	)

	go identityServer.Start()
	defer identityServer.Stop()

	// Stop server gracefully
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
