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
		w := httptest.NewRecorder()

		req := testutils.NewRequest(
			http.MethodPost,
			"/match",
			&req.NewGameDto{
				Difficulty: difficulty,
			},
		)

		s.Handler.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf(
				"unexpected status code. got %d wanted %d",
				w.Code,
				http.StatusOK,
			)
		}

		matchstate, err := testutils.Unmarshal[res.MatchstateDto](w.Body)
		if err != nil {
			t.Fatalf(
				"could not unmarshal response. %v",
				err,
			)
		}

		testutils.MustMatchDifficultySettings(
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
