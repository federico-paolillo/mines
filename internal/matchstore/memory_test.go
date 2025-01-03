package matchstore_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/federico-paolillo/mines/internal/matchstore"
	"github.com/federico-paolillo/mines/internal/testutils"
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

func TestMemoryStoreLoadsMatches(t *testing.T) {
	m := testutils.SomeMatch()

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

func TestMemoryStoreReturnsErrorWhenMatchDoesNotExist(t *testing.T) {
	memstore := matchstore.NewMemoryStore()

	_, err := memstore.Fetch("abc")

	if !errors.Is(err, matchmaking.ErrNoSuchMatch) {
		t.Fatalf(
			"missing match returned wrong error. %v",
			err,
		)
	}
}

func TestMemoryStoreRefusesToSaveMatchWithDifferentVersionThanStored(t *testing.T) {
	m := testutils.SomeMatch()

	memstore := matchstore.NewMemoryStore()

	err := memstore.Save(m)

	if err != nil {
		t.Fatalf(
			"failed to store match. %v",
			err,
		)
	}

	matchFromStore, err := memstore.Fetch("abc")

	if err != nil {
		t.Fatalf(
			"failed to retrieve match. %v",
			err,
		)
	}

	_ = matchFromStore.Apply(
		matchmaking.Move{
			Type: matchmaking.MoveOpen,
			X:    1,
			Y:    1,
		},
	)

	err = memstore.Save(matchFromStore)

	if err != nil {
		t.Fatalf(
			"failed to save match. %v",
			err,
		)
	}

	// Now we change the original match and try to save it again

	_ = matchFromStore.Apply(
		matchmaking.Move{
			Type: matchmaking.MoveOpen,
			X:    3,
			Y:    1,
		},
	)

	err = memstore.Save(matchFromStore)

	if !errors.Is(err, matchmaking.ErrConcurrentUpdate) {
		t.Fatalf(
			"memorystore did not detect concurrent update. %v",
			err,
		)
	}
}
