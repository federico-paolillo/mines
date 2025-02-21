package gc

import (
	"fmt"

	"github.com/federico-paolillo/mines/internal/gc/cron"
	"github.com/federico-paolillo/mines/internal/reaper"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

func Program(
	mines *mines.Mines,
	cfg *config.Root,
) error {
	reaper := reaper.NewReaper(
		mines.ReaperStore,
	)

	cronShutdown, err := cron.Start(mines, cfg, reaper)
	if err != nil {
		return fmt.Errorf(
			"gc: failed to start reaper job. %w",
			err,
		)
	}

	defer cronShutdown()

	mines.Logger.Info(
		"gc: reaper started",
	)

	return nil
}
