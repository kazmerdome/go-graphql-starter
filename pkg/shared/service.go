package shared

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/config"
	"github.com/kazmerdome/go-graphql-starter/pkg/logger"
)

type SharedService struct {
	Logger logger.Logger
	Config config.ConfigService
}

func NewSharedService(l logger.Logger, c config.ConfigService) *SharedService {
	return &SharedService{l, c}
}
