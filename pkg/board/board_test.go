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

func TestUnopenSafeCellsCountIsCorrect2(t *testing.T) {
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

	// Flagging 1,1 will not reduce the unopened safe cells by one

	cell := b.Retrieve(dimensions.Location{X: 1, Y: 1})

	cell.Flag()

	unopenCells = b.CountUnopenSafeCells()
	expectedUnopenCells = 3

	if unopenCells != expectedUnopenCells {
		t.Fatalf(
			"expected to have %d unopen cells. we have %d instead",
			expectedUnopenCells,
			unopenCells,
		)
	}
}

func TestCountAdjacentCellOfStatusIsCorrect(t *testing.T) {
	bb := board.NewBuilder(dimensions.Size{Width: 2, Height: 2})

	bb.PlaceSafe(1, 1)
	bb.PlaceSafe(2, 1)
	bb.PlaceSafe(1, 2)
	bb.PlaceMine(2, 2)

	b := bb.Build()

	c := b.Retrieve(dimensions.Location{X: 2, Y: 1})

	c.Flag()

	count := b.CountAdjacentCellsOfStatus(board.Flagged, dimensions.Location{X: 1, Y: 1})

	if count != 1 {
		t.Fatalf(
			"expected to count %d flagged cells. counted %d",
			1,
			count,
		)
	}
}

func TestRetrieveAdjacentCellOfStatusIsCorrect(t *testing.T) {
	bb := board.NewBuilder(dimensions.Size{Width: 2, Height: 2})

	bb.PlaceSafe(1, 1)
	bb.PlaceSafe(2, 1)
	bb.PlaceSafe(1, 2)
	bb.PlaceMine(2, 2)

	b := bb.Build()

	c := b.Retrieve(dimensions.Location{X: 2, Y: 1})

	c.Flag()

	cells := b.RetrieveAdjacentCellsOfStatus(board.Flagged, dimensions.Location{X: 1, Y: 1})

	actualLen := len(cells)
	expectedLen := 1

	if actualLen != expectedLen {
		t.Fatalf(
			"expected cells to be %d len. instead got %d",
			expectedLen,
			actualLen,
		)
	}

	expectedPosition := dimensions.Location{X: 2, Y: 1}
	actualPosition := cells[0].Position()

	if actualPosition != expectedPosition {
		t.Fatalf(
			"expected cells retrieved to be at %v. instead is %v",
			expectedPosition,
			actualPosition,
		)
	}
}
