package menus

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
)

func NewMainMenu(
	console *console.Console,
	dispatcher *dispatcher.Dispatcher,
) *dialog.Dialog {
	return NewMenu(
		console,
		[]Entry{
			{
				Prompt: "quit",
				Dialog: NewQuitMenu(
					console,
					dispatcher,
				),
			},
			{
				Prompt: "move",
				Dialog: NewMoveMenu(
					console,
					dispatcher,
				),
			},
		},
	)
}
