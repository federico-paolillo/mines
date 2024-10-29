package matchmaking_test

import (
	"testing"

	"github.com/federico-paolillo/mines/pkg/matchmaking"
)

func TestMatchVersionsAreIncremental(t *testing.T) {
	v1 := matchmaking.NextVersion()
	v2 := matchmaking.NextVersion()

	if v1 > v2 {
		t.Fatalf(
			"expected v2 to be greater than v1. v1: '%d' v2: '%d'",
			v1,
			v2,
		)
	}
}
