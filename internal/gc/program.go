package gc

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

func Program(
	mines *mines.Mines,
	cfg *config.Root,
) error {
	scheduler, err := NewScheduler(
		mines,
		cfg,
	)
	if err != nil {
		return fmt.Errorf(
			"gc: failed to init scheduler. %w",
			err,
		)
	}

	scheduler.Start()

	mines.Logger.Info(
		"gc: scheduler started",
	)

	defer scheduler.Stop()

	sigc := make(chan os.Signal, 1)

	signal.Notify(sigc, os.Interrupt)

	<-sigc

	return nil
}
