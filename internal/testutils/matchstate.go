package testutils

import (
	"testing"

	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/game"
)

func MustMatchDifficultySettings(
	t *testing.T,
	matchstate *res.MatchstateDto,
	difficulty game.Difficulty,
) {
	t.Helper()

	settings := game.GetDifficultySettings(difficulty)

	if matchstate.Lives != settings.Lives {
		t.Errorf(
			"unexpected number of lives. wanted %d got %d",
			settings.Lives,
			matchstate.Lives,
		)
	}

	if matchstate.State != game.PlayingGame {
		t.Errorf(
			"unexpected game state. wanted '%s' got '%s'",
			game.PlayingGame,
			matchstate.State,
		)
	}

	if matchstate.Height != settings.BoardSize.Height {
		t.Errorf(
			"unexpected height. wanted '%d' got '%d'",
			settings.BoardSize.Height,
			matchstate.Height,
		)
	}

	if matchstate.Width != settings.BoardSize.Width {
		t.Errorf(
			"unexpected width. wanted '%d' got '%d'",
			settings.BoardSize.Width,
			matchstate.Width,
		)
	}
}
