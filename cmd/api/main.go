package main

import (
	"fmt"
	"os"

	"github.com/federico-paolillo/mines/internal/runner"
	"github.com/federico-paolillo/mines/internal/server"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

func main() {
	statusCode := runner.Run(serverProgram)

	if statusCode == runner.NotOk {
		os.Exit(1)
	}
}

func serverProgram(
	mines *mines.Mines,
	cfg config.Root,
) error {
	server := server.NewServer(
		mines,
		cfg.Server,
	)

	err := server.ListenAndServe()
	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak errors that are not under our control
		return fmt.Errorf(
			"api: failed to listen and serve. %v",
			err,
		)
	}

	return nil
}
