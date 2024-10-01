package commands

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/spf13/cobra"
)

func NewRootCmd(
	console *console.Console,
) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mines [command]",
		Short: "A Minesweeper server with built-in TUI",
		Long: `A Minesweeper server with built-in TUI.

You can run this CLI and play directly using a TUI with 'mines play'.
You run a UNIX Socket-based server and plug you own UI with 'mines serve'`,
		SilenceUsage: true,
	}

	rootCmd.SetOut(console.Stdout)
	rootCmd.SetIn(console.Stdin)

	rootCmd.AddCommand(NewPlayCommand(console))

	return rootCmd
}
