package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federico-paolillo/mines/internal/testutils"
	"github.com/federico-paolillo/mines/pkg/game"
)

func TestGetGameHandlerReturnsMatchWhenProvidingSameId(t *testing.T) {
	t.Skip()

	s := testutils.NewServer()

	matchstate := testutils.MustMakeNewGame(t, s, game.BeginnerDifficulty)

	mustGetGame(t, s, matchstate.Id)
}

func mustGetGame(t *testing.T, s *http.Server, matchId string) {
	t.Helper()

	w := httptest.NewRecorder()

	getGameReq := testutils.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/match/%s", matchId),
		nil,
	)

	s.Handler.ServeHTTP(w, getGameReq)

	if w.Code != http.StatusOK {
		t.Fatalf(
			"unexpected status code on 'GET /match/%s'. got %d wanted %d",
			matchId,
			w.Code,
			http.StatusOK,
		)
	}
}
