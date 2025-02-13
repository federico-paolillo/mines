package cron

import (
	"errors"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

type CronShutdown func()

var ErrCronDisabled = errors.New("server: go-cron is disabled")

var ErrCronSchedulingFailure = errors.New("server: could not schedule a go-cron job")

var noOpShutdown CronShutdown = func() {}

func Start(
	mines *mines.Mines,
	cfg *config.Root,
) (CronShutdown, error) {
	if mines.Cron == nil {
		return noOpShutdown, ErrCronDisabled
	}

	err := scheduleReaperJob(mines, cfg)
	if err != nil {
		return noOpShutdown, err
	}

	mines.Cron.Start()

	// We purposefully ignore issues during shutdown. I don't really care
	silentShutdown := func() {
		_ = mines.Cron.Shutdown()
	}

	return silentShutdown, nil
}
