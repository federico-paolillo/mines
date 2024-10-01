package main

import (
	"os"

	"github.com/federico-paolillo/mines/cmd/cli/commands"
	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/spf13/cobra"
)

func main() {
	c := console.NewConsole(
		os.Stdin,
		os.Stdout,
	)

	r := commands.NewRootCmd(
		c,
	)

	err := r.Execute()

	cobra.CheckErr(err)
}
