package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/internal/testutils"
	"github.com/federico-paolillo/mines/pkg/game"
)

func TestNewGameHandlerReturnsNewMatchWithProperConfiguration(t *testing.T) {
	s := testutils.NewServer()

	difficulties := []game.Difficulty{
		game.BeginnerDifficulty,
		game.ExpertDifficulty,
		game.IntermediateDifficulty,
	}

	for _, difficulty := range difficulties {
		matchstate := testutils.MustMakeNewGame(
			t,
			s,
			difficulty,
		)

		mustMatchDifficultySettings(
			t,
			matchstate,
			difficulty,
		)
	}
}

func TestNewGameHandlerRejectsUnknownDifficultyValues(t *testing.T) {
	s := testutils.NewServer()

	w := httptest.NewRecorder()

	req := testutils.NewRequest(
		http.MethodPost,
		"/match",
		&req.NewGameDto{
			Difficulty: "pippo",
		},
	)

	s.Handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf(
			"unexpected status code. got '%d' wanted '%d'",
			w.Code,
			http.StatusBadRequest,
		)
	}
}

func mustMatchDifficultySettings(
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
