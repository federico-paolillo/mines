package matchmaking

import (
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/game"
)

type Cells = [][]Cell

// Simplified Cell view of a Minesweeper board
type Cell struct {
	X, Y  int
	State board.Cellstate
}

// Summary of how a Match is doing
type Matchstate struct {
	Id     string
	Lives  int
	State  game.Gamestate
	Width  int
	Height int
	Cells  Cells
}