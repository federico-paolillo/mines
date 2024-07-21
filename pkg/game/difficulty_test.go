package game_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/mines"
)

func TestDifficultiesHaveProperSettings(t *testing.T) {
	expectedSettings := [3]struct {
		difficulty game.Difficulty
		lives      int
		size       mines.Size
		mines      int
	}{
		{
			difficulty: game.Beginner,
			size:       mines.Size{Width: 9, Height: 9},
			mines:      10,
			lives:      2,
		},
		{
			difficulty: game.Intermediate,
			lives:      1,
			size:       mines.Size{Width: 16, Height: 16},
			mines:      40,
		},
		{
			difficulty: game.Expert,
			lives:      0,
			size:       mines.Size{Width: 30, Height: 16},
			mines:      99,
		},
	}

	for _, expectation := range expectedSettings {
		settings := game.GetDifficultySettings(expectation.difficulty)

		if settings.Lives != expectation.lives {
			t.Errorf(
				"expected difficulty '%s' to have %d lives. instead it has %d",
				expectation.difficulty,
				expectation.lives,
				settings.Lives,
			)
		}

		if settings.BoardSize != expectation.size {
			t.Errorf(
				"expected difficulty '%s' to have board size %v. instead it has size %v",
				expectation.difficulty,
				expectation.size,
				settings.BoardSize,
			)
		}

		if settings.NumberOfMines != expectation.mines {
			t.Errorf(
				"expected difficulty '%s' to have %d mines. instead it has %d",
				expectation.difficulty,
				expectation.mines,
				settings.NumberOfMines,
			)
		}
	}
}
