package matchstore_test

import (
	"reflect"
	"testing"

	"github.com/federico-paolillo/mines/internal/matchstore"
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

func TestHydrationRestoresMatchProperly(t *testing.T) {
	bb := board.NewBuilder(
		dimensions.Size{
			Width:  2,
			Height: 1,
		},
	)

	bb.PlaceSafe(1, 1)
	bb.PlaceMine(2, 1)

	b := bb.Build()

	g := game.NewGame(12, b)

	m := matchmaking.NewMatch(
		"abc",
		123,
		b,
		g,
	)

	state := m.Status()

	mReborn := matchstore.HydrateMatch(state)

	rebornState := mReborn.Status()

	stateMatches := reflect.DeepEqual(state, rebornState)

	if !stateMatches {
		t.Fatalf(
			"hydration produced a different state. wanted\n%+v\ngot\n%+v\n",
			state,
			rebornState,
		)
	}
}
