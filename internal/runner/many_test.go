package runner_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/federico-paolillo/mines/internal/runner"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/stretchr/testify/require"
)

func TestRunnerReturnsNotOkWhenProgramFails(t *testing.T) {
	ctx := t.Context()

	testProgram := func(
		_ context.Context,
		_ *mines.Mines,
		_ *config.Root,
	) error {
		t.Helper()

		return errors.New("failure is expected")
	}

	statusCode := runner.RunMany(ctx, testProgram)

	require.Equal(t, runner.NotOk, statusCode)
}

func TestRunnerReturnsOkWhenProgramSucceeds(t *testing.T) {
	ctx := t.Context()

	testProgram := func(
		_ context.Context,
		_ *mines.Mines,
		_ *config.Root,
	) error {
		t.Helper()
		return nil
	}

	statusCode := runner.RunMany(ctx, testProgram)

	require.Equal(t, runner.Ok, statusCode)
}

func TestRunnerReturnsOkWhenAllProgramSucceeds(t *testing.T) {
	ctx := t.Context()

	testProgram1 := func(
		_ context.Context,
		_ *mines.Mines,
		_ *config.Root,
	) error {
		t.Helper()
		return nil
	}

	testProgram2 := func(
		_ context.Context,
		_ *mines.Mines,
		_ *config.Root,
	) error {
		t.Helper()
		return nil
	}

	statusCode := runner.RunMany(ctx, testProgram1, testProgram2)

	require.Equal(t, runner.Ok, statusCode)
}

func TestRunnerReturnsOkWhenOneProgramFails(t *testing.T) {
	ctx := t.Context()

	testProgram1 := func(
		_ context.Context,
		_ *mines.Mines,
		_ *config.Root,
	) error {
		t.Helper()
		time.Sleep(1 * time.Second) // Simulate work to let the other program fail
		return nil
	}

	testProgram2 := func(
		_ context.Context,
		_ *mines.Mines,
		_ *config.Root,
	) error {
		t.Helper()
		return errors.New("failure is expected")
	}

	statusCode := runner.RunMany(ctx, testProgram1, testProgram2)

	require.Equal(t, runner.NotOk, statusCode)
}
