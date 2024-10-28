package matchmaking

import (
	"errors"
	"fmt"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/game"
)

type Matchversion = int64

var ErrIllegalMove = errors.New("unrecognized move")
var ErrGameHasEnded = errors.New("game has ended")

// A Match is a particular instance of a Game that is addressable by an unique identifier
type Match struct {
	Id      string
	Version Matchversion
	board   *board.Board
	game    *game.Game
}

func NewMatch(
	id string,
	version Matchversion,
	board *board.Board,
	game *game.Game,
) *Match {
	return &Match{
		id,
		version,
		board,
		game,
	}
}

func (m *Match) Status() *Matchstate {
	bSize := m.board.Size()

	cells := ExportCells(m.board)

	return &Matchstate{
		Id:     m.Id,
		Lives:  m.game.Lives(),
		Width:  bSize.Width,
		Height: bSize.Height,
		State:  m.game.Status(),
		Cells:  cells,
	}
}

func (m *Match) Apply(move Move) error {
	if m.game.Ended() {
		return fmt.Errorf(
			"match: move '%s' cannot be applied to match '%s'. %w",
			move.Type,
			m.Id,
			ErrGameHasEnded,
		)
	}

	switch move.Type {
	case MoveOpen:
		m.game.Open(move.X, move.Y)
		break
	case MoveFlag:
		m.game.Flag(move.X, move.Y)
		break
	case MoveChord:
		m.game.Chord(move.X, move.Y)
	default:
		return fmt.Errorf(
			"match: move '%s' cannot be applied to match '%s'. %w",
			move.Type,
			m.Id,
			ErrIllegalMove,
		)
	}

	return nil
}
