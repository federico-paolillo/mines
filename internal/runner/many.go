package runner

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
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

func runManyPrograms(
	ctx context.Context,
	logger *slog.Logger,
	mines *mines.Mines,
	cfg *config.Root,
	programs ...ProgramE,
) StatusCode {
	programsCount := len(programs)

	// Allow cancellation via SIGTERM

	ctx, cancel := context.WithCancel(ctx)

	sigtermChan := make(chan os.Signal, 1)

	signal.Notify(sigtermChan, os.Interrupt)

	// We buffer for n-programs. So that each program can report a failure without blocking
	// TODO: We currently consume just the first program error

	errChan := make(chan error, programsCount)

	var wg sync.WaitGroup

	wg.Add(programsCount)

	for _, p := range programs {
		go runOneProgram(
			ctx,
			&wg,
			errChan,
			mines,
			cfg,
			p,
		)
	}

	err := waitForTermination(
		cancel,
		errChan,
		sigtermChan,
		&wg,
	)
	if err != nil {
		logger.Error(
			"runner: failed to run a program",
			slog.Any("err", err),
		)

		return NotOk
	}

	return Ok
}

func waitForTermination(
	cancel context.CancelFunc,
	errChan chan error,
	sigtermChan chan os.Signal,
	wg *sync.WaitGroup,
) error {
	var err error

	// As soon as one program fails, finishes, or the context is cancelled we give up execution
	// TODO: We should consume all errors. The channel has a finite number of values

	select {
	case err = <-errChan:
		cancel()
	case <-sigtermChan:
		cancel()
	}

	// Once we signalled cancellation we wait for all goroutines to clean up before returning

	wg.Wait()

	return err
}

func runOneProgram(
	ctx context.Context,
	wg *sync.WaitGroup,
	errChan chan<- error,
	mines *mines.Mines,
	cfg *config.Root,
	program ProgramE,
) {
	defer wg.Done()

	// errChan has one buffer space for each program. We don't have to worry about blocking write

	errChan <- program(ctx, mines, cfg)
}
