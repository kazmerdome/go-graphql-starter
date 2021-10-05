package health

import "github.com/kazmerdome/go-graphql-starter/pkg/shared"

type HealthService interface {
	GetHealthz() string
	GetReadyz() string
}

type healthService struct {
	shared.SharedService
}

func NewHealthService(s shared.SharedService) HealthService {
	return &healthService{
		SharedService: s,
	}
}

func (r *healthService) GetHealthz() string {
	return "healthy"
}

func (r *healthService) GetReadyz() string {
	return "ready"
}
