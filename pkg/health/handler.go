package health

import (
	"net/http"

	"github.com/kazmerdome/go-graphql-starter/pkg/server"
	"github.com/kazmerdome/go-graphql-starter/pkg/shared"

	"github.com/labstack/echo"
)

type healthHandler struct {
	shared.SharedService
	s HealthService
}

func NewHealthHandler(ss shared.SharedService, hs HealthService) server.Handler {
	return &healthHandler{
		s:             hs,
		SharedService: ss,
	}
}

func (r *healthHandler) GetRoutes(e *echo.Echo) {
	e.GET("/healthz", func(c echo.Context) error {
		data := r.s.GetHealthz()
		return c.String(http.StatusOK, data)
	})

	e.GET("/readyz", func(c echo.Context) error {
		data := r.s.GetReadyz()
		return c.String(http.StatusOK, data)
	})
}
