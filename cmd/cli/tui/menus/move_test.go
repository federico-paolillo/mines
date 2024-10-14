package menus_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
	"github.com/federico-paolillo/mines/cmd/cli/tui/intents"
	"github.com/federico-paolillo/mines/cmd/cli/tui/menus"
)

func TestMoveMenuDispatchesCorrectIntent(t *testing.T) {
	moveMenuDispatchesProperIntent[intents.OpenIntent](
		t,
		"2\n2\n3\nY\n",
		2,
		3,
	)

	moveMenuDispatchesProperIntent[intents.FlagIntent](
		t,
		"1\n2\n3\nY\n",
		2,
		3,
	)

	moveMenuDispatchesProperIntent[intents.ChordIntent](
		t,
		"3\n2\n3\nY\n",
		2,
		3,
	)
}

func moveMenuDispatchesProperIntent[Intent intents.Position](
	t *testing.T,
	stdinSequence string,
	expectedX int,
	expectedY int,
) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		stdinSequence,
	)

	c := console.NewConsole(
		stdin,
		&stdout,
	)

	d := dispatcher.NewDispatcher()

	m := menus.NewMoveMenu(
		c,
		d,
	)

	var intent interface{}

	_ = d.Subscribe(func(i any) {
		intent = i
	})

	m.Interact(dialog.NoInputs)

	if move, ok := intent.(Intent); ok {
		if move.XCoord() != expectedX {
			t.Errorf(
				"wrong move x coord. expected '%d' got '%d'",
				move.XCoord(),
				expectedX,
			)
		}

		if move.YCoord() != expectedY {
			t.Errorf(
				"wrong move y coord. expected '%d' got '%d'",
				move.YCoord(),
				expectedY,
			)
		}
	} else {
		t.Errorf(
			"dispatched wrong intent. wanted '%s' got '%s'",
			reflect.TypeFor[Intent]().String(),
			reflect.TypeOf(intent).Name(),
		)
	}
}
func TestMoveMenuDoesNotReturnChosenCoordinatesOnDecline(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"1\n1\n1\nn\n",
	)

	c := console.NewConsole(
		stdin,
		&stdout,
	)

	d := dispatcher.NewDispatcher()

	m := menus.NewMoveMenu(
		c,
		d,
	)

	didDispatch := false

	_ = d.Subscribe(func(i any) {
		didDispatch = true
	})

	m.Interact(dialog.NoInputs)

	if didDispatch {
		t.Fatal("dispatched unexpected intent")
	}
}

func TestMoveMenuRejectsNonNumericalCoordinates(t *testing.T) {
	var stdout strings.Builder

	var stdin = strings.NewReader(
		"1\n1\na\n1\nn\n",
	)

	c := console.NewConsole(
		stdin,
		&stdout,
	)

	d := dispatcher.NewDispatcher()

	m := menus.NewMoveMenu(
		c,
		d,
	)

	didDispatch := false

	_ = d.Subscribe(func(i any) {
		didDispatch = true
	})

	m.Interact(dialog.NoInputs)

	if didDispatch {
		t.Fatal("dispatched unexpected intent")
	}
}
