package controller

import (
	"general-bff/config/yml"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxyConfig struct {
	ReverseProxies []ServiceProxy
}

type ServiceProxy struct {
	name         string
	svcUrl       string
	reverseProxy *httputil.ReverseProxy
	url          *url.URL
	Proxy        func(ctx echo.Context) error
	Paths        []yml.Path
}

func NewReverseProxyConfig(servicesCfg *yml.ServicesConfig) (*ReverseProxyConfig, error) {
	reverseProxies, err := getReverseProxies(servicesCfg)
	if err != nil {
		return nil, err
	}

	return &ReverseProxyConfig{
		ReverseProxies: reverseProxies,
	}, nil
}

func getReverseProxies(servicesCfg *yml.ServicesConfig) ([]ServiceProxy, error) {
	serviceProxies := make([]ServiceProxy, 0)

	for _, service := range servicesCfg.Services {
		if !service.Enabled {
			continue
		}

		serviceProxy := ServiceProxy{
			name:   service.Name,
			svcUrl: service.ServiceUrl,
			Paths:  service.Paths,
		}

		svcURL, err := url.Parse(service.ServiceUrl)
		if err != nil {
			return nil, err
		}

		serviceProxy.url = svcURL
		serviceProxy.reverseProxy = createReverseProxy(serviceProxy.url)
		serviceProxy.Proxy = func(ctx echo.Context) error {
			serviceProxy.reverseProxy.ServeHTTP(ctx.Response(), ctx.Request())

			return nil
		}

		serviceProxies = append(serviceProxies, serviceProxy)
	}

	return serviceProxies, nil
}

func createReverseProxy(url *url.URL, requestsModifiers ...func(*http.Request)) *httputil.ReverseProxy {
	proxy := httputil.NewSingleHostReverseProxy(url)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = url.Host

		for _, requestsModifier := range requestsModifiers {
			requestsModifier(req)
		}
	}

	return proxy
}
