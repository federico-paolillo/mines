package commands

import (
	"github.com/federico-paolillo/mines/internal/console"
	"github.com/spf13/cobra"
)

func NewPlayCommand(
	console *console.Console,
) *cobra.Command {
	return &cobra.Command{
		Use:   "play",
		Short: "Play a game of Minesweeper right in the terminal",
		Run: func(cmd *cobra.Command, args []string) {
			runPlayCommand(console)
		},
	}
}

func runPlayCommand(
	console *console.Console,
) {
	// TODO: Do nothing for now
}
