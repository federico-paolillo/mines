package config_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/stretchr/testify/require"
)

func TestConfigurationLoadsConfigurationFromEnvVars(t *testing.T) {
	t.Setenv("MINES_SEED", "9999")
	t.Setenv("MINES_TTL", "2")
	t.Setenv("MINES_SERVER_HOST", "192.192.192.192")
	t.Setenv("MINES_SERVER_PORT", "3333")
	t.Setenv("MINES_MEMCACHED_SERVERS", "pippolone.com:1234,pistolone.it:4321")
	t.Setenv("MINES_MEMCACHED_ENABLED", "true")

	cfg, err := config.Load()

	require.NoError(t, err)

	require.Equal(t, 9999, cfg.Seed)
	require.Equal(t, 2, cfg.TTL)
	require.Equal(t, "192.192.192.192", cfg.Server.Host)
	require.Equal(t, "3333", cfg.Server.Port)
	require.True(t, cfg.Memcached.Enabled)

	require.Len(t, cfg.Memcached.Servers, 2)

	require.Equal(t, "pippolone.com:1234", cfg.Memcached.Servers[0])
	require.Equal(t, "pistolone.it:4321", cfg.Memcached.Servers[1])
}
