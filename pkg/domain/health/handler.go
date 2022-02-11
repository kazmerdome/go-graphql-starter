package health

import (
	"net/http"

	"github.com/kazmerdome/go-graphql-starter/pkg/module/provider/handler"

	"github.com/labstack/echo"
)

type HealthHandler interface {
	AddSubroute(e *echo.Echo)
}

type healthHandler struct {
	handler.HandlerConfig
	healthService HealthService
}

func newHealthHandler(c handler.HandlerConfig, healthService HealthService) HealthHandler {
	h := healthHandler{
		healthService: healthService,
		HandlerConfig: c,
	}
	return &h
}

func (r *healthHandler) AddSubroute(e *echo.Echo) {
	e.GET("/healthz", func(c echo.Context) error {
		data := r.healthService.GetHealthz()
		return c.String(http.StatusOK, data)
	})

	e.GET("/readyz", func(c echo.Context) error {
		data := r.healthService.GetReadyz()
		return c.String(http.StatusOK, data)
	})
}
