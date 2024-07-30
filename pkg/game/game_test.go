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

func TestCanPlayAndWinGameExample(t *testing.T) {
	/* Grid 9x9:
	 * 1 M 1 0 0 1 1 1 0
	 * 2 2 1 0 0 2 M 2 0
	 * M 2 1 0 0 2 M 2 0
	 * 2 M 1 0 0 2 2 2 0
	 * 2 3 2 1 0 1 M 1 0
	 * M 2 M 1 0 1 2 2 1
	 * 2 3 1 1 0 0 1 M 1
	 * M 1 0 0 0 0 1 1 1
	 * 1 1 0 0 0 0 0 0 0
	 */

	bb := board.NewBuilder(dimensions.Size{Width: 9, Height: 9})

	bb.PlaceSafe(1, 1)
	bb.PlaceMine(2, 1)
	bb.PlaceSafe(3, 1)
	bb.PlaceSafe(4, 1)
	bb.PlaceSafe(5, 1)
	bb.PlaceSafe(6, 1)
	bb.PlaceSafe(7, 1)
	bb.PlaceSafe(8, 1)
	bb.PlaceSafe(9, 1)

	bb.PlaceSafe(1, 2)
	bb.PlaceSafe(2, 2)
	bb.PlaceSafe(3, 2)
	bb.PlaceSafe(4, 2)
	bb.PlaceSafe(5, 2)
	bb.PlaceSafe(6, 2)
	bb.PlaceMine(7, 2)
	bb.PlaceSafe(8, 2)
	bb.PlaceSafe(9, 2)

	bb.PlaceMine(1, 3)
	bb.PlaceSafe(2, 3)
	bb.PlaceSafe(3, 3)
	bb.PlaceSafe(4, 3)
	bb.PlaceSafe(5, 3)
	bb.PlaceSafe(6, 3)
	bb.PlaceMine(7, 3)
	bb.PlaceSafe(8, 3)
	bb.PlaceSafe(9, 3)

	bb.PlaceSafe(1, 4)
	bb.PlaceMine(2, 4)
	bb.PlaceSafe(3, 4)
	bb.PlaceSafe(4, 4)
	bb.PlaceSafe(5, 4)
	bb.PlaceSafe(6, 4)
	bb.PlaceSafe(7, 4)
	bb.PlaceSafe(8, 4)
	bb.PlaceSafe(9, 4)

	bb.PlaceSafe(1, 5)
	bb.PlaceSafe(2, 5)
	bb.PlaceSafe(3, 5)
	bb.PlaceSafe(4, 5)
	bb.PlaceSafe(5, 5)
	bb.PlaceSafe(6, 5)
	bb.PlaceMine(7, 5)
	bb.PlaceSafe(8, 5)
	bb.PlaceSafe(9, 5)

	bb.PlaceMine(1, 6)
	bb.PlaceSafe(2, 6)
	bb.PlaceMine(3, 6)
	bb.PlaceSafe(4, 6)
	bb.PlaceSafe(5, 6)
	bb.PlaceSafe(6, 6)
	bb.PlaceSafe(7, 6)
	bb.PlaceSafe(8, 6)
	bb.PlaceSafe(9, 6)

	bb.PlaceMine(1, 6)
	bb.PlaceSafe(2, 6)
	bb.PlaceMine(3, 6)
	bb.PlaceSafe(4, 6)
	bb.PlaceSafe(5, 6)
	bb.PlaceSafe(6, 6)
	bb.PlaceSafe(7, 6)
	bb.PlaceSafe(8, 6)
	bb.PlaceSafe(9, 6)

	bb.PlaceSafe(1, 7)
	bb.PlaceSafe(2, 7)
	bb.PlaceSafe(3, 7)
	bb.PlaceSafe(4, 7)
	bb.PlaceSafe(5, 7)
	bb.PlaceSafe(6, 7)
	bb.PlaceSafe(7, 7)
	bb.PlaceMine(8, 7)
	bb.PlaceSafe(9, 7)

	bb.PlaceMine(1, 8)
	bb.PlaceSafe(2, 8)
	bb.PlaceSafe(3, 8)
	bb.PlaceSafe(4, 8)
	bb.PlaceSafe(5, 8)
	bb.PlaceSafe(6, 8)
	bb.PlaceSafe(7, 8)
	bb.PlaceSafe(8, 8)
	bb.PlaceSafe(9, 8)

	bb.PlaceSafe(1, 9)
	bb.PlaceSafe(2, 9)
	bb.PlaceSafe(3, 9)
	bb.PlaceSafe(4, 9)
	bb.PlaceSafe(5, 9)
	bb.PlaceSafe(6, 9)
	bb.PlaceSafe(7, 9)
	bb.PlaceSafe(8, 9)
	bb.PlaceSafe(9, 9)

	b := bb.Build()

	g := game.NewGame(0, b)

	g.Open(5, 5)
	g.Open(7, 4)
	g.Open(2, 6)
	g.Open(8, 6)
	g.Open(7, 1)
	g.Open(1, 7)
	g.Open(1, 5)
	g.Open(1, 4)
	g.Open(2, 5)
	g.Open(1, 9)
	g.Open(2, 2)
	g.Open(2, 3)
	g.Open(1, 1)
	g.Open(1, 2)
	g.Open(8, 5)
	g.Open(9, 5)
	g.Open(9, 6)
	g.Open(9, 7)

	if g.Status() != game.Won {
		t.Fatalf("expected game to be won after the moves played")
	}
}
