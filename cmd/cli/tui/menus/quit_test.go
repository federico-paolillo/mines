package menus_test

import (
	"strings"
	"testing"

	"github.com/federico-paolillo/mines/cmd/cli/tui"
	"github.com/federico-paolillo/mines/cmd/cli/tui/intents"
	"github.com/federico-paolillo/mines/cmd/cli/tui/menus"
)

func TestQuitMenuDispatchesQuitApplicationOnConfirm(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"Y\n",
	)

	c := tui.NewConsole(
		stdin,
		&stdout,
	)

	didQuit := false

	d := tui.NewDispatcher()

	_ = d.Subscribe(func(intent any) {
		if _, ok := intent.(intents.QuitApplicationIntent); ok {
			didQuit = true
		}
	})

	m := menus.NewQuitMenu(
		c,
		d,
	)

	m.Interact(tui.NoInputs)

	if !didQuit {
		t.Fatalf("quit menu did not dispatch quit application intent on exit confirmation")
	}
}

func TestQuitMenuDoesNotDispatchQuitApplicationOnDeny(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"n\n",
	)

	c := tui.NewConsole(
		stdin,
		&stdout,
	)

	didQuit := false

	d := tui.NewDispatcher()

	_ = d.Subscribe(func(intent any) {
		if _, ok := intent.(intents.QuitApplicationIntent); ok {
			didQuit = true
		}
	})

	m := menus.NewQuitMenu(
		c,
		d,
	)

	m.Interact(tui.NoInputs)

	if didQuit {
		t.Fatalf("quit menu did dispatch quit application intent on exit denial")
	}
}

func TestQuitMenuRejectsAnswersThatAreNotYN(t *testing.T) {
}
