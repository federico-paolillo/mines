package main

import (
	"os"

	"github.com/federico-paolillo/mines/cmd/cli/tui"
)

func main() {
	t := tui.InitTui(
		os.Stdin,
		os.Stdout,
	)

	t.Run()
}
