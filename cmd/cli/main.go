package main

import (
	"os"

	"github.com/federico-paolillo/mines/cmd/cli/tui"
	"github.com/federico-paolillo/mines/cmd/cli/tui/menus"
)

func main() {
	c := tui.NewConsole(os.Stdin, os.Stdout)

	d := tui.NewDispatcher()

	m := menus.NewMainMenu(c, d)

	t := tui.NewTui(d, m)

	t.Loop()
}
