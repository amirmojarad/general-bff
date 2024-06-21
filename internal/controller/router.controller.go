package controller

import "github.com/labstack/echo/v4"

func SetReverseProxyRoutes(router *echo.Group, reverseProxy []ServiceProxy) {
	for _, rp := range reverseProxy {
		rp.toRequests(router)
	}
}
