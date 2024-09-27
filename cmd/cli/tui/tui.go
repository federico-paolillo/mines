package tui

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
	"github.com/federico-paolillo/mines/cmd/cli/tui/intents"
)

type Tui struct {
	menu       *dialog.Dialog
	dispatcher *dispatcher.Dispatcher
	quitting   bool
}

func NewTui(
	dispatcher *dispatcher.Dispatcher,
	menu *dialog.Dialog,
) *Tui {
	return &Tui{
		menu,
		dispatcher,
		false,
	}
}

func (t *Tui) Run() {
	t.quitting = false

	unsub := t.dispatcher.Subscribe(t.handleDispatch)

	defer unsub()

	for {
		if t.quitting {
			break
		}
		//render board
		t.menu.Interact(dialog.NoInputs)
	}
}

func (t *Tui) handleDispatch(intent any) {
	switch intent.(type) {
	case intents.QuitApplicationIntent:
		t.quitting = true
		break
		// case intents.Move:
		// board.DoMove(intent.coords)
	}
}
