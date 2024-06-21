package middlewares

import (
	"general-bff/config/yml"
)

type Middleware struct {
	servicesCfg *yml.ServicesConfig
}

func NewMiddleware(servicesCfg *yml.ServicesConfig) *Middleware {
	return &Middleware{servicesCfg: servicesCfg}
}
