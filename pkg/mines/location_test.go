package mines_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/mines"
)

func TestOriginAdjacentLocationsAreCorrect(t *testing.T) {
	origin := mines.Location{X: 5, Y: 5}

	expectedLocations := [8]mines.Location{
		{X: 5, Y: 6},
		{X: 6, Y: 6},
		{X: 6, Y: 5},
		{X: 6, Y: 4},
		{X: 5, Y: 4},
		{X: 4, Y: 4},
		{X: 4, Y: 5},
		{X: 4, Y: 6},
	}

	locations := origin.AdjacentLocations()

	if locations != expectedLocations {
		t.Fatalf(
			"wrong adjacent locations for origin %v. expected %v but got %v",
			origin,
			expectedLocations,
			locations,
		)
	}
}
