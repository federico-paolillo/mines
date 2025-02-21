package cron

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/federico-paolillo/mines/internal/reaper"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/go-co-op/gocron/v2"
)

func ScheduleReaperJob(
	mines *mines.Mines,
	cfg *config.Root,
	reaper *reaper.Reaper,
) error {
	reap := func() {
		reapStats := reaper.Reap()

		mines.Logger.Info(
			"gc: cleanup completed",
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
			err,
			ErrCronSchedulingFailure,
		)
	}

	return nil
}
