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
