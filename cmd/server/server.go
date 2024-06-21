package server

import (
	"fmt"
	"general-bff/config/env"
	"general-bff/config/yml"
	"general-bff/internal/controller"
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
	reverseProxies, err := controller.NewReverseProxyConfig(servicesCfg)
	if err != nil {
		return nil, err
	}

	router := echo.New()

	group := router.Group("/api")
	controller.SetReverseProxyRoutes(group, reverseProxies.ReverseProxies)

	return router, nil
}
