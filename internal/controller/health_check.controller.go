package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheck struct{}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (hc *HealthCheck) Health(ctx echo.Context) error {
	if err := ctx.NoContent(http.StatusOK); err != nil {
		ctx.Error(err)
	}

	return nil
}
