package runner

import (
	"log/slog"
	"os"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

type ProgramE = func(
	*mines.Mines,
	config.Root,
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

	// TODO: Read it from file and env.
	cfg := config.Root{
		Seed: 1234,
		Server: config.Server{
			Host: "",
			Port: "65000",
		},
	}

	mines := mines.NewMines(
		logger,
		cfg,
	)

	err := program(
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

	return Ok
}
