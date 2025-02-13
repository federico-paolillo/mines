package cron

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/go-co-op/gocron/v2"
)

func scheduleReaperJob(mines *mines.Mines, cfg *config.Root) error {
	reap := func() {
		reapStats := mines.Reaper.Reap()

		mines.Logger.Info(
			"reaper: task complete",
			slog.Int("expired", reapStats.Expired),
			slog.Int("completed", reapStats.Completed),
			slog.Int("ok", reapStats.Ok),
		)
	}

	frequency := time.Duration(cfg.Reaper.FrequencySeconds) * time.Second

	_, err := mines.Cron.NewJob(
		gocron.DurationJob(frequency),
		gocron.NewTask(reap),
	)
	if err != nil {
		return fmt.Errorf(
			"server: could not schedule reaper job. %w. %w",
			ErrCronSchedulingFailure,
			err,
		)
	}

	return nil
}
