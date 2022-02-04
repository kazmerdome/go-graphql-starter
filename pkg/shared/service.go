package shared

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/observe/logger"
)

type SharedService struct {
	Logger logger.Logger
	Config config.Config
}

func NewSharedService(l logger.Logger, c config.Config) *SharedService {
	return &SharedService{l, c}
}
