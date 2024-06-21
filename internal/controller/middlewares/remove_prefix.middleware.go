package middlewares

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func removePrefix(c echo.Context, prefix string) {
	path := c.Request().URL.Path
	path = strings.Replace(path, prefix, "", 1)
	c.SetPath(path)
	c.Request().URL.Path = path

	rawPath := c.Request().URL.RawPath
	rawPath = strings.Replace(rawPath, prefix, "", 1)
	c.Request().URL.RawPath = rawPath

	uri := c.Request().URL.RequestURI()
	uri = strings.Replace(uri, prefix, "", 1)
	c.Request().RequestURI = uri
}

func (m *Middleware) RemovePrefix(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		removePrefix(c, m.servicesCfg.Middlewares.Prefixes.API)
		removePrefix(c, m.servicesCfg.Middlewares.Prefixes.APIBo)

		if err := next(c); err != nil {
			return err
		}
		return nil
	}
}
