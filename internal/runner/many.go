package runner

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

/*
 * Channels aren't like files; you don't usually need to close them.
 * Closing is only necessary when the receiver must be told there are no more values coming.
 * We take advantage of this fact to share errChan among all program goroutines.
 * Even though a program goroutine should own the output channel (to close it) we make the main goroutin has owner.
 * There is no need for a program to tell main that no new errors will come. The channel is effectively shared.
 * As soon as one program terminates all programs will terminate. They share the same lifespan
 */

func RunMany(
	ctx context.Context,
	programs ...ProgramE,
) StatusCode {
	logger := slog.New(
		slog.NewJSONHandler(
			os.Stdout,
			nil,
		),
	)

	cfg, err := config.Load()
	if err != nil {
		logger.Error(
			"runner: failed to load configuration",
			slog.Any("err", err),
		)

		return NotOk
	}

	mines, err := mines.NewMines(
		logger,
		cfg,
	)
	if err != nil {
		logger.Error(
			"runner: failed to construct dependencies",
			slog.Any("err", err),
		)

		return NotOk
	}

	starTime := time.Now()

	statusCode := runManyPrograms(ctx, logger, mines, cfg, programs...)

	endTime := time.Now()
	runtimeDuration := endTime.Sub(starTime)
	runtimeInSeconds := runtimeDuration.Seconds()

	logger.Info(
		"runner: program has completed. This does not indicate success",
		slog.Float64("runtime_s", runtimeInSeconds),
		slog.Any("status_code", statusCode),
	)

	return statusCode
}
