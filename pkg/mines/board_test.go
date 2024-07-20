package mines_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/mines"
)

func TestAdjacentMinesRecalculateWhenPlacingNewCells(t *testing.T) {
	board := mines.NewBoard(mines.Size{2, 2})

	board.PlaceCell(mines.Location{1, 1})
	board.PlaceMine(mines.Location{2, 1})
	board.PlaceMine(mines.Location{1, 2})
	board.PlaceCell(mines.Location{2, 2})

	expectedMines := 2

	adjacentMinesAt22 := board.AdjacentMines(mines.Location{2, 2})

	if adjacentMinesAt22 != expectedMines {
		t.Fatalf(
			"expected cell at 2,2 to have %d adjacent mines. instead it has %d adjacent mines",
			expectedMines,
			adjacentMinesAt22,
		)
	}
}

func TestEmptyCellsArePlaced(t *testing.T) {
	board := mines.NewBoard(mines.Size{2, 2})

	expectedLocation := mines.Location{1, 2}

	board.PlaceCell(expectedLocation)

	cell := board.Retrieve(expectedLocation)

	if cell.Position != expectedLocation {
		t.Errorf(
			"retrieved cell is not in the expected position. it is %v instead of %v",
			cell.Position,
			expectedLocation,
		)
	}

	if cell.Mined {
		t.Error("retrieved cell is mined")
	}

	if cell.Status != mines.Closed {
		t.Error("retrieved cell is not closed")
	}
}

func TestMinedCellsArePlaced(t *testing.T) {
	board := mines.NewBoard(mines.Size{2, 2})

	expectedLocation := mines.Location{1, 2}

	board.PlaceMine(expectedLocation)

	cell := board.Retrieve(expectedLocation)

	if cell.Position != expectedLocation {
		t.Errorf(
			"retrieved cell is not in the expected position. it is %v instead of %v",
			cell.Position,
			expectedLocation,
		)
	}

	if !cell.Mined {
		t.Error("retrieved cell is not mined")
	}

	if cell.Status != mines.Closed {
		t.Error("retrieved cell is not closed")
	}
}

func TestCellsCanBeVoided(t *testing.T) {
	board := mines.NewBoard(mines.Size{2, 2})

	location := mines.Location{1, 2}

	board.PlaceCell(location)

	cell := board.Retrieve(location)

	if cell == mines.Void {
		t.Fatalf(
			"should have been a cell at %v",
			location,
		)
	}

	board.PlaceVoid(location)

	cell = board.Retrieve(location)

	if cell != mines.Void {
		t.Fatalf(
			"should NOT have been a cell at %v",
			location,
		)
	}
}
