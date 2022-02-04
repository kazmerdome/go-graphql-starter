package health

import (
	"net/http"

	"github.com/kazmerdome/go-graphql-starter/pkg/provider/handler"
	"github.com/kazmerdome/go-graphql-starter/pkg/server"

	"github.com/labstack/echo"
)

type healthHandler struct {
	*handler.HandlerConfig
	healthService HealthService
}

func newHealthHandler(c *handler.HandlerConfig, healthService HealthService) server.Handler {
	return &healthHandler{
		healthService: healthService,
		HandlerConfig: c,
	}
}

func (r *healthHandler) GetRoutes(e *echo.Echo) {
	e.GET("/healthz", func(c echo.Context) error {
		data := r.healthService.GetHealthz()
		return c.String(http.StatusOK, data)
	})

	e.GET("/readyz", func(c echo.Context) error {
		data := r.healthService.GetReadyz()
		return c.String(http.StatusOK, data)
	})
}
