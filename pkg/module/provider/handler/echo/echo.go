package echoHandler

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider"
	"github.com/labstack/echo"
)

type EchoHandler interface {
	AddSubroute(e *echo.Echo)
}

type EchoHandlerConfig interface {
	provider.ProviderConfig
}

type echoHandlerConfig struct {
	provider.ProviderConfig
}

func NewHandlerConfig(config provider.ProviderConfig) EchoHandlerConfig {
	return &echoHandlerConfig{ProviderConfig: config}
}
