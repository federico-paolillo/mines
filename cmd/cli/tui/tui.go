package tui

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
	"github.com/federico-paolillo/mines/cmd/cli/tui/intents"
	"github.com/federico-paolillo/mines/cmd/cli/tui/menus"
	"github.com/federico-paolillo/mines/pkg/game"
)

type Tui struct {
	quitting   bool
	menu       *dialog.Dialog
	viewer     *menus.BoardView
	dispatcher *dispatcher.Dispatcher
	game       *game.Game
}

func NewTui(
	dispatcher *dispatcher.Dispatcher,
	menu *dialog.Dialog,
	viewer *menus.BoardView,
	game *game.Game,
) *Tui {
	return &Tui{
		false,
		menu,
		viewer,
		dispatcher,
		game,
	}
}

func (t *Tui) Run() {
	t.quitting = false

	unsub := t.dispatcher.Subscribe(t.handleDispatch)

	defer unsub()

	for {
		s := t.game.Status()

		if s == game.Won {
			t.menu.Console.Printline("you won :)")
			break
		}

		if s == game.Won {
			t.menu.Console.Printline("you lost :(")
			break
		}

		if t.quitting {
			break
		}

		t.viewer.Render()
		t.menu.Interact(dialog.NoInputs)
	}
}

func (t *Tui) handleDispatch(intent any) {
	switch i := intent.(type) {
	case intents.QuitApplicationIntent:
		t.quitting = true
		break
	case intents.FlagIntent:
		t.game.Flag(i.X, i.Y)
		break
	case intents.OpenIntent:
		t.game.Open(i.X, i.Y)
		break
	case intents.ChordIntent:
		t.game.Chord(i.X, i.Y)
		break
	}
}
