package menus

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui"
	"github.com/federico-paolillo/mines/cmd/cli/tui/intents"
)

var uppercaseY = "Y"
var lowercaseN = "n"

func NewQuitMenu(
	console *tui.Console,
	dispatcher *tui.Dispatcher,
) *tui.Dialog {
	return &tui.Dialog{
		Console: console,
		Steps: []tui.Step{
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
		OnCompleteInteraction: func(inputs tui.Inputs) {
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
