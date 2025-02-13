package server

import (
	"fmt"
	"log/slog"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/gin-gonic/gin"
)

func Program(
	mines *mines.Mines,
	cfg *config.Root,
) error {
	gin.SetMode(gin.ReleaseMode)

	server := NewServer(
		mines,
		cfg.Server,
	)

	mines.Logger.Info(
		"server: listening",
		slog.String("endpoint", cfg.Server.Endpoint()),
	)

	shutdownCron := maybeStartCron(mines)

	defer shutdownCron()

	err := server.ListenAndServe()
	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak errors that are not under our control
		return fmt.Errorf(
			"server: failed to listen and serve. %v",
			err,
		)
	}

	return nil
}
