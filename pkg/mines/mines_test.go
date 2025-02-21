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
}

func TestCompositionRootDoesNotSetupCronWhenNotEmbedded(t *testing.T) {
	logger := slogt.New(t)

	cfg := config.Root{
		Reaper: config.Reaper{
			Bundled: false,
		},
	}

	mines, err := mines.NewMines(logger, &cfg)

	require.NoError(t, err, "failed to boostrap dependencies")

	require.Nil(t, mines.Cron, "cron was setup even if it was disabled")
}

func TestCompositionRootSetupsCronWhenEmbedded(t *testing.T) {
	logger := slogt.New(t)

	cfg := config.Root{
		Reaper: config.Reaper{
			Bundled:          true,
			TimeoutSeconds:   10,
			FrequencySeconds: 1,
		},
	}

	mines, err := mines.NewMines(logger, &cfg)

	require.NoError(t, err, "failed to boostrap dependencies")

	require.NotNil(t, mines.Cron, "cron was not setup even if it was enabled")

	_ = mines.Cron.Shutdown()
}
