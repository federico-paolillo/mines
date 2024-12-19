package handlers_test

import (
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
