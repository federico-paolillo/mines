package tui

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui/intents"
)

type Tui struct {
	menu       *Dialog
	dispatcher *Dispatcher
	quitting   bool
}

func NewTui(
	dispatcher *Dispatcher,
	menu *Dialog,
) *Tui {
	return &Tui{
		menu,
		dispatcher,
		false,
	}
}

func (t *Tui) Loop() {
	t.quitting = false

	unsub := t.dispatcher.Subscribe(t.handleDispatch)

	defer unsub()

	for {
		if t.quitting {
			break
		}
		//render board
		t.menu.Interact(NoInputs)
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
