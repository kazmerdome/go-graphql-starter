package echoExposure

import (
	"context"
	"time"

	"github.com/kazmerdome/go-graphql-starter/pkg/exposure"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	LOG_TYPE_FANCY  = "fancy"
	LOG_TYPE_LOGGER = "logger"
)

type EchoExposure interface {
	exposure.ExposureConfig
	Start()
	Stop()
}

type Handler interface {
	AddSubroute(*echo.Echo)
}

type echoExposure struct {
	exposure.ExposureConfig
	port        string
	subroutes   []Handler
	middlewares []echo.MiddlewareFunc
	e           *echo.Echo
	logType     string
}

func NewEchoExposure(
	c exposure.ExposureConfig,
	middlewares []echo.MiddlewareFunc,
	subroutes []Handler,
	port string,
	withFancyLog bool,
) EchoExposure {
	logType := LOG_TYPE_LOGGER
	if withFancyLog {
		logType = LOG_TYPE_FANCY
	}

	return &echoExposure{
		ExposureConfig: c,
		subroutes:      subroutes,
		middlewares:    middlewares,
		port:           port,
		e:              echo.New(),
		logType:        logType,
	}
}

func (r *echoExposure) Start() {
	// Setup & configure server
	// more info -> https://echo.labstack.com/
	r.e.HideBanner = true

	// Handle Middlewares
	r.e.Use(middleware.CORS())
	r.e.Use(middleware.Recover())

	if r.logType == LOG_TYPE_FANCY {
		r.e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "\033[0;34mlatency: \033[0;97m${latency_human} " +
				"\033[0;34mremote_ip: \033[0;97m${remote_ip} " +
				"\033[0;34mmethod: \033[0;97m${method} " +
				"\033[0;34mstatus: \033[0;97m${status} " +
				"\033[0;34mdate: \033[0;97m${time_rfc3339} " +
				"\033[0;34mpath: \033[0;97m${path} " +
				"\033[0;34mquery: \033[0;97m${query} " +
				"\033[0;34muri: \033[0;97m${uri} " +
				"\n"}))
	} else {
		r.e.Use(middleware.Logger())
	}

	// Init Additional Middlewares
	if r.middlewares != nil && len(r.middlewares) > 0 {
		for _, mw := range r.middlewares {
			r.e.Use(mw)
		}
	}

	// Init Handlers (Subrouters)
	if r.subroutes != nil && len(r.subroutes) > 0 {
		for _, subroute := range r.subroutes {
			subroute.AddSubroute(r.e)
		}
	}

	// Start server routes
	go func() {
		if err := r.e.Start(":" + r.port); err != nil {
			r.GetLogger().Warn("shutting down the server")
		}
	}()
}

func (r *echoExposure) Stop() {
	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.e.Shutdown(ctx); err != nil {
		r.GetLogger().Fatal(err.Error())
	}
}
