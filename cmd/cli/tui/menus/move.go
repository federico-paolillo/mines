package menus

import (
	"github.com/federico-paolillo/mines/cmd/cli/tui/console"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dialog"
	"github.com/federico-paolillo/mines/cmd/cli/tui/dispatcher"
	"github.com/federico-paolillo/mines/cmd/cli/tui/intents"
	"github.com/federico-paolillo/mines/cmd/cli/tui/validation"
)

type PositionSelectionHandler = func(x, y int)

func dispatchFlagIntent(d *dispatcher.Dispatcher) PositionSelectionHandler {
	return func(x, y int) {
		d.Dispatch(intents.FlagIntent{Coords: intents.Coords{X: x, Y: y}})
	}
}

func dispatchChordIntent(d *dispatcher.Dispatcher) PositionSelectionHandler {
	return func(x, y int) {
		d.Dispatch(intents.ChordIntent{Coords: intents.Coords{X: x, Y: y}})
	}
}

func dispatchOpenIntent(d *dispatcher.Dispatcher) PositionSelectionHandler {
	return func(x, y int) {
		d.Dispatch(intents.OpenIntent{Coords: intents.Coords{X: x, Y: y}})
	}
}

func NewMoveMenu(
	console *console.Console,
	dispatcher *dispatcher.Dispatcher,
) *dialog.Dialog {
	return NewMenu(
		console,
		[]Entry{
			{
				Prompt: "flag",
				Dialog: NewMoveDialog(
					console,
					dispatchFlagIntent(dispatcher),
				),
			},
			{
				Prompt: "open",
				Dialog: NewMoveDialog(
					console,
					dispatchOpenIntent(dispatcher),
				),
			},
			{
				Prompt: "chord",
				Dialog: NewMoveDialog(
					console,
					dispatchChordIntent(dispatcher),
				),
			},
		},
	)
}

func NewMoveDialog(
	console *console.Console,
	onPositionSelection PositionSelectionHandler,
) *dialog.Dialog {
	return &dialog.Dialog{
		Console: console,
		Steps: []dialog.Step{
			{
				Name: "x",
				Prompt: []string{
					"enter 'X' coordinate",
				},
				Validate: validation.IsNumber,
			},
			{
				Name: "y",
				Prompt: []string{
					"enter 'Y' coordinate",
				},
				Validate: validation.IsNumber,
			},
			{
				Name: "confirm",
				Prompt: []string{
					"are you sure ? (Y/n)",
				},
				Validate: validation.IsYN,
			},
		},
		OnCompleteInteraction: func(inputs dialog.Inputs) {
			didConfirm := inputs["confirm"]

			if validation.IsN(didConfirm) {
				return
			}

			x := validation.ToNumberUnsafely(inputs["x"])
			y := validation.ToNumberUnsafely(inputs["y"])

			onPositionSelection(x, y)
		},
	}
}
