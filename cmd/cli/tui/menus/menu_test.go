package menus_test

import (
	"strings"
	"testing"

	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
	"github.com/federico-paolillo/mines/cmd/cli/tui/menus"
)

func TestRendersMenuSelectionsCorrectly(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"2\n",
	)

	c := console.NewConsole(
		stdin,
		&stdout,
	)

	m := menus.NewMenu(
		c,
		[]menus.Entry{
			{
				Prompt: "pippo",
				Dialog: menus.NewInhertMenu(c),
			},
			{
				Prompt: "pluto",
				Dialog: menus.NewInhertMenu(c),
			},
			{
				Prompt: "topolino",
				Dialog: menus.NewInhertMenu(c),
			},
		},
	)

	m.Interact(dialog.NoInputs)

	screen := stdout.String()
	screenExpected := "1 pippo\n2 pluto\n3 topolino\n"

	if screen != screenExpected {
		t.Errorf("dialog did not render expected output. wanted '%q' got '%q'", screenExpected, screen)
	}
}

func TestRendersChosesMenuCorrectly(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"2\n",
	)

	c := console.NewConsole(
		stdin,
		&stdout,
	)

	didCall := false

	m := menus.NewMenu(
		c,
		[]menus.Entry{
			{
				Prompt: "pippo",
				Dialog: menus.NewInhertMenu(c),
			},
			{
				Prompt: "pluto",
				Dialog: &dialog.Dialog{
					Console: c,
					Steps:   []dialog.Step{},
					OnCompleteInteraction: func(_ dialog.Inputs) {
						didCall = true
					},
				},
			},
			{
				Prompt: "topolino",
				Dialog: menus.NewInhertMenu(c),
			},
		},
	)

	m.Interact(dialog.NoInputs)

	if !didCall {
		t.Errorf("dialog did not call menu entry")
	}
}
