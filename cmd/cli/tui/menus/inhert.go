package menus

import "github.com/federico-paolillo/mines/cmd/cli/tui"

func NewInhertMenu(console *tui.Console) *tui.Dialog {
	return &tui.Dialog{
		Console:               console,
		Steps:                 make([]tui.Step, 0),
		OnCompleteInteraction: func(inputs tui.Inputs) {},
	}
}
