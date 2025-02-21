package reaper_test

import (
	"testing"

	"github.com/federico-paolillo/mines/internal/reaper"
	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/internal/testutils"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/stretchr/testify/require"
)

func TestReaperReapsCompletedMatches(t *testing.T) {
	memstore := storage.NewInMemoryStore()

	reaper := reaper.NewReaper(
		storage.NewReaperStore(
			memstore,
		),
	)

	m1 := testutils.SomeRandomMatch()
	m2 := testutils.SomeRandomMatch()

	s1 := m1.Status()
	s2 := m2.Status()

	s1.State = game.LostGame // We expect this match to expire

	memstore.Save(s1)
	memstore.Save(s2)

	result := reaper.Reap()

	_, err := memstore.Fetch(s1.Id)

	require.Error(t, storage.ErrNoSuchItem, err)

	require.Equal(t, 0, result.Expired)
	require.Equal(t, 1, result.Completed)
	require.Equal(t, 1, result.Ok)
}
