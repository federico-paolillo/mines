package gc

import (
	"context"
	"fmt"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

func Program(
	ctx context.Context,
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

	<-ctx.Done()

	return nil
}
