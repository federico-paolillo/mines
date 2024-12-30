package runner_test

import (
	"errors"
	"testing"

	"github.com/federico-paolillo/mines/internal/runner"
	"github.com/federico-paolillo/mines/pkg/mines"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

func TestRunnerLoadsConfigurationFromEnvVars(t *testing.T) {
	t.Setenv("MINES_SEED", "9999")
	t.Setenv("MINES_SERVER_HOST", "192.192.192.192")
	t.Setenv("MINES_SERVER_PORT", "3333")

	testRunner := func(_ *mines.Mines, cfg *config.Root) error {
		t.Helper()

		if cfg.Seed != 9999 {
			t.Errorf(
				"seed is not correct. wanted '%d' got '%d'",
				9999,
				cfg.Seed,
			)
		}

		if cfg.Server.Host != "192.192.192.192" {
			t.Errorf(
				"server host is not correct. wanted '%s' got '%s'",
				"192.192.192.192",
				cfg.Server.Host,
			)
		}

		if cfg.Server.Port != "3333" {
			t.Errorf(
				"server port is not correct. wanted '%s' got '%s'",
				"3333",
				cfg.Server.Port,
			)
		}

		return nil
	}

	statusCode := runner.Run(testRunner)

	if statusCode != runner.Ok {
		t.Fatalf("program has failed with 'not ok'")
	}
}

func TestRunnerReturnsNotOkWhenProgramFails(t *testing.T) {
	testRunner := func(_ *mines.Mines, _ *config.Root) error {
		t.Helper()

		return errors.New("failure is expected")
	}

	statusCode := runner.Run(testRunner)

	if statusCode != runner.NotOk {
		t.Fatalf("runner was expected to return 'not ok'")
	}
}

func TestRunnerReturnsOkWhenProgramSucceeds(t *testing.T) {
	testRunner := func(_ *mines.Mines, _ *config.Root) error {
		t.Helper()
		return nil
	}

	statusCode := runner.Run(testRunner)

	if statusCode != runner.Ok {
		t.Fatalf("runner was expected to return 'ok'")
	}
}
