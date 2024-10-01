package tui

import (
	"io"

	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
	"github.com/federico-paolillo/mines/cmd/cli/tui/menus"
)

func InitTui(
	stdin io.Reader,
	stdout io.Writer,
) *Tui {
	c := console.NewConsole(
		stdin,
		stdout,
	)

	d := dispatcher.NewDispatcher()

	m := menus.NewMainMenu(
		c,
		d,
	)

	return NewTui(
		d,
		m,
	)
}
