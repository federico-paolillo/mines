package game_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/game"
)

func TestChordingOpensAppropriateCell(t *testing.T) {
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

	game.Open(dimensions.Location{X: 1, Y: 1})

	openLocations := [6]dimensions.Location{
		{X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 3},
		{X: 1, Y: 4},
	}

	closeLocations := [6]dimensions.Location{
		{X: 2, Y: 2},
		{X: 3, Y: 2},
		{X: 2, Y: 3},
		{X: 3, Y: 3},
		{X: 2, Y: 4},
		{X: 3, Y: 4},
	}

	for _, location := range openLocations {
		cell := b.Retrieve(location)
		if cell.Status(board.Opened) != true {
			t.Errorf("expected cell at %v to be open", location)
		}
	}

	for _, location := range closeLocations {
		cell := b.Retrieve(location)
		if cell.Status(board.Closed) != true {
			t.Errorf("expected cell at %v to be closed", location)
		}
	}
}

func TestWhenOpeningAllSafeCellsTheGameIsWon(t *testing.T) {
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

	b := bb.Build()

	currentGame := game.NewGame(1, b)

	// We can leverage chording for this board. That is why we don't need all moves

	currentGame.Open(dimensions.Location{X: 1, Y: 1})
	currentGame.Open(dimensions.Location{X: 2, Y: 2})
	currentGame.Open(dimensions.Location{X: 3, Y: 2})
	currentGame.Open(dimensions.Location{X: 2, Y: 3})
	currentGame.Open(dimensions.Location{X: 2, Y: 4})
	currentGame.Open(dimensions.Location{X: 3, Y: 4})

	if currentGame.Status != game.Won {
		t.Fatal("expected game to be won after clearing all safe empty cells")
	}
}

func TestWhenFlaggingAllMineCellsTheGameIsWon(t *testing.T) {
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

	b := bb.Build()

	currentGame := game.NewGame(1, b)

	currentGame.Flag(dimensions.Location{X: 3, Y: 3})

	if currentGame.Status != game.Won {
		t.Fatal("expected game to be won after flagging all mined cells")
	}
}
