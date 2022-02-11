package exposure

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/observer/logger"
)

type ExposureType string

type ExposureConfig interface {
	GetConfig() config.Config
	GetLogger() logger.Logger
}

type exposureConfig struct {
	logger logger.Logger
	config config.Config
}

func NewExposureConfig(l logger.Logger, c config.Config) ExposureConfig {
	return &exposureConfig{l, c}
}

func (r *exposureConfig) GetConfig() config.Config {
	return r.config
}

func (r *exposureConfig) GetLogger() logger.Logger {
	return r.logger
}
