package server

import (
	"general-bff/config/env"
	"general-bff/config/yml"
	"log/slog"
)

func Run() error {
	cfg, err := env.NewConfig()
	if err != nil {
		return err
	}

	services := yml.NewServicesConfig(cfg)
	slog.Info(slog.Any("service-name", services.Services[0].Name).String())
	return nil
}
