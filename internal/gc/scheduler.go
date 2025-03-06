package gc

import (
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/go-co-op/gocron/v2"
)

var ErrSchedulerFailure = errors.New("scheduler: could not schedule")

type Scheduler struct {
	logger *slog.Logger
	cron   gocron.Scheduler
}

func NewScheduler(
	mines *mines.Mines,
	cfg *config.Root,
) (*Scheduler, error) {
	cron, err := gocron.NewScheduler(
		gocron.WithLogger(mines.Logger),
		gocron.WithLimitConcurrentJobs(1, gocron.LimitModeReschedule),
		gocron.WithStopTimeout(time.Duration(cfg.Reaper.TimeoutSeconds)*time.Second),
	)
	if err != nil {
		//nolint:errorlint // We do not want to wrap and leak errors that are not under our control
		return nil, fmt.Errorf(
			"scheduler: could not setup gocron. %v. %w",
			err,
			ErrSchedulerFailure,
		)
	}

	err = scheduleReaperJob(mines, cfg, cron)
	if err != nil {
		return nil, fmt.Errorf(
			"scheduler: could not schedule a job. %w",
			err,
		)
	}

	return &Scheduler{
		mines.Logger,
		cron,
	}, nil
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	err := s.cron.Shutdown()
	if err != nil {
		s.logger.Error(
			"scheduler: could not stop gocron",
			slog.Any("err", err),
		)
	}
}
