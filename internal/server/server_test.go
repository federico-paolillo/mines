package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/federico-paolillo/mines/internal/server/req"
	"github.com/federico-paolillo/mines/internal/server/res"
	"github.com/federico-paolillo/mines/internal/testutils"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/gavv/httpexpect/v2"
)

func TestE2EYouCanCompleteWholeMatchUsingApi(t *testing.T) {
	minesServer := testutils.NewCustomServer(
		func(mines *mines.Mines) {
			mines.Matchmaker = matchmaking.NewMatchmaker(
				mines.Store,
				testutils.NewFixedBoardGenerator(),
			)
		},
	)

	httpServer := httptest.NewServer(minesServer.Handler)

	defer httpServer.Close()

	e := httpexpect.Default(t, httpServer.URL)

	// Create match
	var matchstate res.MatchstateDto

	e.POST("/match").
		WithJSON(req.NewGameDto{
			Difficulty: game.BeginnerDifficulty,
		}).
		Expect().
		Status(http.StatusOK).
		JSON().
		Decode(&matchstate)

	// Win match (see fixed board generator)
	e.POST("/match/" + matchstate.Id + "/move").
		WithJSON(req.MoveDto{
			Type: matchmaking.MoveOpen,
			X:    1,
			Y:    1,
		}).
		Expect().
		Status(http.StatusOK)

	// Query the match
	e.GET("/match/" + matchstate.Id).
		Expect().
		Status(http.StatusOK).
		JSON().
		Decode(&matchstate)

	if matchstate.State != game.WonGame {
		t.Fatal("match is still on-going")
	}
}
