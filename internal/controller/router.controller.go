package controller

import "github.com/labstack/echo/v4"

func SetReverseProxyRoutes(router *echo.Group, reverseProxy []ServiceProxy) {
	for _, rp := range reverseProxy {
		rp.toRequests(router)
	}
}

func SetHealthCheck(router *echo.Echo, healthCheck *HealthCheck) {
	router.GET("/health", healthCheck.Health)
}
