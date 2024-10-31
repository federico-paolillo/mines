package matchstore

import (
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/game"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

func HydrateMatch(state *matchmaking.Matchstate) *matchmaking.Match {
	bb := board.NewBuilder(
		dimensions.Size{
			Width:  state.Width,
			Height: state.Height,
		},
	)

	for _, row := range state.Cells {
		for _, col := range row {
			hydrateCell(bb, col)
		}
	}

	b := bb.Build()

	return matchmaking.NewMatch(
		state.Id,
		state.Version,
		b,
		game.NewGame(state.Lives, b),
	)
}

func hydrateCell(
	bb *board.Builder,
	cell matchmaking.Cell,
) {
	if cell.Mined {
		_ = bb.PlaceMine(cell.X, cell.Y)
	} else {
		_ = bb.PlaceSafe(cell.X, cell.Y)
	}

	if cell.State == board.OpenCell {
		_ = bb.MarkOpen(cell.X, cell.Y)
	} else if cell.State == board.FlaggedCell {
		_ = bb.MarkFlagged(cell.X, cell.Y)
	}
}
