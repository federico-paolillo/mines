package server

import "github.com/federico-paolillo/mines/pkg/mines"

type CronShutdown = func()

func maybeStartCron(mines *mines.Mines) CronShutdown {
	if mines.Cron == nil {
		return func() {}
	}

	mines.Logger.Info(
		"server: starting embedded cron",
	)

	mines.Cron.Start()

	silentShutdown := func() {
		// We purposefully ignore issues during shutdown. It don't really matter
		_ = mines.Cron.Shutdown()
	}

	return silentShutdown
}
