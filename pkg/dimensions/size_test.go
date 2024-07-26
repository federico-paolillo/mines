package dimensions_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/dimensions"
)

func TestLocationsWithinBoundsAreDetected(t *testing.T) {
	size := dimensions.Size{Width: 10, Height: 10}
	location := dimensions.Location{X: 2, Y: 3}

	contains := size.Contains(location)

	if !contains {
		t.Fatalf(
			"expected size %v to contain location %v",
			size,
			location,
		)
	}
}

func TestLocationsOutOfBoundsAreDetected(t *testing.T) {
	size := dimensions.Size{Width: 10, Height: 10}
	location := dimensions.Location{X: 11, Y: -3}

	contains := size.Contains(location)

	if contains {
		t.Fatalf(
			"expected size %v to NOT contain location %v",
			size,
			location,
		)
	}
}
