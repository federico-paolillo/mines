package runner

import (
	"context"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

type ProgramE = func(
	context.Context,
	*mines.Mines,
	*config.Root,
) error

type StatusCode = string

const (
	Ok    StatusCode = "ok"
	NotOk StatusCode = "nok"
)
