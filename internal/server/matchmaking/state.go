package matchmaking

import (
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/game"
)

type Cell struct {
	X, Y  int
	State board.Cellstate
}

type Matchstate struct {
	Id     string
	Lives  int
	State  game.Gamestate
	Width  int
	Height int
	Cells  [][]Cell
}
