package testutils

import (
	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

type FixedBoardGenerator struct {
	builder *board.Builder
}

func NewFixedBoardGenerator() *FixedBoardGenerator {
	builder := board.NewBuilder(
		dimensions.Size{
			Width:  3,
			Height: 3,
		},
	)

	prepareFixedBoard(builder)

	return &FixedBoardGenerator{
		builder,
	}
}

func (f *FixedBoardGenerator) Generate(size dimensions.Size, mines int) *board.Board {
	return f.builder.Build()
}

var _ matchmaking.BoardGenerator = (*FixedBoardGenerator)(nil)

func prepareFixedBoard(builder *board.Builder) {
	/*
	 * Assume a board like:
	 * 0 0 0
	 * 0 1 1
	 * 0 1 M
	 * where x is a closed empty cell
	 * 			 1 is a cell with adjacent mines
	 *       M is a mined cell
	 */
	_ = builder.PlaceSafe(1, 1)
	_ = builder.PlaceSafe(2, 1)
	_ = builder.PlaceSafe(3, 1)

	_ = builder.PlaceSafe(1, 2)
	_ = builder.PlaceSafe(2, 2)
	_ = builder.PlaceSafe(3, 2)

	_ = builder.PlaceSafe(1, 3)
	_ = builder.PlaceSafe(2, 3)
	_ = builder.PlaceMine(3, 3)
}
