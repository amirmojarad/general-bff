package main

import (
	"fmt"
	"general-bff/cmd/server"
	"log/slog"
)

func main() {
	if err := server.Run(); err != nil {
		slog.Error(fmt.Sprintf("failed to run server %s", err.Error()))
	}
}
