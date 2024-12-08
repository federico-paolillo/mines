package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/game"
)

func EnsureNewGameResponseMatchesDifficultySettings(
	t *testing.T,
	difficulty game.Difficulty,
	w *httptest.ResponseRecorder,
) {
	t.Helper()

	if w.Code != http.StatusOK {
		t.Fatalf(
			"unexpected status code. got %d wanted %d",
			w.Code,
			http.StatusOK,
		)
	}

	matchstate, err := Unmarshal[res.MatchstateDto](w.Body)
	if err != nil {
		t.Fatalf(
			"could not unmarshal response. %v",
			err,
		)
	}

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
