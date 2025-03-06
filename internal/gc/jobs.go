package gc

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/go-co-op/gocron/v2"
)

func scheduleReaperJob(
	mines *mines.Mines,
	cfg *config.Root,
	cron gocron.Scheduler,
) error {
	reap := func() {
		now := time.Now().Unix()

		reapStats := mines.Reaper.Reap(now)

		mines.Logger.Info(
			"gc: cleanup completed",
			slog.Int("expired", reapStats.Expired),
			slog.Int("completed", reapStats.Completed),
			slog.Int("ok", reapStats.Ok),
		)
	}

	frequency := time.Duration(cfg.Reaper.FrequencySeconds) * time.Second

	_, err := cron.NewJob(
		gocron.DurationJob(frequency),
		gocron.NewTask(reap),
		gocron.WithName("reaper"),
	)
	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak errors that are not under our control
		return fmt.Errorf(
			"server: could not schedule 'reaper' job. %v. %w",
			err,
			ErrSchedulerFailure,
		)
	}

	return nil
}
