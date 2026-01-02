package memory_test

import (
	"testing"

	"github.com/federico-paolillo/mines/internal/storage/memory"
)

func TestMatchVersionsAreIncremental(t *testing.T) {
	v1 := memory.NextVersion()
	v2 := memory.NextVersion()

	if v1 > v2 {
		t.Fatalf(
			"expected v2 to be greater than v1. v1: '%d' v2: '%d'",
			v1,
			v2,
		)
	}
}
