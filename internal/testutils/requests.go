package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/game"
)

func MustMakeNewGame(
	t *testing.T,
	s *http.Server,
	difficulty game.Difficulty,
) *res.MatchstateDto {
	t.Helper()

	newGameReq := NewRequest(
		http.MethodPost,
		"/match",
		&req.NewGameDto{
			Difficulty: difficulty,
		},
	)

	w := httptest.NewRecorder()

	s.Handler.ServeHTTP(w, newGameReq)

	if w.Code != http.StatusOK {
		t.Fatalf(
			"unexpected status code on 'POST /match'. got %d wanted %d",
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

	return matchstate
}
