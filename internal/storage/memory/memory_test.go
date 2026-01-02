package memory_test

import (
	"errors"
	"slices"
	"testing"

	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/internal/storage/memory"
	"github.com/federico-paolillo/mines/internal/testutils"
)

func TestMemoryStoreDeletesMatchstate(t *testing.T) {
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

	_, err = memstore.Fetch(testutils.SomeMatchId)

	if err != nil {
		t.Fatalf(
			"failed to retrieve match. %v",
			err,
		)
	}

	memstore.Delete(testutils.SomeMatchId)

	_, err = memstore.Fetch(testutils.SomeMatchId)

	if !errors.Is(err, storage.ErrNoSuchItem) {
		if err != nil {
			t.Fatalf(
				"failed to retrieve match, but for the wrong reason. %v",
				err,
			)
		}
	}
}

func TestMemoryStoreRangesOverAllMatchstates(t *testing.T) {
	m1 := testutils.SomeCustomMatch("blabla", 1234, 4567)
	m2 := testutils.SomeCustomMatch("gnegne", 1234, 4567)

	memstore := memory.NewInMemoryStore()

	s1 := m1.Status()
	s2 := m2.Status()

	var err error

	err = errors.Join(memstore.Save(s1))
	err = errors.Join(memstore.Save(s2))

	if err != nil {
		t.Fatalf(
			"failed to store match. %v",
			err,
		)
	}

	idsCollected := make([]string, 0, 2)

	for v := range memstore.All() {
		idsCollected = append(idsCollected, v.Id)
	}

	idsCount := len(idsCollected)

	idsExpected := []string{
		s1.Id,
		s2.Id,
	}

	slices.Sort(idsCollected)
	slices.Sort(idsExpected)

	idsMatch := slices.Compare(idsCollected, idsExpected) == 0

	if !idsMatch {
		t.Errorf(
			"expected '%+v' to match '%+v'",
			idsCollected,
			idsExpected,
		)
	}

	if idsCount != 2 {
		t.Fatalf(
			"expected '2' matches. got '%d'",
			idsCount,
		)
	}
}

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
