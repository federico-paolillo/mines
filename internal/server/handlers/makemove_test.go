package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/internal/testutils"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

func TestMakeMoveHandlerReturnsNewMatchstate(t *testing.T) {
	s := testutils.NewServer()

	matchstate := testutils.MustMakeNewGame(
		t,
		s,
		game.BeginnerDifficulty,
	)

	matchstate = testutils.MustMakeMove(
		t,
		s,
		matchstate.Id,
		matchmaking.MoveFlag,
		1,
		1,
	)

	currentState := matchstate.Cells[0][0].State // Remember that in Dto board origin is 0,0
	expectedState := board.FlaggedCell

	if currentState != expectedState {
		t.Errorf(
			"did not apply move. expected state to be '%s', it is '%s'",
			expectedState,
			currentState,
		)
	}
}

func TestMakeMoveHandlerReturnsNotFoundWhenProvidingJunkId(t *testing.T) {
	s := testutils.NewServer()

	w := httptest.NewRecorder()

	makeMoveReq := testutils.NewRequest(
		http.MethodPost,
		"/match/non-existing-id/move",
		req.MoveDto{
			Type: matchmaking.MoveChord,
			X:    1,
			Y:    1,
		},
	)

	s.Handler.ServeHTTP(w, makeMoveReq)

	if w.Code != http.StatusNotFound {
		t.Fatalf(
			"unexpected status code on 'GET /match/non-existing-id/move'. got '%d' wanted '%d'",
			w.Code,
			http.StatusOK,
		)
	}
}
