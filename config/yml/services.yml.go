package yml

import (
	"general-bff/config/env"
	"gopkg.in/yaml.v2"
	"log/slog"
	"os"
)

type Path struct {
	POST  *string `yaml:"POST,omitempty"`
	GET   *string `yaml:"GET,omitempty"`
	PATCH *string `yaml:"PATCH,omitempty"`
	PUT   *string `yaml:"PUT,omitempty"`
}

type Service struct {
	Name       string `yaml:"name"`
	Enabled    bool   `yaml:"enabled"`
	ServiceUrl string `yaml:"service_url"`
	Paths      []Path `yaml:"paths"`
}

type ServicesConfig struct {
	Services []Service `yaml:"services"`
}

func NewServicesConfig(cfg *env.AppConfig) *ServicesConfig {
	servicesConfig := read(cfg.Services.ServicesPath)
	return &servicesConfig
}

func read(path string) ServicesConfig {
	data, err := os.ReadFile(path)

	var config ServicesConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		slog.Error(err.Error())
	}

	return config
}
