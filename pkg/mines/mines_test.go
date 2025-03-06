package mines_test

import (
	"testing"

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

	require.NotNil(t, mines.MemoryStore)
	require.NotNil(t, mines.Matchmaker)
	require.NotNil(t, mines.MatchStore)
	require.NotNil(t, mines.Generator)
	require.NotNil(t, mines.ReaperStore)
	require.NotNil(t, mines.Reaper)
}
