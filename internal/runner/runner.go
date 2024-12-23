package runner

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/sherifabdlnaby/configuro"
)

type ProgramE = func(
	*mines.Mines,
	*config.Root,
) error

type StatusCode = int

const (
	Ok    StatusCode = 0
	NotOk StatusCode = 1
)

func Run(program ProgramE) StatusCode {
	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			nil,
		),
	)

	cfg, err := loadConfiguration()
	if err != nil {
		logger.Error(
			"runner: failed to load configuration",
			slog.Any("err", err),
		)

		return NotOk
	}

	mines := mines.NewMines(
		logger,
		cfg,
	)

	err = program(
		mines,
		cfg,
	)
	if err != nil {
		logger.Error(
			"runner: failed to run program",
			slog.Any("err", err),
		)

		return NotOk
	}

	logger.Info("runner: program completed 'ok'")

	return Ok
}

func loadConfiguration() (*config.Root, error) {
	cfguro, err := configuro.NewConfig(
		configuro.WithLoadFromEnvVars("MINES_"),
		configuro.WithLoadFromConfigFile("config.yml", false),
		configuro.WithoutEnvConfigPathOverload(),
		configuro.WithoutLoadDotEnv(),
	)
	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak errors that are not under our control
		return nil, fmt.Errorf(
			"runner: failed to setup configuro. %v",
			err,
		)
	}

	cfg := &config.Root{
		Seed: 1234,
		Server: config.Server{
			Host: "",
			Port: "65000",
		},
	}

	err = cfguro.Load(cfg)
	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak errors that are not under our control
		return nil, fmt.Errorf(
			"runner: failed to bind configuration. %v",
			err,
		)
	}

	return cfg, nil
}
