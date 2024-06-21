package logger

import (
	"log/slog"
	"os"
)

//nolint:gochecknoinits
func init() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})

	slog.SetDefault(slog.New(logHandler))
}
