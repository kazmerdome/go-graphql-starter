package health

import (
	"net/http"

	echoHandler "github.com/kazmerdome/go-graphql-starter/pkg/module/provider/handler/echo"

	"github.com/labstack/echo"
)

type healthHandler struct {
	echoHandler.EchoHandlerConfig
	healthService HealthService
}

func newHealthHandler(c echoHandler.EchoHandlerConfig, healthService HealthService) echoHandler.EchoHandler {
	h := healthHandler{
		healthService:     healthService,
		EchoHandlerConfig: c,
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
