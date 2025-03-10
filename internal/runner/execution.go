package runner

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"

	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

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
