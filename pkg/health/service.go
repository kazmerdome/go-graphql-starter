package health

import (
	"github.com/kazmerdome/go-graphql-starter/pkg/provider/service"
)

type HealthService interface {
	GetHealthz() string
	GetReadyz() string
}

type healthService struct {
	*service.ServiceConfig
}

func newHealthService(c *service.ServiceConfig) HealthService {
	return &healthService{
		ServiceConfig: c,
	}
}

func (r *healthService) GetHealthz() string {
	return "healthy"
}

func (r *healthService) GetReadyz() string {
	return "ready"
}
