package matchmaking

import (
	"fmt"

	"github.com/federico-paolillo/mines/internal/id"
	"github.com/federico-paolillo/mines/pkg/game"
)

type Matchmaker struct {
	storage   Store
	generator BoardGenerator
}

func NewMatchmaker(
	storage Store,
	generator BoardGenerator,
) *Matchmaker {
	return &Matchmaker{
		storage,
		generator,
	}
}

func (m *Matchmaker) New(
	startTime Matchstamp,
	difficulty game.Difficulty,
) (*Matchstate, error) {
	settings := game.GetDifficultySettings(difficulty)

	matchId := id.Generate()

	board := m.generator.Generate(settings.BoardSize, settings.NumberOfMines)

	game := game.NewGame(settings.Lives, board)

	version := NextVersion()

	match := NewMatch(
		matchId,
		version,
		startTime,
		board,
		game,
	)

	err := m.storage.Save(match)
	if err != nil {
		return nil, fmt.Errorf(
			"matchmaker: could not save match. %w",
			err,
		)
	}

	state := match.Status()

	return state, nil
}

func (m *Matchmaker) Apply(id string, move Move) (*Matchstate, error) {
	match, err := m.storage.Fetch(id)
	if err != nil {
		return nil, fmt.Errorf(
			"matchmaker: could not fetch match with id '%s' to apply move '%s'. %w",
			id,
			move.Type,
			err,
		)
	}

	err = match.Apply(move)
	if err != nil {
		return nil, fmt.Errorf(
			"matchmaker: failed to apply move '%s' to match with id '%s'. %w",
			move.Type,
			id,
			err,
		)
	}

	err = m.storage.Save(match)
	if err != nil {
		return nil, fmt.Errorf(
			"matchmaker: failed to save match '%s' after applying move '%s'. %w",
			move.Type,
			id,
			err,
		)
	}

	status := match.Status()

	return status, nil
}

func (m *Matchmaker) Get(id string) (*Matchstate, error) {
	match, err := m.storage.Fetch(id)
	if err != nil {
		return nil, fmt.Errorf(
			"matchmaker: could not fetch match with id '%s'. %w",
			id,
			err,
		)
	}

	state := match.Status()

	return state, nil
}
