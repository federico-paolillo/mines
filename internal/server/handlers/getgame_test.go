package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federico-paolillo/mines/internal/testutils"
	"github.com/federico-paolillo/mines/pkg/game"
)

func TestGetGameHandlerReturnsMatchWhenProvidingSameId(t *testing.T) {
	s := testutils.NewServer()

	matchstate := testutils.MustMakeNewGame(
		t,
		s,
		game.BeginnerDifficulty,
	)

	testutils.MustGetGame(
		t,
		s,
		matchstate.Id,
	)
}

func TestGetGameHandlerReturnsNotFoundWhenProvidingMissingId(t *testing.T) {
	s := testutils.NewServer()

	w := httptest.NewRecorder()

	getGameReq := testutils.NewRequest(
		http.MethodGet,
		"/match/non-existing-id",
		nil,
	)

	s.Handler.ServeHTTP(w, getGameReq)

	if w.Code != http.StatusNotFound {
		t.Fatalf(
			"unexpected status code on 'GET /match/non-existing-id'. got '%d' wanted '%d'",
			w.Code,
			http.StatusOK,
		)
	}
}
