package testutils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
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

func MustGetGame(
	t *testing.T,
	s *http.Server,
	matchId string,
) {
	t.Helper()

	w := httptest.NewRecorder()

	getGameReq := NewRequest(
		http.MethodGet,
		"/match/"+matchId,
		nil,
	)

	s.Handler.ServeHTTP(w, getGameReq)

	if w.Code != http.StatusOK {
		t.Fatalf(
			"unexpected status code on 'GET /match/%s'. got '%d' wanted '%d'",
			matchId,
			w.Code,
			http.StatusOK,
		)
	}
}

func MustMakeMove(
	t *testing.T,
	s *http.Server,
	matchId string,
	movetype matchmaking.Movetype,
	x int,
	y int,
) *res.MatchstateDto {
	t.Helper()

	w := httptest.NewRecorder()

	postMoveReq := NewRequest(
		http.MethodPost,
		"/match/"+matchId+"/move",
		req.MoveDto{
			X:    x,
			Y:    y,
			Type: movetype,
		},
	)

	s.Handler.ServeHTTP(w, postMoveReq)

	if w.Code != http.StatusOK {
		t.Fatalf(
			"unexpected status code on 'POST /match/%s/move'. got '%d' wanted '%d'",
			matchId,
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
