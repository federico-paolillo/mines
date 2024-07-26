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

	bb := mines.NewBuilder(mines.Size{Width: 2, Height: 2})

	bb.PlaceSafe(1, 1)
	bb.PlaceMine(2, 1)
	bb.PlaceMine(1, 2)
	bb.PlaceSafe(2, 2)

	board := bb.Build()

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

	bb := mines.NewBuilder(mines.Size{Width: 3, Height: 4})

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

	board := bb.Build()

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

func TestUnopenSafeCellsCountIsCorrect(t *testing.T) {
	bb := mines.NewBuilder(mines.Size{Width: 2, Height: 2})

	bb.PlaceSafe(1, 1)
	bb.PlaceSafe(2, 1)
	bb.PlaceSafe(1, 2)
	bb.PlaceMine(2, 2)

	board := bb.Build()

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
	bb := mines.NewBuilder(mines.Size{Width: 2, Height: 2})

	bb.PlaceSafe(1, 1)
	bb.PlaceSafe(2, 1)
	bb.PlaceSafe(1, 2)
	bb.PlaceMine(2, 2)

	board := bb.Build()

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
