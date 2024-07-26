package game_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/mines"
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

	game := game.NewGame(1, board)

	game.Open(mines.Location{X: 1, Y: 1})

	openLocations := [6]mines.Location{
		{X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 1, Y: 2},
		{X: 1, Y: 3},
		{X: 1, Y: 4},
	}

	closeLocations := [6]mines.Location{
		{X: 2, Y: 2},
		{X: 3, Y: 2},
		{X: 2, Y: 3},
		{X: 3, Y: 3},
		{X: 2, Y: 4},
		{X: 3, Y: 4},
	}

	for _, location := range openLocations {
		cell := board.Retrieve(location)
		if cell.Status(mines.Opened) != true {
			t.Errorf("expected cell at %v to be open", location)
		}
	}

	for _, location := range closeLocations {
		cell := board.Retrieve(location)
		if cell.Status(mines.Closed) != true {
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

	currentGame := game.NewGame(1, board)

	// We can leverage chording for this board. That is why we don't need all moves

	currentGame.Open(mines.Location{X: 1, Y: 1})
	currentGame.Open(mines.Location{X: 2, Y: 2})
	currentGame.Open(mines.Location{X: 3, Y: 2})
	currentGame.Open(mines.Location{X: 2, Y: 3})
	currentGame.Open(mines.Location{X: 2, Y: 4})
	currentGame.Open(mines.Location{X: 3, Y: 4})

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

	currentGame := game.NewGame(1, board)

	currentGame.Flag(mines.Location{X: 3, Y: 3})

	if currentGame.Status != game.Won {
		t.Fatal("expected game to be won after flagging all mined cells")
	}
}
