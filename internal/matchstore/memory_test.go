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

type versionlessState struct {
	Id     string
	Lives  int
	State  game.Gamestate
	Width  int
	Height int
	Cells  matchmaking.Cells
}

func TestMemorystoreLoadsMatches(t *testing.T) {
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

	memstore := matchstore.NewMemoryStore()

	err := memstore.Save(m)

	if err != nil {
		t.Fatalf(
			"failed to store match. %v",
			err,
		)
	}

	mRetrieved, err := memstore.Fetch("abc")

	if err != nil {
		t.Fatalf(
			"failed to load match. %v",
			err,
		)
	}

	state := m.Status()

	retrievedState := mRetrieved.Status()

	expectationWithoutVersion := versionlessState{
		state.Id,
		state.Lives,
		state.State,
		state.Width,
		state.Height,
		state.Cells,
	}

	actualWithoutVersion := versionlessState{
		retrievedState.Id,
		retrievedState.Lives,
		retrievedState.State,
		retrievedState.Width,
		retrievedState.Height,
		retrievedState.Cells,
	}

	stateEquals := reflect.DeepEqual(expectationWithoutVersion, actualWithoutVersion)

	if !stateEquals {
		t.Fatalf(
			"retrieving from memstore produced a different state. wanted\n%+v\ngot\n%+v\n",
			expectationWithoutVersion,
			actualWithoutVersion,
		)
	}
}
