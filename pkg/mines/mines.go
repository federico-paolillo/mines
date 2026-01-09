package mines

import (
	"log/slog"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/federico-paolillo/mines/internal/generators"
	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/internal/storage/memcached"
	"github.com/federico-paolillo/mines/internal/storage/memory"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

// The composition root.
type Mines struct {
	Logger     *slog.Logger
	Store      storage.Store
	MatchStore matchmaking.Store
	Matchmaker *matchmaking.Matchmaker
	Generator  matchmaking.BoardGenerator
}

func NewMines(
	logger *slog.Logger,
	cfg *config.Root,
) (*Mines, error) {
	mines := &Mines{
		Logger: logger,
	}

	initGenerator(mines, cfg)
	initStores(mines, cfg)
	initMatchmaker(mines)

	return mines, nil
}

func initGenerator(mines *Mines, cfg *config.Root) {
	mines.Generator = generators.NewRngBoardGenerator(
		cfg.Seed,
	)
}

func initStores(mines *Mines, cfg *config.Root) {
	if cfg.Memcached.Enabled {
		client := memcache.New(
			cfg.Memcached.Servers...,
		)
		mines.Store = memcached.NewMemcached(
			client,
			time.Duration(cfg.TTL)*time.Hour,
		)
	} else {
		mines.Store = memory.NewInMemoryStore()
	}

	mines.MatchStore = storage.NewMatchStore(
		mines.Store,
	)
}

func initMatchmaker(mines *Mines) {
	mines.Matchmaker = matchmaking.NewMatchmaker(
		storage.NewMatchStore(
			mines.Store,
		),
		mines.Generator,
	)
}
