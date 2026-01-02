package memcached

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/internal/testutils"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/memcached"
)

var MemcachedImageVersion = "memcached:1.6.40"

func spinMemcachedCI(ctx context.Context, t *testing.T) *memcached.Container {
	t.Helper()
	t.Log("Spinning Memcached for CI env.")

	container, err := memcached.Run(
		ctx,
		MemcachedImageVersion,
	)
	if err != nil {
		t.Fatalf("Could not start memcached: %s", err)
	}

	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			t.Fatalf("Could not stop memcached: %s", err)
		}
	})

	return container
}

func spinMemcachedLocal(ctx context.Context, t *testing.T) *memcached.Container {
	t.Helper()
	t.Log("Spinning Memcached for MacOS/Podman env.")

	// See: https://golang.testcontainers.org/system_requirements/using_podman/#macos

	t.Setenv("TESTCONTAINERS_RYUK_CONTAINER_PRIVILEGED", "true")

	container, err := memcached.Run(
		ctx,
		MemcachedImageVersion,
		testcontainers.WithProvider(
			testcontainers.ProviderPodman,
		),
	)
	if err != nil {
		t.Fatalf("Could not start memcached: %s", err)
	}

	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			t.Fatalf("Could not stop memcached: %s", err)
		}
	})

	return container
}

func spinMemcached(t *testing.T) *Memcached {
	t.Helper()

	ctx := context.Background()

	var container *memcached.Container

	if os.Getenv("CI") == "true" {
		container = spinMemcachedCI(ctx, t)
	} else {
		container = spinMemcachedLocal(ctx, t)
	}

	endpoint, err := container.HostPort(ctx)

	require.NoError(t, err)

	client := memcache.New(endpoint)

	err = client.Ping()

	require.NoError(t, err)

	return NewMemcached(client, 2*time.Hour)
}

func TestMemcached(t *testing.T) {
	store := spinMemcached(t)

	t.Run("Fetch", func(t *testing.T) {
		t.Run("returns not found if match does not exist", func(t *testing.T) {
			_, err := store.Fetch("non-existing-id")

			require.ErrorIs(t, err, storage.ErrNoSuchItem)
		})

		t.Run("can fetch a just saved match", func(t *testing.T) {
			match := testutils.NewMatchState(t, "new-match", 0)

			err := store.Save(match)

			require.NoError(t, err)

			fetched, err := store.Fetch(match.Id)

			require.NoError(t, err)
			require.NotEqual(t, 0, fetched.Version)

			match.Version = fetched.Version // HACK: We put back the old version to use .Equal to compare contents

			require.Equal(t, match, fetched)
		})
	})

	t.Run("Save", func(t *testing.T) {
		t.Run("can save a new match", func(t *testing.T) {
			match := testutils.NewMatchState(t, "new-match-2", 0)

			err := store.Save(match)

			require.NoError(t, err)

			fetched, err := store.Fetch(match.Id)

			require.NoError(t, err)
			require.NotNil(t, fetched)
		})

		t.Run("can update an existing match", func(t *testing.T) {
			match := testutils.NewMatchState(t, "existing-match", 0)

			err := store.Save(match)

			require.NoError(t, err)

			fetched, err := store.Fetch(match.Id)

			require.NoError(t, err)
			require.NotEqual(t, 0, fetched.Version)

			originalVersion := fetched.Version

			fetched.Lives = 2

			err = store.Save(fetched)

			require.NoError(t, err)

			updated, err := store.Fetch(match.Id)

			require.NoError(t, err)
			require.Equal(t, fetched.Lives, updated.Lives)
			require.NotEqual(t, originalVersion, updated.Version)
		})

		t.Run("returns concurrent update error on cas mismatch", func(t *testing.T) {
			match := testutils.NewMatchState(t, "cas-mismatch", 0)

			err := store.Save(match)

			require.NoError(t, err)

			matchWithOldVersion, err := store.Fetch(match.Id)

			require.NoError(t, err)

			matchWithNewVersion, err := store.Fetch(match.Id)

			require.NoError(t, err)

			matchWithNewVersion.Lives = 2

			err = store.Save(matchWithNewVersion)

			require.NoError(t, err)

			// Check versions are in fact different

			require.NotEqual(t, matchWithOldVersion.Version, matchWithNewVersion.Version)

			// Now we try to update with the old version

			matchWithOldVersion.Lives = 4

			err = store.Save(matchWithOldVersion)

			require.ErrorIs(t, err, storage.ErrConcurrentUpdate)
		})
	})
}
