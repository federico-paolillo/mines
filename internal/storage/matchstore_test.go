package storage_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/internal/storage/memory"
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

func TestMatchStoreFetchesMatch(t *testing.T) {
	m := testutils.SomeMatch()

	memstore := storage.NewMatchStore(
		memory.NewInMemoryStore(),
	)

	err := memstore.Save(m)

	if err != nil {
		t.Fatalf(
			"failed to store match. %v",
			err,
		)
	}

	mRetrieved, err := memstore.Fetch(testutils.SomeMatchId)

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

func TestMatchStoreReturnsErrorWhenMatchDoesNotExist(t *testing.T) {
	memstore := storage.NewMatchStore(
		memory.NewInMemoryStore(),
	)
	_, err := memstore.Fetch(testutils.SomeMatchId)

	if !errors.Is(err, matchmaking.ErrNoSuchMatch) {
		t.Fatalf(
			"missing match returned wrong error. %v",
			err,
		)
	}
}

func TestMatchStoreRefusesToSaveMatchWithDifferentVersionThanStored(t *testing.T) {
	m := testutils.SomeMatch()

	memstore := storage.NewMatchStore(
		memory.NewInMemoryStore(),
	)

	err := memstore.Save(m)

	if err != nil {
		t.Fatalf(
			"failed to store match. %v",
			err,
		)
	}

	matchFromStore, err := memstore.Fetch(testutils.SomeMatchId)

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

	t.Log(matchFromStore.Version)

	if !errors.Is(err, matchmaking.ErrConcurrentUpdate) {
		t.Fatalf(
			"memorystore did not detect concurrent update. %v",
			err,
		)
	}
}
