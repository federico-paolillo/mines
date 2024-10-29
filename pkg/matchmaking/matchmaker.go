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

func (m *Matchmaker) New(difficulty game.Difficulty) (*Match, error) {
	settings := game.GetDifficultySettings(difficulty)

	matchId, err := id.Generate()

	if err != nil {
		return nil, fmt.Errorf(
			"matchmaker: could not generate match. %w",
			err,
		)
	}

	board := m.generator.Generate(settings.BoardSize, settings.NumberOfMines)

	game := game.NewGame(settings.Lives, board)

	version := NextVersion()

	match := NewMatch(
		matchId,
		version,
		board,
		game,
	)

	err = m.storage.Save(match)

	if err != nil {
		return nil, fmt.Errorf(
			"matchmaker: could not save match. %w",
			err,
		)
	}

	return match, nil
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

	status := match.Status()

	return status, nil
}
