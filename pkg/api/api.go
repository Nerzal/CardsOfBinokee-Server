package api

import (
	"net/http"
	"time"

	"github.com/heptiolabs/healthcheck"
	"github.com/labstack/echo"
)

type API interface {
	AddRoutes(router *echo.Echo)
}

type api struct {
	cardAPI CardAPI
}

func NewAPI(cardAPI CardAPI) API {
	return &api{
		cardAPI: cardAPI,
	}
}

func (api *api) AddRoutes(router *echo.Echo) {
	// health checks
	health := healthcheck.NewHandler()
	go http.ListenAndServe("0.0.0.0:8086", health)

	router.GET("/ping", func(c echo.Context) error {
		return c.JSON(200, "pong")
	})

	apiGroup := router.Group("/api/v1")
	cardGroup := apiGroup.Group("/cards")
	cardGroup.GET("", api.cardAPI.GetCards)
	cardGroup.POST("", api.cardAPI.PostCards)
}

func initHealthCheck() {
	health := healthcheck.NewHandler()
	// go routine check
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100000))
	health.AddReadinessCheck("readiness", healthcheck.Timeout(func() error {
		return nil
	}, 1*time.Second))

	go http.ListenAndServe("0.0.0.0:8086", health)
}
