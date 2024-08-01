package generators

import (
	"math/rand"

	"github.com/federico-paolillo/mines/pkg/board"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

type rngBoardGenerator struct {
	seed int
	rng  *rand.Rand
}

func NewRngBoardGenerator(seed int) *rngBoardGenerator {
	rng := rand.New(rand.NewSource(int64(seed)))

	return &rngBoardGenerator{
		seed: seed,
		rng:  rng,
	}
}

func (gen *rngBoardGenerator) Generate(size dimensions.Size, mines int) *board.Board {
	// Dumb algorithm: fill the board with safe cells then replace those with n mines
	bb := board.NewBuilder(size)

	for y := range size.Height {
		for x := range size.Width {
			bb.PlaceSafe(x+1, y+1)
		}
	}

	for range mines {
		for {
			x := gen.rng.Intn(size.Width) + 1
			y := gen.rng.Intn(size.Height) + 1

			if bb.IsMine(x, y) { // More dumb stuff: try random cell, if is mined try again
				continue
			}

			bb.PlaceMine(x, y)

			break
		}
	}

	return bb.Build()
}
