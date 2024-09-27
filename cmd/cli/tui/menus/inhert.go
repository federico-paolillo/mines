package menus

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
)

func NewInhertMenu(console *console.Console) *dialog.Dialog {
	return &dialog.Dialog{
		Console:               console,
		Steps:                 make([]dialog.Step, 0),
		OnCompleteInteraction: func(inputs dialog.Inputs) {},
	}
}
