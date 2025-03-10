package config_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/stretchr/testify/require"
)

func TestConfigurationLoadsConfigurationFromEnvVars(t *testing.T) {
	t.Setenv("MINES_SEED", "9999")
	t.Setenv("MINES_SERVER_HOST", "192.192.192.192")
	t.Setenv("MINES_SERVER_PORT", "3333")

	cfg, err := config.Load()

	require.NoError(t, err)

	require.Equal(t, 9999, cfg.Seed)
	require.Equal(t, "192.192.192.192", cfg.Server.Host)
	require.Equal(t, "3333", cfg.Server.Port)
}
