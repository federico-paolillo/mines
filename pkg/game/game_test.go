package game_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/game"
)

func TestCascadingOpensAppropriateCell(t *testing.T) {
	/*
	 * Assume a board like:
	 * x x x
	 * x x x
	 * x x M
	 * x x x
	 * where x is a closed empty cell
	 *       M is a mined cell
	 * opening cell at 1,1 should cascade and produce a board like:
	 * # o o
	 * o 1 1
	 * o 1 M
	 * o 1 x
	 * where o is a cascading empty cell opened
	 *       # is the cell that was opened
	 * 			 1 is an open cell with adjacent mines
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

	game.Open(1, 1)

	openLocations := [10]dimensions.Location{
		{X: 1, Y: 1},
		{X: 2, Y: 1},
		{X: 3, Y: 1},
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 3, Y: 2},
		{X: 1, Y: 3},
		{X: 2, Y: 3},
		{X: 1, Y: 4},
		{X: 2, Y: 4},
	}

	closeLocations := [2]dimensions.Location{
		{X: 3, Y: 3},
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

func TestGameCanBeWon(t *testing.T) {
	/*
	 * Assume a board like:
	 * x x x
	 * x x x
	 * x x M
	 * x x x
	 * where x is a closed empty cell
	 *       M is a mined cell
	 * opening cell at 1,1 should cascade and produce a board like:
	 * # o o
	 * o 1 1
	 * o 1 M
	 * o 1 x
	 * where o is a cascading empty cell opened
	 *       # is the cell that was opened
	 * 			 1 is an open cell with adjacent mines
	 *       M is a mined cell
	 * opening cell at 3,4 should win the game:
	 * o o o
	 * o 1 1
	 * o 1 M
	 * o 1 #
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

	// We can leverage cascading for this board. That is why we don't need all moves

	currentGame.Open(1, 1)
	currentGame.Open(3, 4)

	if currentGame.Status() != game.Won {
		t.Fatal("expected game to be won after opening all safe empty cells")
	}
}

func TestGameCanBeLost(t *testing.T) {
	bb := board.NewBuilder(dimensions.Size{Width: 2, Height: 2})

	bb.PlaceSafe(1, 1)
	bb.PlaceSafe(2, 1)
	bb.PlaceSafe(1, 2)
	bb.PlaceMine(2, 2)

	b := bb.Build()

	currentGame := game.NewGame(0, b)

	// Opening a mine without lives loses the game

	currentGame.Open(2, 2)

	if currentGame.Status() != game.Lost {
		t.Fatal("expected game to be lost after tripping a mine")
	}
}

func TestGameDoesNotAllowFurtherMovesOnceGameIsWon(t *testing.T) {
	/*
	 * Assume a board like:
	 * x x x
	 * x x x
	 * x x M
	 * x x x
	 * where x is a closed empty cell
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

	// We can leverage cascading for this board. That is why we don't need all moves

	currentGame.Open(1, 1)
	currentGame.Open(3, 4)

	if currentGame.Status() != game.Won {
		t.Fatal("expected game to be won after opening all safe empty cells")
	}

	currentGame.Open(3, 3)

	cell := b.Retrieve(dimensions.Location{X: 3, Y: 3})

	if cell.Status(board.Opened) {
		t.Fatalf(
			"cell at %v should not have been openable because the game was won",
			cell.Position(),
		)
	}
}

func TestGameDoesNotAllowFurtherMovesOnceGameIsLost(t *testing.T) {
	/*
	 * Assume a board like:
	 * x x x
	 * x x x
	 * x x M
	 * x x x
	 * where x is a closed empty cell
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

	currentGame := game.NewGame(0, b)

	currentGame.Open(3, 3)

	if currentGame.Status() != game.Lost {
		t.Fatal("expected game to be won after opening all safe empty cells")
	}

	// We can leverage cascading for this board. That is why we don't need all moves

	currentGame.Open(1, 1)

	cell := b.Retrieve(dimensions.Location{X: 1, Y: 1})

	if cell.Status(board.Opened) {
		t.Fatalf(
			"cell at %v should not have been openable because the game was lost",
			cell.Position(),
		)
	}
}
