package board_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

func TestUnopenSafeCellsCountIsCorrect(t *testing.T) {
	bb := board.NewBuilder(dimensions.Size{Width: 2, Height: 2})

	bb.PlaceSafe(1, 1)
	bb.PlaceSafe(2, 1)
	bb.PlaceSafe(1, 2)
	bb.PlaceMine(2, 2)

	b := bb.Build()

	// Starting all safe cells are unopen
	unopenCells := b.CountUnopenSafeCells()
	expectedUnopenCells := 3

	if unopenCells != expectedUnopenCells {
		t.Fatalf(
			"expected to have %d unopen cells. we have %d instead",
			expectedUnopenCells,
			unopenCells,
		)
	}

	// Opening 1,1 will reduce the unopened safe cells by one

	cell := b.Retrieve(dimensions.Location{X: 1, Y: 1})

	cell.Open()

	unopenCells = b.CountUnopenSafeCells()
	expectedUnopenCells = 2

	if unopenCells != expectedUnopenCells {
		t.Fatalf(
			"expected to have %d unopen cells. we have %d instead",
			expectedUnopenCells,
			unopenCells,
		)
	}
}

func TestUnflaggedMinesCountIsCorrect(t *testing.T) {
	bb := board.NewBuilder(dimensions.Size{Width: 2, Height: 2})

	bb.PlaceSafe(1, 1)
	bb.PlaceSafe(2, 1)
	bb.PlaceSafe(1, 2)
	bb.PlaceMine(2, 2)

	b := bb.Build()

	// Starting all mine cells are unflagged
	unflaggedMines := b.CountUnflaggedMines()
	expectedUnflaggedCells := 1

	if unflaggedMines != expectedUnflaggedCells {
		t.Fatalf(
			"expected to have %d unflagged mines. we have %d instead",
			expectedUnflaggedCells,
			unflaggedMines,
		)
	}

	// Flagging 2,2 will reduce the unflagged mines to none

	cell := b.Retrieve(dimensions.Location{X: 2, Y: 2})

	cell.Flag()

	unflaggedMines = b.CountUnflaggedMines()
	expectedUnflaggedCells = 0

	if unflaggedMines != expectedUnflaggedCells {
		t.Fatalf(
			"expected to have %d unflagged mines. we have %d instead",
			expectedUnflaggedCells,
			unflaggedMines,
		)
	}
}