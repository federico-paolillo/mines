package mines

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/federico-paolillo/mines/internal/gc"
	"github.com/federico-paolillo/mines/internal/generators"
	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
	"github.com/federico-paolillo/mines/pkg/mines/config"
	"github.com/go-co-op/gocron/v2"
)

// The composition root
type Mines struct {
	Logger      *slog.Logger
	MemoryStore *storage.InMemoryStore
	MatchStore  matchmaking.Store
	Matchmaker  *matchmaking.Matchmaker
	Generator   matchmaking.BoardGenerator
	Reaper      *gc.Reaper
	GcStore     gc.Store
	Cron        gocron.Scheduler
}

func NewMines(
	logger *slog.Logger,
	cfg *config.Root,
) (*Mines, error) {
	mines := &Mines{
		Logger: logger,
	}

	initGenerator(mines, cfg)
	initMemoryStore(mines)
	initMatchmaker(mines)
	initReaper(mines)

	err := initCron(mines, cfg)
	if err != nil {
		return nil, fmt.Errorf(
			"init: could not setup composition root. %w",
			err,
		)
	}

	return mines, nil
}

func initGenerator(mines *Mines, cfg *config.Root) {
	mines.Generator = generators.NewRngBoardGenerator(
		cfg.Seed,
	)
}

func initMemoryStore(mines *Mines) {
	mines.MemoryStore = storage.NewInMemoryStore()

	mines.MatchStore = storage.NewMatchStore(
		mines.MemoryStore,
	)

	mines.GcStore = storage.NewGcStore(
		mines.MemoryStore,
	)
}

func initMatchmaker(mines *Mines) {
	mines.Matchmaker = matchmaking.NewMatchmaker(
		mines.MatchStore,
		mines.Generator,
	)
}

func initReaper(mines *Mines) {
	mines.Reaper = gc.NewReaper(
		mines.GcStore,
	)
}

func initCron(mines *Mines, cfg *config.Root) error {
	if !cfg.Reaper.Bundled {
		mines.Logger.Info("init: go-cron is disabled")

		return nil
	}

	cron, err := gocron.NewScheduler(
		gocron.WithLogger(mines.Logger),
		gocron.WithLimitConcurrentJobs(1, gocron.LimitModeReschedule),
		gocron.WithStopTimeout(time.Duration(cfg.Reaper.TimeoutSeconds)*time.Second),
	)
	if err != nil {
		return fmt.Errorf(
			"init: could not setup go-cron. %w",
			err,
		)
	}

	mines.Cron = cron

	return nil
}
