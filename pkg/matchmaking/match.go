package matchmaking

import (
	"errors"
	"fmt"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/game"
)

var (
	ErrIllegalMove  = errors.New("match: unrecognized move")
	ErrGameHasEnded = errors.New("match: game has ended")
)

// A Match is a particular instance of a Game that is addressable by an unique identifier.
type Match struct {
	Id        string
	Version   uint64
	StartTime int64
	board     *board.Board
	game      *game.Game
}

func NewMatch(
	id string,
	version uint64,
	startTime int64,
	board *board.Board,
	game *game.Game,
) *Match {
	return &Match{
		id,
		version,
		startTime,
		board,
		game,
	}
}

func (m *Match) Status() *Matchstate {
	bSize := m.board.Size()

	cells := ExportCells(m.board)

	return &Matchstate{
		m.Id,
		m.Version,
		m.game.Lives(),
		m.game.Status(),
		bSize.Width,
		bSize.Height,
		cells,
		m.StartTime,
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
	case MoveFlag:
		m.game.Flag(move.X, move.Y)
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
