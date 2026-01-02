package memory_test

import (
	"errors"
	"testing"

	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/internal/storage/memory"
	"github.com/federico-paolillo/mines/internal/testutils"
)

func TestMemoryStoreSavesMatchstate(t *testing.T) {
	m := testutils.SomeMatch()

	memstore := memory.NewInMemoryStore()

	s := m.Status()

	err := memstore.Save(s)

	if err != nil {
		t.Fatalf(
			"failed to store match. %v",
			err,
		)
	}

	fm, err := memstore.Fetch(testutils.SomeMatchId)

	if err != nil {
		t.Fatalf(
			"failed to retrieve match. %v",
			err,
		)
	}

	if fm.Id != testutils.SomeMatchId {
		t.Fatalf(
			"retrieved wrong match. expected id '%s' but got '%s'",
			testutils.SomeMatchId,
			fm.Id,
		)
	}
}

func TestMemoryStoreForbidsConcurrentUpdates(t *testing.T) {
	m1 := testutils.SomeMatch()

	// Save once, change the same instance and save again

	memstore := memory.NewInMemoryStore()

	s1 := m1.Status()

	err := memstore.Save(s1)

	if err != nil {
		t.Fatalf(
			"failed to store match. %v",
			err,
		)
	}

	s1.Lives = 14

	err = memstore.Save(s1)

	if err == nil {
		t.Fatalf(
			"store did not detect concurrent updates",
		)
	}

	if !errors.Is(err, storage.ErrConcurrentUpdate) {
		t.Fatalf(
			"failed to store match. %v",
			err,
		)
	}
}

func TestMemoryStoreErrorsWhenMatchstateIsMissing(t *testing.T) {
	memstore := memory.NewInMemoryStore()

	fm, err := memstore.Fetch("doesnotexist")

	if err == nil {
		t.Fatalf(
			"non existing match fetched without errors",
		)
	}

	if fm != nil {
		t.Fatalf(
			"returned a match that does not exist",
		)
	}

	if !errors.Is(err, storage.ErrNoSuchItem) {
		t.Fatalf(
			"match fetching returned wrong error. %v",
			err,
		)
	}
}
