package mines

import (
	"log/slog"

	"github.com/federico-paolillo/mines/internal/generators"
	"github.com/federico-paolillo/mines/internal/storage"
	"github.com/federico-paolillo/mines/internal/storage/memory"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

// The composition root.
type Mines struct {
	Logger      *slog.Logger
	MemoryStore storage.Store
	MatchStore  matchmaking.Store
	Matchmaker  *matchmaking.Matchmaker
	Generator   matchmaking.BoardGenerator
}

func NewMines(
	logger *slog.Logger,
	cfg *config.Root,
) (*Mines, error) {
	mines := &Mines{
		Logger: logger,
	}

	initGenerator(mines, cfg)
	initStores(mines)
	initMatchmaker(mines)

	return mines, nil
}

func initGenerator(mines *Mines, cfg *config.Root) {
	mines.Generator = generators.NewRngBoardGenerator(
		cfg.Seed,
	)
}

func initStores(mines *Mines) {
	mines.MemoryStore = memory.NewInMemoryStore()

	mines.MatchStore = storage.NewMatchStore(
		mines.MemoryStore,
	)
}

func initMatchmaker(mines *Mines) {
	mines.Matchmaker = matchmaking.NewMatchmaker(
		storage.NewMatchStore(
			mines.MemoryStore,
		),
		mines.Generator,
	)
}
