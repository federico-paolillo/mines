package game_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/game"
)

func TestDifficultiesHaveProperSettings(t *testing.T) {
	expectedSettings := [3]struct {
		difficulty game.Difficulty
		lives      int
		size       dimensions.Size
		mines      int
	}{
		{
			difficulty: game.BeginnerDifficulty,
			size:       dimensions.Size{Width: 9, Height: 9},
			mines:      10,
			lives:      2,
		},
		{
			difficulty: game.IntermediateDifficulty,
			lives:      1,
			size:       dimensions.Size{Width: 16, Height: 16},
			mines:      40,
		},
		{
			difficulty: game.ExpertDifficulty,
			lives:      0,
			size:       dimensions.Size{Width: 30, Height: 16},
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
