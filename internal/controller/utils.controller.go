package controller

import (
	"general-bff/config/yml"
	"github.com/labstack/echo/v4"
)

type HttpMethod string

const (
	HttpGetMethod    HttpMethod = "GET"
	HttpPostMethod   HttpMethod = "POST"
	HttpPutMethod    HttpMethod = "PUT"
	HttpDeleteMethod HttpMethod = "DELETE"
	HttpPatchMethod  HttpMethod = "PATCH"
	UnknownMethod    HttpMethod = "UNKNOWN"
)

func (h HttpMethod) String() string {
	return string(h)
}

func detectHttpMethod(path yml.Path) HttpMethod {
	if path.GET != nil {
		return HttpGetMethod
	}

	if path.POST != nil {
		return HttpPostMethod
	}

	if path.PUT != nil {
		return HttpPutMethod
	}

	if path.PATCH != nil {
		return HttpPatchMethod
	}

	if path.DELETE != nil {
		return HttpDeleteMethod
	}

	return UnknownMethod
}

func (r ServiceProxy) createRoute(router *echo.Group, path yml.Path) {
	method := detectHttpMethod(path)

	if method == UnknownMethod {
		return
	}

	if method == HttpGetMethod {
		router.GET(*path.GET, r.ProxyFn)
	}

	if method == HttpPostMethod {
		router.POST(*path.POST, r.ProxyFn)
	}

	if method == HttpPutMethod {
		router.PUT(*path.PUT, r.ProxyFn)
	}

	if method == HttpPatchMethod {
		router.PATCH(*path.PATCH, r.ProxyFn)
	}

	if method == HttpDeleteMethod {
		router.DELETE(*path.DELETE, r.ProxyFn)
	}
}

func (r ServiceProxy) toRequests(group *echo.Group) {
	for _, path := range r.Paths {
		r.createRoute(group, path)
	}
}
