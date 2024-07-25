package mines_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/mines"
)

func TestAdjacentMinesCalculatesProperly(t *testing.T) {
	/*
	 * Assume a board like:
	 * 2 M
	 * M 2
	 * where x is a closed empty cell
	 * 			 1 is a cell with adjacent mines
	 *       M is a mined cell
	 */

	board := mines.NewBoard(mines.Size{Width: 2, Height: 2})

	board.PlaceCell(mines.Location{X: 1, Y: 1})
	board.PlaceMine(mines.Location{X: 2, Y: 1})
	board.PlaceMine(mines.Location{X: 1, Y: 2})
	board.PlaceCell(mines.Location{X: 2, Y: 2})

	minesExpected := [4]struct {
		location mines.Location
		expected int
	}{
		{mines.Location{X: 1, Y: 1}, 2},
		{mines.Location{X: 2, Y: 1}, 1},
		{mines.Location{X: 1, Y: 2}, 1},
		{mines.Location{X: 2, Y: 2}, 2},
	}

	for _, expectation := range minesExpected {
		adjacentMines := board.AdjacentMines(expectation.location)

		if adjacentMines != expectation.expected {
			t.Errorf(
				"expected cell at %v to have %d mines. instead it has %d",
				expectation.location,
				expectation.expected,
				adjacentMines,
			)
		}
	}
}

func TestAdjacentMinesCalculatesProperly2(t *testing.T) {
	/*
	 * Assume a board like:
	 * x x x
	 * x 1 1
	 * x 1 M
	 * x 1 1
	 * where x is a closed empty cell
	 * 			 1 is a cell with adjacent mines
	 *       M is a mined cell
	 * opening cell a 1,1 should chord and produce a board like:
	 * # o o
	 * o 1 1
	 * o 1 M
	 * o 1 1
	 * where o is a an chording empty cell opened
	 *       # is the cell that was opened
	 * 			 1 is a cell with adjacent mines
	 *       M is a mined cell
	 */

	board := mines.NewBoard(mines.Size{Width: 3, Height: 4})

	board.PlaceCell(mines.Location{X: 1, Y: 1})
	board.PlaceCell(mines.Location{X: 2, Y: 1})
	board.PlaceCell(mines.Location{X: 3, Y: 1})
	board.PlaceCell(mines.Location{X: 1, Y: 2})
	board.PlaceCell(mines.Location{X: 2, Y: 2})
	board.PlaceCell(mines.Location{X: 3, Y: 2})
	board.PlaceCell(mines.Location{X: 1, Y: 3})
	board.PlaceCell(mines.Location{X: 2, Y: 3})
	board.PlaceMine(mines.Location{X: 3, Y: 3})
	board.PlaceCell(mines.Location{X: 1, Y: 4})
	board.PlaceCell(mines.Location{X: 2, Y: 4})
	board.PlaceCell(mines.Location{X: 3, Y: 4})

	minesExpected := [12]struct {
		location mines.Location
		expected int
	}{
		{mines.Location{X: 1, Y: 1}, 0},
		{mines.Location{X: 2, Y: 1}, 0},
		{mines.Location{X: 3, Y: 1}, 0},
		{mines.Location{X: 1, Y: 2}, 0},
		{mines.Location{X: 2, Y: 2}, 1},
		{mines.Location{X: 3, Y: 2}, 1},
		{mines.Location{X: 1, Y: 3}, 0},
		{mines.Location{X: 2, Y: 3}, 1},
		{mines.Location{X: 3, Y: 3}, 0},
		{mines.Location{X: 1, Y: 4}, 0},
		{mines.Location{X: 2, Y: 4}, 1},
		{mines.Location{X: 3, Y: 4}, 1},
	}

	for _, expectation := range minesExpected {
		adjacentMines := board.AdjacentMines(expectation.location)

		if adjacentMines != expectation.expected {
			t.Errorf(
				"expected cell at %v to have %d mines. instead it has %d",
				expectation.location,
				expectation.expected,
				adjacentMines,
			)
		}
	}
}

func TestEmptyCellsArePlaced(t *testing.T) {
	board := mines.NewBoard(mines.Size{Width: 2, Height: 2})

	expectedLocation := mines.Location{X: 1, Y: 2}

	board.PlaceCell(expectedLocation)

	cell := board.Retrieve(expectedLocation)

	if cell.Position() != expectedLocation {
		t.Errorf(
			"retrieved cell is not in the expected position. it is %v instead of %v",
			cell.Position(),
			expectedLocation,
		)
	}

	if cell.Mined() != false {
		t.Error("retrieved cell is mined")
	}

	if cell.Status(mines.Closed) != true {
		t.Error("retrieved cell is not closed")
	}
}

func TestMinedCellsArePlaced(t *testing.T) {
	board := mines.NewBoard(mines.Size{Width: 2, Height: 2})

	expectedLocation := mines.Location{X: 1, Y: 2}

	board.PlaceMine(expectedLocation)

	cell := board.Retrieve(expectedLocation)

	if cell.Position() != expectedLocation {
		t.Errorf(
			"retrieved cell is not in the expected position. it is %v instead of %v",
			cell.Position(),
			expectedLocation,
		)
	}

	if cell.Mined() != true {
		t.Error("retrieved cell is not mined")
	}

	if cell.Status(mines.Closed) != true {
		t.Error("retrieved cell is not closed")
	}
}

func TestCellsCanBeVoided(t *testing.T) {
	board := mines.NewBoard(mines.Size{Width: 2, Height: 2})

	location := mines.Location{X: 1, Y: 2}

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

func TestUnopenSafeCellsCountIsCorrect(t *testing.T) {
	board := mines.NewBoard(mines.Size{Width: 2, Height: 2})

	board.PlaceCell(mines.Location{X: 1, Y: 1})
	board.PlaceCell(mines.Location{X: 2, Y: 1})
	board.PlaceCell(mines.Location{X: 1, Y: 2})
	board.PlaceMine(mines.Location{X: 2, Y: 2})

	// Starting all safe cells are unopen
	unopenCells := board.CountUnopenSafeCells()
	expectedUnopenCells := 3

	if unopenCells != expectedUnopenCells {
		t.Fatalf(
			"expected to have %d unopen cells. we have %d instead",
			expectedUnopenCells,
			unopenCells,
		)
	}

	// Opening 1,1 will reduce the unopened safe cells by one

	cell := board.Retrieve(mines.Location{X: 1, Y: 1})

	cell.Open()

	unopenCells = board.CountUnopenSafeCells()
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
	board := mines.NewBoard(mines.Size{Width: 2, Height: 2})

	board.PlaceCell(mines.Location{X: 1, Y: 1})
	board.PlaceCell(mines.Location{X: 2, Y: 1})
	board.PlaceCell(mines.Location{X: 1, Y: 2})
	board.PlaceMine(mines.Location{X: 2, Y: 2})

	// Starting all mine cells are unflagged
	unflaggedMines := board.CountUnflaggedMines()
	expectedUnflaggedCells := 1

	if unflaggedMines != expectedUnflaggedCells {
		t.Fatalf(
			"expected to have %d unflagged mines. we have %d instead",
			expectedUnflaggedCells,
			unflaggedMines,
		)
	}

	// Flagging 2,2 will reduce the unflagged mines to none

	cell := board.Retrieve(mines.Location{X: 2, Y: 2})

	cell.Flag()

	unflaggedMines = board.CountUnflaggedMines()
	expectedUnflaggedCells = 0

	if unflaggedMines != expectedUnflaggedCells {
		t.Fatalf(
			"expected to have %d unflagged mines. we have %d instead",
			expectedUnflaggedCells,
			unflaggedMines,
		)
	}
}
