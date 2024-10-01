package tui

import (
	"io"

	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
	"github.com/federico-paolillo/mines/cmd/cli/tui/menus"
	"github.com/federico-paolillo/mines/internal/generators"
	"github.com/federico-paolillo/mines/internal/printers"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

func InitTui(
	stdin io.Reader,
	stdout io.Writer,
) *Tui {
	g := generators.NewRngBoardGenerator(1234)

	b := g.Generate(dimensions.Size{Width: 4, Height: 4}, 2)

	c := console.NewConsole(
		stdin,
		stdout,
	)

	d := dispatcher.NewDispatcher()

	p := printers.NewAsciiPrinter()

	v := menus.NewBoardView(
		c,
		b,
		p,
	)

	m := menus.NewMainMenu(
		c,
		d,
	)

	return NewTui(
		d,
		m,
		v,
	)
}
