package menus_test

import (
	"strings"
	"testing"

	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
	"github.com/federico-paolillo/mines/cmd/cli/tui/intents"
	"github.com/federico-paolillo/mines/cmd/cli/tui/menus"
)

func TestQuitMenuDispatchesQuitApplicationOnConfirm(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"Y\n",
	)

	c := console.NewConsole(
		stdin,
		&stdout,
	)

	didQuit := false

	d := dispatcher.NewDispatcher()

	_ = d.Subscribe(func(intent any) {
		if _, ok := intent.(intents.QuitApplicationIntent); ok {
			didQuit = true
		}
	})

	m := menus.NewQuitMenu(
		c,
		d,
	)

	m.Interact(dialog.NoInputs)

	if !didQuit {
		t.Fatalf("quit menu did not dispatch quit application intent on exit confirmation")
	}
}

func TestQuitMenuDoesNotDispatchQuitApplicationOnDeny(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"n\n",
	)

	c := console.NewConsole(
		stdin,
		&stdout,
	)

	didQuit := false

	d := dispatcher.NewDispatcher()

	_ = d.Subscribe(func(intent any) {
		if _, ok := intent.(intents.QuitApplicationIntent); ok {
			didQuit = true
		}
	})

	m := menus.NewQuitMenu(
		c,
		d,
	)

	m.Interact(dialog.NoInputs)

	if didQuit {
		t.Fatalf("quit menu did dispatch quit application intent on exit denial")
	}
}

func TestQuitMenuRejectsAnswersThatAreNotYN(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"balbalba\nY\n",
	)

	c := console.NewConsole(
		stdin,
		&stdout,
	)

	d := dispatcher.NewDispatcher()

	m := menus.NewQuitMenu(
		c,
		d,
	)

	m.Interact(dialog.NoInputs)

	screen := stdout.String()
	expectedScreen := "do you want to quit ? (Y/n)\nhuh?\n"

	if screen != expectedScreen {
		t.Errorf("dialog did not render expected output. wanted '%q' got '%q'", screen, expectedScreen)
	}
}
