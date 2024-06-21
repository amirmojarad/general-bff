package server

import (
	"fmt"
	"general-bff/config/env"
	"general-bff/config/yml"
	"general-bff/internal/controller"
	"general-bff/internal/controller/middlewares"
	"github.com/labstack/echo/v4"
)

func Run() error {
	cfg, err := env.NewConfig()
	if err != nil {
		return err
	}

	servicesCfg := yml.NewServicesConfig(cfg)

	router, err := setupRouter(servicesCfg)
	if err != nil {
		return err
	}

	return router.Start(fmt.Sprintf(":%d", cfg.App.Port))
}

func setupRouter(servicesCfg *yml.ServicesConfig) (*echo.Echo, error) {
	reverseProxyConfig, err := controller.NewReverseProxyConfig(servicesCfg)
	if err != nil {
		return nil, err
	}

	middleware := middlewares.NewMiddleware(servicesCfg)

	router := echo.New()
	controller.SetHealthCheck(router, controller.NewHealthCheck())

	group := router.Group("/api", middleware.RemovePrefix)
	controller.SetReverseProxyRoutes(group, reverseProxyConfig.ServiceProxies)

	return router, nil
}
