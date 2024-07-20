package mines_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/mines"
)

func TestLocationsWithinBoundsAreDetected(t *testing.T) {
	size := mines.Size{10, 10}
	location := mines.Location{2, 3}

	contains := size.Contains(location)

	if !contains {
		t.Fatalf(
			"expected size %v to contain location %v. it did not.",
			size,
			location,
		)
	}
}

func TestLocationsOutOfBoundsAreDetected(t *testing.T) {
	size := mines.Size{10, 10}
	location := mines.Location{11, -3}

	contains := size.Contains(location)

	if contains {
		t.Fatalf(
			"expected size %v to NOT contain location %v. it did.",
			size,
			location,
		)
	}
}
