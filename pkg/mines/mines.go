package mines

import (
	"log/slog"

	"github.com/federico-paolillo/mines/internal/generators"
	"github.com/federico-paolillo/mines/internal/matchstore"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
	"github.com/federico-paolillo/mines/pkg/mines/config"
)

// The composition root
type Mines struct {
	Logger     *slog.Logger
	Store      matchmaking.Store
	Matchmaker *matchmaking.Matchmaker
	Generator  matchmaking.BoardGenerator
}

func NewMines(
	logger *slog.Logger,
	cfg *config.Root,
) *Mines {
	mines := &Mines{
		Logger: logger,
	}

	initGenerator(mines, cfg)
	initMemoryStore(mines)
	initMatchmaker(mines)

	return mines
}

func initGenerator(mines *Mines, cfg *config.Root) {
	mines.Generator = generators.NewRngBoardGenerator(
		cfg.Seed,
	)
}

func initMemoryStore(mines *Mines) {
	mines.Store = matchstore.NewMemoryStore()
}

func initMatchmaker(mines *Mines) {
	mines.Matchmaker = matchmaking.NewMatchmaker(
		mines.Store,
		mines.Generator,
	)
}
