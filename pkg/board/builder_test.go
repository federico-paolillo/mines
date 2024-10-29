package board_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

func TestBuilderBuildsBoardAccordingToInstructions(t *testing.T) {
	/*
	 * 2 M1 2
	 * 3 M2 3
	 * 2 M1 2
	 */

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

	expectations := [9]struct {
		location      dimensions.Location
		mined         bool
		adjacentMines int
	}{
		{dimensions.Location{X: 1, Y: 1}, false, 2},
		{dimensions.Location{X: 2, Y: 1}, true, 1},
		{dimensions.Location{X: 3, Y: 1}, false, 2},
		{dimensions.Location{X: 1, Y: 2}, false, 3},
		{dimensions.Location{X: 2, Y: 2}, true, 2},
		{dimensions.Location{X: 3, Y: 2}, false, 3},
		{dimensions.Location{X: 1, Y: 3}, false, 2},
		{dimensions.Location{X: 2, Y: 3}, true, 1},
		{dimensions.Location{X: 3, Y: 3}, false, 2},
	}

	for _, expectation := range expectations {
		cell := b.Retrieve(expectation.location)

		if cell.AdjacentMines() != expectation.adjacentMines {
			t.Errorf(
				"cell at %v should have %d adjacent mines. it has %d",
				cell.Position(),
				expectation.adjacentMines,
				cell.AdjacentMines(),
			)
		}

		if expectation.mined {
			if cell.Safe() {
				t.Errorf(
					"cell at %v should be mined",
					cell.Position(),
				)
			}
		} else {
			if cell.Mined() {
				t.Errorf(
					"cell at %v should not be mined",
					cell.Position(),
				)
			}
		}
	}
}

func TestAdjacentMinesCalculatesProperly(t *testing.T) {
	/*
	 * Assume a board like:
	 * 2 M
	 * M 2
	 * where x is a closed empty cell
	 * 			 1 is a cell with adjacent mines
	 *       M is a mined cell
	 */

	bb := board.NewBuilder(dimensions.Size{Width: 2, Height: 2})

	bb.PlaceSafe(1, 1)
	bb.PlaceMine(2, 1)
	bb.PlaceMine(1, 2)
	bb.PlaceSafe(2, 2)

	board := bb.Build()

	minesExpected := [4]struct {
		location dimensions.Location
		expected int
	}{
		{dimensions.Location{X: 1, Y: 1}, 2},
		{dimensions.Location{X: 2, Y: 1}, 1},
		{dimensions.Location{X: 1, Y: 2}, 1},
		{dimensions.Location{X: 2, Y: 2}, 2},
	}

	for _, expectation := range minesExpected {
		cell := board.Retrieve(expectation.location)

		if cell.AdjacentMines() != expectation.expected {
			t.Errorf(
				"expected cell at %v to have %d mines. instead it has %d",
				expectation.location,
				expectation.expected,
				cell.AdjacentMines(),
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
	 */

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

	board := bb.Build()

	minesExpected := [12]struct {
		location dimensions.Location
		expected int
	}{
		{dimensions.Location{X: 1, Y: 1}, 0},
		{dimensions.Location{X: 2, Y: 1}, 0},
		{dimensions.Location{X: 3, Y: 1}, 0},
		{dimensions.Location{X: 1, Y: 2}, 0},
		{dimensions.Location{X: 2, Y: 2}, 1},
		{dimensions.Location{X: 3, Y: 2}, 1},
		{dimensions.Location{X: 1, Y: 3}, 0},
		{dimensions.Location{X: 2, Y: 3}, 1},
		{dimensions.Location{X: 3, Y: 3}, 0},
		{dimensions.Location{X: 1, Y: 4}, 0},
		{dimensions.Location{X: 2, Y: 4}, 1},
		{dimensions.Location{X: 3, Y: 4}, 1},
	}

	for _, expectation := range minesExpected {
		cell := board.Retrieve(expectation.location)

		if cell.AdjacentMines() != expectation.expected {
			t.Errorf(
				"expected cell at %v to have %d mines. instead it has %d",
				expectation.location,
				expectation.expected,
				cell.AdjacentMines(),
			)
		}
	}
}

func TestBuilderDoesNotAllowOutOfBoundsLocations(t *testing.T) {
	b := board.NewBuilder(dimensions.Size{Width: 2, Height: 2})

	err := b.PlaceSafe(123, 23)

	if err == nil {
		t.Fatalf("expected builder to reject out of bounds placement. it did not")
	}
}

func TestBuilderDoesNotAllowOutOfBoundsLocations2(t *testing.T) {
	b := board.NewBuilder(dimensions.Size{Width: 2, Height: 2})

	err := b.PlaceSafe(0, 0)

	if err == nil {
		t.Fatalf("expected builder to reject out of bounds placement. it did not")
	}
}

func TestBuilderTogglesStateToOpen(t *testing.T) {
	bb := board.NewBuilder(dimensions.Size{Width: 1, Height: 1})

	bb.PlaceSafe(1, 1)
	bb.MarkOpen(1, 1)

	b := bb.Build()

	c := b.Retrieve(dimensions.Location{X: 1, Y: 1})

	if !c.HasStatus(board.OpenCell) {
		t.Fatalf(
			"cell at 1,1 should have been open. it is '%s'",
			c.Status(),
		)
	}
}

func TestBuilderTogglesStateToClosed(t *testing.T) {
	bb := board.NewBuilder(dimensions.Size{Width: 1, Height: 1})

	bb.PlaceSafe(1, 1)
	bb.MarkOpen(1, 1)

	b := bb.Build()

	c := b.Retrieve(dimensions.Location{X: 1, Y: 1})

	if !c.HasStatus(board.ClosedCell) {
		t.Fatalf(
			"cell at 1,1 should have been closed. it is '%s'",
			c.Status(),
		)
	}
}

func TestBuilderTogglesStateToFlagged(t *testing.T) {
	bb := board.NewBuilder(dimensions.Size{Width: 1, Height: 1})

	bb.PlaceSafe(1, 1)
	bb.MarkOpen(1, 1)

	b := bb.Build()

	c := b.Retrieve(dimensions.Location{X: 1, Y: 1})

	if !c.HasStatus(board.FlaggedCell) {
		t.Fatalf(
			"cell at 1,1 should have been flagged. it is '%s'",
			c.Status(),
		)
	}
}
