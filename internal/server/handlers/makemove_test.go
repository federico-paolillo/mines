package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/internal/testutils"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
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

	currentState := matchstate.Cells[0].State // Remember that in Dto board origin is 0,0
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
			"unexpected status code on 'POST /match/non-existing-id/move'. got '%d' wanted '%d'",
			w.Code,
			http.StatusOK,
		)
	}
}

func TestMakeMoveHandlerRejectsUnknownMoves(t *testing.T) {
	s := testutils.NewServer()

	matchstate := testutils.MustMakeNewGame(
		t,
		s,
		game.BeginnerDifficulty,
	)

	w := httptest.NewRecorder()

	makeMoveReq := testutils.NewRequest(
		http.MethodPost,
		"/match/"+matchstate.Id+"/move",
		req.MoveDto{
			Type: "crap",
			X:    1,
			Y:    1,
		},
	)

	s.Handler.ServeHTTP(w, makeMoveReq)

	if w.Code != http.StatusBadRequest {
		t.Fatalf(
			"unexpected status code on 'POST /match/%s/move'. got '%d' wanted '%d'",
			matchstate.Id,
			w.Code,
			http.StatusBadRequest,
		)
	}
}

func TestMakeMoveHandlerRejectsMovesForMatchesThatHaveEnded(t *testing.T) {
	s := testutils.NewServer()

	matchstate := testutils.MustMakeNewGame(
		t,
		s,
		game.BeginnerDifficulty,
	)

	w := httptest.NewRecorder()

	markMatchAsCompleted(
		t,
		s.Mines.MatchStore,
		matchstate.Id,
	)

	makeMoveReq := testutils.NewRequest(
		http.MethodPost,
		"/match/"+matchstate.Id+"/move",
		req.MoveDto{
			Type: matchmaking.MoveOpen,
			X:    1,
			Y:    1,
		},
	)

	s.Handler.ServeHTTP(w, makeMoveReq)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf(
			"unexpected status code on 'POST /match/%s/move'. got '%d' wanted '%d'",
			matchstate.Id,
			w.Code,
			http.StatusUnprocessableEntity,
		)
	}
}

func markMatchAsCompleted(
	t *testing.T,
	matchstore matchmaking.Store,
	matchId string,
) {
	t.Helper()

	m, _ := matchstore.Fetch(matchId)

	bb := board.NewBuilder(dimensions.Size{Width: 1, Height: 1})
	bb.PlaceSafe(1, 1)

	b := bb.Build()

	g := game.NewGame(1, b)
	g.Open(1, 1)

	if !g.Ended() {
		t.Fatalf("game did not end, you must have forgotten to apply some moves")
	}

	// We make a new match with a different board and game that will replace the original match

	m = matchmaking.NewMatch(
		matchId,
		m.Version,
		m.StartTime,
		b,
		g,
	)

	err := matchstore.Save(m)

	if err != nil {
		t.Fatalf(
			"could not store match. %v",
			err,
		)
	}
}
