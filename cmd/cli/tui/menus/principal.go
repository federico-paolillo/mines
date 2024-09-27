package menus

import "github.com/federico-paolillo/mines/cmd/cli/tui"

func NewMainMenu(
	console *tui.Console,
	dispatcher *tui.Dispatcher,
) *tui.Dialog {
	return NewMenu(
		console,
		dispatcher,
		[]Entry{
			{
				Prompt: "quit",
				Dialog: NewQuitMenu(
					console,
					dispatcher,
				),
			},
		},
	)
}
