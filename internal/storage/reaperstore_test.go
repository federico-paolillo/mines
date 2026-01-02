package storage_test

import (
	"testing"

	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/internal/storage/memory"
	"github.com/federico-paolillo/mines/internal/testutils"
	"github.com/stretchr/testify/require"
)

func TestGcStoreReturnsAllMatchstates(t *testing.T) {
	memstore := memory.NewInMemoryStore()

	gcstore := storage.NewReaperStore(
		memstore,
	)

	m1 := testutils.SomeRandomMatch()
	m2 := testutils.SomeRandomMatch()

	s1 := m1.Status()
	s2 := m2.Status()

	memstore.Save(s1)
	memstore.Save(s2)

	actualIds := make([]string, 0, 2)

	for v := range gcstore.All() {
		actualIds = append(actualIds, v.Id)
	}

	expectedIds := []string{
		s1.Id,
		s2.Id,
	}

	require.Len(t, actualIds, 2)
	require.Subset(t, expectedIds, actualIds)
}

func TestGcDeletesMatchstates(t *testing.T) {
	memstore := memory.NewInMemoryStore()

	gcstore := storage.NewReaperStore(
		memstore,
	)

	m1 := testutils.SomeRandomMatch()
	m2 := testutils.SomeRandomMatch()

	s1 := m1.Status()
	s2 := m2.Status()

	memstore.Save(s1)
	memstore.Save(s2)

	gcstore.Delete(s2.Id)

	actualIds := make([]string, 0, 1)

	for v := range gcstore.All() {
		actualIds = append(actualIds, v.Id)
	}

	expectedIds := []string{
		s1.Id,
	}

	require.Len(t, actualIds, 1)
	require.Subset(t, expectedIds, actualIds)
}
