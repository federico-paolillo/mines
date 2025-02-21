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

	r := reaper.NewReaper(
		storage.NewReaperStore(
			memstore,
		),
	)

	m1 := testutils.SomeCustomMatch("pluto", 123, 9000)
	m2 := testutils.SomeCustomMatch("pippo", 123, 2000) // We expect this match to be expired

	s1 := m1.Status()
	s2 := m2.Status()

	s1.State = game.LostGame // We expect this match to be complete

	memstore.Save(s1)
	memstore.Save(s2)

	result := r.Reap(10_000)

	_, err := memstore.Fetch(s1.Id)

	require.Error(t, storage.ErrNoSuchItem, err)

	require.Equal(t, 1, result.Expired, "wrong number of 'expired' matches")
	require.Equal(t, 1, result.Completed, "wrong number of 'completed' matches")
	require.Equal(t, 0, result.Ok, "wrong number of 'ok' matches")
}
