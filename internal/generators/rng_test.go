package generators_test

import (
	"math/rand"
	"testing"

	"github.com/federico-paolillo/mines/internal/generators"
	"github.com/federico-paolillo/mines/pkg/dimensions"
)

func TestRngBoardGeneratorPlacesRequestedMines(t *testing.T) {
	iterations := 5
	minesWanted := 10

	size := dimensions.Size{Width: 9, Height: 9}

	for range iterations {
		seed := rand.Int()

		g := generators.NewRngBoardGenerator(seed)
		b := g.Generate(size, minesWanted)

		safeCells := b.CountUnopenSafeCells()
		minesGenerated := size.Area() - safeCells

		if minesGenerated != minesWanted {
			t.Fatalf(
				"expected %d mines, got %d. seed is %d",
				minesGenerated,
				minesWanted,
				seed,
			)
		}
	}
}
