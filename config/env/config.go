package env

import (
	"os"
	"strconv"
)

type AppConfig struct {
	App struct {
		Name string
		Port int
		Env  string
	}

	Services struct {
		ServicesPath string
	}
}

func NewConfig() (*AppConfig, error) {
	cfg := &AppConfig{}

	if err := setAppEnv(cfg); err != nil {
		return nil, err
	}

	setServicesPath(cfg)

	return cfg, nil
}

func setAppEnv(cfg *AppConfig) error {
	cfg.App.Env = os.Getenv("APP_ENV")

	port := os.Getenv("APP_PORT")
	portAsInt, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	cfg.App.Port = portAsInt
	cfg.App.Env = os.Getenv("APP_ENV")

	return nil
}

func setServicesPath(cfg *AppConfig) {
	cfg.Services.ServicesPath = os.Getenv("SERVICES_PATH")
}
