package menus

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
	"github.com/federico-paolillo/mines/cmd/cli/tui/intents"
	"github.com/federico-paolillo/mines/cmd/cli/tui/validation"
)

func NewQuitMenu(
	console *console.Console,
	dispatcher *dispatcher.Dispatcher,
) *dialog.Dialog {
	return &dialog.Dialog{
		Console: console,
		Steps: []dialog.Step{
			{
				Prompt: []string{
					"do you want to quit ? (Y/n)",
				},
				Name: "quit",
				Validate: func(value string) bool {
					return validation.IsYN(value)
				},
			},
		},
		OnCompleteInteraction: func(inputs dialog.Inputs) {
			if validation.IsY(inputs["quit"]) {
				dispatcher.Dispatch(intents.QuitApplication)
			}
		},
	}
}
