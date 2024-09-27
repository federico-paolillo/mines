package menus

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
	"github.com/federico-paolillo/mines/cmd/cli/tui/intents"
)

var uppercaseY = "Y"
var lowercaseN = "n"

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
					return isYn(value)
				},
			},
		},
		OnCompleteInteraction: func(inputs dialog.Inputs) {
			if isY(inputs["quit"]) {
				dispatcher.Dispatch(intents.QuitApplication)
			}
		},
	}
}

func isYn(value string) bool {
	return isY(value) || isn(value)
}

func isY(value string) bool {
	return value == uppercaseY
}

func isn(value string) bool {
	return value == lowercaseN
}
