package matchmaking

import (
	"sync"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/game"
)

// A Match is a particular instance of a Game that is addressable by and unique identifier
type Match struct {
	Id    string
	mux   *sync.Mutex
	board *board.Board
	game  *game.Game
}

func NewMatch(
	id string,
	board *board.Board,
	game *game.Game,
) *Match {
	return &Match{
		id,
		&sync.Mutex{},
		board,
		game,
	}
}

func (m *Match) Status() Matchstate {
	m.mux.Lock()

	defer m.mux.Unlock()

	bSize := m.board.Size()

	cells := ExportCells(m.board)

	return Matchstate{
		Id:     m.Id,
		Lives:  m.game.Lives(),
		Width:  bSize.Width,
		Height: bSize.Height,
		State:  m.game.Status(),
		Cells:  cells,
	}
}
