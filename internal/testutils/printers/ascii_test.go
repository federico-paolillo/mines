package printers_test

import (
	"testing"

	"github.com/federico-paolillo/mines/internal/testutils/printers"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/game"
)

func TestPrintsBoardProperlyWithDefaultSymbols(t *testing.T) {
	bb := board.NewBuilder(dimensions.Size{Width: 3, Height: 3})

	bb.PlaceSafe(1, 1)
	bb.PlaceMine(2, 1)
	bb.PlaceSafe(3, 1)
	bb.PlaceSafe(1, 2)
	bb.PlaceMine(2, 2)
	bb.PlaceSafe(3, 2)
	bb.PlaceSafe(1, 3)
	bb.PlaceMine(2, 3)
	bb.PlaceSafe(3, 3)

	b := bb.Build()

	expected :=
		`ooo
ooo
ooo
`

	result := printers.Render(b)

	if result != expected {
		t.Fatalf(
			"board was rendered as %s. we wanted %s",
			result,
			expected,
		)
	}
}
func TestPrintsBoardProperlyWithDefaultSymbols2(t *testing.T) {
	bb := board.NewBuilder(dimensions.Size{Width: 3, Height: 4})

	bb.PlaceSafe(1, 1)
	bb.PlaceSafe(2, 1)
	bb.PlaceSafe(3, 1)
	bb.PlaceSafe(1, 2)
	bb.PlaceSafe(2, 2)
	bb.PlaceSafe(3, 2)
	bb.PlaceSafe(1, 3)
	bb.PlaceSafe(2, 3)
	bb.PlaceMine(3, 3)
	bb.PlaceSafe(1, 4)
	bb.PlaceSafe(2, 4)
	bb.PlaceSafe(3, 4)

	b := bb.Build()

	game := game.NewGame(1, b)

	game.Open(1, 1)

	expected :=
		`000
011
01o
01o
`
	result := printers.Render(b)

	if result != expected {
		t.Fatalf(
			"board was rendered as %s. we wanted %s",
			result,
			expected,
		)
	}
}
