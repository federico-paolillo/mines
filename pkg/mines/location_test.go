package mines_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/mines"
)

func TestOriginAdjacentLocationsAreCorrect(t *testing.T) {
	origin := mines.Location{5, 5}

	expectedLocations := [8]mines.Location{
		{5, 6},
		{6, 6},
		{6, 5},
		{6, 4},
		{5, 4},
		{4, 4},
		{4, 5},
		{4, 6},
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
