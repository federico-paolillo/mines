package mines_test

import (
	"testing"

	"github.com/federico-paolillo/mines/internal/storage/memcached"
	"github.com/federico-paolillo/mines/internal/storage/memory"
	"github.com/federico-paolillo/mines/internal/testutils/slogt"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/stretchr/testify/require"
)

func TestCompositionRootBoostrapsDependencies(t *testing.T) {
	logger := slogt.New(t)

	cfg := config.Root{}

	mines, err := mines.NewMines(logger, &cfg)

	require.NoError(t, err, "failed to boostrap dependencies")

	require.Same(t, logger, mines.Logger, "bootstrapper did not use provided logger")

	require.NotNil(t, mines.Store)
	require.NotNil(t, mines.Matchmaker)
	require.NotNil(t, mines.MatchStore)
	require.NotNil(t, mines.Generator)
}

func TestCompositionRootUsesMemcachedWhenConfigured(t *testing.T) {
	logger := slogt.New(t)

	cfg := config.Root{
		Memcached: config.Memcached{
			Enabled: true,
			Servers: []string{"localhost:11211"},
		},
	}

	mines, err := mines.NewMines(logger, &cfg)

	require.NoError(t, err, "failed to boostrap dependencies")

	require.IsType(t, &memcached.Memcached{}, mines.Store)
}

func TestCompositionRootUsesMemoryWhenConfigured(t *testing.T) {
	logger := slogt.New(t)

	cfg := config.Root{
		Memcached: config.Memcached{
			Enabled: false,
			Servers: []string{},
		},
	}

	mines, err := mines.NewMines(logger, &cfg)

	require.NoError(t, err, "failed to boostrap dependencies")

	require.IsType(t, &memory.InMemoryStore{}, mines.Store)
}
