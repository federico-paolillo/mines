package menus_test

import (
	"strings"
	"testing"

	"github.com/federico-paolillo/mines/cmd/cli/tui"
	"github.com/federico-paolillo/mines/cmd/cli/tui/menus"
)

func TestRendersMenuSelectionsCorrectly(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"2\n",
	)

	c := tui.NewConsole(
		stdin,
		&stdout,
	)

	d := tui.NewDispatcher()

	m := menus.NewMenu(
		c,
		d,
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

	m.Interact(tui.NoInputs)

	screen := stdout.String()
	screenExpected := "1 pippo\n2 pluto\n3 topolino\n"

	if screen != screenExpected {
		t.Errorf("dialog did not render expected output. wanted '%q' got '%q'", screen, screenExpected)
	}
}

func TestRendersChosesMenuCorrectly(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"2\n",
	)

	c := tui.NewConsole(
		stdin,
		&stdout,
	)

	d := tui.NewDispatcher()

	didCall := false

	m := menus.NewMenu(
		c,
		d,
		[]menus.Entry{
			{
				Prompt: "pippo",
				Dialog: menus.NewInhertMenu(c),
			},
			{
				Prompt: "pluto",
				Dialog: &tui.Dialog{
					Console: c,
					Steps:   []tui.Step{},
					OnCompleteInteraction: func(_ tui.Inputs) {
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

	m.Interact(tui.NoInputs)

	if !didCall {
		t.Errorf("dialog did not call menu entry")
	}
}
